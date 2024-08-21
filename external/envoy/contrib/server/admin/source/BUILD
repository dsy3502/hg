load(
    "//bazel:envoy_build_system.bzl",
    "envoy_cc_library",
    "envoy_contrib_package",
)

licenses(["notice"])  # Apache 2

envoy_contrib_package()

envoy_cc_library(
    name = "cluster_endpoints_handler_lib",
    srcs = ["cluster_endpoints_handler.cc"],
    hdrs = ["cluster_endpoints_handler.h"],
    visibility = ["//visibility:public"],
    deps = [
        "//envoy/http:codes_interface",
        "//envoy/server:admin_interface",
        "//envoy/server:instance_interface",
        "//source/common/buffer:buffer_lib",
        "//source/common/http:codes_lib",
        "//source/common/http:header_map_lib",
        "//source/common/upstream:host_utility_lib",
        "//source/server/admin:handler_ctx_lib",
        "//source/server/admin:utils_lib",
        "@envoy_api//envoy/admin/v3:pkg_cc_proto",
    ],
)
