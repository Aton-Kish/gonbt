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

// Code generated by testgen.go; DO NOT EDIT.

package nbt

import (
	"bytes"
	"testing"

	"github.com/Aton-Kish/gonbt/snbt"
	"github.com/stretchr/testify/assert"
)

func TestNewDoublePayload(t *testing.T) {
	type Case struct {
		name     string
		value    float64
		expected *DoublePayload
	}

	cases := []Case{}

	for _, c := range doubleTagCases {
		cases = append(cases, Case{
			name:     c.name,
			value:    float64(*c.nbt.payload),
			expected: c.nbt.payload,
		})
	}

	for _, tt := range cases {
		actual := NewDoublePayload(tt.value)
		assert.Equal(t, tt.expected, actual)
	}
}

func TestDoublePayload_TypeId(t *testing.T) {
	payload := new(DoublePayload)
	expected := TagTypeDouble
	actual := payload.TypeId()
	assert.Equal(t, expected, actual)
}

func TestDoublePayload_encode(t *testing.T) {
	type Case struct {
		name        string
		payload     *DoublePayload
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range doubleTagCases {
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
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr, err)
			}
		})
	}
}

func TestDoublePayload_decode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    *DoublePayload
		expectedErr error
	}

	cases := []Case{}

	for _, c := range doubleTagCases {
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

			payload := new(DoublePayload)
			err := payload.decode(buf)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, payload)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr, err)
			}
		})
	}
}

func TestDoublePayload_stringify_default(t *testing.T) {
	type Case struct {
		name     string
		payload  *DoublePayload
		expected string
	}

	cases := []Case{}

	for _, c := range doubleTagCases {
		cases = append(cases, Case{
			name:     c.name,
			payload:  c.nbt.payload,
			expected: c.snbt.payload.typeDefault,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.payload.stringify(" ", "", 0)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestDoublePayload_stringify_compact(t *testing.T) {
	type Case struct {
		name     string
		payload  *DoublePayload
		expected string
	}

	cases := []Case{}

	for _, c := range doubleTagCases {
		cases = append(cases, Case{
			name:     c.name,
			payload:  c.nbt.payload,
			expected: c.snbt.payload.typeCompact,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.payload.stringify("", "", 0)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestDoublePayload_stringify_pretty(t *testing.T) {
	type Case struct {
		name     string
		payload  *DoublePayload
		expected string
	}

	cases := []Case{}

	for _, c := range doubleTagCases {
		cases = append(cases, Case{
			name:     c.name,
			payload:  c.nbt.payload,
			expected: c.snbt.payload.typePretty,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.payload.stringify(" ", "  ", 0)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestDoublePayload_parse(t *testing.T) {
	type Case struct {
		name        string
		snbt        string
		expected    *DoublePayload
		expectedErr error
	}

	cases := []Case{}

	for _, c := range doubleTagCases {
		cases = append(cases, Case{
			name:        c.name,
			snbt:        c.snbt.payload.typeDefault,
			expected:    c.nbt.payload,
			expectedErr: nil,
		})
		cases = append(cases, Case{
			name:        c.name,
			snbt:        c.snbt.payload.typeCompact,
			expected:    c.nbt.payload,
			expectedErr: nil,
		})
		cases = append(cases, Case{
			name:        c.name,
			snbt:        c.snbt.payload.typePretty,
			expected:    c.nbt.payload,
			expectedErr: nil,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			parser := snbt.NewParser(tt.snbt)
			err := parser.Compact()
			assert.NoError(t, err)

			payload := new(DoublePayload)
			err = payload.parse(parser)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, payload)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr, err)
			}
		})
	}
}

func TestDoublePayload_json_default(t *testing.T) {
	type Case struct {
		name     string
		payload  *DoublePayload
		expected string
	}

	cases := []Case{}

	for _, c := range doubleTagCases {
		cases = append(cases, Case{
			name:     c.name,
			payload:  c.nbt.payload,
			expected: c.json.payload.typeDefault,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.payload.json(" ", "", 0)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestDoublePayload_json_compact(t *testing.T) {
	type Case struct {
		name     string
		payload  *DoublePayload
		expected string
	}

	cases := []Case{}

	for _, c := range doubleTagCases {
		cases = append(cases, Case{
			name:     c.name,
			payload:  c.nbt.payload,
			expected: c.json.payload.typeCompact,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.payload.json("", "", 0)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestDoublePayload_json_pretty(t *testing.T) {
	type Case struct {
		name     string
		payload  *DoublePayload
		expected string
	}

	cases := []Case{}

	for _, c := range doubleTagCases {
		cases = append(cases, Case{
			name:     c.name,
			payload:  c.nbt.payload,
			expected: c.json.payload.typePretty,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.payload.json(" ", "  ", 0)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
