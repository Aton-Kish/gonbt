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

var endTagCases = []tagTestCase[Payload]{
	{
		name: "positive case: EndTag",
		nbt: nbtTestCase[Payload]{
			tagType: EndType,
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: End(=0)
				0x00,
			},
		},
	},
}

func TestNewEndTag(t *testing.T) {
	type Case struct {
		name     string
		expected *EndTag
	}

	cases := []Case{}

	for _, c := range byteTagCases {
		cases = append(cases, Case{
			name:     c.name,
			expected: &EndTag{},
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewEndTag()
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestEndTag_TypeId(t *testing.T) {
	tag := new(EndTag)
	expected := EndType
	actual := tag.TypeId()
	assert.Equal(t, expected, actual)
}

func TestEndTag_encode(t *testing.T) {
	type Case struct {
		name        string
		tag         *EndTag
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range endTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tag:         NewEndTag(),
			expected:    c.raw.tagType,
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

func TestEndTag_decode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    Tag
		expectedErr error
	}

	cases := []Case{}

	for _, c := range endTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         c.raw.tagType,
			expected:    NewEndTag(),
			expectedErr: nil,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer(tt.raw)

			tag := new(EndTag)
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
