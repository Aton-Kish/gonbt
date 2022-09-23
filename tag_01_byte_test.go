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

var byteTagCases = []struct {
	name string
	tag  Tag
	raw  []byte
}{
	{
		name: "positive case",
		tag:  NewByteTag(NewTagName("Byte"), NewBytePayload(123)),
		raw: []byte{
			// Type: Byte(=1)
			0x01,
			// Name Length: 4
			0x00, 0x04,
			// Name: "Byte"
			0x42, 0x79, 0x74, 0x65,
			// Payload: 123b
			0x7B,
		},
	},
}

func TestNewByteTag(t *testing.T) {
	cases := []struct {
		name     string
		tagName  *TagName
		payload  *BytePayload
		expected Tag
	}{
		{
			name:    "positive case",
			tagName: NewTagName("Byte"),
			payload: NewBytePayload(123),
			expected: &ByteTag{
				tagName: NewTagName("Byte"),
				payload: NewBytePayload(123),
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewByteTag(tt.tagName, tt.payload)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestByteTag_TypeId(t *testing.T) {
	tag := new(ByteTag)
	expected := ByteType
	actual := tag.TypeId()
	assert.Equal(t, expected, actual)
}

func TestByteTag_Encode(t *testing.T) {
	type Case struct {
		name        string
		tag         Tag
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range byteTagCases {
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

func TestByteTag_Decode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    Tag
		expectedErr error
	}

	cases := []Case{}

	for _, c := range byteTagCases {
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

			tag := new(ByteTag)
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

var bytePayloadCases = []struct {
	name    string
	payload Payload
	raw     []byte
}{
	{
		name:    "positive case",
		payload: NewBytePayload(123),
		raw: []byte{
			// Payload: 123b
			0x7B,
		},
	},
}

func TestNewBytePayload(t *testing.T) {
	cases := []struct {
		name     string
		value    int8
		expected Payload
	}{
		{
			name:     "positive case",
			value:    123,
			expected: pointer.Pointer(BytePayload(123)),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewBytePayload(tt.value)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestBytePayload_TypeId(t *testing.T) {
	payload := new(BytePayload)
	expected := ByteType
	actual := payload.TypeId()
	assert.Equal(t, expected, actual)
}

func TestBytePayload_Encode(t *testing.T) {
	type Case struct {
		name        string
		payload     Payload
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range bytePayloadCases {
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

func TestBytePayload_Decode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    Payload
		expectedErr error
	}

	cases := []Case{}

	for _, c := range bytePayloadCases {
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

			payload := new(BytePayload)
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
