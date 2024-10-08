#pragma once

#include <memory>
#include <typeinfo>

#include "envoy/config/route/v3/scoped_route.pb.h"
#include "envoy/extensions/filters/network/http_connection_manager/v3/http_connection_manager.pb.h"
#include "envoy/router/rds.h"
#include "envoy/router/router.h"
#include "envoy/router/scopes.h"
#include "envoy/thread_local/thread_local.h"

#include "source/common/common/hash.h"
#include "source/common/protobuf/utility.h"
#include "source/common/router/config_impl.h"

#include "absl/numeric/int128.h"
#include "absl/strings/str_format.h"

namespace Envoy {
namespace Router {

using envoy::extensions::filters::network::http_connection_manager::v3::ScopedRoutes;

#if defined(ALIMESH)
using ReComputeCb = std::function<std::unique_ptr<ScopeKeyFragmentBase>()>;
using ReComputeCbPtr = std::shared_ptr<ReComputeCb>;
using ReComputeCbWeakPtr = std::weak_ptr<ReComputeCb>;
#endif

/**
 * Base class for fragment builders.
 */
class FragmentBuilderBase {
public:
  explicit FragmentBuilderBase(ScopedRoutes::ScopeKeyBuilder::FragmentBuilder&& config)
      : config_(std::move(config)) {}
  virtual ~FragmentBuilderBase() = default;

#if defined(ALIMESH)
  virtual std::unique_ptr<ScopeKeyFragmentBase>
  computeFragment(const Http::HeaderMap& headers, const StreamInfo::StreamInfo* info,
                  ReComputeCbPtr& recompute) const PURE;
#else
  // Returns a fragment if the fragment rule applies, a nullptr indicates no fragment could be
  // generated from the headers.
  virtual std::unique_ptr<ScopeKeyFragmentBase>
  computeFragment(const Http::HeaderMap& headers) const PURE;
#endif

protected:
  const ScopedRoutes::ScopeKeyBuilder::FragmentBuilder config_;
};

class HeaderValueExtractorImpl : public FragmentBuilderBase {
public:
  explicit HeaderValueExtractorImpl(ScopedRoutes::ScopeKeyBuilder::FragmentBuilder&& config);

#if defined(ALIMESH)
  std::unique_ptr<ScopeKeyFragmentBase> computeFragment(const Http::HeaderMap& headers,
                                                        const StreamInfo::StreamInfo* info,
                                                        ReComputeCbPtr& recompute) const override;
  std::unique_ptr<ScopeKeyFragmentBase> computeFragment(const Http::HeaderMap& headers) const;
#else
  std::unique_ptr<ScopeKeyFragmentBase>
  computeFragment(const Http::HeaderMap& headers) const override;

#endif

private:
  const ScopedRoutes::ScopeKeyBuilder::FragmentBuilder::HeaderValueExtractor&
      header_value_extractor_config_;
};

#if defined(ALIMESH)
class HostValueExtractorImpl : public FragmentBuilderBase {
public:
  explicit HostValueExtractorImpl(ScopedRoutes::ScopeKeyBuilder::FragmentBuilder&& config);

  std::unique_ptr<ScopeKeyFragmentBase> computeFragment(const Http::HeaderMap& headers,
                                                        const StreamInfo::StreamInfo* info,
                                                        ReComputeCbPtr& recompute) const override;

private:
  std::unique_ptr<ScopeKeyFragmentBase> reComputeHelper(const std::string& host,
                                                        ReComputeCbWeakPtr& weak_next_recompute,
                                                        uint32_t recompute_seq) const;

  static constexpr uint32_t DefaultMaxRecomputeNum = 100;

  const ScopedRoutes::ScopeKeyBuilder::FragmentBuilder::HostValueExtractor&
      host_value_extractor_config_;
  const uint32_t max_recompute_num_;
};

class LocalPortValueExtractorImpl : public FragmentBuilderBase {
public:
  explicit LocalPortValueExtractorImpl(ScopedRoutes::ScopeKeyBuilder::FragmentBuilder&& config);

