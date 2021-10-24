package gonbt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTagNamePtr(t *testing.T) {
	n := TagName(`Test`)
	assert.Equal(t, &n, TagNamePtr(n))
}

func TestBytePayloadPtr(t *testing.T) {
	p := BytePayload(123)
	assert.Equal(t, &p, BytePayloadPtr(p))
}

func TestShortPayloadPtr(t *testing.T) {
	p := ShortPayload(12345)
	assert.Equal(t, &p, ShortPayloadPtr(p))
}

func TestIntPayloadPtr(t *testing.T) {
	p := IntPayload(123456789)
	assert.Equal(t, &p, IntPayloadPtr(p))
}

func TestLongPayloadPtr(t *testing.T) {
	p := LongPayload(123456789123456789)
	assert.Equal(t, &p, LongPayloadPtr(p))
}

func TestFloatPayloadPtr(t *testing.T) {
	p := FloatPayload(0.12345678)
	assert.Equal(t, &p, FloatPayloadPtr(p))
}

func TestDoublePayloadPtr(t *testing.T) {
	p := DoublePayload(0.123456789)
	assert.Equal(t, &p, DoublePayloadPtr(p))
}

func TestStringPayloadPtr(t *testing.T) {
	p := StringPayload(`Test`)
	assert.Equal(t, &p, StringPayloadPtr(p))
}
