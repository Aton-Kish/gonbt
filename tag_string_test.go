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

// Code generated by testgen.go; DO NOT EDIT.

package nbt

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/Aton-Kish/gonbt/slices"
	"github.com/stretchr/testify/assert"
)

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

func TestStringTag_encode(t *testing.T) {
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
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr, err)
			}
		})
	}
}

func TestStringTag_decode(t *testing.T) {
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
			raw:         slices.Concat(c.raw.tagType, c.raw.tagName, c.raw.payload),
			expected:    NewStringTag(&c.nbt.tagName, c.nbt.payload),
			expectedErr: nil,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer(tt.raw)

			tag := new(StringTag)
			err := tag.decode(buf)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, tag)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr, err)
			}
		})
	}
}

func TestStringTag_stringify_default(t *testing.T) {
	type Case struct {
		name     string
		tag      *StringTag
		expected string
	}

	cases := []Case{}

	for _, c := range stringTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tag:      NewStringTag(&c.nbt.tagName, c.nbt.payload),
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

func TestStringTag_stringify_compact(t *testing.T) {
	type Case struct {
		name     string
		tag      *StringTag
		expected string
	}

	cases := []Case{}

	for _, c := range stringTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tag:      NewStringTag(&c.nbt.tagName, c.nbt.payload),
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

func TestStringTag_stringify_pretty(t *testing.T) {
	type Case struct {
		name     string
		tag      *StringTag
		expected string
	}

	cases := []Case{}

	for _, c := range stringTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tag:      NewStringTag(&c.nbt.tagName, c.nbt.payload),
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

func TestStringTag_json_default(t *testing.T) {
	type Case struct {
		name     string
		tag      *StringTag
		expected string
	}

	cases := []Case{}

	for _, c := range stringTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tag:      NewStringTag(&c.nbt.tagName, c.nbt.payload),
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

func TestStringTag_json_compact(t *testing.T) {
	type Case struct {
		name     string
		tag      *StringTag
		expected string
	}

	cases := []Case{}

	for _, c := range stringTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tag:      NewStringTag(&c.nbt.tagName, c.nbt.payload),
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

func TestStringTag_json_pretty(t *testing.T) {
	type Case struct {
		name     string
		tag      *StringTag
		expected string
	}

	cases := []Case{}

	for _, c := range stringTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tag:      NewStringTag(&c.nbt.tagName, c.nbt.payload),
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
