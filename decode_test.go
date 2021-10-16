package gonbt

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeNBT(t *testing.T) {
	type decodeTestCase struct {
		name    string
		raw     []byte
		want    Tag
		wantErr error
	}

	cases := []decodeTestCase{
		{
			name:    `Invalid Case: EOF error`,
			raw:     []byte{},
			want:    new(CompoundTag),
			wantErr: errors.New(`EOF`),
		},
		{
			name:    `Invalid Case: GZip`,
			raw:     []byte{0x1F, 0x8B},
			want:    new(CompoundTag),
			wantErr: errors.New(`invalid tag type 31`),
		},
	}

	for _, v := range NBTEncodingValidCases {
		c := decodeTestCase{
			name:    v.name,
			raw:     v.raw,
			want:    v.nbt,
			wantErr: nil,
		}
		cases = append(cases, c)
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			buf := bytes.NewBuffer(c.raw)

			tag := new(Tag)
			err := Decode(buf, tag)

			if c.wantErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, c.want, *tag)
			} else {
				assert.EqualError(t, err, c.wantErr.Error())
			}
		})
	}
}

func TestTagName_Decode(t *testing.T) {
	type decodeTestCase struct {
		name    string
		raw     []byte
		want    *TagName
		wantErr error
	}

	cases := []decodeTestCase{
		{
			name:    `Invalid Case: EOF error`,
			raw:     []byte{},
			want:    nil,
			wantErr: errors.New(`EOF`),
		},
	}

	for _, v := range TagNameEncodingValidCases {
		c := decodeTestCase{
			name:    v.name,
			raw:     v.raw,
			want:    v.nbt,
			wantErr: nil,
		}
		cases = append(cases, c)
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			buf := bytes.NewBuffer(c.raw)

			var n TagName
			err := n.Decode(buf)

			if c.wantErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, *c.want, n)
			} else {
				assert.EqualError(t, err, c.wantErr.Error())
			}
		})
	}
}

func TestPayload_Decode(t *testing.T) {
	type decodeTestCase struct {
		name    string
		raw     []byte
		want    Payload
		wantErr error
	}

	cases := []decodeTestCase{
		{
			name:    `Invalid Case: BytePayload - EOF error`,
			raw:     []byte{},
			want:    new(BytePayload),
			wantErr: errors.New(`EOF`),
		},
		{
			name:    `Invalid Case: ShortPayload - EOF error`,
			raw:     []byte{},
			want:    new(ShortPayload),
			wantErr: errors.New(`EOF`),
		},
		{
			name:    `Invalid Case: IntPayload - EOF error`,
			raw:     []byte{},
			want:    new(IntPayload),
			wantErr: errors.New(`EOF`),
		},
		{
			name:    `Invalid Case: LongPayload - EOF error`,
			raw:     []byte{},
			want:    new(LongPayload),
			wantErr: errors.New(`EOF`),
		},
		{
			name:    `Invalid Case: FloatPayload - EOF error`,
			raw:     []byte{},
			want:    new(FloatPayload),
			wantErr: errors.New(`EOF`),
		},
		{
			name:    `Invalid Case: DoublePayload - EOF error`,
			raw:     []byte{},
			want:    new(DoublePayload),
			wantErr: errors.New(`EOF`),
		},
		{
			name:    `Invalid Case: ByteArrayPayload - EOF error`,
			raw:     []byte{},
			want:    new(ByteArrayPayload),
			wantErr: errors.New(`EOF`),
		},
		{
			name:    `Invalid Case: StringPayload - EOF error`,
			raw:     []byte{},
			want:    new(StringPayload),
			wantErr: errors.New(`EOF`),
		},
		{
			name:    `Invalid Case: ListPayload - EOF error`,
			raw:     []byte{},
			want:    new(ListPayload),
			wantErr: errors.New(`EOF`),
		},
		{
			name:    `Invalid Case: CompoundPayload - EOF error`,
			raw:     []byte{},
			want:    new(CompoundPayload),
			wantErr: errors.New(`EOF`),
		},
		{
			name:    `Invalid Case: IntArrayPayload - EOF error`,
			raw:     []byte{},
			want:    new(IntArrayPayload),
			wantErr: errors.New(`EOF`),
		},
		{
			name:    `Invalid Case: LongArrayPayload - EOF error`,
			raw:     []byte{},
			want:    new(LongArrayPayload),
			wantErr: errors.New(`EOF`),
		},
	}

	for _, v := range PayloadEncodingValidCases {
		c := decodeTestCase{
			name:    v.name,
			raw:     v.raw,
			want:    v.nbt,
			wantErr: nil,
		}
		cases = append(cases, c)
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			buf := bytes.NewBuffer(c.raw)

			payload, err := NewPayload(c.want.TypeId())
			assert.NoError(t, err)

			err = payload.Decode(buf)

			if c.wantErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, c.want, payload)
			} else {
				assert.EqualError(t, err, c.wantErr.Error())
			}
		})
	}
}

