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

func TestIntArrayTag_TypeId(t *testing.T) {
	cases := []struct {
		name     string
		tag      Tag
		expected TagType
	}{
		{
			name:     "positive case",
			tag:      NewIntArrayTag(),
			expected: IntArrayType,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.tag.TypeId()
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestIntArrayTag_Encode(t *testing.T) {
	cases := []struct {
		name        string
		tag         Tag
		expected    []byte
		expectedErr error
	}{
		{
			name: "positive case",
			tag: &IntArrayTag{
				TagName:         TagName("IntArray"),
				IntArrayPayload: IntArrayPayload{0, 1, 2, 3},
			},
			expected: []byte{
				// Name Length: 8
				0x00, 0x08,
				// Name: "IntArray"
				0x49, 0x6E, 0x74, 0x41, 0x72, 0x72, 0x61, 0x79,
				// Payload Length: 4
				0x00, 0x00, 0x00, 0x04,
				// Payload: [I; 0, 1, 2, 3]
				0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x01,
				0x00, 0x00, 0x00, 0x02,
				0x00, 0x00, 0x00, 0x03,
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

func TestIntArrayPayload_TypeId(t *testing.T) {
	cases := []struct {
		name     string
		payload  Payload
		expected TagType
	}{
		{
			name:     "positive case",
			payload:  NewIntArrayPayload(),
			expected: IntArrayType,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.payload.TypeId()
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestIntArrayPayload_Encode(t *testing.T) {
	cases := []struct {
		name        string
		payload     Payload
		expected    []byte
		expectedErr error
	}{
		{
			name:    "positive case: has items",
			payload: &IntArrayPayload{0, 1, 2, 3},
			expected: []byte{
				// Payload Length: 4
				0x00, 0x00, 0x00, 0x04,
				// Payload: [I; 0, 1, 2, 3]
				0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x01,
				0x00, 0x00, 0x00, 0x02,
				0x00, 0x00, 0x00, 0x03,
			},
			expectedErr: nil,
		},
		{
			name:    "positive case: empty",
			payload: &IntArrayPayload{},
			expected: []byte{
				// Payload Length: 0
				0x00, 0x00, 0x00, 0x00,
				// Payload: [I; ]
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
