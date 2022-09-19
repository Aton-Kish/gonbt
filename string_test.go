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

func TestStringTag_TypeId(t *testing.T) {
	cases := []struct {
		name     string
		tag      Tag
		expected TagType
	}{
		{
			name:     "positive case",
			tag:      NewStringTag(),
			expected: StringType,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.tag.TypeId()
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestStringTag_Encode(t *testing.T) {
	cases := []struct {
		name        string
		tag         Tag
		expected    []byte
		expectedErr error
	}{
		{
			name: "positive case",
			tag: &StringTag{
				TagName:       TagName("String"),
				StringPayload: StringPayload("Hello World"),
			},
			expected: []byte{
				// Name Length: 6
				0x00, 0x06,
				// Name: "String"
				0x53, 0x74, 0x72, 0x69, 0x6E, 0x67,
				// Payload Length: 11
				0x00, 0x0B,
				// Payload: "Hello World"
				0x48, 0x65, 0x6C, 0x6C, 0x6F, 0x20, 0x57, 0x6F, 0x72, 0x6C, 0x64,
			},
			expectedErr: nil,
		},
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

func TestStringPayload_TypeId(t *testing.T) {
	cases := []struct {
		name     string
		payload  Payload
		expected TagType
	}{
		{
			name:     "positive case",
			payload:  NewStringPayload(),
			expected: StringType,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.payload.TypeId()
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestStringPayload_Encode(t *testing.T) {
	cases := []struct {
		name        string
		payload     Payload
		expected    []byte
		expectedErr error
	}{
		{
			name:    "positive case: \"Test\"",
			payload: PayloadPointer(StringPayload("Test")),
			expected: []byte{
				// Payload Length: 4
				0x00, 0x04,
				// Payload: "Test"
				0x54, 0x65, 0x73, 0x74,
			},
			expectedErr: nil,
		},
		{
			name:    "positive case: \"minecraft:the_end\"",
			payload: PayloadPointer(StringPayload("minecraft:the_end")),
			expected: []byte{
				// Payload Length: 17
				0x00, 0x11,
				// Payload: "minecraft:the_end"
				0x6D, 0x69, 0x6E, 0x65, 0x63, 0x72, 0x61, 0x66, 0x74, 0x3A,
				0x74, 0x68, 0x65, 0x5F, 0x65, 0x6E, 0x64,
			},
			expectedErr: nil,
		},
		{
			name:    "positive case: \"\"",
			payload: PayloadPointer(StringPayload("")),
			expected: []byte{
				// Payload Length: 0
				0x00, 0x00,
				// Payload: ""
			},
			expectedErr: nil,
		},
		{
			name:    "positive case: \"マインクラフト\"",
			payload: PayloadPointer(StringPayload("マインクラフト")),
			expected: []byte{
				// Payload Length: 21
				0x00, 0x15,
				// Payload: "マインクラフト"
				0xE3, 0x83, 0x9E, 0xE3, 0x82, 0xA4, 0xE3, 0x83, 0xB3,
				0xE3, 0x82, 0xAF, 0xE3, 0x83, 0xA9, 0xE3, 0x83, 0x95,
				0xE3, 0x83, 0x88,
			},
			expectedErr: nil,
		},
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
