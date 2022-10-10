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

var intArrayTagCases = []tagTestCase[*IntArrayPayload]{
	{
		name: `positive case: IntArrayTag - has items`,
		nbt: nbtTestCase[*IntArrayPayload]{
			tagType: IntArrayType,
			tagName: `IntArray`,
			payload: NewIntArrayPayload(0, 1, 2, 3),
		},
		snbt: snbtTestCase{
			tagName: `IntArray`,
			payload: stringifyType{
				typeDefault: `[I; 0, 1, 2, 3]`,
				typeCompact: `[I;0,1,2,3]`,
				typePretty:  `[I; 0, 1, 2, 3]`,
			},
		},
		json: jsonTestCase{
			tagName: `"IntArray"`,
			payload: stringifyType{
				typeDefault: `[0, 1, 2, 3]`,
				typeCompact: `[0,1,2,3]`,
				typePretty: `[
  0,
  1,
  2,
  3
]`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: IntArray(=11)
				0x0B,
			},
			tagName: []byte{
				// Name Length: 8
				0x00, 0x08,
				// Name: IntArray
				0x49, 0x6E, 0x74, 0x41, 0x72, 0x72, 0x61, 0x79,
			},
			payload: []byte{
				// Payload Length: 4
				0x00, 0x00, 0x00, 0x04,
				// Payload: [I; 0, 1, 2, 3]
				0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x01,
				0x00, 0x00, 0x00, 0x02,
				0x00, 0x00, 0x00, 0x03,
			},
		},
	},
	{
		name: `positive case: IntArrayTag - empty`,
		nbt: nbtTestCase[*IntArrayPayload]{
			tagType: IntArrayType,
			tagName: `IntArray`,
			payload: NewIntArrayPayload(),
		},
		snbt: snbtTestCase{
			tagName: `IntArray`,
			payload: stringifyType{
				typeDefault: `[I; ]`,
				typeCompact: `[I;]`,
				typePretty:  `[I; ]`,
			},
		},
		json: jsonTestCase{
			tagName: `"IntArray"`,
			payload: stringifyType{
				typeDefault: `[]`,
				typeCompact: `[]`,
				typePretty:  `[]`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: IntArray(=11)
				0x0B,
			},
			tagName: []byte{
				// Name Length: 8
				0x00, 0x08,
				// Name: IntArray
				0x49, 0x6E, 0x74, 0x41, 0x72, 0x72, 0x61, 0x79,
			},
			payload: []byte{
				// Payload Length: 0
				0x00, 0x00, 0x00, 0x00,
				// Payload: [I; ]
			},
		},
	},
}

