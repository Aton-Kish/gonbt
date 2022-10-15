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

	"github.com/Aton-Kish/gonbt/pointer"
	"github.com/stretchr/testify/assert"
)

func TestNewTagName(t *testing.T) {
	type Case struct {
		name     string
		value    string
		expected *TagName
	}

	cases := []Case{}

	for _, c := range tagCases {
		cases = append(cases, Case{
			name:     c.name,
			value:    string(c.nbt.tagName),
			expected: pointer.Pointer(c.nbt.tagName),
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

	for _, c := range excludeEndTagCases {
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
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr, err)
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

	for _, c := range excludeEndTagCases {
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
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr, err)
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
			name:     `positive case: quotation - Test`,
			tagName:  TagName(`Test`),
			expected: `Test`,
		},
		{
			name:     `positive case: quotation - '"Test'`,
			tagName:  TagName(`"Test`),
			expected: `'"Test'`,
		},
		{
			name:     `positive case: quotation - "'Test"`,
			tagName:  TagName(`'Test`),
			expected: `"'Test"`,
		},
		{
			name:     `positive case: quotation - "\"'Test"`,
			tagName:  TagName(`"'Test`),
			expected: `"\"'Test"`,
		},
	}

	for _, c := range excludeEndTagCases {
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

func TestTagName_json(t *testing.T) {
	type Case struct {
		name     string
		tagName  TagName
		expected string
	}

	cases := []Case{
		{
			name:     `positive case: quotation - Test`,
			tagName:  TagName(`Test`),
			expected: `"Test"`,
		},
		{
			name:     `positive case: quotation - '"Test'`,
			tagName:  TagName(`"Test`),
			expected: `"\"Test"`,
		},
		{
			name:     `positive case: quotation - "'Test"`,
			tagName:  TagName(`'Test`),
			expected: `"'Test"`,
		},
		{
			name:     `positive case: quotation - "\"'Test"`,
			tagName:  TagName(`"'Test`),
			expected: `"\"'Test"`,
		},
	}

	for _, c := range excludeEndTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tagName:  c.nbt.tagName,
			expected: c.json.tagName,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.tagName.json()
			assert.Equal(t, tt.expected, actual)
		})
	}
}
