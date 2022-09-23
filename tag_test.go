// Copyright (c) 2022 Aton-Kish
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package nbt

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type tagTestCase[T any, U Payload] struct {
	name string
	data T
	nbt  nbtTestCase[U]
	raw  rawTestCase
}

type nbtTestCase[T Payload] struct {
	tagType TagType
	tagName TagName
	payload T
}

type rawTestCase struct {
	tagType []byte
	tagName []byte
	payload []byte
}

func TestTagType_Encode(t *testing.T) {
	type Case struct {
		name        string
		tagType     TagType
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range endTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tagType:     c.nbt.tagType,
			expected:    c.raw.tagType,
			expectedErr: nil,
		})
	}

	for _, c := range byteTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tagType:     c.nbt.tagType,
			expected:    c.raw.tagType,
			expectedErr: nil,
		})
	}

	for _, c := range shortTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tagType:     c.nbt.tagType,
			expected:    c.raw.tagType,
			expectedErr: nil,
		})
	}

	for _, c := range intTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tagType:     c.nbt.tagType,
			expected:    c.raw.tagType,
			expectedErr: nil,
		})
	}

	for _, c := range longTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tagType:     c.nbt.tagType,
			expected:    c.raw.tagType,
			expectedErr: nil,
		})
	}

	for _, c := range floatTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tagType:     c.nbt.tagType,
			expected:    c.raw.tagType,
			expectedErr: nil,
		})
	}

	for _, c := range doubleTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tagType:     c.nbt.tagType,
			expected:    c.raw.tagType,
			expectedErr: nil,
		})
	}

	for _, c := range byteArrayTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tagType:     c.nbt.tagType,
			expected:    c.raw.tagType,
			expectedErr: nil,
		})
	}

	for _, c := range stringTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tagType:     c.nbt.tagType,
			expected:    c.raw.tagType,
			expectedErr: nil,
		})
	}

	for _, c := range listTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tagType:     c.nbt.tagType,
			expected:    c.raw.tagType,
			expectedErr: nil,
		})
	}

	for _, c := range compoundTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tagType:     c.nbt.tagType,
			expected:    c.raw.tagType,
			expectedErr: nil,
		})
	}

	for _, c := range intArrayTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tagType:     c.nbt.tagType,
			expected:    c.raw.tagType,
			expectedErr: nil,
		})
	}

	for _, c := range longArrayTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tagType:     c.nbt.tagType,
			expected:    c.raw.tagType,
			expectedErr: nil,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			err := tt.tagType.Encode(buf)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, buf.Bytes())
			} else {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}

func TestTagType_Decode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    TagType
		expectedErr error
	}

	cases := []Case{}

	for _, c := range endTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         c.raw.tagType,
			expected:    c.nbt.tagType,
			expectedErr: nil,
		})
	}

	for _, c := range byteTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         c.raw.tagType,
			expected:    c.nbt.tagType,
			expectedErr: nil,
		})
	}

	for _, c := range shortTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         c.raw.tagType,
			expected:    c.nbt.tagType,
			expectedErr: nil,
		})
	}

	for _, c := range intTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         c.raw.tagType,
			expected:    c.nbt.tagType,
			expectedErr: nil,
		})
	}

	for _, c := range longTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         c.raw.tagType,
			expected:    c.nbt.tagType,
			expectedErr: nil,
		})
	}

	for _, c := range floatTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         c.raw.tagType,
			expected:    c.nbt.tagType,
			expectedErr: nil,
		})
	}

	for _, c := range doubleTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         c.raw.tagType,
			expected:    c.nbt.tagType,
			expectedErr: nil,
		})
	}

	for _, c := range byteArrayTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         c.raw.tagType,
			expected:    c.nbt.tagType,
			expectedErr: nil,
		})
	}

	for _, c := range stringTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         c.raw.tagType,
			expected:    c.nbt.tagType,
			expectedErr: nil,
		})
	}

	for _, c := range listTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         c.raw.tagType,
			expected:    c.nbt.tagType,
			expectedErr: nil,
		})
	}

	for _, c := range compoundTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         c.raw.tagType,
			expected:    c.nbt.tagType,
			expectedErr: nil,
		})
	}

	for _, c := range intArrayTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         c.raw.tagType,
			expected:    c.nbt.tagType,
			expectedErr: nil,
		})
	}

	for _, c := range longArrayTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         c.raw.tagType,
			expected:    c.nbt.tagType,
			expectedErr: nil,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer(tt.raw)

			var typ TagType
			err := typ.Decode(buf)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, typ)
			} else {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}

