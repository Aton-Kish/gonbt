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

	"github.com/Aton-Kish/gonbt/slices"
	"github.com/stretchr/testify/assert"
)

var intTagCases = []tagTestCase[*IntPayload]{
	{
		name: "positive case: IntTag",
		nbt: nbtTestCase[*IntPayload]{
			tagType: IntType,
			tagName: "Int",
			payload: NewIntPayload(123456789),
		},
		snbt: snbtTestCase{
			tagName: "Int",
			payload: "123456789",
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: Int(=3)
				0x03,
			},
			tagName: []byte{
				// Name Length: 3
				0x00, 0x03,
				// Name: "Int"
				0x49, 0x6E, 0x74,
			},
			payload: []byte{
				// Payload: 123456789
				0x07, 0x5B, 0xCD, 0x15,
			},
		},
	},
}

func TestNewIntTag(t *testing.T) {
	type Case struct {
		name     string
		tagName  *TagName
		payload  *IntPayload
		expected *IntTag
	}

	cases := []Case{}

	for _, c := range intTagCases {
		cases = append(cases, Case{
			name:    c.name,
			tagName: &c.nbt.tagName,
			payload: c.nbt.payload,
			expected: &IntTag{
				tagName: &c.nbt.tagName,
				payload: c.nbt.payload,
			},
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewIntTag(tt.tagName, tt.payload)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestIntTag_TypeId(t *testing.T) {
	tag := new(IntTag)
	expected := IntType
	actual := tag.TypeId()
	assert.Equal(t, expected, actual)
}

func TestIntTag_encode(t *testing.T) {
	type Case struct {
		name        string
		tag         *IntTag
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range intTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tag:         NewIntTag(&c.nbt.tagName, c.nbt.payload),
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

func TestIntTag_decode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    *IntTag
		expectedErr error
	}

	cases := []Case{}

	for _, c := range intTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         slices.Concat(c.raw.tagType, c.raw.tagName, c.raw.payload),
			expected:    NewIntTag(&c.nbt.tagName, c.nbt.payload),
			expectedErr: nil,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer(tt.raw)

			tag := new(IntTag)
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

func TestNewIntPayload(t *testing.T) {
	type Case struct {
		name     string
		value    int32
		expected *IntPayload
	}

	cases := []Case{}

	for _, c := range intTagCases {
		cases = append(cases, Case{
			name:     c.name,
			value:    int32(*c.nbt.payload),
			expected: c.nbt.payload,
		})
	}

	for _, tt := range cases {
		actual := NewIntPayload(tt.value)
		assert.Equal(t, tt.expected, actual)
	}
}

func TestIntPayload_TypeId(t *testing.T) {
	payload := new(IntPayload)
	expected := IntType
	actual := payload.TypeId()
	assert.Equal(t, expected, actual)
}

func TestIntPayload_encode(t *testing.T) {
	type Case struct {
		name        string
		payload     *IntPayload
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range intTagCases {
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

func TestIntPayload_decode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    *IntPayload
		expectedErr error
	}

	cases := []Case{}

	for _, c := range intTagCases {
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

			payload := new(IntPayload)
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
