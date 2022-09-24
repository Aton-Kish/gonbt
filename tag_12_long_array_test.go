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

var longArrayTagCases = []tagTestCase[*LongArrayPayload]{
	{
		name: `positive case: LongArrayTag - has items`,
		nbt: nbtTestCase[*LongArrayPayload]{
			tagType: LongArrayType,
			tagName: `LongArray`,
			payload: NewLongArrayPayload(0, 1, 2, 3),
		},
		snbt: snbtTestCase{
			tagName: `LongArray`,
			payload: stringifyType{
				typeDefault: `[L; 0L, 1L, 2L, 3L]`,
				typeCompact: `[L;0L,1L,2L,3L]`,
				typePretty:  `[L; 0L, 1L, 2L, 3L]`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: LongArray(=12)
				0x0C,
			},
			tagName: []byte{
				// Name Length: 9
				0x00, 0x09,
				// Name: LongArray
				0x4C, 0x6F, 0x6E, 0x67, 0x41, 0x72, 0x72, 0x61, 0x79,
			},
			payload: []byte{
				// Payload Length: 4
				0x00, 0x00, 0x00, 0x04,
				// Payload: [L; 0L, 1L, 2L, 3L]
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03,
			},
		},
	},
	{
		name: `positive case: LongArrayTag - empty`,
		nbt: nbtTestCase[*LongArrayPayload]{
			tagType: LongArrayType,
			tagName: `LongArray`,
			payload: NewLongArrayPayload(),
		},
		snbt: snbtTestCase{
			tagName: `LongArray`,
			payload: stringifyType{
				typeDefault: `[L; ]`,
				typeCompact: `[L;]`,
				typePretty:  `[L; ]`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: LongArray(=12)
				0x0C,
			},
			tagName: []byte{
				// Name Length: 9
				0x00, 0x09,
				// Name: LongArray
				0x4C, 0x6F, 0x6E, 0x67, 0x41, 0x72, 0x72, 0x61, 0x79,
			},
			payload: []byte{
				// Payload Length: 0
				0x00, 0x00, 0x00, 0x00,
				// Payload: [L; ]
			},
		},
	},
}

func TestNewLongArrayTag(t *testing.T) {
	type Case struct {
		name     string
		tagName  *TagName
		payload  *LongArrayPayload
		expected *LongArrayTag
	}

	cases := []Case{}

	for _, c := range longArrayTagCases {
		cases = append(cases, Case{
			name:    c.name,
			tagName: &c.nbt.tagName,
			payload: c.nbt.payload,
			expected: &LongArrayTag{
				tagName: &c.nbt.tagName,
				payload: c.nbt.payload,
			},
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewLongArrayTag(tt.tagName, tt.payload)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestLongArrayTag_TypeId(t *testing.T) {
	tag := new(LongArrayTag)
	expected := LongArrayType
	actual := tag.TypeId()
	assert.Equal(t, expected, actual)
}

func TestLongArrayTag_encode(t *testing.T) {
	type Case struct {
		name        string
		tag         *LongArrayTag
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range longArrayTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tag:         NewLongArrayTag(&c.nbt.tagName, c.nbt.payload),
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

func TestLongArrayTag_decode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    *LongArrayTag
		expectedErr error
	}

	cases := []Case{}

	for _, c := range longArrayTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         slices.Concat(c.raw.tagType, c.raw.tagName, c.raw.payload),
			expected:    NewLongArrayTag(&c.nbt.tagName, c.nbt.payload),
			expectedErr: nil,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer(tt.raw)

			tag := new(LongArrayTag)
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

func TestLongArrayTag_stringify_default(t *testing.T) {
	type Case struct {
		name     string
		tag      *LongArrayTag
		expected string
	}

	cases := []Case{}

	for _, c := range longArrayTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tag:      NewLongArrayTag(&c.nbt.tagName, c.nbt.payload),
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

func TestLongArrayTag_stringify_compact(t *testing.T) {
	type Case struct {
		name     string
		tag      *LongArrayTag
		expected string
	}

	cases := []Case{}

	for _, c := range longArrayTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tag:      NewLongArrayTag(&c.nbt.tagName, c.nbt.payload),
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

func TestLongArrayTag_stringify_pretty(t *testing.T) {
	type Case struct {
		name     string
		tag      *LongArrayTag
		expected string
	}

	cases := []Case{}

	for _, c := range longArrayTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tag:      NewLongArrayTag(&c.nbt.tagName, c.nbt.payload),
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

func TestNewLongArrayPayload(t *testing.T) {
	type Case struct {
		name     string
		values   []int64
		expected *LongArrayPayload
	}

	cases := []Case{}

	for _, c := range longArrayTagCases {
		cases = append(cases, Case{
			name:     c.name,
			values:   []int64(*c.nbt.payload),
			expected: c.nbt.payload,
		})
	}

	for _, tt := range cases {
		actual := NewLongArrayPayload(tt.values...)
		assert.Equal(t, tt.expected, actual)
	}
}

func TestLongArrayPayload_TypeId(t *testing.T) {
	payload := new(LongArrayPayload)
	expected := LongArrayType
	actual := payload.TypeId()
	assert.Equal(t, expected, actual)
}

func TestLongArrayPayload_encode(t *testing.T) {
	type Case struct {
		name        string
		payload     *LongArrayPayload
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range longArrayTagCases {
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

func TestLongArrayPayload_decode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    *LongArrayPayload
		expectedErr error
	}

	cases := []Case{}

	for _, c := range longArrayTagCases {
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

			payload := new(LongArrayPayload)
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

func TestLongArrayPayload_stringify_default(t *testing.T) {
	type Case struct {
		name     string
		payload  *LongArrayPayload
		expected string
	}

	cases := []Case{}

	for _, c := range longArrayTagCases {
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

func TestLongArrayPayload_stringify_compact(t *testing.T) {
	type Case struct {
		name     string
		payload  *LongArrayPayload
		expected string
	}

	cases := []Case{}

	for _, c := range longArrayTagCases {
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

func TestLongArrayPayload_stringify_pretty(t *testing.T) {
	type Case struct {
		name     string
		payload  *LongArrayPayload
		expected string
	}

	cases := []Case{}

	for _, c := range longArrayTagCases {
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
