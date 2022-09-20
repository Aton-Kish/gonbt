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

var intTagCases = []struct {
	name string
	tag  Tag
	raw  []byte
}{
	{
		name: "positive case",
		tag: &IntTag{
			TagName:    TagName("Int"),
			IntPayload: IntPayload(123456789),
		},
		raw: []byte{
			// Name Length: 3
			0x00, 0x03,
			// Name: "Int"
			0x49, 0x6E, 0x74,
			// Payload: 123456789
			0x07, 0x5B, 0xCD, 0x15,
		},
	},
}

func TestIntTag_TypeId(t *testing.T) {
	tag := NewIntTag()
	expected := IntType
	actual := tag.TypeId()
	assert.Equal(t, expected, actual)
}

func TestIntTag_Encode(t *testing.T) {
	type Case struct {
		name        string
		tag         Tag
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range intTagCases {
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

var intPayloadCases = []struct {
	name    string
	payload Payload
	raw     []byte
}{
	{
		name:    "positive case",
		payload: PayloadPointer(IntPayload(123456789)),
		raw: []byte{
			// Payload: 123456789
			0x07, 0x5B, 0xCD, 0x15,
		},
	},
}

func TestIntPayload_TypeId(t *testing.T) {
	payload := NewIntPayload()
	expected := IntType
	actual := payload.TypeId()
	assert.Equal(t, expected, actual)
}

func TestIntPayload_Encode(t *testing.T) {
	type Case struct {
		name        string
		payload     Payload
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range intPayloadCases {
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
