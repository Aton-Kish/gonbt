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

var byteArrayTagCases = []struct {
	name string
	tag  Tag
	raw  []byte
}{
	{
		name: "positive case",
		tag: &ByteArrayTag{
			TagName:          TagName("ByteArray"),
			ByteArrayPayload: ByteArrayPayload{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		raw: []byte{
			// Name Length: 9
			0x00, 0x09,
			// Name: "ByteArray"
			0x42, 0x79, 0x74, 0x65, 0x41, 0x72, 0x72, 0x61, 0x79,
			// Payload Length: 10
			0x00, 0x00, 0x00, 0x0A,
			// Payload: [B; 0b, 1b, 2b, 3b, 4b, 5b, 6b, 7b, 8b, 9b]
			0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09,
		},
	},
}

func TestByteArrayTag_TypeId(t *testing.T) {
	tag := NewByteArrayTag()
	expected := ByteArrayType
	actual := tag.TypeId()
	assert.Equal(t, expected, actual)
}

func TestByteArrayTag_Encode(t *testing.T) {
	type Case struct {
		name        string
		tag         Tag
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range byteArrayTagCases {
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

				raw := buf.Bytes()
				assert.Equal(t, byte(tt.tag.TypeId()), raw[0])
				assert.Equal(t, tt.expected, raw[1:])
			} else {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}

var byteArrayPayloadCases = []struct {
	name    string
	payload Payload
	raw     []byte
}{
	{
		name:    "positive case: has items",
		payload: &ByteArrayPayload{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		raw: []byte{
			// Payload Length: 10
			0x00, 0x00, 0x00, 0x0A,
			// Payload: [B; 0b, 1b, 2b, 3b, 4b, 5b, 6b, 7b, 8b, 9b]
			0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09,
		},
	},
	{
		name:    "positive case: empty",
		payload: &ByteArrayPayload{},
		raw: []byte{
			// Payload Length: 0
			0x00, 0x00, 0x00, 0x00,
			// Payload: [B; ]
		},
	},
}

func TestByteArrayPayload_TypeId(t *testing.T) {
	payload := NewByteArrayPayload()
	expected := ByteArrayType
	actual := payload.TypeId()
	assert.Equal(t, expected, actual)
}

func TestByteArrayPayload_Encode(t *testing.T) {
	type Case struct {
		name        string
		payload     Payload
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range byteArrayPayloadCases {
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
