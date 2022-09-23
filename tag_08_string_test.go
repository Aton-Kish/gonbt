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

var stringTagCases = []tagTestCase[string, *StringPayload]{
	{
		name: "positive case: StringTag - \"Hello World\"",
		data: "Hello World",
		nbt: nbtTestCase[*StringPayload]{
			tagType: StringType,
			tagName: "String",
			payload: NewStringPayload("Hello World"),
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: String(=8)
				0x08,
			},
			tagName: []byte{
				// Name Length: 6
				0x00, 0x06,
				// Name: "String"
				0x53, 0x74, 0x72, 0x69, 0x6E, 0x67,
			},
			payload: []byte{
				// Payload Length: 11
				0x00, 0x0B,
				// Payload: "Hello World"
				0x48, 0x65, 0x6C, 0x6C, 0x6F, 0x20, 0x57, 0x6F, 0x72, 0x6C, 0x64,
			},
		},
	},
	{
		name: "positive case: StringTag - \"Test\"",
		data: "Test",
		nbt: nbtTestCase[*StringPayload]{
			tagType: StringType,
			tagName: "String",
			payload: NewStringPayload("Test"),
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: String(=8)
				0x08,
			},
			tagName: []byte{
				// Name Length: 6
				0x00, 0x06,
				// Name: "String"
				0x53, 0x74, 0x72, 0x69, 0x6E, 0x67,
			},
			payload: []byte{
				// Payload Length: 4
				0x00, 0x04,
				// Payload: "Test"
				0x54, 0x65, 0x73, 0x74,
			},
		},
	},
	{
		name: "positive case: StringTag - \"minecraft:the_end\"",
		data: "minecraft:the_end",
		nbt: nbtTestCase[*StringPayload]{
			tagType: StringType,
			tagName: "String",
			payload: NewStringPayload("minecraft:the_end"),
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: String(=8)
				0x08,
			},
			tagName: []byte{
				// Name Length: 6
				0x00, 0x06,
				// Name: "String"
				0x53, 0x74, 0x72, 0x69, 0x6E, 0x67,
			},
			payload: []byte{
				// Payload Length: 17
				0x00, 0x11,
				// Payload: "minecraft:the_end"
				0x6D, 0x69, 0x6E, 0x65, 0x63, 0x72, 0x61, 0x66, 0x74, 0x3A,
				0x74, 0x68, 0x65, 0x5F, 0x65, 0x6E, 0x64,
			},
		},
	},
	{
		name: "positive case: StringTag - \"\"",
		data: "",
		nbt: nbtTestCase[*StringPayload]{
			tagType: StringType,
			tagName: "String",
			payload: NewStringPayload(""),
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: String(=8)
				0x08,
			},
			tagName: []byte{
				// Name Length: 6
				0x00, 0x06,
				// Name: "String"
				0x53, 0x74, 0x72, 0x69, 0x6E, 0x67,
			},
			payload: []byte{
				// Payload Length: 0
				0x00, 0x00,
				// Payload: ""
			},
		},
	},
	{
		name: "positive case: StringTag - \"マインクラフト\"",
		data: "マインクラフト",
		nbt: nbtTestCase[*StringPayload]{
			tagType: StringType,
			tagName: "String",
			payload: NewStringPayload("マインクラフト"),
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: String(=8)
				0x08,
			},
			tagName: []byte{
				// Name Length: 6
				0x00, 0x06,
				// Name: "String"
				0x53, 0x74, 0x72, 0x69, 0x6E, 0x67,
			},
			payload: []byte{
				// Payload Length: 21
				0x00, 0x15,
				// Payload: "マインクラフト"
				0xE3, 0x83, 0x9E, 0xE3, 0x82, 0xA4, 0xE3, 0x83, 0xB3,
				0xE3, 0x82, 0xAF, 0xE3, 0x83, 0xA9, 0xE3, 0x83, 0x95,
				0xE3, 0x83, 0x88,
			},
		},
	},
}

func TestNewStringTag(t *testing.T) {
	type Case struct {
		name     string
		tagName  *TagName
		payload  *StringPayload
		expected *StringTag
	}

	cases := []Case{}

	for _, c := range stringTagCases {
		cases = append(cases, Case{
			name:    c.name,
			tagName: &c.nbt.tagName,
			payload: c.nbt.payload,
			expected: &StringTag{
				tagName: &c.nbt.tagName,
				payload: c.nbt.payload,
			},
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewStringTag(tt.tagName, tt.payload)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestStringTag_TypeId(t *testing.T) {
	tag := new(StringTag)
	expected := StringType
	actual := tag.TypeId()
	assert.Equal(t, expected, actual)
}

func TestStringTag_Encode(t *testing.T) {
	type Case struct {
		name        string
		tag         *StringTag
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range stringTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tag:         NewStringTag(&c.nbt.tagName, c.nbt.payload),
			expected:    append(append(c.raw.tagType, c.raw.tagName...), c.raw.payload...),
			expectedErr: nil,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			err := tt.tag.Encode(buf)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, buf.Bytes())
			} else {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}

func TestStringTag_Decode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    *StringTag
		expectedErr error
	}

	cases := []Case{}

	for _, c := range stringTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         append(append(c.raw.tagType, c.raw.tagName...), c.raw.payload...),
			expected:    NewStringTag(&c.nbt.tagName, c.nbt.payload),
			expectedErr: nil,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer(tt.raw)

			tag := new(StringTag)
			err := tag.Decode(buf)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, tag)
			} else {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}

func TestNewStringPayload(t *testing.T) {
	type Case struct {
		name     string
		value    string
		expected *StringPayload
	}

	cases := []Case{}

	for _, c := range stringTagCases {
		cases = append(cases, Case{
			name:     c.name,
			value:    c.data,
			expected: c.nbt.payload,
		})
	}

	for _, tt := range cases {
		actual := NewStringPayload(tt.value)
		assert.Equal(t, tt.expected, actual)
	}
}

func TestStringPayload_TypeId(t *testing.T) {
	payload := new(StringPayload)
	expected := StringType
	actual := payload.TypeId()
	assert.Equal(t, expected, actual)
}

func TestStringPayload_Encode(t *testing.T) {
	type Case struct {
		name        string
		payload     *StringPayload
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range stringTagCases {
		cases = append(cases, Case{
			name:        c.name,
			payload:     c.nbt.payload,
			expected:    c.raw.payload,
			expectedErr: nil,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			err := tt.payload.Encode(buf)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, buf.Bytes())
			} else {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}

func TestStringPayload_Decode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    *StringPayload
		expectedErr error
	}

	cases := []Case{}

	for _, c := range stringTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         c.raw.payload,
			expected:    c.nbt.payload,
			expectedErr: nil,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer(tt.raw)

			payload := new(StringPayload)
			err := payload.Decode(buf)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, payload)
			} else {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}
