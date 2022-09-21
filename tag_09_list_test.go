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

var listTagCases = []struct {
	name string
	tag  Tag
	raw  []byte
}{
	{
		name: "positive case",
		tag: &ListTag{
			tagName: TagName("List"),
			payload: ListPayload{PrimitivePayloadPointer[ShortPayload](12345), PrimitivePayloadPointer[ShortPayload](6789)},
		},
		raw: []byte{
			// Name Length: 4
			0x00, 0x04,
			// Name: "List"
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
	},
}

func TestListTag_TypeId(t *testing.T) {
	tag := NewListTag()
	expected := ListType
	actual := tag.TypeId()
	assert.Equal(t, expected, actual)
}

func TestListTag_Encode(t *testing.T) {
	type Case struct {
		name        string
		tag         Tag
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range listTagCases {
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

var listPayloadCases = []struct {
	name    string
	payload Payload
	raw     []byte
}{
	{
		name:    "positive case: Short",
		payload: &ListPayload{PrimitivePayloadPointer[ShortPayload](12345), PrimitivePayloadPointer[ShortPayload](6789)},
		raw: []byte{
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
	},
	{
		name:    "positive case: ByteArray",
		payload: &ListPayload{&ByteArrayPayload{0, 1}, &ByteArrayPayload{2, 3}},
		raw: []byte{
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
	},
	{
		name:    "positive case: String",
		payload: &ListPayload{PrimitivePayloadPointer[StringPayload]("Hello"), PrimitivePayloadPointer[StringPayload]("World")},
		raw: []byte{
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
	},
	{
		name: "positive case: List",
		payload: &ListPayload{
			&ListPayload{PrimitivePayloadPointer[BytePayload](123)},
			&ListPayload{PrimitivePayloadPointer[StringPayload]("Test")},
		},
		raw: []byte{
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
	},
	{
		name: "positive case: Compound",
		payload: &ListPayload{
			&CompoundPayload{&ByteTag{TagName("Byte"), BytePayload(123)}, &EndTag{}},
			&CompoundPayload{&StringTag{TagName("String"), StringPayload("Hello")}, &EndTag{}},
		},
		raw: []byte{
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
	},
	{
		name:    "positive case: empty",
		payload: &ListPayload{},
		raw: []byte{
			// Payload Type: TagEnd(=0)
			0x00,
			// Payload Length: 0
			0x00, 0x00, 0x00, 0x00,
			// Payload: []
		},
	},
}

func TestListPayload_TypeId(t *testing.T) {
	payload := NewListPayload()
	expected := ListType
	actual := payload.TypeId()
	assert.Equal(t, expected, actual)
}

func TestListPayload_Encode(t *testing.T) {
	type Case struct {
		name        string
		payload     Payload
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range listPayloadCases {
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

func TestListPayload_Decode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    Payload
		expectedErr error
	}

	cases := []Case{}

	// for _, c := range listPayloadCases {
	// 	cases = append(cases, Case{
	// 		name:        c.name,
	// 		raw:         c.raw,
	// 		expected:    c.payload,
	// 		expectedErr: nil,
	// 	})
	// }

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer(tt.raw)

			p := new(ListPayload)
			err := p.Decode(buf)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, p)
			} else {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}
