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

var compoundTagCases = []tagTestCase[*CompoundPayload]{
	{
		name: "positive case: CompoundTag - has items",
		nbt: nbtTestCase[*CompoundPayload]{
			tagType: CompoundType,
			tagName: "Compound",
			payload: NewCompoundPayload(
				NewShortTag(NewTagName("Short"), NewShortPayload(12345)),
				NewByteArrayTag(NewTagName("ByteArray"), NewByteArrayPayload(0, 1)),
				NewStringTag(NewTagName("String"), NewStringPayload("Hello")),
				NewListTag(NewTagName("List"), NewListPayload(NewBytePayload(123))),
				NewCompoundTag(NewTagName("Compound"), NewCompoundPayload(NewStringTag(NewTagName("String"), NewStringPayload("World")), NewEndTag())),
				NewEndTag(),
			),
		},
		snbt: snbtTestCase{
			tagName: "Compound",
			payload: "{ByteArray: [B; 0b, 1b], Compound: {String: \"World\"}, List: [123b], Short: 12345s, String: \"Hello\"}",
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: Compound(=10)
				0x0A,
			},
			tagName: []byte{
				// Name Length: 8
				0x00, 0x08,
				// Name: "Compound"
				0x43, 0x6F, 0x6D, 0x70, 0x6F, 0x75, 0x6E, 0x64,
			},
			payload: []byte{
				// Payload:
				//   - Type: Short(=2)
				//     Name: "Short"
				//     Payload: 12345s
				0x02,
				0x00, 0x05,
				0x53, 0x68, 0x6f, 0x72, 0x74,
				0x30, 0x39,
				//   - Type: ByteArray(=7)
				//     Name: "ByteArray"
				//     Payload: [B; 0b, 1b]
				0x07,
				0x00, 0x09,
				0x42, 0x79, 0x74, 0x65, 0x41, 0x72, 0x72, 0x61, 0x79,
				0x00, 0x00, 0x00, 0x02,
				0x00, 0x01,
				//   - Type: String(=8)
				//     Name: "String"
				//     Payload: "Hello"
				0x08,
				0x00, 0x06,
				0x53, 0x74, 0x72, 0x69, 0x6E, 0x67,
				0x00, 0x05,
				0x48, 0x65, 0x6C, 0x6C, 0x6F,
				//   - Type: List(=9)
				//     Name: "List"
				//     Payload: [123b]
				0x09,
				0x00, 0x04,
				0x4C, 0x69, 0x73, 0x74,
				0x01,
				0x00, 0x00, 0x00, 0x01,
				0x7B,
				//   - Type: Compound(=10)
				//     Name: "Compound"
				0x0A,
				0x00, 0x08,
				0x43, 0x6F, 0x6D, 0x70, 0x6F, 0x75, 0x6E, 0x64,
				//     Payload:
				//       - Type: String(=8)
				//         Name: "String"
				//         Payload: "World"
				0x08,
				0x00, 0x06,
				0x53, 0x74, 0x72, 0x69, 0x6E, 0x67,
				0x00, 0x05,
				0x57, 0x6F, 0x72, 0x6C, 0x64,
				//       - Type: End(=0)
				0x00,
				//   - Type: End(=0)
				0x00,
			},
		},
	},
	{
		name: "positive case: CompoundTag - empty",
		nbt: nbtTestCase[*CompoundPayload]{
			tagType: CompoundType,
			tagName: "Compound",
			payload: NewCompoundPayload(NewEndTag()),
		},
		snbt: snbtTestCase{
			tagName: "Compound",
			payload: "{}",
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: Compound(=10)
				0x0A,
			},
			tagName: []byte{
				// Name Length: 8
				0x00, 0x08,
				// Name: "Compound"
				0x43, 0x6F, 0x6D, 0x70, 0x6F, 0x75, 0x6E, 0x64,
			},
			payload: []byte{
				// Payload:
				//   - Type: End(=0)
				0x00,
			},
		},
	},
}

func TestNewCompoundTag(t *testing.T) {
	type Case struct {
		name     string
		tagName  *TagName
		payload  *CompoundPayload
		expected *CompoundTag
	}

	cases := []Case{}

	for _, c := range compoundTagCases {
		cases = append(cases, Case{
			name:    c.name,
			tagName: &c.nbt.tagName,
			payload: c.nbt.payload,
			expected: &CompoundTag{
				tagName: &c.nbt.tagName,
				payload: c.nbt.payload,
			},
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewCompoundTag(tt.tagName, tt.payload)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestCompoundTag_TypeId(t *testing.T) {
	tag := new(CompoundTag)
	expected := CompoundType
	actual := tag.TypeId()
	assert.Equal(t, expected, actual)
}

func TestCompoundTag_encode(t *testing.T) {
	type Case struct {
		name        string
		tag         *CompoundTag
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range compoundTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tag:         NewCompoundTag(&c.nbt.tagName, c.nbt.payload),
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

func TestCompoundTag_decode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    *CompoundTag
		expectedErr error
	}

	cases := []Case{}

	for _, c := range compoundTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         slices.Concat(c.raw.tagType, c.raw.tagName, c.raw.payload),
			expected:    NewCompoundTag(&c.nbt.tagName, c.nbt.payload),
			expectedErr: nil,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer(tt.raw)

			tag := new(CompoundTag)
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

func TestCompoundTag_stringify_default(t *testing.T) {
	type Case struct {
		name     string
		tag      *CompoundTag
		expected string
	}

	cases := []Case{}

	for _, c := range compoundTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tag:      NewCompoundTag(&c.nbt.tagName, c.nbt.payload),
			expected: fmt.Sprintf("%s: %s", c.snbt.tagName, c.snbt.payload),
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.tag.stringify(" ", "", 0)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestNewCompoundPayload(t *testing.T) {
	type Case struct {
		name     string
		values   []Tag
		expected *CompoundPayload
	}

	cases := []Case{}

	for _, c := range compoundTagCases {
		cases = append(cases, Case{
			name:     c.name,
			values:   []Tag(*c.nbt.payload),
			expected: c.nbt.payload,
		})
	}

	for _, tt := range cases {
		actual := NewCompoundPayload(tt.values...)
		assert.Equal(t, tt.expected, actual)
	}
}

func TestCompoundPayload_TypeId(t *testing.T) {
	payload := new(CompoundPayload)
	expected := CompoundType
	actual := payload.TypeId()
	assert.Equal(t, expected, actual)
}

func TestCompoundPayload_encode(t *testing.T) {
	type Case struct {
		name        string
		payload     *CompoundPayload
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range compoundTagCases {
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

func TestCompoundPayload_decode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    *CompoundPayload
		expectedErr error
	}

	cases := []Case{}

	for _, c := range compoundTagCases {
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

			payload := new(CompoundPayload)
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

func TestCompoundPayload_stringify_default(t *testing.T) {
	type Case struct {
		name     string
		payload  *CompoundPayload
		expected string
	}

	cases := []Case{}

	for _, c := range compoundTagCases {
		cases = append(cases, Case{
			name:     c.name,
			payload:  c.nbt.payload,
			expected: c.snbt.payload,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.payload.stringify(" ", "", 0)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
