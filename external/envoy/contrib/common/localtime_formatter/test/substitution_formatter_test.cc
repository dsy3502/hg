#include <chrono>
#include <cstdint>
#include <cstdlib>
#include <string>
#include <vector>

#include "envoy/config/core/v3/base.pb.h"
#include "envoy/event/timer.h"

#include "source/common/common/logger.h"
#include "source/common/common/utility.h"
#include "source/common/formatter/substitution_formatter.h"
#include "source/common/http/header_map_impl.h"
#include "source/common/json/json_loader.h"
#include "source/common/network/address_impl.h"
#include "source/common/protobuf/utility.h"
#include "source/common/router/string_accessor_impl.h"

#include "test/mocks/api/mocks.h"
#include "test/mocks/http/mocks.h"
#include "test/mocks/ssl/mocks.h"
#include "test/mocks/stream_info/mocks.h"
#include "test/test_common/printers.h"
#include "test/test_common/threadsafe_singleton_injector.h"
#include "test/test_common/utility.h"

#include "contrib/common/localtime_formatter/source/substitution_formatter.h"
#include "gmock/gmock.h"
#include "gtest/gtest.h"

using testing::NiceMock;
using testing::Return;

namespace Envoy {
namespace Formatter {
namespace {

TEST(AlimeshSubstitutionFormatterTest, localStartTimeFormatter) {
  NiceMock<StreamInfo::MockStreamInfo> stream_info;
  Http::TestRequestHeaderMapImpl request_headers{{":method", "GET"}, {":path", "/"}};
  Http::TestResponseHeaderMapImpl response_headers;
  Http::TestResponseTrailerMapImpl response_trailers;
  std::string body;

  {
    LocalStartTimeFormatter local_start_time_format("");
    int rand = std::rand();
    std::shared_ptr<SystemTime> time =
        std::make_shared<SystemTime>(Envoy::Event::TimeSystem::Microseconds(rand));
    EXPECT_CALL(stream_info, startTime()).WillRepeatedly(Return(*time));
    EXPECT_EQ(LocalStartTimeFormatter::DateTimeFormatter::fromTime(*time, absl::LocalTimeZone()),
              local_start_time_format.format(request_headers, response_headers, response_trailers,
                                             stream_info, body));
  }

  {
    absl::Time abslStartTime =
        TestUtility::parseTime("Dec 18 01:50:34 2018 GMT", "%b %e %H:%M:%S %Y GMT");
    SystemTime startTime = absl::ToChronoTime(abslStartTime);

    EXPECT_EQ("2018-12-18T01:50:34.000Z", LocalStartTimeFormatter::DateTimeFormatter::fromTime(
                                              startTime, absl::LocalTimeZone()));
  }
}

} // namespace
} // namespace Formatter
} // namespace Envoy
