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
	"testing"

	"github.com/stretchr/testify/assert"
)

var tagTypeCases = []struct {
	name    string
	tagType TagType
	raw     []byte
}{
	{
		name:    "positive case: EndType",
		tagType: EndType,
		raw: []byte{
			// Type: End(=0)
			0x00,
		},
	},
	{
		name:    "positive case: ByteType",
		tagType: ByteType,
		raw: []byte{
			// Type: Byte(=1)
			0x01,
		},
	},
	{
		name:    "positive case: ShortType",
		tagType: ShortType,
		raw: []byte{
			// Type: Short(=2)
			0x02,
		},
	},
	{
		name:    "positive case: IntType",
		tagType: IntType,
		raw: []byte{
			// Type: Int(=3)
			0x03,
		},
	},
	{
		name:    "positive case: LongType",
		tagType: LongType,
		raw: []byte{
			// Type: Long(=4)
			0x04,
		},
	},
	{
		name:    "positive case: FloatType",
		tagType: FloatType,
		raw: []byte{
			// Type: Float(=5)
			0x05,
		},
	},
	{
		name:    "positive case: DoubleType",
		tagType: DoubleType,
		raw: []byte{
			// Type: Double(=6)
			0x06,
		},
	},
	{
		name:    "positive case: ByteArrayType",
		tagType: ByteArrayType,
		raw: []byte{
			// Type: ByteArray(=7)
			0x07,
		},
	},
	{
		name:    "positive case: StringType",
		tagType: StringType,
		raw: []byte{
			// Type: String(=8)
			0x08,
		},
	},
	{
		name:    "positive case: ListType",
		tagType: ListType,
		raw: []byte{
			// Type: List(=9)
			0x09,
		},
	},
	{
		name:    "positive case: CompoundType",
		tagType: CompoundType,
		raw: []byte{
			// Type: Compound(=10)
			0x0A,
		},
	},
	{
		name:    "positive case: IntArrayType",
		tagType: IntArrayType,
		raw: []byte{
			// Type: IntArray(=11)
			0x0B,
		},
	},
	{
		name:    "positive case: LongArrayType",
		tagType: LongArrayType,
		raw: []byte{
			// Type: LongArray(=12)
			0x0C,
		},
	},
}

func TestTagType_Encode(t *testing.T) {
	type Case struct {
		name        string
		tagType     TagType
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range tagTypeCases {
		cases = append(cases, Case{
			name:        c.name,
			tagType:     c.tagType,
			expected:    c.raw,
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

	for _, c := range tagTypeCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         c.raw,
			expected:    c.tagType,
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

var tagNameCases = []struct {
	name    string
	tagName TagName
	raw     []byte
}{
	{
		name:    "positive case: \"Test\"",
		tagName: TagName("Test"),
		raw: []byte{
			// Name Length: 4
			0x00, 0x04,
			// Name: "Test"
			0x54, 0x65, 0x73, 0x74,
		},
	},
	{
		name:    "positive case: \"minecraft:the_end\"",
		tagName: TagName("minecraft:the_end"),
		raw: []byte{
			// Name Length: 17
			0x00, 0x11,
			// Name: "minecraft:the_end"
			0x6D, 0x69, 0x6E, 0x65, 0x63, 0x72, 0x61, 0x66, 0x74, 0x3A,
			0x74, 0x68, 0x65, 0x5F, 0x65, 0x6E, 0x64,
		},
	},
	{
		name:    "positive case: \"\"",
		tagName: TagName(""),
		raw: []byte{
			// Name Length: 0
			0x00, 0x00,
			// Name: ""
		},
	},
	{
		name:    "positive case: \"マインクラフト\"",
		tagName: TagName("マインクラフト"),
		raw: []byte{
			// Name Length: 21
			0x00, 0x15,
			// Name: "マインクラフト"
			0xE3, 0x83, 0x9E, 0xE3, 0x82, 0xA4, 0xE3, 0x83, 0xB3,
			0xE3, 0x82, 0xAF, 0xE3, 0x83, 0xA9, 0xE3, 0x83, 0x95,
			0xE3, 0x83, 0x88,
		},
	},
}

func TestTagName_Encode(t *testing.T) {
	type Case struct {
		name        string
		tagName     TagName
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range tagNameCases {
		cases = append(cases, Case{
			name:        c.name,
			tagName:     c.tagName,
			expected:    c.raw,
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

	for _, c := range tagNameCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         c.raw,
			expected:    c.tagName,
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

func TestPayload_BytePayload(t *testing.T) {
	value := BytePayload(123)
	expected := &value
	actual := PrimitivePayloadPointer(value)
	assert.Equal(t, expected, actual)
}

func TestPayload_ShortPayload(t *testing.T) {
	value := ShortPayload(12345)
	expected := &value
	actual := PrimitivePayloadPointer(value)
	assert.Equal(t, expected, actual)
}

func TestPayload_IntPayload(t *testing.T) {
	value := IntPayload(123456789)
	expected := &value
	actual := PrimitivePayloadPointer(value)
	assert.Equal(t, expected, actual)
}

func TestPayload_LongPayload(t *testing.T) {
	value := LongPayload(123456789123456789)
	expected := &value
	actual := PrimitivePayloadPointer(value)
	assert.Equal(t, expected, actual)
}

func TestPayload_FloatPayload(t *testing.T) {
	value := FloatPayload(0.12345678)
	expected := &value
	actual := PrimitivePayloadPointer(value)
	assert.Equal(t, expected, actual)
}

func TestPayload_DoublePayload(t *testing.T) {
	value := DoublePayload(0.123456789)
	expected := &value
	actual := PrimitivePayloadPointer(value)
	assert.Equal(t, expected, actual)
}

func TestPayload_StringPayload(t *testing.T) {
	value := StringPayload("Test")
	expected := &value
	actual := PrimitivePayloadPointer(value)
	assert.Equal(t, expected, actual)
}
