#pragma once

#include "envoy/buffer/buffer.h"
#include "envoy/http/codes.h"
#include "envoy/http/header_map.h"
#include "envoy/server/admin.h"
#include "envoy/server/instance.h"

#include "source/server/admin/handler_ctx.h"

#include "absl/strings/string_view.h"

namespace Envoy {
namespace Server {

class ClusterEndpointsHandler : public HandlerContextBase {

public:
  ClusterEndpointsHandler(Server::Instance& server);

  Http::Code handlerClusterEndpoints(absl::string_view path_and_query,
                                     Http::HeaderMap& response_headers, Buffer::Instance& response,
                                     AdminStream&);

private:
  bool isFormUrlEncoded(const Http::HeaderEntry* content_type) const;
};

} // namespace Server
} // namespace Envoy
