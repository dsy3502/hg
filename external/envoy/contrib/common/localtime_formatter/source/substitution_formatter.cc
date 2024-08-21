#include "contrib/common/localtime_formatter/source/substitution_formatter.h"

#include <climits>
#include <cstdint>
#include <regex>
#include <string>
#include <vector>

#include "envoy/config/core/v3/base.pb.h"

#include "source/common/api/os_sys_calls_impl.h"
#include "source/common/common/assert.h"
#include "source/common/common/empty_string.h"
#include "source/common/common/fmt.h"
#include "source/common/common/utility.h"
#include "source/common/config/metadata.h"
#include "source/common/grpc/common.h"
#include "source/common/grpc/status.h"
#include "source/common/http/utility.h"
#include "source/common/protobuf/message_validator_impl.h"
#include "source/common/protobuf/utility.h"
#include "source/common/stream_info/utility.h"

#include "absl/strings/str_split.h"
#include "fmt/format.h"

namespace Envoy {
namespace Formatter {

LocalStartTimeFormatter::LocalStartTimeFormatter(const std::string& token)
    : date_formatter_(parseFormat(token, sizeof("LOCAL_START_TIME(") - 1)) {}

absl::optional<std::string> LocalStartTimeFormatter::format(
    const Http::RequestHeaderMap&, const Http::ResponseHeaderMap&, const Http::ResponseTrailerMap&,
    const StreamInfo::StreamInfo& stream_info, absl::string_view) const {
  return DateTimeFormatter::fromTime(stream_info.startTime(), absl::LocalTimeZone());
}

ProtobufWkt::Value LocalStartTimeFormatter::formatValue(
    const Http::RequestHeaderMap& request_headers, const Http::ResponseHeaderMap& response_headers,
    const Http::ResponseTrailerMap& response_trailers, const StreamInfo::StreamInfo& stream_info,
    absl::string_view local_reply_body) const {
  return ValueUtil::optionalStringValue(
      format(request_headers, response_headers, response_trailers, stream_info, local_reply_body));
}

std::string LocalStartTimeFormatter::DateTimeFormatter::fromTime(const SystemTime& system_time,
                                                                 absl::TimeZone time_zone) {
  static const std::string DefaultDateFormat = "%Y-%m-%dT%H:%M:%E3SZ";

  struct CachedTime {
    std::chrono::seconds epoch_time_seconds;
    std::string formatted_time;
  };
  static thread_local CachedTime cached_time;

  const std::chrono::milliseconds epoch_time_ms =
      std::chrono::duration_cast<std::chrono::milliseconds>(system_time.time_since_epoch());

  const std::chrono::seconds epoch_time_seconds =
      std::chrono::duration_cast<std::chrono::seconds>(epoch_time_ms);

  if (cached_time.formatted_time.empty() || cached_time.epoch_time_seconds != epoch_time_seconds) {
    cached_time.formatted_time =
        absl::FormatTime(DefaultDateFormat, absl::FromChrono(system_time), time_zone);
    cached_time.epoch_time_seconds = epoch_time_seconds;
  } else {
    // Overwrite the digits in the ".000Z" at the end of the string with the
    // millisecond count from the input time.
    ASSERT(cached_time.formatted_time.length() == 24);
    size_t offset = cached_time.formatted_time.length() - 4;
    uint32_t msec = epoch_time_ms.count() % 1000;
    cached_time.formatted_time[offset++] = ('0' + (msec / 100));
    msec %= 100;
    cached_time.formatted_time[offset++] = ('0' + (msec / 10));
    msec %= 10;
    cached_time.formatted_time[offset++] = ('0' + msec);
  }

  return cached_time.formatted_time;
}

} // namespace Formatter
} // namespace Envoy
