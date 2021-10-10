package gonbt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPayload_TypeId(t *testing.T) {
	cases := []struct {
		name    string
		payload Payload
		want    TagType
	}{
		{
			name:    `Valid Case: BytePayload`,
			payload: BytePayloadPtr(0),
			want:    TagByte,
		},
		{
			name:    `Valid Case: ShortPayload`,
			payload: ShortPayloadPtr(0),
			want:    TagShort,
		},
		{
			name:    `Valid Case: IntPayload`,
			payload: IntPayloadPtr(0),
			want:    TagInt,
		},
		{
			name:    `Valid Case: LongPayload`,
			payload: LongPayloadPtr(0),
			want:    TagLong,
		},
		{
			name:    `Valid Case: FloatPayload`,
			payload: FloatPayloadPtr(0),
			want:    TagFloat,
		},
		{
			name:    `Valid Case: DoublePayload`,
			payload: DoublePayloadPtr(0),
			want:    TagDouble,
		},
		{
			name:    `Valid Case: ByteArrayPayload`,
			payload: &ByteArrayPayload{},
			want:    TagByteArray,
		},
		{
			name:    `Valid Case: StringPayload`,
			payload: StringPayloadPtr(``),
			want:    TagString,
		},
		{
			name:    `Valid Case: ListPayload`,
			payload: &ListPayload{},
			want:    TagList,
		},
		{
			name:    `Valid Case: CompoundPayload`,
			payload: &CompoundPayload{},
			want:    TagCompound,
		},
		{
			name:    `Valid Case: IntArrayPayload`,
			payload: &IntArrayPayload{},
			want:    TagIntArray,
		},
		{
			name:    `Valid Case: LongArrayPayload`,
			payload: &LongArrayPayload{},
			want:    TagLongArray,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.want, c.payload.TypeId())
		})
	}
}

func TestTag_TypeId(t *testing.T) {
	cases := []struct {
		name string
		tag  Tag
		want TagType
	}{
		{
			name: `Valid Case: EndTag`,
			tag:  &EndTag{},
			want: TagEnd,
		},
		{
			name: `Valid Case: ByteTag`,
			tag:  &ByteTag{},
			want: TagByte,
		},
		{
			name: `Valid Case: ShortTag`,
			tag:  &ShortTag{},
			want: TagShort,
		},
		{
			name: `Valid Case: IntTag`,
			tag:  &IntTag{},
			want: TagInt,
		},
		{
			name: `Valid Case: LongTag`,
			tag:  &LongTag{},
			want: TagLong,
		},
		{
			name: `Valid Case: FloatTag`,
			tag:  &FloatTag{},
			want: TagFloat,
		},
		{
			name: `Valid Case: DoubleTag`,
			tag:  &DoubleTag{},
			want: TagDouble,
		},
		{
			name: `Valid Case: ByteArrayTag`,
			tag:  &ByteArrayTag{},
			want: TagByteArray,
		},
		{
			name: `Valid Case: StringTag`,
			tag:  &StringTag{},
			want: TagString,
		},
		{
			name: `Valid Case: ListTag`,
			tag:  &ListTag{},
			want: TagList,
		},
		{
			name: `Valid Case: CompoundTag`,
			tag:  &CompoundTag{},
			want: TagCompound,
		},
		{
			name: `Valid Case: IntArrayTag`,
			tag:  &IntArrayTag{},
			want: TagIntArray,
		},
		{
			name: `Valid Case: LongArrayTag`,
			tag:  &LongArrayTag{},
			want: TagLongArray,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.want, c.tag.TypeId())
		})
	}
}