  std::unique_ptr<ScopeKeyFragmentBase> computeFragment(const Http::HeaderMap& headers,
                                                        const StreamInfo::StreamInfo* info,
                                                        ReComputeCbPtr& recompute) const override;
};
#endif
/**
 * Base class for ScopeKeyBuilder implementations.
 */
class ScopeKeyBuilderBase {
public:
  explicit ScopeKeyBuilderBase(ScopedRoutes::ScopeKeyBuilder&& config)
      : config_(std::move(config)) {}
  virtual ~ScopeKeyBuilderBase() = default;

#if defined(ALIMESH)
  virtual ScopeKeyPtr computeScopeKey(const Http::HeaderMap& headers,
                                      const StreamInfo::StreamInfo* info,
                                      std::function<ScopeKeyPtr()>& recompute) const PURE;
#else
  // Computes scope key for given headers, returns nullptr if a key can't be computed.
  virtual ScopeKeyPtr computeScopeKey(const Http::HeaderMap& headers) const PURE;

#endif

protected:
  const ScopedRoutes::ScopeKeyBuilder config_;
};

class ScopeKeyBuilderImpl : public ScopeKeyBuilderBase {
public:
  explicit ScopeKeyBuilderImpl(ScopedRoutes::ScopeKeyBuilder&& config);

#if defined(ALIMESH)
  ScopeKeyPtr computeScopeKey(const Http::HeaderMap& headers, const StreamInfo::StreamInfo* info,
                              std::function<ScopeKeyPtr()>& recompute) const override;
  // only for test
  ScopeKeyPtr computeScopeKey(const Http::HeaderMap& headers) const;
#else
  ScopeKeyPtr computeScopeKey(const Http::HeaderMap& headers) const override;
#endif

private:
  std::vector<std::unique_ptr<FragmentBuilderBase>> fragment_builders_;
};

// ScopedRouteConfiguration and corresponding RouteConfigProvider.
class ScopedRouteInfo {
public:
  ScopedRouteInfo(envoy::config::route::v3::ScopedRouteConfiguration&& config_proto,
                  ConfigConstSharedPtr&& route_config);

  const ConfigConstSharedPtr& routeConfig() const { return route_config_; }
  const ScopeKey& scopeKey() const { return scope_key_; }
  const envoy::config::route::v3::ScopedRouteConfiguration& configProto() const {
    return config_proto_;
  }
  const std::string& scopeName() const { return config_proto_.name(); }

private:
  envoy::config::route::v3::ScopedRouteConfiguration config_proto_;
  ScopeKey scope_key_;
  ConfigConstSharedPtr route_config_;
};
using ScopedRouteInfoConstSharedPtr = std::shared_ptr<const ScopedRouteInfo>;
// Ordered map for consistent config dumping.
using ScopedRouteMap = std::map<std::string, ScopedRouteInfoConstSharedPtr>;

/**
 * Each Envoy worker is assigned an instance of this type. When config updates are received,
 * addOrUpdateRoutingScope() and removeRoutingScope() are called to update the set of scoped routes.
 *
 * ConnectionManagerImpl::refreshCachedRoute() will call getRouterConfig() to obtain the
 * Router::ConfigConstSharedPtr to use for route selection.
 */
class ScopedConfigImpl : public ScopedConfig {
public:
  ScopedConfigImpl(ScopedRoutes::ScopeKeyBuilder&& scope_key_builder)
      : scope_key_builder_(std::move(scope_key_builder)) {}

  void
  addOrUpdateRoutingScopes(const std::vector<ScopedRouteInfoConstSharedPtr>& scoped_route_infos);

  void removeRoutingScopes(const std::vector<std::string>& scope_names);

  // Envoy::Router::ScopedConfig
#if defined(ALIMESH)
  Router::ConfigConstSharedPtr
  getRouteConfig(const Http::HeaderMap& headers,
                 const StreamInfo::StreamInfo* info = nullptr) const override;
#else
  Router::ConfigConstSharedPtr getRouteConfig(const Http::HeaderMap& headers) const override;
#endif

#if defined(ALIMESH)
  ScopeKeyPtr computeScopeKey(const Http::HeaderMap& headers,
                              const StreamInfo::StreamInfo* info) const override;
  // only for test
  ScopeKeyPtr computeScopeKey(const Http::HeaderMap& headers) const;
#else
  // The return value is not null only if the scope corresponding to the header exists.
  ScopeKeyPtr computeScopeKey(const Http::HeaderMap& headers) const override;
#endif

private:
  ScopeKeyBuilderImpl scope_key_builder_;
  // From scope name to cached ScopedRouteInfo.
  absl::flat_hash_map<std::string, ScopedRouteInfoConstSharedPtr> scoped_route_info_by_name_;
  // Hash by ScopeKey hash to lookup in constant time.
  absl::flat_hash_map<uint64_t, ScopedRouteInfoConstSharedPtr> scoped_route_info_by_key_;
};

/**
 * A NULL implementation of the scoped routing configuration.
 */
class NullScopedConfigImpl : public ScopedConfig {
public:
#if defined(ALIMESH)
  Router::ConfigConstSharedPtr getRouteConfig(const Http::HeaderMap&,
                                              const StreamInfo::StreamInfo*) const override {
    return std::make_shared<const NullConfigImpl>();
  }
#else
  Router::ConfigConstSharedPtr getRouteConfig(const Http::HeaderMap&) const override {
    return std::make_shared<const NullConfigImpl>();
  }
#endif
};

} // namespace Router
} // namespace Envoy
