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
	"fmt"
	"testing"

	"github.com/Aton-Kish/gonbt/slices"
	"github.com/Aton-Kish/gonbt/snbt"
	"github.com/stretchr/testify/assert"
)

var floatTagCases = []tagTestCase[*FloatPayload]{
	{
		name: `positive case: FloatTag - %f`,
		nbt: nbtTestCase[*FloatPayload]{
			tagType: FloatType,
			tagName: `Float`,
			payload: NewFloatPayload(0.12345678),
		},
		snbt: snbtTestCase{
			tagName: `Float`,
			payload: stringifyType{
				typeDefault: `0.12345678f`,
				typeCompact: `0.12345678f`,
				typePretty:  `0.12345678f`,
			},
		},
		json: jsonTestCase{
			tagName: `"Float"`,
			payload: stringifyType{
				typeDefault: `0.12345678`,
				typeCompact: `0.12345678`,
				typePretty:  `0.12345678`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: Float(=5)
				0x05,
			},
			tagName: []byte{
				// Name Length: 5
				0x00, 0x05,
				// Name: Float
				0x46, 0x6C, 0x6F, 0x61, 0x74,
			},
			payload: []byte{
				// Payload: 0.12345678f
				0x3D, 0xFC, 0xD6, 0xE9,
			},
		},
	},
	{
		name: `positive case: FloatTag - +%e`,
		nbt: nbtTestCase[*FloatPayload]{
			tagType: FloatType,
			tagName: `Float`,
			payload: NewFloatPayload(1234567.8),
		},
		snbt: snbtTestCase{
			tagName: `Float`,
			payload: stringifyType{
				typeDefault: `1.2345678e+06f`,
				typeCompact: `1.2345678e+06f`,
				typePretty:  `1.2345678e+06f`,
			},
		},
		json: jsonTestCase{
			tagName: `"Float"`,
			payload: stringifyType{
				typeDefault: `1.2345678e+06`,
				typeCompact: `1.2345678e+06`,
				typePretty:  `1.2345678e+06`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: Float(=5)
				0x05,
			},
			tagName: []byte{
				// Name Length: 5
				0x00, 0x05,
				// Name: Float
				0x46, 0x6C, 0x6F, 0x61, 0x74,
			},
			payload: []byte{
				// Payload: 1.2345678e+06f
				0x49, 0x96, 0xB4, 0x3E,
			},
		},
	},
	{
		name: `positive case: FloatTag - -%e`,
		nbt: nbtTestCase[*FloatPayload]{
			tagType: FloatType,
			tagName: `Float`,
			payload: NewFloatPayload(0.000012345678),
		},
		snbt: snbtTestCase{
			tagName: `Float`,
			payload: stringifyType{
				typeDefault: `1.2345678e-05f`,
				typeCompact: `1.2345678e-05f`,
				typePretty:  `1.2345678e-05f`,
			},
		},
		json: jsonTestCase{
			tagName: `"Float"`,
			payload: stringifyType{
				typeDefault: `1.2345678e-05`,
				typeCompact: `1.2345678e-05`,
				typePretty:  `1.2345678e-05`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: Float(=5)
				0x05,
			},
			tagName: []byte{
				// Name Length: 5
				0x00, 0x05,
				// Name: Float
				0x46, 0x6C, 0x6F, 0x61, 0x74,
			},
			payload: []byte{
				// Payload: 1.2345678e-05f
				0x37, 0x4F, 0x20, 0x49,
			},
		},
	},
}

func TestNewFloatTag(t *testing.T) {
	type Case struct {
		name     string
		tagName  *TagName
		payload  *FloatPayload
		expected *FloatTag
	}

	cases := []Case{}

	for _, c := range floatTagCases {
		cases = append(cases, Case{
			name:    c.name,
			tagName: &c.nbt.tagName,
			payload: c.nbt.payload,
			expected: &FloatTag{
				tagName: &c.nbt.tagName,
				payload: c.nbt.payload,
			},
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewFloatTag(tt.tagName, tt.payload)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestFloatTag_TypeId(t *testing.T) {
	tag := new(FloatTag)
	expected := FloatType
	actual := tag.TypeId()
	assert.Equal(t, expected, actual)
}

func TestFloatTag_encode(t *testing.T) {
	type Case struct {
		name        string
		tag         *FloatTag
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range floatTagCases {
		cases = append(cases, Case{
			name:        c.name,
			tag:         NewFloatTag(&c.nbt.tagName, c.nbt.payload),
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
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr, err)
			}
		})
	}
}

func TestFloatTag_decode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    *FloatTag
		expectedErr error
	}

	cases := []Case{}

	for _, c := range floatTagCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         slices.Concat(c.raw.tagType, c.raw.tagName, c.raw.payload),
			expected:    NewFloatTag(&c.nbt.tagName, c.nbt.payload),
			expectedErr: nil,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer(tt.raw)

			tag := new(FloatTag)
			err := tag.decode(buf)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, tag)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr, err)
			}
		})
	}
}