func TestNewTagName(t *testing.T) {
	type Case struct {
		name     string
		value    string
		expected *TagName
	}

	cases := []Case{}

	for _, c := range endTagCases {
		cases = append(cases, Case{
			name:     c.name,
			value:    string(c.nbt.tagName),
			expected: &c.nbt.tagName,
		})
	}

	for _, c := range byteTagCases {
		cases = append(cases, Case{
			name:     c.name,
			value:    string(c.nbt.tagName),
			expected: &c.nbt.tagName,
		})
	}

	for _, c := range shortTagCases {
		cases = append(cases, Case{
			name:     c.name,
			value:    string(c.nbt.tagName),
			expected: &c.nbt.tagName,
		})
	}

	for _, c := range intTagCases {
		cases = append(cases, Case{
			name:     c.name,
			value:    string(c.nbt.tagName),
			expected: &c.nbt.tagName,
		})
	}

	for _, c := range longTagCases {
		cases = append(cases, Case{
			name:     c.name,
			value:    string(c.nbt.tagName),
			expected: &c.nbt.tagName,
		})
	}

	for _, c := range floatTagCases {
		cases = append(cases, Case{
			name:     c.name,
			value:    string(c.nbt.tagName),
			expected: &c.nbt.tagName,
		})
	}

	for _, c := range doubleTagCases {
		cases = append(cases, Case{
			name:     c.name,
			value:    string(c.nbt.tagName),
			expected: &c.nbt.tagName,
		})
	}

	for _, c := range byteArrayTagCases {
		cases = append(cases, Case{
			name:     c.name,
			value:    string(c.nbt.tagName),
			expected: &c.nbt.tagName,
		})
	}

	for _, c := range stringTagCases {
		cases = append(cases, Case{
			name:     c.name,
			value:    string(c.nbt.tagName),
			expected: &c.nbt.tagName,
		})
	}

	for _, c := range listTagCases {
		cases = append(cases, Case{
			name:     c.name,
			value:    string(c.nbt.tagName),
			expected: &c.nbt.tagName,
		})
	}

	for _, c := range compoundTagCases {
		cases = append(cases, Case{
			name:     c.name,
			value:    string(c.nbt.tagName),
			expected: &c.nbt.tagName,
		})
	}

	for _, c := range intArrayTagCases {
		cases = append(cases, Case{
			name:     c.name,
			value:    string(c.nbt.tagName),
			expected: &c.nbt.tagName,
		})
	}

	for _, c := range longArrayTagCases {
		cases = append(cases, Case{
			name:     c.name,
			value:    string(c.nbt.tagName),
			expected: &c.nbt.tagName,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewTagName(tt.value)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestTagName_Encode(t *testing.T) {
	type Case struct {
		name        string
		tagName     TagName
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range byteTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tagName:     c.nbt.tagName,
			expected:    c.raw.tagName,
			expectedErr: nil,
		})
	}

	for _, c := range shortTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tagName:     c.nbt.tagName,
			expected:    c.raw.tagName,
			expectedErr: nil,
		})
	}

	for _, c := range intTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tagName:     c.nbt.tagName,
			expected:    c.raw.tagName,
			expectedErr: nil,
		})
	}

	for _, c := range longTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tagName:     c.nbt.tagName,
			expected:    c.raw.tagName,
			expectedErr: nil,
		})
	}

	for _, c := range floatTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tagName:     c.nbt.tagName,
			expected:    c.raw.tagName,
			expectedErr: nil,
		})
	}

	for _, c := range doubleTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tagName:     c.nbt.tagName,
			expected:    c.raw.tagName,
			expectedErr: nil,
		})
	}

	for _, c := range byteArrayTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tagName:     c.nbt.tagName,
			expected:    c.raw.tagName,
			expectedErr: nil,
		})
	}

	for _, c := range stringTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tagName:     c.nbt.tagName,
			expected:    c.raw.tagName,
			expectedErr: nil,
		})
	}

	for _, c := range listTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tagName:     c.nbt.tagName,
			expected:    c.raw.tagName,
			expectedErr: nil,
		})
	}

	for _, c := range compoundTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tagName:     c.nbt.tagName,
			expected:    c.raw.tagName,
			expectedErr: nil,
		})
	}

	for _, c := range intArrayTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tagName:     c.nbt.tagName,
			expected:    c.raw.tagName,
			expectedErr: nil,
		})
	}

	for _, c := range longArrayTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tagName:     c.nbt.tagName,
			expected:    c.raw.tagName,
			expectedErr: nil,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			err := tt.tagName.Encode(buf)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, buf.Bytes())
			} else {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}

