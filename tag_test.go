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

type tagTestCase[T Payload] struct {
	name string
	nbt  nbtTestCase[T]
	snbt snbtTestCase
	raw  rawTestCase
}

type nbtTestCase[T Payload] struct {
	tagType TagType
	tagName TagName
	payload T
}

type snbtTestCase struct {
	tagName string
	payload string
}

type rawTestCase struct {
	tagType []byte
	tagName []byte
	payload []byte
}

func TestTagType_encode(t *testing.T) {
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
			err := tt.tagType.encode(buf)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, buf.Bytes())
			} else {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}

func TestTagType_decode(t *testing.T) {
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
			err := typ.decode(buf)

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

func TestTagName_encode(t *testing.T) {
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
			err := tt.tagName.encode(buf)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, buf.Bytes())
			} else {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}

func TestTagName_decode(t *testing.T) {
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
			err := n.decode(buf)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, n)
			} else {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}

func TestTagName_stringify(t *testing.T) {
	type Case struct {
		name     string
		tagName  TagName
		expected string
	}

	cases := []Case{
		{
			name:     "positive case: quotation - Test",
			tagName:  TagName("Test"),
			expected: "Test",
		},
		{
			name:     "positive case: quotation - '\"Test'",
			tagName:  TagName("\"Test"),
			expected: "'\"Test'",
		},
		{
			name:     "positive case: quotation - \"'Test\"",
			tagName:  TagName("'Test"),
			expected: "\"'Test\"",
		},
		{
			name:     "positive case: quotation - \"\\\"'Test\"",
			tagName:  TagName("\"'Test"),
			expected: "\"\\\"'Test\"",
		},
	}

	for _, c := range byteTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tagName:  c.nbt.tagName,
			expected: c.snbt.tagName,
		})
	}

	for _, c := range shortTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tagName:  c.nbt.tagName,
			expected: c.snbt.tagName,
		})
	}

	for _, c := range intTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tagName:  c.nbt.tagName,
			expected: c.snbt.tagName,
		})
	}

	for _, c := range longTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tagName:  c.nbt.tagName,
			expected: c.snbt.tagName,
		})
	}

	for _, c := range floatTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tagName:  c.nbt.tagName,
			expected: c.snbt.tagName,
		})
	}

	for _, c := range doubleTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tagName:  c.nbt.tagName,
			expected: c.snbt.tagName,
		})
	}

	for _, c := range byteArrayTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tagName:  c.nbt.tagName,
			expected: c.snbt.tagName,
		})
	}

	for _, c := range stringTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tagName:  c.nbt.tagName,
			expected: c.snbt.tagName,
		})
	}

	for _, c := range listTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tagName:  c.nbt.tagName,
			expected: c.snbt.tagName,
		})
	}

	for _, c := range compoundTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tagName:  c.nbt.tagName,
			expected: c.snbt.tagName,
		})
	}

	for _, c := range intArrayTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tagName:  c.nbt.tagName,
			expected: c.snbt.tagName,
		})
	}

	for _, c := range longArrayTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tagName:  c.nbt.tagName,
			expected: c.snbt.tagName,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.tagName.stringify()
			assert.Equal(t, tt.expected, actual)
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

var nbtCases = []struct {
	name string
	nbt  Tag
	raw  []byte
}{
	{
		name: "positive case: Simple",
		nbt: NewCompoundTag(NewTagName(""), NewCompoundPayload(
			NewCompoundTag(NewTagName("Hello World"), NewCompoundPayload(
				NewStringTag(NewTagName("Name"), NewStringPayload("Steve")),
				NewEndTag(),
			)),
			NewEndTag(),
		),
		),
		raw: []byte{
			// CompoundTag(""):
			0x0A,
			0x00, 0x00,
			//   - CompoundTag("Hello World"):
			0x0A,
			0x00, 0x0B,
			0x48, 0x65, 0x6C, 0x6C, 0x6F, 0x20, 0x57, 0x6F, 0x72, 0x6C, 0x64,
			//       - StringTag("Name"): "Steve"
			0x08,
			0x00, 0x04,
			0x4E, 0x61, 0x6D, 0x65,
			0x00, 0x05,
			0x53, 0x74, 0x65, 0x76, 0x65,
			//       - EndTag
			0x00,
			//   - EndTag
			0x00,
		},
	},
	{
		name: "positive case: Tag Check",
		nbt: NewCompoundTag(NewTagName("Compound"), NewCompoundPayload(
			NewShortTag(NewTagName("Short"), NewShortPayload(12345)),
			NewByteArrayTag(NewTagName("ByteArray"), NewByteArrayPayload(0, 1)),
			NewStringTag(NewTagName("String"), NewStringPayload("Hello")),
			NewListTag(NewTagName("List"), NewListPayload(NewBytePayload(123))),
			NewCompoundTag(NewTagName("Compound"), NewCompoundPayload(
				NewStringTag(NewTagName("String"), NewStringPayload("World")),
				NewEndTag(),
			)),
			NewEndTag(),
		),
		),
		raw: []byte{
			// CompoundTag("Compound"):
			0x0A,
			0x00, 0x08,
			0x43, 0x6F, 0x6D, 0x70, 0x6F, 0x75, 0x6E, 0x64,
			//   - ShortTag("Short"): 12345
			0x02,
			0x00, 0x05,
			0x53, 0x68, 0x6f, 0x72, 0x74,
			0x30, 0x39,
			//   - ByteArrayTag("ByteArray"): [B; 0b, 1b]
			0x07,
			0x00, 0x09,
			0x42, 0x79, 0x74, 0x65, 0x41, 0x72, 0x72, 0x61, 0x79,
			0x00, 0x00, 0x00, 0x02,
			0x00, 0x01,
			//   - StringTag("String"): "Hello"
			0x08,
			0x00, 0x06,
			0x53, 0x74, 0x72, 0x69, 0x6E, 0x67,
			0x00, 0x05,
			0x48, 0x65, 0x6C, 0x6C, 0x6F,
			//   - ListTag("List"): <Byte>[123]
			0x09,
			0x00, 0x04,
			0x4C, 0x69, 0x73, 0x74,
			0x01,
			0x00, 0x00, 0x00, 0x01,
			0x7B,
			//   - CompoundTag("Compound"):
			0x0A,
			0x00, 0x08,
			0x43, 0x6F, 0x6D, 0x70, 0x6F, 0x75, 0x6E, 0x64,
			//       - StringTag("String"): "World"
			0x08,
			0x00, 0x06,
			0x53, 0x74, 0x72, 0x69, 0x6E, 0x67,
			0x00, 0x05,
			0x57, 0x6F, 0x72, 0x6C, 0x64,
			//       - EndTag
			0x00,
			//   - EndTag
			0x00,
		},
	},
}

func TestEncode(t *testing.T) {
	type Case struct {
		name        string
		nbt         Tag
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range nbtCases {
		cases = append(cases, Case{
			name:        c.name,
			nbt:         c.nbt,
			expected:    c.raw,
			expectedErr: nil,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			err := Encode(buf, tt.nbt)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, buf.Bytes())
			} else {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}

func TestDecode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    Tag
		expectedErr error
	}

	cases := []Case{}

	for _, c := range nbtCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         c.raw,
			expected:    c.nbt,
			expectedErr: nil,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer(tt.raw)
			actual, err := Decode(buf)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, actual)
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
