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

var doubleTagCases = []tagTestCase[float64, *DoublePayload]{
	{
		name: "positive case: DoubleTag",
		data: 0.123456789,
		nbt: nbtTestCase[*DoublePayload]{
			tagType: DoubleType,
			tagName: "Double",
			payload: NewDoublePayload(0.123456789),
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: Double(=6)
				0x06,
			},
			tagName: []byte{
				// Name Length: 6
				0x00, 0x06,
				// Name: "Double"
				0x44, 0x6F, 0x75, 0x62, 0x6C, 0x65,
			},
			payload: []byte{
				// Payload: 0.123456789d
				0x3F, 0xBF, 0x9A, 0xDD, 0x37, 0x39, 0x63, 0x5F,
			},
		},
	},
}

func TestNewDoubleTag(t *testing.T) {
	type Case struct {
		name     string
		tagName  *TagName
		payload  *DoublePayload
		expected *DoubleTag
	}

	cases := []Case{}

	for _, c := range doubleTagCases {
		cases = append(cases, Case{
			name:    c.name,
			tagName: &c.nbt.tagName,
			payload: c.nbt.payload,
			expected: &DoubleTag{
				tagName: &c.nbt.tagName,
				payload: c.nbt.payload,
			},
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewDoubleTag(tt.tagName, tt.payload)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestDoubleTag_TypeId(t *testing.T) {
	tag := new(DoubleTag)
	expected := DoubleType
	actual := tag.TypeId()
	assert.Equal(t, expected, actual)
}

func TestDoubleTag_Encode(t *testing.T) {
	type Case struct {
		name        string
		tag         *DoubleTag
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range doubleTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tag:         NewDoubleTag(&c.nbt.tagName, c.nbt.payload),
			expected:    append(append(c.raw.tagType, c.raw.tagName...), c.raw.payload...),
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

func TestDoubleTag_Decode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    *DoubleTag
		expectedErr error
	}

	cases := []Case{}

	for _, c := range doubleTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         append(append(c.raw.tagType, c.raw.tagName...), c.raw.payload...),
			expected:    NewDoubleTag(&c.nbt.tagName, c.nbt.payload),
			expectedErr: nil,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer(tt.raw)

			tag := new(DoubleTag)
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

func TestNewDoublePayload(t *testing.T) {
	type Case struct {
		name     string
		value    float64
		expected *DoublePayload
	}

	cases := []Case{}

	for _, c := range doubleTagCases {
		cases = append(cases, Case{
			name:     c.name,
			value:    c.data,
			expected: c.nbt.payload,
		})
	}

	for _, tt := range cases {
		actual := NewDoublePayload(tt.value)
		assert.Equal(t, tt.expected, actual)
	}
}

func TestDoublePayload_TypeId(t *testing.T) {
	payload := new(DoublePayload)
	expected := DoubleType
	actual := payload.TypeId()
	assert.Equal(t, expected, actual)
}

func TestDoublePayload_Encode(t *testing.T) {
	type Case struct {
		name        string
		payload     *DoublePayload
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range doubleTagCases {
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

func TestDoublePayload_Decode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    *DoublePayload
		expectedErr error
	}

	cases := []Case{}

	for _, c := range doubleTagCases {
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

			payload := new(DoublePayload)
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