func TestTagName_Decode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    TagName
		expectedErr error
	}

	cases := []Case{}

	for _, c := range byteTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         c.raw.tagName,
			expected:    c.nbt.tagName,
			expectedErr: nil,
		})
	}

	for _, c := range shortTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         c.raw.tagName,
			expected:    c.nbt.tagName,
			expectedErr: nil,
		})
	}

	for _, c := range intTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         c.raw.tagName,
			expected:    c.nbt.tagName,
			expectedErr: nil,
		})
	}

	for _, c := range longTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         c.raw.tagName,
			expected:    c.nbt.tagName,
			expectedErr: nil,
		})
	}

	for _, c := range floatTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         c.raw.tagName,
			expected:    c.nbt.tagName,
			expectedErr: nil,
		})
	}

	for _, c := range doubleTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         c.raw.tagName,
			expected:    c.nbt.tagName,
			expectedErr: nil,
		})
	}

	for _, c := range byteArrayTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         c.raw.tagName,
			expected:    c.nbt.tagName,
			expectedErr: nil,
		})
	}

	for _, c := range stringTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         c.raw.tagName,
			expected:    c.nbt.tagName,
			expectedErr: nil,
		})
	}

	for _, c := range listTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         c.raw.tagName,
			expected:    c.nbt.tagName,
			expectedErr: nil,
		})
	}

	for _, c := range compoundTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         c.raw.tagName,
			expected:    c.nbt.tagName,
			expectedErr: nil,
		})
	}

	for _, c := range intArrayTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         c.raw.tagName,
			expected:    c.nbt.tagName,
			expectedErr: nil,
		})
	}

	for _, c := range longArrayTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         c.raw.tagName,
			expected:    c.nbt.tagName,
			expectedErr: nil,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer(tt.raw)

			var n TagName
			err := n.Decode(buf)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, n)
			} else {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}

