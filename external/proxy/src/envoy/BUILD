# Copyright 2017 Istio Authors. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
################################################################################
#
load("@rules_pkg//:pkg.bzl", "pkg_tar")
load(
    "@envoy//bazel:envoy_build_system.bzl",
    "envoy_cc_binary",
    "envoy_cc_contrib_extension",
    "envoy_cc_library",
    "envoy_contrib_package",
)
load(
    "@envoy//contrib:all_contrib_extensions.bzl",
    "ARM64_SKIP_CONTRIB_TARGETS",
    "PPC_SKIP_CONTRIB_TARGETS",
    "envoy_all_contrib_extensions",
)

envoy_cc_binary(
    name = "envoy",
    repository = "@envoy",
    visibility = ["//visibility:public"],
    deps = [
        "//extensions/access_log_policy:access_log_policy_lib",
        "//extensions/attributegen:attributegen_plugin",
        "//extensions/metadata_exchange:metadata_exchange_lib",
        "//extensions/stackdriver:stackdriver_plugin",
        "//extensions/stats:stats_plugin",
        "//src/envoy/http/alpn:config_lib",
        "//src/envoy/http/authn:filter_lib",
        "//src/envoy/tcp/forward_downstream_sni:config_lib",
        "//src/envoy/tcp/metadata_exchange:config_lib",
        "//src/envoy/tcp/sni_verifier:config_lib",
        "//src/envoy/tcp/tcp_cluster_rewrite:config_lib",
        "@envoy//source/exe:envoy_main_entry_lib",
    ] + [
        "@envoy//contrib/custom_cluster_plugins/cluster_fallback/source:config",
        "@envoy//contrib/http_dubbo_transcoder/filters/http/source:config",
        "@envoy//contrib/kafka/filters/network/source:kafka_broker_config_lib",
        "@envoy//contrib/kafka/filters/network/source/mesh:config_lib",
        "@envoy//contrib/mysql_proxy/filters/network/source:config",
        "@envoy//contrib/postgres_proxy/filters/network/source:config",
        "@envoy//contrib/rocketmq_proxy/filters/network/source:config",
        "@envoy//contrib/sip_proxy/filters/network/source:config",
        "@envoy//contrib/sip_proxy/filters/network/source/router:config",
        "@envoy//contrib/squash/filters/http/source:config",
        "@envoy//contrib/upstreams/http/dubbo_tcp/source:config"
    ],
)

pkg_tar(
    name = "envoy_tar",
    srcs = [":envoy"],
    extension = "tar.gz",
    mode = "0755",
    package_dir = "/usr/local/bin/",
    tags = ["manual"],
    visibility = ["//visibility:public"],
)
