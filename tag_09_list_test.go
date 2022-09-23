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
	"fmt"
	"testing"

	"github.com/Aton-Kish/gonbt/slices"
	"github.com/stretchr/testify/assert"
)

var listTagCases = []tagTestCase[*ListPayload]{
	{
		name: "positive case: ListTag - Short",
		nbt: nbtTestCase[*ListPayload]{
			tagType: ListType,
			tagName: "List",
			payload: NewListPayload(NewShortPayload(12345), NewShortPayload(6789)),
		},
		snbt: snbtTestCase{
			tagName: "List",
			payload: stringifyType{
				typeDefault: "[12345s, 6789s]",
				typeCompact: "[12345s,6789s]",
				typePretty:  "[\n  12345s,\n  6789s\n]",
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: List(=9)
				0x09,
			},
			tagName: []byte{
				// Name Length: 4
				0x00, 0x04,
				// Name: "List"
				0x4C, 0x69, 0x73, 0x74,
			},
			payload: []byte{
				// Payload Type: TagShort(=2)
				0x02,
				// Payload Length: 2
				0x00, 0x00, 0x00, 0x02,
				// Payload:
				//   - 12345s
				0x30, 0x39,
				//   - 6789s
				0x1A, 0x85,
			},
		},
	},
	{
		name: "positive case: ListTag - ByteArray",
		nbt: nbtTestCase[*ListPayload]{
			tagType: ListType,
			tagName: "List",
			payload: NewListPayload(NewByteArrayPayload(0, 1), NewByteArrayPayload(2, 3)),
		},
		snbt: snbtTestCase{
			tagName: "List",
			payload: stringifyType{
				typeDefault: "[[B; 0b, 1b], [B; 2b, 3b]]",
				typeCompact: "[[B;0b,1b],[B;2b,3b]]",
				typePretty:  "[\n  [B; 0b, 1b],\n  [B; 2b, 3b]\n]",
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: List(=9)
				0x09,
			},
			tagName: []byte{
				// Name Length: 4
				0x00, 0x04,
				// Name: "List"
				0x4C, 0x69, 0x73, 0x74,
			},
			payload: []byte{
				// Payload Type: TagByteArray(=7)
				0x07,
				// Payload Length: 2
				0x00, 0x00, 0x00, 0x02,
				// Payload:
				//   - [B; 0b, 1b]
				0x00, 0x00, 0x00, 0x02,
				0x00, 0x01,
				//   - [B; 2b, 3b]
				0x00, 0x00, 0x00, 0x02,
				0x02, 0x03,
			},
		},
	},
	{
		name: "positive case: ListTag - String",
		nbt: nbtTestCase[*ListPayload]{
			tagType: ListType,
			tagName: "List",
			payload: NewListPayload(NewStringPayload("Hello"), NewStringPayload("World")),
		},
		snbt: snbtTestCase{
			tagName: "List",
			payload: stringifyType{
				typeDefault: "[\"Hello\", \"World\"]",
				typeCompact: "[\"Hello\",\"World\"]",
				typePretty:  "[\n  \"Hello\",\n  \"World\"\n]",
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: List(=9)
				0x09,
			},
			tagName: []byte{
				// Name Length: 4
				0x00, 0x04,
				// Name: "List"
				0x4C, 0x69, 0x73, 0x74,
			},
			payload: []byte{
				// Payload Type: TagString(=8)
				0x08,
				// Payload Length: 2
				0x00, 0x00, 0x00, 0x02,
				// Payload:
				//   - "Hello"
				0x00, 0x05,
				0x48, 0x65, 0x6C, 0x6C, 0x6F,
				//   - "World"
				0x00, 0x05,
				0x57, 0x6F, 0x72, 0x6C, 0x64,
			},
		},
	},
	{
		name: "positive case: ListTag - List",
		nbt: nbtTestCase[*ListPayload]{
			tagType: ListType,
			tagName: "List",
			payload: NewListPayload(
				NewListPayload(NewBytePayload(123)),
				NewListPayload(NewStringPayload("Test")),
			),
		},
		snbt: snbtTestCase{
			tagName: "List",
			payload: stringifyType{
				typeDefault: "[[123b], [\"Test\"]]",
				typeCompact: "[[123b],[\"Test\"]]",
				typePretty:  "[\n  [\n    123b\n  ],\n  [\n    \"Test\"\n  ]\n]",
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: List(=9)
				0x09,
			},
			tagName: []byte{
				// Name Length: 4
				0x00, 0x04,
				// Name: "List"
				0x4C, 0x69, 0x73, 0x74,
			},
			payload: []byte{
				// Payload Type: TagList(=9)
				0x09,
				// Payload Length: 2
				0x00, 0x00, 0x00, 0x02,
				// Payload:
				//   - [123b]
				0x01,
				0x00, 0x00, 0x00, 0x01,
				0x7B,
				//   - ["Test"]
				0x08,
				0x00, 0x00, 0x00, 0x01,
				0x00, 0x04,
				0x54, 0x65, 0x73, 0x74,
			},
		},
	},
	{
		name: "positive case: ListTag - Compound",
		nbt: nbtTestCase[*ListPayload]{
			tagType: ListType,
			tagName: "List",
			payload: NewListPayload(
				NewCompoundPayload(NewByteTag(NewTagName("Byte"), NewBytePayload(123)), NewEndTag()),
				NewCompoundPayload(NewStringTag(NewTagName("String"), NewStringPayload("Hello")), NewEndTag()),
			),
		},
		snbt: snbtTestCase{
			tagName: "List",
			payload: stringifyType{
				typeDefault: "[{Byte: 123b}, {String: \"Hello\"}]",
				typeCompact: "[{Byte:123b},{String:\"Hello\"}]",
				typePretty:  "[\n  {\n    Byte: 123b\n  },\n  {\n    String: \"Hello\"\n  }\n]",
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: List(=9)
				0x09,
			},
			tagName: []byte{
				// Name Length: 4
				0x00, 0x04,
				// Name: "List"
				0x4C, 0x69, 0x73, 0x74,
			},
			payload: []byte{
				// Payload Type: TagCompound(=10)
				0x0A,
				// Payload Length: 2
				0x00, 0x00, 0x00, 0x02,
				// Payload:
				//   - - Type: Byte(=1)
				//       Name: "Byte"
				//       Payload: 123b
				0x01,
				0x00, 0x04,
				0x42, 0x79, 0x74, 0x65,
				0x7B,
				//     - Type: End(=0)
				0x00,
				//   - - Type: String(=8)
				//       Name: "String"
				//       Payload: "Hello"
				0x08,
				0x00, 0x06,
				0x53, 0x74, 0x72, 0x69, 0x6E, 0x67,
				0x00, 0x05,
				0x48, 0x65, 0x6C, 0x6C, 0x6F,
				//     - Type: End(=0)
				0x00,
			},
		},
	},
	{
		name: "positive case: ListTag - empty",
		nbt: nbtTestCase[*ListPayload]{
			tagType: ListType,
			tagName: "List",
			payload: NewListPayload(),
		},
		snbt: snbtTestCase{
			tagName: "List",
			payload: stringifyType{
				typeDefault: "[]",
				typeCompact: "[]",
				typePretty:  "[]",
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: List(=9)
				0x09,
			},
			tagName: []byte{
				// Name Length: 4
				0x00, 0x04,
				// Name: "List"
				0x4C, 0x69, 0x73, 0x74,
			},
			payload: []byte{
				// Payload Type: TagEnd(=0)
				0x00,
				// Payload Length: 0
				0x00, 0x00, 0x00, 0x00,
				// Payload: []
			},
		},
	},
}