func TestNewTag(t *testing.T) {
	cases := []struct {
		name        string
		tagType     TagType
		expected    Tag
		expectedErr error
	}{
		{
			name:        "positive case: EndType",
			tagType:     EndType,
			expected:    NewEndTag(),
			expectedErr: nil,
		},
		{
			name:        "positive case: ByteType",
			tagType:     ByteType,
			expected:    NewByteTag(new(TagName), new(BytePayload)),
			expectedErr: nil,
		},
		{
			name:        "positive case: ShortType",
			tagType:     ShortType,
			expected:    NewShortTag(new(TagName), new(ShortPayload)),
			expectedErr: nil,
		},
		{
			name:        "positive case: IntType",
			tagType:     IntType,
			expected:    NewIntTag(new(TagName), new(IntPayload)),
			expectedErr: nil,
		},
		{
			name:        "positive case: LongType",
			tagType:     LongType,
			expected:    NewLongTag(new(TagName), new(LongPayload)),
			expectedErr: nil,
		},
		{
			name:        "positive case: FloatType",
			tagType:     FloatType,
			expected:    NewFloatTag(new(TagName), new(FloatPayload)),
			expectedErr: nil,
		},
		{
			name:        "positive case: DoubleType",
			tagType:     DoubleType,
			expected:    NewDoubleTag(new(TagName), new(DoublePayload)),
			expectedErr: nil,
		},
		{
			name:        "positive case: ByteArrayType",
			tagType:     ByteArrayType,
			expected:    NewByteArrayTag(new(TagName), new(ByteArrayPayload)),
			expectedErr: nil,
		},
		{
			name:        "positive case: StringType",
			tagType:     StringType,
			expected:    NewStringTag(new(TagName), new(StringPayload)),
			expectedErr: nil,
		},
		{
			name:        "positive case: ListType",
			tagType:     ListType,
			expected:    NewListTag(new(TagName), new(ListPayload)),
			expectedErr: nil,
		},
		{
			name:        "positive case: CompoundType",
			tagType:     CompoundType,
			expected:    NewCompoundTag(new(TagName), new(CompoundPayload)),
			expectedErr: nil,
		},
		{
			name:        "positive case: IntArrayType",
			tagType:     IntArrayType,
			expected:    NewIntArrayTag(new(TagName), new(IntArrayPayload)),
			expectedErr: nil,
		},
		{
			name:        "positive case: LongArrayType",
			tagType:     LongArrayType,
			expected:    NewLongArrayTag(new(TagName), new(LongArrayPayload)),
			expectedErr: nil,
		},
		{
			name:        "negative case: out of range",
			tagType:     TagType(0x0D),
			expected:    nil,
			expectedErr: errors.New("invalid tag type id 13"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tag, err := NewTag(tt.tagType)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, tag)
			} else {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}

func TestNewPayload(t *testing.T) {
	cases := []struct {
		name        string
		tagType     TagType
		expected    Payload
		expectedErr error
	}{
		{
			name:        "positive case: ByteType",
			tagType:     ByteType,
			expected:    new(BytePayload),
			expectedErr: nil,
		},
		{
			name:        "positive case: ShortType",
			tagType:     ShortType,
			expected:    new(ShortPayload),
			expectedErr: nil,
		},
		{
			name:        "positive case: IntType",
			tagType:     IntType,
			expected:    new(IntPayload),
			expectedErr: nil,
		},
		{
			name:        "positive case: LongType",
			tagType:     LongType,
			expected:    new(LongPayload),
			expectedErr: nil,
		},
		{
			name:        "positive case: FloatType",
			tagType:     FloatType,
			expected:    new(FloatPayload),
			expectedErr: nil,
		},
		{
			name:        "positive case: DoubleType",
			tagType:     DoubleType,
			expected:    new(DoublePayload),
			expectedErr: nil,
		},
		{
			name:        "positive case: ByteArrayType",
			tagType:     ByteArrayType,
			expected:    new(ByteArrayPayload),
			expectedErr: nil,
		},
		{
			name:        "positive case: StringType",
			tagType:     StringType,
			expected:    new(StringPayload),
			expectedErr: nil,
		},
		{
			name:        "positive case: ListType",
			tagType:     ListType,
			expected:    new(ListPayload),
			expectedErr: nil,
		},
		{
			name:        "positive case: CompoundType",
			tagType:     CompoundType,
			expected:    new(CompoundPayload),
			expectedErr: nil,
		},
		{
			name:        "positive case: IntArrayType",
			tagType:     IntArrayType,
			expected:    new(IntArrayPayload),
			expectedErr: nil,
		},
		{
			name:        "positive case: LongArrayType",
			tagType:     LongArrayType,
			expected:    new(LongArrayPayload),
			expectedErr: nil,
		},
		{
			name:        "negative case",
			tagType:     TagType(0x0D),
			expected:    nil,
			expectedErr: errors.New("invalid tag type id 13"),
		},
		{
			name:        "negative case: EndType",
			tagType:     EndType,
			expected:    nil,
			expectedErr: errors.New("invalid tag type id 0"),
		},

		{
			name:        "negative case: out of range",
			tagType:     TagType(0x0D),
			expected:    nil,
			expectedErr: errors.New("invalid tag type id 13"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			payload, err := NewPayload(tt.tagType)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, payload)
			} else {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}
