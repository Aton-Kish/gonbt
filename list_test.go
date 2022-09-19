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

func TestListTag_TypeId(t *testing.T) {
	cases := []struct {
		name     string
		tag      Tag
		expected TagType
	}{
		{
			name:     "positive case",
			tag:      NewListTag(),
			expected: ListType,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.tag.TypeId()
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestListTag_Encode(t *testing.T) {
	cases := []struct {
		name        string
		tag         Tag
		expected    []byte
		expectedErr error
	}{
		{
			name: "positive case",
			tag: &ListTag{
				TagName:     TagName("List"),
				ListPayload: ListPayload{PayloadPointer(ShortPayload(12345)), PayloadPointer(ShortPayload(6789))},
			},
			expected: []byte{
				// Name Length: 4
				0x00, 0x04,
				// Name: List
				0x4C, 0x69, 0x73, 0x74,
				// Payload Type: TagShort(=2)
				0x02,
				// Payload Length: 2
				0x00, 0x00, 0x00, 0x02,
				// Payload:
				//   - 12345s
				0x30, 0x39,
				//   - 6789s
				0x1A, 0x85,
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

func TestListPayload_TypeId(t *testing.T) {
	cases := []struct {
		name     string
		payload  Payload
		expected TagType
	}{
		{
			name:     "positive case",
			payload:  NewListPayload(),
			expected: ListType,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.payload.TypeId()
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestListPayload_Encode(t *testing.T) {
	cases := []struct {
		name        string
		payload     Payload
		expected    []byte
		expectedErr error
	}{
		{
			name:    "positive case: Short",
			payload: &ListPayload{PayloadPointer(ShortPayload(12345)), PayloadPointer(ShortPayload(6789))},
			expected: []byte{
				// Payload Type: TagShort(=2)
				0x02,
				// Payload Length: 2
				0x00, 0x00, 0x00, 0x02,
				// Payload:
				//   - 12345s
				0x30, 0x39,
				//   - 6789s
				0x1A, 0x85,
			},
			expectedErr: nil,
		},
		{
			name:    "positive case: ByteArray",
			payload: &ListPayload{&ByteArrayPayload{0, 1}, &ByteArrayPayload{2, 3}},
			expected: []byte{
				// Payload Type: TagByteArray(=7)
				0x07,
				// Payload Length: 2
				0x00, 0x00, 0x00, 0x02,
				// Payload:
				//   - [B; 0b, 1b]
				0x00, 0x00, 0x00, 0x02,
				0x00, 0x01,
				//   - [B; 2b, 3b]
				0x00, 0x00, 0x00, 0x02,
				0x02, 0x03,
			},
			expectedErr: nil,
		},
		{
			name:    "positive case: String",
			payload: &ListPayload{PayloadPointer(StringPayload("Hello")), PayloadPointer(StringPayload("World"))},
			expected: []byte{
				// Payload Type: TagString(=8)
				0x08,
				// Payload Length: 2
				0x00, 0x00, 0x00, 0x02,
				// Payload:
				//   - "Hello"
				0x00, 0x05,
				0x48, 0x65, 0x6C, 0x6C, 0x6F,
				//   - "World"
				0x00, 0x05,
				0x57, 0x6F, 0x72, 0x6C, 0x64,
			},
			expectedErr: nil,
		},
		{
			name: "positive case: List",
			payload: &ListPayload{
				&ListPayload{PayloadPointer(BytePayload(123))},
				&ListPayload{PayloadPointer(StringPayload("Test"))},
			},
			expected: []byte{
				// Payload Type: TagList(=9)
				0x09,
				// Payload Length: 2
				0x00, 0x00, 0x00, 0x02,
				// Payload:
				//   - [123b]
				0x01,
				0x00, 0x00, 0x00, 0x01,
				0x7B,
				//   - ["Test"]
				0x08,
				0x00, 0x00, 0x00, 0x01,
				0x00, 0x04,
				0x54, 0x65, 0x73, 0x74,
			},
			expectedErr: nil,
		},
		{
			name: "positive case: Compound",
			payload: &ListPayload{
				&CompoundPayload{&ByteTag{TagName("Byte"), BytePayload(123)}, &EndTag{}},
				&CompoundPayload{&StringTag{TagName("String"), StringPayload("Hello")}, &EndTag{}},
			},
			expected: []byte{
				// Payload Type: TagCompound(=10)
				0x0A,
				// Payload Length: 2
				0x00, 0x00, 0x00, 0x02,
				// Payload:
				//   - - Type: Byte(=1)
				//       Name: "Byte"
				//       Payload: 123b
				0x01,
				0x00, 0x04,
				0x42, 0x79, 0x74, 0x65,
				0x7B,
				//     - Type: End(=0)
				0x00,
				//   - - Type: String(=8)
				//       Name: "String"
				//       Payload: "Hello"
				0x08,
				0x00, 0x06,
				0x53, 0x74, 0x72, 0x69, 0x6E, 0x67,
				0x00, 0x05,
				0x48, 0x65, 0x6C, 0x6C, 0x6F,
				//     - Type: End(=0)
				0x00,
			},
			expectedErr: nil,
		},
		{
			name:    "positive case: empty",
			payload: &ListPayload{},
			expected: []byte{
				// Payload Type: TagEnd(=0)
				0x00,
				// Payload Length: 0
				0x00, 0x00, 0x00, 0x00,
				// Payload: []
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
