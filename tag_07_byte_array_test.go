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

var byteArrayTagCases = []tagTestCase[*ByteArrayPayload]{
	{
		name: "positive case: ByteArrayTag - has items",
		nbt: nbtTestCase[*ByteArrayPayload]{
			tagType: ByteArrayType,
			tagName: "ByteArray",
			payload: NewByteArrayPayload(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
		},
		snbt: snbtTestCase{
			tagName: "ByteArray",
			payload: "[B; 0b, 1b, 2b, 3b, 4b, 5b, 6b, 7b, 8b, 9b]",
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: ByteArray(=7)
				0x07,
			},
			tagName: []byte{
				// Name Length: 9
				0x00, 0x09,
				// Name: "ByteArray"
				0x42, 0x79, 0x74, 0x65, 0x41, 0x72, 0x72, 0x61, 0x79,
			},
			payload: []byte{
				// Payload Length: 10
				0x00, 0x00, 0x00, 0x0A,
				// Payload: [B; 0b, 1b, 2b, 3b, 4b, 5b, 6b, 7b, 8b, 9b]
				0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09,
			},
		},
	},
	{
		name: "positive case: ByteArrayTag - empty",
		nbt: nbtTestCase[*ByteArrayPayload]{
			tagType: ByteArrayType,
			tagName: "ByteArray",
			payload: NewByteArrayPayload(),
		},
		snbt: snbtTestCase{
			tagName: "ByteArray",
			payload: "[B; ]",
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: ByteArray(=7)
				0x07,
			},
			tagName: []byte{
				// Name Length: 9
				0x00, 0x09,
				// Name: "ByteArray"
				0x42, 0x79, 0x74, 0x65, 0x41, 0x72, 0x72, 0x61, 0x79,
			},
			payload: []byte{
				// Payload Length: 0
				0x00, 0x00, 0x00, 0x00,
				// Payload: [B; ]
			},
		},
	},
}

func TestNewByteArrayTag(t *testing.T) {
	type Case struct {
		name     string
		tagName  *TagName
		payload  *ByteArrayPayload
		expected *ByteArrayTag
	}

	cases := []Case{}

	for _, c := range byteArrayTagCases {
		cases = append(cases, Case{
			name:    c.name,
			tagName: &c.nbt.tagName,
			payload: c.nbt.payload,
			expected: &ByteArrayTag{
				tagName: &c.nbt.tagName,
				payload: c.nbt.payload,
			},
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewByteArrayTag(tt.tagName, tt.payload)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestByteArrayTag_TypeId(t *testing.T) {
	tag := new(ByteArrayTag)
	expected := ByteArrayType
	actual := tag.TypeId()
	assert.Equal(t, expected, actual)
}

func TestByteArrayTag_encode(t *testing.T) {
	type Case struct {
		name        string
		tag         *ByteArrayTag
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range byteArrayTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tag:         NewByteArrayTag(&c.nbt.tagName, c.nbt.payload),
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

func TestByteArrayTag_decode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    *ByteArrayTag
		expectedErr error
	}

	cases := []Case{}

	for _, c := range byteArrayTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         slices.Concat(c.raw.tagType, c.raw.tagName, c.raw.payload),
			expected:    NewByteArrayTag(&c.nbt.tagName, c.nbt.payload),
			expectedErr: nil,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer(tt.raw)

			tag := new(ByteArrayTag)
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

func TestByteArrayTag_stringify_default(t *testing.T) {
	type Case struct {
		name     string
		tag      *ByteArrayTag
		expected string
	}

	cases := []Case{}

	for _, c := range byteArrayTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tag:      NewByteArrayTag(&c.nbt.tagName, c.nbt.payload),
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

func TestNewByteArrayPayload(t *testing.T) {
	type Case struct {
		name     string
		values   []int8
		expected *ByteArrayPayload
	}

	cases := []Case{}

	for _, c := range byteArrayTagCases {
		cases = append(cases, Case{
			name:     c.name,
			values:   []int8(*c.nbt.payload),
			expected: c.nbt.payload,
		})
	}

	for _, tt := range cases {
		actual := NewByteArrayPayload(tt.values...)
		assert.Equal(t, tt.expected, actual)
	}
}

func TestByteArrayPayload_TypeId(t *testing.T) {
	payload := new(ByteArrayPayload)
	expected := ByteArrayType
	actual := payload.TypeId()
	assert.Equal(t, expected, actual)
}

func TestByteArrayPayload_encode(t *testing.T) {
	type Case struct {
		name        string
		payload     *ByteArrayPayload
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range byteArrayTagCases {
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

func TestByteArrayPayload_decode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    *ByteArrayPayload
		expectedErr error
	}

	cases := []Case{}

	for _, c := range byteArrayTagCases {
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

			payload := new(ByteArrayPayload)
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

func TestByteArrayPayload_stringify_default(t *testing.T) {
	type Case struct {
		name     string
		payload  *ByteArrayPayload
		expected string
	}

	cases := []Case{}

	for _, c := range byteArrayTagCases {
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