func TestNewListTag(t *testing.T) {
	type Case struct {
		name     string
		tagName  *TagName
		payload  *ListPayload
		expected *ListTag
	}

	cases := []Case{}

	for _, c := range listTagCases {
		cases = append(cases, Case{
			name:    c.name,
			tagName: &c.nbt.tagName,
			payload: c.nbt.payload,
			expected: &ListTag{
				tagName: &c.nbt.tagName,
				payload: c.nbt.payload,
			},
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewListTag(tt.tagName, tt.payload)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestListTag_TypeId(t *testing.T) {
	tag := new(ListTag)
	expected := ListType
	actual := tag.TypeId()
	assert.Equal(t, expected, actual)
}

func TestListTag_encode(t *testing.T) {
	type Case struct {
		name        string
		tag         *ListTag
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range listTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tag:         NewListTag(&c.nbt.tagName, c.nbt.payload),
			expected:    slices.Concat(c.raw.tagType, c.raw.tagName, c.raw.payload),
			expectedErr: nil,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			err := tt.tag.encode(buf)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, buf.Bytes())
			} else {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}

func TestListTag_decode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    *ListTag
		expectedErr error
	}

	cases := []Case{}

	for _, c := range listTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         slices.Concat(c.raw.tagType, c.raw.tagName, c.raw.payload),
			expected:    NewListTag(&c.nbt.tagName, c.nbt.payload),
			expectedErr: nil,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer(tt.raw)

			tag := new(ListTag)
			err := tag.decode(buf)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, tag)
			} else {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}

func TestListTag_stringify_default(t *testing.T) {
	type Case struct {
		name     string
		tag      *ListTag
		expected string
	}

	cases := []Case{}

	for _, c := range listTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tag:      NewListTag(&c.nbt.tagName, c.nbt.payload),
			expected: fmt.Sprintf("%s: %s", c.snbt.tagName, c.snbt.payload.typeDefault),
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.tag.stringify(" ", "", 0)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestListTag_stringify_compact(t *testing.T) {
	type Case struct {
		name     string
		tag      *ListTag
		expected string
	}

	cases := []Case{}

	for _, c := range listTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tag:      NewListTag(&c.nbt.tagName, c.nbt.payload),
			expected: fmt.Sprintf("%s:%s", c.snbt.tagName, c.snbt.payload.typeCompact),
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.tag.stringify("", "", 0)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestListTag_stringify_pretty(t *testing.T) {
	type Case struct {
		name     string
		tag      *ListTag
		expected string
	}

	cases := []Case{}

	for _, c := range listTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tag:      NewListTag(&c.nbt.tagName, c.nbt.payload),
			expected: fmt.Sprintf("%s: %s", c.snbt.tagName, c.snbt.payload.typePretty),
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.tag.stringify(" ", "  ", 0)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestNewListPayload(t *testing.T) {
	type Case struct {
		name     string
		values   []Payload
		expected *ListPayload
	}

	cases := []Case{}

	for _, c := range listTagCases {
		cases = append(cases, Case{
			name:     c.name,
			values:   []Payload(*c.nbt.payload),
			expected: c.nbt.payload,
		})
	}

	for _, tt := range cases {
		actual := NewListPayload(tt.values...)
		assert.Equal(t, tt.expected, actual)
	}
}

func TestListPayload_TypeId(t *testing.T) {
	payload := new(ListPayload)
	expected := ListType
	actual := payload.TypeId()
	assert.Equal(t, expected, actual)
}

func TestListPayload_encode(t *testing.T) {
	type Case struct {
		name        string
		payload     *ListPayload
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range listTagCases {
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
			err := tt.payload.encode(buf)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, buf.Bytes())
			} else {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}

func TestListPayload_decode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    *ListPayload
		expectedErr error
	}

	cases := []Case{}

	for _, c := range listTagCases {
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

			payload := new(ListPayload)
			err := payload.decode(buf)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, payload)
			} else {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}

func TestListPayload_stringify_default(t *testing.T) {
	type Case struct {
		name     string
		payload  *ListPayload
		expected string
	}

	cases := []Case{}

	for _, c := range listTagCases {
		cases = append(cases, Case{
			name:     c.name,
			payload:  c.nbt.payload,
			expected: c.snbt.payload.typeDefault,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.payload.stringify(" ", "", 0)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestListPayload_stringify_compact(t *testing.T) {
	type Case struct {
		name     string
		payload  *ListPayload
		expected string
	}

	cases := []Case{}

	for _, c := range listTagCases {
		cases = append(cases, Case{
			name:     c.name,
			payload:  c.nbt.payload,
			expected: c.snbt.payload.typeCompact,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.payload.stringify("", "", 0)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestListPayload_stringify_pretty(t *testing.T) {
	type Case struct {
		name     string
		payload  *ListPayload
		expected string
	}

	cases := []Case{}

	for _, c := range listTagCases {
		cases = append(cases, Case{
			name:     c.name,
			payload:  c.nbt.payload,
			expected: c.snbt.payload.typePretty,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.payload.stringify(" ", "  ", 0)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
