#pragma once

#include <functional>
#include <string>
#include <vector>

#include "envoy/common/time.h"
#include "envoy/config/core/v3/base.pb.h"
#include "envoy/formatter/substitution_formatter.h"
#include "envoy/stream_info/stream_info.h"

#include "source/common/common/utility.h"

#include "absl/container/flat_hash_map.h"
#include "absl/types/optional.h"

namespace Envoy {
namespace Formatter {

/**
 * Access log format parser.
 */
class LocalStartTimeFormatter : public FormatterProvider {
public:
  LocalStartTimeFormatter(const std::string& format);

  // FormatterProvider
  absl::optional<std::string> format(const Http::RequestHeaderMap&, const Http::ResponseHeaderMap&,
                                     const Http::ResponseTrailerMap&, const StreamInfo::StreamInfo&,
                                     absl::string_view) const override;
  ProtobufWkt::Value formatValue(const Http::RequestHeaderMap&, const Http::ResponseHeaderMap&,
                                 const Http::ResponseTrailerMap&, const StreamInfo::StreamInfo&,
                                 absl::string_view) const override;

  /**
   * Utility class for access log date/time format with milliseconds support.
   */
  class DateTimeFormatter {
  public:
    static std::string fromTime(const SystemTime& time, absl::TimeZone time_zone);
  };

  std::string parseFormat(const std::string& token, size_t parameters_start) {
    const size_t parameters_length = token.length() - (parameters_start + 1);
    return token[parameters_start - 1] == '(' ? token.substr(parameters_start, parameters_length)
                                              : "";
  }

private:
  const DateFormatter date_formatter_;
};

} // namespace Formatter
} // namespace Envoy