func TestNewIntArrayTag(t *testing.T) {
	type Case struct {
		name     string
		tagName  *TagName
		payload  *IntArrayPayload
		expected *IntArrayTag
	}

	cases := []Case{}

	for _, c := range intArrayTagCases {
		cases = append(cases, Case{
			name:    c.name,
			tagName: &c.nbt.tagName,
			payload: c.nbt.payload,
			expected: &IntArrayTag{
				tagName: &c.nbt.tagName,
				payload: c.nbt.payload,
			},
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewIntArrayTag(tt.tagName, tt.payload)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestIntArrayTag_TypeId(t *testing.T) {
	tag := new(IntArrayTag)
	expected := IntArrayType
	actual := tag.TypeId()
	assert.Equal(t, expected, actual)
}

func TestIntArrayTag_encode(t *testing.T) {
	type Case struct {
		name        string
		tag         *IntArrayTag
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range intArrayTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tag:         NewIntArrayTag(&c.nbt.tagName, c.nbt.payload),
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

func TestIntArrayTag_decode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    *IntArrayTag
		expectedErr error
	}

	cases := []Case{}

	for _, c := range intArrayTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         slices.Concat(c.raw.tagType, c.raw.tagName, c.raw.payload),
			expected:    NewIntArrayTag(&c.nbt.tagName, c.nbt.payload),
			expectedErr: nil,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer(tt.raw)

			tag := new(IntArrayTag)
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

func TestIntArrayTag_stringify_default(t *testing.T) {
	type Case struct {
		name     string
		tag      *IntArrayTag
		expected string
	}

	cases := []Case{}

	for _, c := range intArrayTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tag:      NewIntArrayTag(&c.nbt.tagName, c.nbt.payload),
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

func TestIntArrayTag_stringify_compact(t *testing.T) {
	type Case struct {
		name     string
		tag      *IntArrayTag
		expected string
	}

	cases := []Case{}

	for _, c := range intArrayTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tag:      NewIntArrayTag(&c.nbt.tagName, c.nbt.payload),
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

func TestIntArrayTag_stringify_pretty(t *testing.T) {
	type Case struct {
		name     string
		tag      *IntArrayTag
		expected string
	}

	cases := []Case{}

	for _, c := range intArrayTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tag:      NewIntArrayTag(&c.nbt.tagName, c.nbt.payload),
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

func TestIntArrayTag_json_default(t *testing.T) {
	type Case struct {
		name     string
		tag      *IntArrayTag
		expected string
	}

	cases := []Case{}

	for _, c := range intArrayTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tag:      NewIntArrayTag(&c.nbt.tagName, c.nbt.payload),
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

func TestIntArrayTag_json_compact(t *testing.T) {
	type Case struct {
		name     string
		tag      *IntArrayTag
		expected string
	}

	cases := []Case{}

	for _, c := range intArrayTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tag:      NewIntArrayTag(&c.nbt.tagName, c.nbt.payload),
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

func TestIntArrayTag_json_pretty(t *testing.T) {
	type Case struct {
		name     string
		tag      *IntArrayTag
		expected string
	}

	cases := []Case{}

	for _, c := range intArrayTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tag:      NewIntArrayTag(&c.nbt.tagName, c.nbt.payload),
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

func TestNewIntArrayPayload(t *testing.T) {
	type Case struct {
		name     string
		values   []int32
		expected *IntArrayPayload
	}

	cases := []Case{}

	for _, c := range intArrayTagCases {
		cases = append(cases, Case{
			name:     c.name,
			values:   []int32(*c.nbt.payload),
			expected: c.nbt.payload,
		})
	}

	for _, tt := range cases {
		actual := NewIntArrayPayload(tt.values...)
		assert.Equal(t, tt.expected, actual)
	}
}

func TestIntArrayPayload_TypeId(t *testing.T) {
	payload := new(IntArrayPayload)
	expected := IntArrayType
	actual := payload.TypeId()
	assert.Equal(t, expected, actual)
}

func TestIntArrayPayload_encode(t *testing.T) {
	type Case struct {
		name        string
		payload     *IntArrayPayload
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range intArrayTagCases {
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

func TestIntArrayPayload_decode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    *IntArrayPayload
		expectedErr error
	}

	cases := []Case{}

	for _, c := range intArrayTagCases {
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

			payload := new(IntArrayPayload)
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

func TestIntArrayPayload_stringify_default(t *testing.T) {
	type Case struct {
		name     string
		payload  *IntArrayPayload
		expected string
	}

	cases := []Case{}

	for _, c := range intArrayTagCases {
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

func TestIntArrayPayload_stringify_compact(t *testing.T) {
	type Case struct {
		name     string
		payload  *IntArrayPayload
		expected string
	}

	cases := []Case{}

	for _, c := range intArrayTagCases {
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

func TestIntArrayPayload_stringify_pretty(t *testing.T) {
	type Case struct {
		name     string
		payload  *IntArrayPayload
		expected string
	}

	cases := []Case{}

	for _, c := range intArrayTagCases {
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

func TestIntArrayPayload_parse(t *testing.T) {
	type Case struct {
		name        string
		snbt        string
		expected    *IntArrayPayload
		expectedErr error
	}

	cases := []Case{}

	for _, c := range intArrayTagCases {
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

			payload := NewIntArrayPayload()
			err = payload.parse(parser)

			if tt.expectedErr == nil {
				assert.ErrorIs(t, err, snbt.StopIterationError)
				assert.Equal(t, tt.expected, payload)
			} else {
				assert.Error(t, err, tt.expectedErr.Error())
			}
		})
	}
}

func TestIntArrayPayload_json_default(t *testing.T) {
	type Case struct {
		name     string
		payload  *IntArrayPayload
		expected string
	}

	cases := []Case{}

	for _, c := range intArrayTagCases {
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

func TestIntArrayPayload_json_compact(t *testing.T) {
	type Case struct {
		name     string
		payload  *IntArrayPayload
		expected string
	}

	cases := []Case{}

	for _, c := range intArrayTagCases {
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

func TestIntArrayPayload_json_pretty(t *testing.T) {
	type Case struct {
		name     string
		payload  *IntArrayPayload
		expected string
	}

	cases := []Case{}

	for _, c := range intArrayTagCases {
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
