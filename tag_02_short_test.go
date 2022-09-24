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

var shortTagCases = []tagTestCase[*ShortPayload]{
	{
		name: `positive case: ShortTag`,
		nbt: nbtTestCase[*ShortPayload]{
			tagType: ShortType,
			tagName: `Short`,
			payload: NewShortPayload(12345),
		},
		snbt: snbtTestCase{
			tagName: `Short`,
			payload: stringifyType{
				typeDefault: `12345s`,
				typeCompact: `12345s`,
				typePretty:  `12345s`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: Short(=2)
				0x02,
			},
			tagName: []byte{
				// Name Length: 5
				0x00, 0x05,
				// Name: Short
				0x53, 0x68, 0x6f, 0x72, 0x74,
			},
			payload: []byte{
				// Payload: 12345s
				0x30, 0x39,
			},
		},
	},
}

func TestNewShortTag(t *testing.T) {
	type Case struct {
		name     string
		tagName  *TagName
		payload  *ShortPayload
		expected *ShortTag
	}

	cases := []Case{}

	for _, c := range shortTagCases {
		cases = append(cases, Case{
			name:    c.name,
			tagName: &c.nbt.tagName,
			payload: c.nbt.payload,
			expected: &ShortTag{
				tagName: &c.nbt.tagName,
				payload: c.nbt.payload,
			},
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewShortTag(tt.tagName, tt.payload)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestShortTag_TypeId(t *testing.T) {
	tag := new(ShortTag)
	expected := ShortType
	actual := tag.TypeId()
	assert.Equal(t, expected, actual)
}

func TestShortTag_encode(t *testing.T) {
	type Case struct {
		name        string
		tag         *ShortTag
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range shortTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tag:         NewShortTag(&c.nbt.tagName, c.nbt.payload),
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

func TestShortTag_decode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    *ShortTag
		expectedErr error
	}

	cases := []Case{}

	for _, c := range shortTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         slices.Concat(c.raw.tagType, c.raw.tagName, c.raw.payload),
			expected:    NewShortTag(&c.nbt.tagName, c.nbt.payload),
			expectedErr: nil,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer(tt.raw)

			tag := new(ShortTag)
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

func TestShortTag_stringify_default(t *testing.T) {
	type Case struct {
		name     string
		tag      *ShortTag
		expected string
	}

	cases := []Case{}

	for _, c := range shortTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tag:      NewShortTag(&c.nbt.tagName, c.nbt.payload),
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

func TestShortTag_stringify_compact(t *testing.T) {
	type Case struct {
		name     string
		tag      *ShortTag
		expected string
	}

	cases := []Case{}

	for _, c := range shortTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tag:      NewShortTag(&c.nbt.tagName, c.nbt.payload),
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

func TestShortTag_stringify_pretty(t *testing.T) {
	type Case struct {
		name     string
		tag      *ShortTag
		expected string
	}

	cases := []Case{}

	for _, c := range shortTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tag:      NewShortTag(&c.nbt.tagName, c.nbt.payload),
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

func TestNewShortPayload(t *testing.T) {
	type Case struct {
		name     string
		value    int16
		expected *ShortPayload
	}

	cases := []Case{}

	for _, c := range shortTagCases {
		cases = append(cases, Case{
			name:     c.name,
			value:    int16(*c.nbt.payload),
			expected: c.nbt.payload,
		})
	}

	for _, tt := range cases {
		actual := NewShortPayload(tt.value)
		assert.Equal(t, tt.expected, actual)
	}
}

func TestShortPayload_TypeId(t *testing.T) {
	payload := new(ShortPayload)
	expected := ShortType
	actual := payload.TypeId()
	assert.Equal(t, expected, actual)
}

func TestShortPayload_encode(t *testing.T) {
	type Case struct {
		name        string
		payload     *ShortPayload
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range shortTagCases {
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

func TestShortPayload_decode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    *ShortPayload
		expectedErr error
	}

	cases := []Case{}

	for _, c := range shortTagCases {
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

			payload := new(ShortPayload)
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

func TestShortPayload_stringify_default(t *testing.T) {
	type Case struct {
		name     string
		payload  *ShortPayload
		expected string
	}

	cases := []Case{}

	for _, c := range shortTagCases {
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

func TestShortPayload_stringify_compact(t *testing.T) {
	type Case struct {
		name     string
		payload  *ShortPayload
		expected string
	}

	cases := []Case{}

	for _, c := range shortTagCases {
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

func TestShortPayload_stringify_pretty(t *testing.T) {
	type Case struct {
		name     string
		payload  *ShortPayload
		expected string
	}

	cases := []Case{}

	for _, c := range shortTagCases {
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
