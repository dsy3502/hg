load(
    "//bazel:envoy_build_system.bzl",
    "envoy_cc_library",
    "envoy_package",
)

licenses(["notice"])  # Apache 2

envoy_package()

envoy_cc_library(
    name = "context_interface",
    hdrs = ["context.h"],
    deps = [
        ":router_interface",
    ],
)

envoy_cc_library(
    name = "rds_interface",
    hdrs = ["rds.h"],
    deps = [
        ":router_interface",
        "//envoy/http:filter_interface",
        "@envoy_api//envoy/config/route/v3:pkg_cc_proto",
    ],
)

envoy_cc_library(
    name = "route_config_provider_manager_interface",
    hdrs = ["route_config_provider_manager.h"],
    deps = [
        ":rds_interface",
        "//envoy/event:dispatcher_interface",
        "//envoy/json:json_object_interface",
        "//envoy/local_info:local_info_interface",
        "//envoy/runtime:runtime_interface",
        "//envoy/server:filter_config_interface",
        "//envoy/stats:stats_interface",
        "//envoy/thread_local:thread_local_interface",
        "//envoy/upstream:cluster_manager_interface",
        "@envoy_api//envoy/config/route/v3:pkg_cc_proto",
        "@envoy_api//envoy/extensions/filters/network/http_connection_manager/v3:pkg_cc_proto",
    ],
)

envoy_cc_library(
    name = "route_config_update_info_interface",
    hdrs = ["route_config_update_receiver.h"],
    external_deps = ["abseil_optional"],
    deps = [
        ":rds_interface",
        "//envoy/common:time_interface",
        "//source/common/protobuf",
        "@envoy_api//envoy/config/route/v3:pkg_cc_proto",
        "@envoy_api//envoy/service/discovery/v3:pkg_cc_proto",
    ],
)

envoy_cc_library(
    name = "router_interface",
    hdrs = ["router.h"],
    alimesh_deps = [
        "//contrib/envoy/http:active_redirect_policy_interface",
    ],
    external_deps = ["abseil_optional"],
    deps = [
        ":internal_redirect_interface",
        "//envoy/access_log:access_log_interface",
        "//envoy/common:conn_pool_interface",
        "//envoy/common:matchers_interface",
        "//envoy/common:time_interface",
        "//envoy/config:typed_metadata_interface",
        "//envoy/http:codec_interface",
        "//envoy/http:codes_interface",
        "//envoy/http:conn_pool_interface",
        "//envoy/http:hash_policy_interface",
        "//envoy/http:header_map_interface",
        "//envoy/tcp:conn_pool_interface",
        "//envoy/tracing:http_tracer_interface",
        "//envoy/upstream:resource_manager_interface",
        "//envoy/upstream:retry_interface",
        "//source/common/protobuf",
        "//source/common/protobuf:utility_lib",
        "@envoy_api//envoy/config/core/v3:pkg_cc_proto",
        "@envoy_api//envoy/config/route/v3:pkg_cc_proto",
        "@envoy_api//envoy/type/v3:pkg_cc_proto",
    ],
)

envoy_cc_library(
    name = "scopes_interface",
    hdrs = ["scopes.h"],
    deps = [
        ":router_interface",
        "//envoy/config:config_provider_interface",
        "//envoy/http:header_map_interface",
    ],
)

envoy_cc_library(
    name = "router_ratelimit_interface",
    hdrs = ["router_ratelimit.h"],
    deps = [
        "//envoy/http:filter_interface",
        "//envoy/ratelimit:ratelimit_interface",
        "@envoy_api//envoy/config/core/v3:pkg_cc_proto",
    ],
)

envoy_cc_library(
    name = "shadow_writer_interface",
    hdrs = ["shadow_writer.h"],
    deps = [
        "//envoy/http:async_client_interface",
        "//envoy/http:message_interface",
    ],
)

envoy_cc_library(
    name = "string_accessor_interface",
    hdrs = ["string_accessor.h"],
    external_deps = ["abseil_optional"],
    deps = [
        "//envoy/stream_info:filter_state_interface",
    ],
)

envoy_cc_library(
    name = "internal_redirect_interface",
    hdrs = ["internal_redirect.h"],
    deps = [
        "//envoy/config:typed_config_interface",
        "//envoy/stream_info:filter_state_interface",
        "//source/common/common:minimal_logger_lib",
    ],
)

envoy_cc_library(
    name = "cluster_specifier_plugin_interface",
    hdrs = ["cluster_specifier_plugin.h"],
    deps = [
        ":router_interface",
        "//envoy/config:typed_config_interface",
        "//envoy/server:factory_context_interface",
    ],
)
