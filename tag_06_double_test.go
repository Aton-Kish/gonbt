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

var doubleTagCases = []struct {
	name string
	tag  Tag
	raw  []byte
}{
	{
		name: "positive case",
		tag:  NewDoubleTag(NewTagName("Double"), NewDoublePayload(0.123456789)),
		raw: []byte{
			// Type: Double(=6)
			0x06,
			// Name Length: 6
			0x00, 0x06,
			// Name: "Double"
			0x44, 0x6F, 0x75, 0x62, 0x6C, 0x65,
			// Payload: 0.123456789d
			0x3F, 0xBF, 0x9A, 0xDD, 0x37, 0x39, 0x63, 0x5F,
		},
	},
}

func TestNewDoubleTag(t *testing.T) {
	cases := []struct {
		name     string
		tagName  *TagName
		payload  *DoublePayload
		expected Tag
	}{
		{
			name:    "positive case",
			tagName: NewTagName("Double"),
			payload: NewDoublePayload(0.123456789),
			expected: &DoubleTag{
				tagName: NewTagName("Double"),
				payload: NewDoublePayload(0.123456789),
			},
		},
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
		tag         Tag
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range doubleTagCases {
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

func TestDoubleTag_Decode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    Tag
		expectedErr error
	}

	cases := []Case{}

	for _, c := range doubleTagCases {
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

var doublePayloadCases = []struct {
	name    string
	payload Payload
	raw     []byte
}{
	{
		name:    "positive case",
		payload: NewDoublePayload(0.123456789),
		raw: []byte{
			// Payload: 0.123456789d
			0x3F, 0xBF, 0x9A, 0xDD, 0x37, 0x39, 0x63, 0x5F,
		},
	},
}

func TestNewDoublePayload(t *testing.T) {
	cases := []struct {
		name     string
		value    float64
		expected Payload
	}{
		{
			name:     "positive case",
			value:    0.123456789,
			expected: pointer.Pointer[DoublePayload](0.123456789),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewDoublePayload(tt.value)
			assert.Equal(t, tt.expected, actual)
		})
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
		payload     Payload
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range doublePayloadCases {
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

func TestDoublePayload_Decode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    Payload
		expectedErr error
	}

	cases := []Case{}

	for _, c := range doublePayloadCases {
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
