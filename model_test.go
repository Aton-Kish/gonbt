package gonbt

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPayload(t *testing.T) {
	cases := []struct {
		name    string
		tagType TagType
		want    Payload
		wantErr error
	}{
		{
			name:    `Valid Case: BytePayload`,
			tagType: TagByte,
			want:    new(BytePayload),
			wantErr: nil,
		},
		{
			name:    `Valid Case: ShortPayload`,
			tagType: TagShort,
			want:    new(ShortPayload),
			wantErr: nil,
		},
		{
			name:    `Valid Case: IntPayload`,
			tagType: TagInt,
			want:    new(IntPayload),
			wantErr: nil,
		},
		{
			name:    `Valid Case: LongPayload`,
			tagType: TagLong,
			want:    new(LongPayload),
			wantErr: nil,
		},
		{
			name:    `Valid Case: FloatPayload`,
			tagType: TagFloat,
			want:    new(FloatPayload),
			wantErr: nil,
		},
		{
			name:    `Valid Case: DoublePayload`,
			tagType: TagDouble,
			want:    new(DoublePayload),
			wantErr: nil,
		},
		{
			name:    `Valid Case: ByteArrayPayload`,
			tagType: TagByteArray,
			want:    new(ByteArrayPayload),
			wantErr: nil,
		},
		{
			name:    `Valid Case: StringPayload`,
			tagType: TagString,
			want:    new(StringPayload),
			wantErr: nil,
		},
		{
			name:    `Valid Case: ListPayload`,
			tagType: TagList,
			want:    new(ListPayload),
			wantErr: nil,
		},
		{
			name:    `Valid Case: CompoundPayload`,
			tagType: TagCompound,
			want:    new(CompoundPayload),
			wantErr: nil,
		},
		{
			name:    `Valid Case: IntArrayPayload`,
			tagType: TagIntArray,
			want:    new(IntArrayPayload),
			wantErr: nil,
		},
		{
			name:    `Valid Case: LongArrayPayload`,
			tagType: TagLongArray,
			want:    new(LongArrayPayload),
			wantErr: nil,
		},
		{
			name:    `Invalid Case: invalid tag type 0`,
			tagType: TagEnd,
			want:    nil,
			wantErr: errors.New(`invalid tag type 0`),
		},
		{
			name:    `Invalid Case: invalid tag type 13`,
			tagType: 13,
			want:    nil,
			wantErr: errors.New(`invalid tag type 13`),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			p, err := NewPayload(c.tagType)

			if c.wantErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, c.want, p)
			} else {
				assert.EqualError(t, err, c.wantErr.Error())
			}
		})
	}
}

func TestNewTag(t *testing.T) {
	cases := []struct {
		name    string
		tagType TagType
		want    Tag
		wantErr error
	}{
		{
			name:    `Valid Case: EndTag`,
			tagType: TagEnd,
			want:    new(EndTag),
			wantErr: nil,
		},
		{
			name:    `Valid Case: ByteTag`,
			tagType: TagByte,
			want:    new(ByteTag),
			wantErr: nil,
		},
		{
			name:    `Valid Case: ShortTag`,
			tagType: TagShort,
			want:    new(ShortTag),
			wantErr: nil,
		},
		{
			name:    `Valid Case: IntTag`,
			tagType: TagInt,
			want:    new(IntTag),
			wantErr: nil,
		},
		{
			name:    `Valid Case: LongTag`,
			tagType: TagLong,
			want:    new(LongTag),
			wantErr: nil,
		},
		{
			name:    `Valid Case: FloatTag`,
			tagType: TagFloat,
			want:    new(FloatTag),
			wantErr: nil,
		},
		{
			name:    `Valid Case: DoubleTag`,
			tagType: TagDouble,
			want:    new(DoubleTag),
			wantErr: nil,
		},
		{
			name:    `Valid Case: ByteArrayTag`,
			tagType: TagByteArray,
			want:    new(ByteArrayTag),
			wantErr: nil,
		},
		{
			name:    `Valid Case: StringTag`,
			tagType: TagString,
			want:    new(StringTag),
			wantErr: nil,
		},
		{
			name:    `Valid Case: ListTag`,
			tagType: TagList,
			want:    new(ListTag),
			wantErr: nil,
		},
		{
			name:    `Valid Case: CompoundTag`,
			tagType: TagCompound,
			want:    new(CompoundTag),
			wantErr: nil,
		},
		{
			name:    `Valid Case: IntArrayTag`,
			tagType: TagIntArray,
			want:    new(IntArrayTag),
			wantErr: nil,
		},
		{
			name:    `Valid Case: LongArrayTag`,
			tagType: TagLongArray,
			want:    new(LongArrayTag),
			wantErr: nil,
		},
		{
			name:    `Invalid Case: invalid type`,
			tagType: 13,
			want:    nil,
			wantErr: errors.New(`invalid tag type 13`),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			p, err := NewTag(c.tagType)

			if c.wantErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, c.want, p)
			} else {
				assert.EqualError(t, err, c.wantErr.Error())
			}
		})
	}
}