func TestTag_Decode(t *testing.T) {
	type decodeTestCase struct {
		name    string
		raw     []byte
		want    Tag
		wantErr error
	}

	cases := []decodeTestCase{
		{
			name:    `Invalid Case: ByteTag - EOF error`,
			raw:     []byte{},
			want:    new(ByteTag),
			wantErr: errors.New(`EOF`),
		},
		{
			name:    `Invalid Case: ShortTag - EOF error`,
			raw:     []byte{},
			want:    new(ShortTag),
			wantErr: errors.New(`EOF`),
		},
		{
			name:    `Invalid Case: IntTag - EOF error`,
			raw:     []byte{},
			want:    new(IntTag),
			wantErr: errors.New(`EOF`),
		},
		{
			name:    `Invalid Case: LongTag - EOF error`,
			raw:     []byte{},
			want:    new(LongTag),
			wantErr: errors.New(`EOF`),
		},
		{
			name:    `Invalid Case: FloatTag - EOF error`,
			raw:     []byte{},
			want:    new(FloatTag),
			wantErr: errors.New(`EOF`),
		},
		{
			name:    `Invalid Case: DoubleTag - EOF error`,
			raw:     []byte{},
			want:    new(DoubleTag),
			wantErr: errors.New(`EOF`),
		},
		{
			name:    `Invalid Case: ByteArrayTag - EOF error`,
			raw:     []byte{},
			want:    new(ByteArrayTag),
			wantErr: errors.New(`EOF`),
		},
		{
			name:    `Invalid Case: StringTag - EOF error`,
			raw:     []byte{},
			want:    new(StringTag),
			wantErr: errors.New(`EOF`),
		},
		{
			name:    `Invalid Case: ListTag - EOF error`,
			raw:     []byte{},
			want:    new(ListTag),
			wantErr: errors.New(`EOF`),
		},
		{
			name:    `Invalid Case: CompoundTag - EOF error`,
			raw:     []byte{},
			want:    new(CompoundTag),
			wantErr: errors.New(`EOF`),
		},
		{
			name:    `Invalid Case: IntArrayTag - EOF error`,
			raw:     []byte{},
			want:    new(IntArrayTag),
			wantErr: errors.New(`EOF`),
		},
		{
			name:    `Invalid Case: LongArrayTag - EOF error`,
			raw:     []byte{},
			want:    new(LongArrayTag),
			wantErr: errors.New(`EOF`),
		},
	}

	for _, v := range TagEncodingValidCases {
		c := decodeTestCase{
			name:    v.name,
			raw:     v.raw,
			want:    v.nbt,
			wantErr: nil,
		}
		cases = append(cases, c)
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			buf := bytes.NewBuffer(c.raw)

			tag, err := NewTag(c.want.TypeId())
			assert.NoError(t, err)

			err = tag.Decode(buf)

			if c.wantErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, c.want, tag)
			} else {
				assert.EqualError(t, err, c.wantErr.Error())
			}
		})
	}
}
