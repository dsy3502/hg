load(
    "//bazel:envoy_build_system.bzl",
    "envoy_cc_test",
    "envoy_contrib_package",
)

licenses(["notice"])  # Apache 2

envoy_contrib_package()

envoy_cc_test(
    name = "substitution_formatter_test",
    srcs = ["substitution_formatter_test.cc"],
    deps = [
        "//contrib/common/localtime_formatter/source:substitution_formatter_lib",
        "//source/common/common:utility_lib",
        "//source/common/http:header_map_lib",
        "//source/common/json:json_loader_lib",
        "//source/common/network:address_lib",
        "//source/common/router:string_accessor_lib",
        "//test/mocks/api:api_mocks",
        "//test/mocks/http:http_mocks",
        "//test/mocks/ssl:ssl_mocks",
        "//test/mocks/stream_info:stream_info_mocks",
        "//test/test_common:threadsafe_singleton_injector_lib",
        "//test/test_common:utility_lib",
        "@envoy_api//envoy/config/core/v3:pkg_cc_proto",
    ],
)
