#pragma once

#include <cstdint>
#include <forward_list>
#include <string>
#include <vector>

#include "source/common/common/logger.h"
#include "source/extensions/filters/network/common/redis/codec.h"

namespace Envoy {
namespace Extensions {
namespace NetworkFilters {
namespace Common {
namespace Redis {

/**
 * Decoder implementation of https://redis.io/topics/protocol
 *
 * This implementation buffers when needed and will always consume all bytes passed for decoding.
 */
class DecoderImpl : public Decoder, Logger::Loggable<Logger::Id::redis> {
public:
  DecoderImpl(DecoderCallbacks& callbacks) : callbacks_(callbacks) {}

  // RedisProxy::Decoder
  void decode(Buffer::Instance& data) override;

private:
  enum class State {
    ValueRootStart,
    ValueStart,
    IntegerStart,
    Integer,
    IntegerLF,
    BulkStringBody,
    CR,
    LF,
    SimpleString,
    ValueComplete
  };

  struct PendingInteger {
    void reset() {
      integer_ = 0;
      negative_ = false;
    }

    uint64_t integer_;
    bool negative_;
  };

  struct PendingValue {
    RespValue* value_;
    uint64_t current_array_element_;
  };

  void parseSlice(const Buffer::RawSlice& slice);

  DecoderCallbacks& callbacks_;
  State state_{State::ValueRootStart};
  PendingInteger pending_integer_;
  RespValuePtr pending_value_root_;
  std::forward_list<PendingValue> pending_value_stack_;
};

#if defined(ALIMESH)
class RawDecoderImpl : public Decoder, Logger::Loggable<Logger::Id::redis> {
public:
  RawDecoderImpl(RawDecoderCallbacks& callbacks) : callbacks_(callbacks) {}

  // RedisProxy::Decoder
  void decode(Buffer::Instance& data) override;

private:
  enum class State {
    ValueRootStart,
    ValueStart,
    IntegerStart,
    Integer,
    IntegerLF,
    BulkStringBody,
    CR,
    LF,
    SimpleString,
    ValueComplete
  };

  struct PendingInteger {
    void reset() {
      integer_ = 0;
      negative_ = false;
    }

    uint64_t integer_;
    bool negative_;
  };

  struct PendingValue {
    RespType type;
    std::string value;
    uint64_t current_array_element;
    uint64_t total_array_element;
  };

  void parseSlice(const Buffer::RawSlice& slice);

  RawDecoderCallbacks& callbacks_;
  State state_{State::ValueRootStart};
  PendingInteger pending_integer_;
  std::string pending_value_root_;
  std::forward_list<PendingValue> pending_value_stack_;
};
#endif
/**
 * A factory implementation that returns a real decoder.
 */
class DecoderFactoryImpl : public DecoderFactory {
public:
  // RedisProxy::DecoderFactory
  DecoderPtr create(DecoderCallbacks& callbacks) override {
    return DecoderPtr{new DecoderImpl(callbacks)};
  }
};
#if defined(ALIMESH)
class RawDecoderFactoryImpl : public RawDecoderFactory {
public:
  // RedisProxy::RawDecoderFactory
  DecoderPtr create(RawDecoderCallbacks& callbacks) override {
    return DecoderPtr{new RawDecoderImpl(callbacks)};
  }
};
#endif
/**
 * Encoder implementation of https://redis.io/topics/protocol
 */
class EncoderImpl : public Encoder {
public:
  // RedisProxy::Encoder
  void encode(const RespValue& value, Buffer::Instance& out) override;

private:
  void encodeArray(const std::vector<RespValue>& array, Buffer::Instance& out);
  void encodeCompositeArray(const RespValue::CompositeArray& array, Buffer::Instance& out);
  void encodeBulkString(const std::string& string, Buffer::Instance& out);
  void encodeError(const std::string& string, Buffer::Instance& out);
  void encodeInteger(int64_t integer, Buffer::Instance& out);
  void encodeSimpleString(const std::string& string, Buffer::Instance& out);
};
#if defined(ALIMESH)
class RawEncoderImpl : public RawEncoder {
public:
  // RedisProxy::RawEncoder
  void encode(absl::string_view value, Buffer::Instance& out) override;
};
#endif
} // namespace Redis
} // namespace Common
} // namespace NetworkFilters
} // namespace Extensions
} // namespace Envoy
