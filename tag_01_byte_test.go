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
	"github.com/Aton-Kish/gonbt/snbt"
	"github.com/stretchr/testify/assert"
)

var byteTagCases = []tagTestCase[*BytePayload]{
	{
		name: `positive case: ByteTag`,
		nbt: nbtTestCase[*BytePayload]{
			tagType: ByteType,
			tagName: `Byte`,
			payload: NewBytePayload(123),
		},
		snbt: snbtTestCase{
			tagName: `Byte`,
			payload: stringifyType{
				typeDefault: `123b`,
				typeCompact: `123b`,
				typePretty:  `123b`,
			},
		},
		json: jsonTestCase{
			tagName: `"Byte"`,
			payload: stringifyType{
				typeDefault: `123`,
				typeCompact: `123`,
				typePretty:  `123`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: Byte(=1)
				0x01,
			},
			tagName: []byte{
				// Name Length: 4
				0x00, 0x04,
				// Name: Byte
				0x42, 0x79, 0x74, 0x65,
			},
			payload: []byte{
				// Payload: 123b
				0x7B,
			},
		},
	},
}

func TestNewByteTag(t *testing.T) {
	type Case struct {
		name     string
		tagName  *TagName
		payload  *BytePayload
		expected *ByteTag
	}

	cases := []Case{}

	for _, c := range byteTagCases {
		cases = append(cases, Case{
			name:    c.name,
			tagName: &c.nbt.tagName,
			payload: c.nbt.payload,
			expected: &ByteTag{
				tagName: &c.nbt.tagName,
				payload: c.nbt.payload,
			},
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewByteTag(tt.tagName, tt.payload)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestByteTag_TypeId(t *testing.T) {
	tag := new(ByteTag)
	expected := ByteType
	actual := tag.TypeId()
	assert.Equal(t, expected, actual)
}

func TestByteTag_encode(t *testing.T) {
	type Case struct {
		name        string
		tag         *ByteTag
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range byteTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tag:         NewByteTag(&c.nbt.tagName, c.nbt.payload),
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

func TestByteTag_decode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    *ByteTag
		expectedErr error
	}

	cases := []Case{}

	for _, c := range byteTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         slices.Concat(c.raw.tagType, c.raw.tagName, c.raw.payload),
			expected:    NewByteTag(&c.nbt.tagName, c.nbt.payload),
			expectedErr: nil,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer(tt.raw)

			tag := new(ByteTag)
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

func TestByteTag_stringify_default(t *testing.T) {
	type Case struct {
		name     string
		tag      *ByteTag
		expected string
	}

	cases := []Case{}

	for _, c := range byteTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tag:      NewByteTag(&c.nbt.tagName, c.nbt.payload),
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

func TestByteTag_stringify_compact(t *testing.T) {
	type Case struct {
		name     string
		tag      *ByteTag
		expected string
	}

	cases := []Case{}

	for _, c := range byteTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tag:      NewByteTag(&c.nbt.tagName, c.nbt.payload),
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

func TestByteTag_stringify_pretty(t *testing.T) {
	type Case struct {
		name     string
		tag      *ByteTag
		expected string
	}

	cases := []Case{}

	for _, c := range byteTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tag:      NewByteTag(&c.nbt.tagName, c.nbt.payload),
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

func TestByteTag_json_default(t *testing.T) {
	type Case struct {
		name     string
		tag      *ByteTag
		expected string
	}

	cases := []Case{}

	for _, c := range byteTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tag:      NewByteTag(&c.nbt.tagName, c.nbt.payload),
			expected: fmt.Sprintf("%s: %s", c.json.tagName, c.json.payload.typeDefault),
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.tag.json(" ", "", 0)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestByteTag_json_compact(t *testing.T) {
	type Case struct {
		name     string
		tag      *ByteTag
		expected string
	}

	cases := []Case{}

	for _, c := range byteTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tag:      NewByteTag(&c.nbt.tagName, c.nbt.payload),
			expected: fmt.Sprintf("%s:%s", c.json.tagName, c.json.payload.typeCompact),
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.tag.json("", "", 0)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestByteTag_json_pretty(t *testing.T) {
	type Case struct {
		name     string
		tag      *ByteTag
		expected string
	}

	cases := []Case{}

	for _, c := range byteTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tag:      NewByteTag(&c.nbt.tagName, c.nbt.payload),
			expected: fmt.Sprintf("%s: %s", c.json.tagName, c.json.payload.typePretty),
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.tag.json(" ", "  ", 0)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestNewBytePayload(t *testing.T) {
	type Case struct {
		name     string
		value    int8
		expected *BytePayload
	}

	cases := []Case{}

	for _, c := range byteTagCases {
		cases = append(cases, Case{
			name:     c.name,
			value:    int8(*c.nbt.payload),
			expected: c.nbt.payload,
		})
	}

	for _, tt := range cases {
		actual := NewBytePayload(tt.value)
		assert.Equal(t, tt.expected, actual)
	}
}

func TestBytePayload_TypeId(t *testing.T) {
	payload := new(BytePayload)
	expected := ByteType
	actual := payload.TypeId()
	assert.Equal(t, expected, actual)
}

func TestBytePayload_encode(t *testing.T) {
	type Case struct {
		name        string
		payload     *BytePayload
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range byteTagCases {
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

func TestBytePayload_decode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    *BytePayload
		expectedErr error
	}

	cases := []Case{}

	for _, c := range byteTagCases {
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

			payload := new(BytePayload)
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

func TestBytePayload_stringify_default(t *testing.T) {
	type Case struct {
		name     string
		payload  *BytePayload
		expected string
	}

	cases := []Case{}

	for _, c := range byteTagCases {
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

func TestBytePayload_stringify_compact(t *testing.T) {
	type Case struct {
		name     string
		payload  *BytePayload
		expected string
	}

	cases := []Case{}

	for _, c := range byteTagCases {
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

func TestBytePayload_stringify_pretty(t *testing.T) {
	type Case struct {
		name     string
		payload  *BytePayload
		expected string
	}

	cases := []Case{}

	for _, c := range byteTagCases {
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

func TestBytePayload_parse(t *testing.T) {
	type Case struct {
		name        string
		snbt        string
		expected    *BytePayload
		expectedErr error
	}

	cases := []Case{}

	for _, c := range byteTagCases {
		cases = append(cases, Case{
			name:        c.name,
			snbt:        c.snbt.payload.typeDefault,
			expected:    c.nbt.payload,
			expectedErr: nil,
		})
		cases = append(cases, Case{
			name:        c.name,
			snbt:        c.snbt.payload.typeCompact,
			expected:    c.nbt.payload,
			expectedErr: nil,
		})
		cases = append(cases, Case{
			name:        c.name,
			snbt:        c.snbt.payload.typePretty,
			expected:    c.nbt.payload,
			expectedErr: nil,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			parser := snbt.NewParser(tt.snbt)
			err := parser.Compact()
			assert.NoError(t, err)

			payload := new(BytePayload)
			err = payload.parse(parser)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, payload)
			} else {
				assert.Error(t, err, tt.expectedErr.Error())
			}
		})
	}
}

func TestBytePayload_json_default(t *testing.T) {
	type Case struct {
		name     string
		payload  *BytePayload
		expected string
	}

	cases := []Case{}

	for _, c := range byteTagCases {
		cases = append(cases, Case{
			name:     c.name,
			payload:  c.nbt.payload,
			expected: c.json.payload.typeDefault,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.payload.json(" ", "", 0)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestBytePayload_json_compact(t *testing.T) {
	type Case struct {
		name     string
		payload  *BytePayload
		expected string
	}

	cases := []Case{}

	for _, c := range byteTagCases {
		cases = append(cases, Case{
			name:     c.name,
			payload:  c.nbt.payload,
			expected: c.json.payload.typeCompact,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.payload.json("", "", 0)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestBytePayload_json_pretty(t *testing.T) {
	type Case struct {
		name     string
		payload  *BytePayload
		expected string
	}

	cases := []Case{}

	for _, c := range byteTagCases {
		cases = append(cases, Case{
			name:     c.name,
			payload:  c.nbt.payload,
			expected: c.json.payload.typePretty,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.payload.json(" ", "  ", 0)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
