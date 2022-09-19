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

func TestCompoundTag_TypeId(t *testing.T) {
	cases := []struct {
		name     string
		tag      Tag
		expected TagType
	}{
		{
			name:     "positive case",
			tag:      NewCompoundTag(),
			expected: CompoundType,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.tag.TypeId()
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestCompoundTag_Encode(t *testing.T) {
	cases := []struct {
		name        string
		tag         Tag
		expected    []byte
		expectedErr error
	}{
		{
			name: "positive case",
			tag: &CompoundTag{
				TagName: TagName("Compound"),
				CompoundPayload: CompoundPayload{
					&ShortTag{TagName("Short"), ShortPayload(12345)},
					&ByteArrayTag{TagName("ByteArray"), ByteArrayPayload{0, 1}},
					&StringTag{TagName("String"), StringPayload("Hello")},
					&ListTag{TagName("List"), ListPayload{PayloadPointer(BytePayload(123))}},
					&CompoundTag{TagName("Compound"), CompoundPayload{&StringTag{TagName("String"), StringPayload("World")}, &EndTag{}}},
					&EndTag{},
				},
			},
			expected: []byte{
				// Name Length: 8
				0x00, 0x08,
				// Name: "Compound"
				0x43, 0x6F, 0x6D, 0x70, 0x6F, 0x75, 0x6E, 0x64,
				// Payload:
				//   - Type: Short(=2)
				//     Name: "Short"
				//     Payload: 12345s
				0x02,
				0x00, 0x05,
				0x53, 0x68, 0x6f, 0x72, 0x74,
				0x30, 0x39,
				//   - Type: ByteArray(=7)
				//     Name: "ByteArray"
				//     Payload: [B; 0b, 1b]
				0x07,
				0x00, 0x09,
				0x42, 0x79, 0x74, 0x65, 0x41, 0x72, 0x72, 0x61, 0x79,
				0x00, 0x00, 0x00, 0x02,
				0x00, 0x01,
				//   - Type: String(=8)
				//     Name: "String"
				//     Payload: "Hello"
				0x08,
				0x00, 0x06,
				0x53, 0x74, 0x72, 0x69, 0x6E, 0x67,
				0x00, 0x05,
				0x48, 0x65, 0x6C, 0x6C, 0x6F,
				//   - Type: List(=9)
				//     Name: "List"
				//     Payload: [123b]
				0x09,
				0x00, 0x04,
				0x4C, 0x69, 0x73, 0x74,
				0x01,
				0x00, 0x00, 0x00, 0x01,
				0x7B,
				//   - Type: Compound(=10)
				//     Name: "Compound"
				0x0A,
				0x00, 0x08,
				0x43, 0x6F, 0x6D, 0x70, 0x6F, 0x75, 0x6E, 0x64,
				//     Payload:
				//       - Type: String(=8)
				//         Name: "String"
				//         Payload: "World"
				0x08,
				0x00, 0x06,
				0x53, 0x74, 0x72, 0x69, 0x6E, 0x67,
				0x00, 0x05,
				0x57, 0x6F, 0x72, 0x6C, 0x64,
				//       - Type: End(=0)
				0x00,
				//   - Type: End(=0)
				0x00,
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

func TestCompoundPayload_TypeId(t *testing.T) {
	cases := []struct {
		name     string
		payload  Payload
		expected TagType
	}{
		{
			name:     "positive case",
			payload:  NewCompoundPayload(),
			expected: CompoundType,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.payload.TypeId()
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestCompoundPayload_Encode(t *testing.T) {
	cases := []struct {
		name        string
		payload     Payload
		expected    []byte
		expectedErr error
	}{
		{
			name: "positive case: has items",
			payload: &CompoundPayload{
				&ShortTag{TagName("Short"), ShortPayload(12345)},
				&ByteArrayTag{TagName("ByteArray"), ByteArrayPayload{0, 1}},
				&StringTag{TagName("String"), StringPayload(`Hello`)},
				&ListTag{TagName("List"), ListPayload{PayloadPointer(BytePayload(123))}},
				&CompoundTag{TagName("Compound"), CompoundPayload{&StringTag{TagName("String"), StringPayload("World")}, &EndTag{}}},
				&EndTag{},
			},
			expected: []byte{
				// Payload:
				//   - Type: Short(=2)
				//     Name: "Short"
				//     Payload: 12345s
				0x02,
				0x00, 0x05,
				0x53, 0x68, 0x6f, 0x72, 0x74,
				0x30, 0x39,
				//   - Type: ByteArray(=7)
				//     Name: "ByteArray"
				//     Payload: [B; 0b, 1b]
				0x07,
				0x00, 0x09,
				0x42, 0x79, 0x74, 0x65, 0x41, 0x72, 0x72, 0x61, 0x79,
				0x00, 0x00, 0x00, 0x02,
				0x00, 0x01,
				//   - Type: String(=8)
				//     Name: "String"
				//     Payload: "Hello"
				0x08,
				0x00, 0x06,
				0x53, 0x74, 0x72, 0x69, 0x6E, 0x67,
				0x00, 0x05,
				0x48, 0x65, 0x6C, 0x6C, 0x6F,
				//   - Type: List(=9)
				//     Name: "List"
				//     Payload: [123b]
				0x09,
				0x00, 0x04,
				0x4C, 0x69, 0x73, 0x74,
				0x01,
				0x00, 0x00, 0x00, 0x01,
				0x7B,
				//   - Type: Compound(=10)
				//     Name: "Compound"
				0x0A,
				0x00, 0x08,
				0x43, 0x6F, 0x6D, 0x70, 0x6F, 0x75, 0x6E, 0x64,
				//     Payload:
				//       - Type: String(=8)
				//         Name: "String"
				//         Payload: "World"
				0x08,
				0x00, 0x06,
				0x53, 0x74, 0x72, 0x69, 0x6E, 0x67,
				0x00, 0x05,
				0x57, 0x6F, 0x72, 0x6C, 0x64,
				//       - Type: End(=0)
				0x00,
				//   - Type: End(=0)
				0x00,
			},
			expectedErr: nil,
		},
		{
			name:    "positive case: empty",
			payload: &CompoundPayload{&EndTag{}},
			expected: []byte{
				// Payload:
				//   - Type: End(=0)
				0x00,
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