func TestFloatTag_stringify_default(t *testing.T) {
	type Case struct {
		name     string
		tag      *FloatTag
		expected string
	}

	cases := []Case{}

	for _, c := range floatTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tag:      NewFloatTag(&c.nbt.tagName, c.nbt.payload),
			expected: fmt.Sprintf("%s: %s", c.snbt.tagName, c.snbt.payload.typeDefault),
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.tag.stringify(" ", "", 0)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestFloatTag_stringify_compact(t *testing.T) {
	type Case struct {
		name     string
		tag      *FloatTag
		expected string
	}

	cases := []Case{}

	for _, c := range floatTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tag:      NewFloatTag(&c.nbt.tagName, c.nbt.payload),
			expected: fmt.Sprintf("%s:%s", c.snbt.tagName, c.snbt.payload.typeCompact),
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.tag.stringify("", "", 0)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestFloatTag_stringify_pretty(t *testing.T) {
	type Case struct {
		name     string
		tag      *FloatTag
		expected string
	}

	cases := []Case{}

	for _, c := range floatTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tag:      NewFloatTag(&c.nbt.tagName, c.nbt.payload),
			expected: fmt.Sprintf("%s: %s", c.snbt.tagName, c.snbt.payload.typePretty),
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.tag.stringify(" ", "  ", 0)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestFloatTag_json_default(t *testing.T) {
	type Case struct {
		name     string
		tag      *FloatTag
		expected string
	}

	cases := []Case{}

	for _, c := range floatTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tag:      NewFloatTag(&c.nbt.tagName, c.nbt.payload),
			expected: fmt.Sprintf("%s: %s", c.json.tagName, c.json.payload.typeDefault),
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.tag.json(" ", "", 0)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestFloatTag_json_compact(t *testing.T) {
	type Case struct {
		name     string
		tag      *FloatTag
		expected string
	}

	cases := []Case{}

	for _, c := range floatTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tag:      NewFloatTag(&c.nbt.tagName, c.nbt.payload),
			expected: fmt.Sprintf("%s:%s", c.json.tagName, c.json.payload.typeCompact),
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.tag.json("", "", 0)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestFloatTag_json_pretty(t *testing.T) {
	type Case struct {
		name     string
		tag      *FloatTag
		expected string
	}

	cases := []Case{}

	for _, c := range floatTagCases {
		cases = append(cases, Case{
			name:     c.name,
			tag:      NewFloatTag(&c.nbt.tagName, c.nbt.payload),
			expected: fmt.Sprintf("%s: %s", c.json.tagName, c.json.payload.typePretty),
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.tag.json(" ", "  ", 0)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestNewFloatPayload(t *testing.T) {
	type Case struct {
		name     string
		value    float32
		expected *FloatPayload
	}

	cases := []Case{}

	for _, c := range floatTagCases {
		cases = append(cases, Case{
			name:     c.name,
			value:    float32(*c.nbt.payload),
			expected: c.nbt.payload,
		})
	}

	for _, tt := range cases {
		actual := NewFloatPayload(tt.value)
		assert.Equal(t, tt.expected, actual)
	}
}

func TestFloatPayload_TypeId(t *testing.T) {
	payload := new(FloatPayload)
	expected := FloatType
	actual := payload.TypeId()
	assert.Equal(t, expected, actual)
}

func TestFloatPayload_encode(t *testing.T) {
	type Case struct {
		name        string
		payload     *FloatPayload
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range floatTagCases {
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

func TestFloatPayload_decode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    *FloatPayload
		expectedErr error
	}

	cases := []Case{}

	for _, c := range floatTagCases {
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

			payload := new(FloatPayload)
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

func TestFloatPayload_stringify_default(t *testing.T) {
	type Case struct {
		name     string
		payload  *FloatPayload
		expected string
	}

	cases := []Case{}

	for _, c := range floatTagCases {
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

func TestFloatPayload_stringify_compact(t *testing.T) {
	type Case struct {
		name     string
		payload  *FloatPayload
		expected string
	}

	cases := []Case{}

	for _, c := range floatTagCases {
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

func TestFloatPayload_stringify_pretty(t *testing.T) {
	type Case struct {
		name     string
		payload  *FloatPayload
		expected string
	}

	cases := []Case{}

	for _, c := range floatTagCases {
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

func TestFloatPayload_parse(t *testing.T) {
	type Case struct {
		name        string
		snbt        string
		expected    *FloatPayload
		expectedErr error
	}

	cases := []Case{}

	for _, c := range floatTagCases {
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

			payload := new(FloatPayload)
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

func TestFloatPayload_json_default(t *testing.T) {
	type Case struct {
		name     string
		payload  *FloatPayload
		expected string
	}

	cases := []Case{}

	for _, c := range floatTagCases {
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

func TestFloatPayload_json_compact(t *testing.T) {
	type Case struct {
		name     string
		payload  *FloatPayload
		expected string
	}

	cases := []Case{}

	for _, c := range floatTagCases {
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

func TestFloatPayload_json_pretty(t *testing.T) {
	type Case struct {
		name     string
		payload  *FloatPayload
		expected string
	}

	cases := []Case{}

	for _, c := range floatTagCases {
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
