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

var intTagCases = []struct {
	name string
	tag  Tag
	raw  []byte
}{
	{
		name: "positive case",
		tag:  NewIntTag(NewTagName("Int"), NewIntPayload(123456789)),
		raw: []byte{
			// Type: Int(=3)
			0x03,
			// Name Length: 3
			0x00, 0x03,
			// Name: "Int"
			0x49, 0x6E, 0x74,
			// Payload: 123456789
			0x07, 0x5B, 0xCD, 0x15,
		},
	},
}

func TestNewIntTag(t *testing.T) {
	cases := []struct {
		name     string
		tagName  *TagName
		payload  *IntPayload
		expected Tag
	}{
		{
			name:    "positive case",
			tagName: NewTagName("Int"),
			payload: NewIntPayload(123456789),
			expected: &IntTag{
				tagName: NewTagName("Int"),
				payload: NewIntPayload(123456789),
			},
		},
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

func TestIntTag_Encode(t *testing.T) {
	type Case struct {
		name        string
		tag         Tag
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range intTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tag:         c.tag,
			expected:    c.raw,
			expectedErr: nil,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			err := tt.tag.Encode(buf)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, buf.Bytes())
			} else {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}

func TestIntTag_Decode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    Tag
		expectedErr error
	}

	cases := []Case{}

	for _, c := range intTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         c.raw,
			expected:    c.tag,
			expectedErr: nil,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer(tt.raw)

			tag := new(IntTag)
			err := tag.Decode(buf)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, tag)
			} else {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}

var intPayloadCases = []struct {
	name    string
	payload Payload
	raw     []byte
}{
	{
		name:    "positive case",
		payload: NewIntPayload(123456789),
		raw: []byte{
			// Payload: 123456789
			0x07, 0x5B, 0xCD, 0x15,
		},
	},
}

func TestNewIntPayload(t *testing.T) {
	cases := []struct {
		name     string
		value    int32
		expected Payload
	}{
		{
			name:     "positive case",
			value:    123456789,
			expected: pointer.Pointer[IntPayload](123456789),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewIntPayload(tt.value)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestIntPayload_TypeId(t *testing.T) {
	payload := new(IntPayload)
	expected := IntType
	actual := payload.TypeId()
	assert.Equal(t, expected, actual)
}

func TestIntPayload_Encode(t *testing.T) {
	type Case struct {
		name        string
		payload     Payload
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range intPayloadCases {
		cases = append(cases, Case{
			name:        c.name,
			payload:     c.payload,
			expected:    c.raw,
			expectedErr: nil,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			err := tt.payload.Encode(buf)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, buf.Bytes())
			} else {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}

func TestIntPayload_Decode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    Payload
		expectedErr error
	}

	cases := []Case{}

	for _, c := range intPayloadCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         c.raw,
			expected:    c.payload,
			expectedErr: nil,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer(tt.raw)

			payload := new(IntPayload)
			err := payload.Decode(buf)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, payload)
			} else {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}
