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

var intArrayTagCases = []tagTestCase[*IntArrayPayload]{
	{
		name: "positive case: IntArrayTag - has items",
		nbt: nbtTestCase[*IntArrayPayload]{
			tagType: IntArrayType,
			tagName: "IntArray",
			payload: NewIntArrayPayload(0, 1, 2, 3),
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: IntArray(=11)
				0x0B,
			},
			tagName: []byte{
				// Name Length: 8
				0x00, 0x08,
				// Name: "IntArray"
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
		name: "positive case: IntArrayTag - empty",
		nbt: nbtTestCase[*IntArrayPayload]{
			tagType: IntArrayType,
			tagName: "IntArray",
			payload: NewIntArrayPayload(),
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: IntArray(=11)
				0x0B,
			},
			tagName: []byte{
				// Name Length: 8
				0x00, 0x08,
				// Name: "IntArray"
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

func TestIntArrayTag_Encode(t *testing.T) {
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

func TestIntArrayTag_Decode(t *testing.T) {
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

func TestIntArrayPayload_Encode(t *testing.T) {
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

func TestIntArrayPayload_Decode(t *testing.T) {
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
