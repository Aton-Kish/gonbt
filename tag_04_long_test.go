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

	"github.com/Aton-Kish/gonbt/slices"
	"github.com/stretchr/testify/assert"
)

var longTagCases = []tagTestCase[*LongPayload]{
	{
		name: "positive case: LongTag",
		nbt: nbtTestCase[*LongPayload]{
			tagType: LongType,
			tagName: "Long",
			payload: NewLongPayload(123456789123456789),
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: Long(=4)
				0x04,
			},
			tagName: []byte{
				// Name Length: 4
				0x00, 0x04,
				// Name: "Long"
				0x4C, 0x6F, 0x6E, 0x67,
			},
			payload: []byte{
				// Payload: 123456789123456789L
				0x01, 0xB6, 0x9B, 0x4B, 0xAC, 0xD0, 0x5F, 0x15,
			},
		},
	},
}

func TestNewLongTag(t *testing.T) {
	type Case struct {
		name     string
		tagName  *TagName
		payload  *LongPayload
		expected *LongTag
	}

	cases := []Case{}

	for _, c := range longTagCases {
		cases = append(cases, Case{
			name:    c.name,
			tagName: &c.nbt.tagName,
			payload: c.nbt.payload,
			expected: &LongTag{
				tagName: &c.nbt.tagName,
				payload: c.nbt.payload,
			},
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewLongTag(tt.tagName, tt.payload)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestLongTag_TypeId(t *testing.T) {
	tag := new(LongTag)
	expected := LongType
	actual := tag.TypeId()
	assert.Equal(t, expected, actual)
}

func TestLongTag_encode(t *testing.T) {
	type Case struct {
		name        string
		tag         *LongTag
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range longTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tag:         NewLongTag(&c.nbt.tagName, c.nbt.payload),
			expected:    slices.Concat(c.raw.tagType, c.raw.tagName, c.raw.payload),
			expectedErr: nil,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			err := tt.tag.encode(buf)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, buf.Bytes())
			} else {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}

func TestLongTag_decode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    *LongTag
		expectedErr error
	}

	cases := []Case{}

	for _, c := range longTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         slices.Concat(c.raw.tagType, c.raw.tagName, c.raw.payload),
			expected:    NewLongTag(&c.nbt.tagName, c.nbt.payload),
			expectedErr: nil,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer(tt.raw)

			tag := new(LongTag)
			err := tag.decode(buf)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, tag)
			} else {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}

func TestNewLongPayload(t *testing.T) {
	type Case struct {
		name     string
		value    int64
		expected *LongPayload
	}

	cases := []Case{}

	for _, c := range longTagCases {
		cases = append(cases, Case{
			name:     c.name,
			value:    int64(*c.nbt.payload),
			expected: c.nbt.payload,
		})
	}

	for _, tt := range cases {
		actual := NewLongPayload(tt.value)
		assert.Equal(t, tt.expected, actual)
	}
}

func TestLongPayload_TypeId(t *testing.T) {
	payload := new(LongPayload)
	expected := LongType
	actual := payload.TypeId()
	assert.Equal(t, expected, actual)
}

func TestLongPayload_encode(t *testing.T) {
	type Case struct {
		name        string
		payload     *LongPayload
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range longTagCases {
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
			err := tt.payload.encode(buf)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, buf.Bytes())
			} else {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}

func TestLongPayload_decode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    *LongPayload
		expectedErr error
	}

	cases := []Case{}

	for _, c := range longTagCases {
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

			payload := new(LongPayload)
			err := payload.decode(buf)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, payload)
			} else {
				assert.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}
