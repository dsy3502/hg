#include "contrib/server/admin/source/cluster_endpoints_handler.h"

#include "envoy/admin/v3/clusters.pb.h"

#include "source/common/http/headers.h"
#include "source/common/http/utility.h"
#include "source/common/network/utility.h"
#include "source/common/upstream/host_utility.h"
#include "source/server/admin/utils.h"

namespace Envoy {
namespace Server {

namespace {} // namespace

ClusterEndpointsHandler::ClusterEndpointsHandler(Server::Instance& server)
    : HandlerContextBase(server) {}

Http::Code ClusterEndpointsHandler::handlerClusterEndpoints(absl::string_view path,
                                                            Http::HeaderMap&,
                                                            Buffer::Instance& response,
                                                            AdminStream& admin_stream) {
  Http::Utility::QueryParams params = Http::Utility::parseQueryString(path);
  if (params.empty()) {
    // Check if the params are in the request's body.}{}
    if (admin_stream.getRequestBody() != nullptr &&
        isFormUrlEncoded(admin_stream.getRequestHeaders().ContentType())) {
      params = Http::Utility::parseFromBody(admin_stream.getRequestBody()->toString());
    }
  }

  absl::optional<absl::string_view> service_name =
      params.find("service") != params.end()
          ? absl::optional<absl::string_view>{params.at("service")}
          : absl::nullopt;

  if (params.empty() || !service_name.has_value()) {
    response.add("usage: /endpoints?service=value1&protocol=value2\n");
    response.add("       or send the parameters as form values\n");
    response.add("use an empty value to remove a previously added override");
    return Http::Code::BadRequest;
  }

  auto all_clusters = server_.clusterManager().clusters();
  for (const auto& [name, cluster_ref] : all_clusters.active_clusters_) {
    // const Upstream::Cluster& cluster = cluster_pair.second.get();
    const Upstream::Cluster& cluster = cluster_ref.get();
    Upstream::ClusterInfoConstSharedPtr cluster_info = cluster.info();
    if (absl::StrContains(cluster_info->name(), service_name.value())) {
      response.add(fmt::format("ClusterName: {}\n", cluster_info->name()));
      for (auto& host_set : cluster.prioritySet().hostSetsPerPriority()) {
        for (auto& host : host_set->hosts()) {
          response.add(fmt::format("\tendpoint: {}", host->address()->asString()));
          response.add(
              fmt::format("\tmetadata: {}\n",
                          host->metadata() == nullptr
                              ? std::string()
                              : MessageUtil::getJsonStringFromMessageOrDie(*host->metadata())));
        }
      }
    }
  }
  return Http::Code::OK;
}

bool ClusterEndpointsHandler::isFormUrlEncoded(const Http::HeaderEntry* content_type) const {
  if (content_type == nullptr) {
    return false;
  }

  return content_type->value().getStringView() ==
         Http::Headers::get().ContentTypeValues.FormUrlEncoded;
}

} // namespace Server
} // namespace Envoy
