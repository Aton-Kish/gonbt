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

func TestNewTag(t *testing.T) {
	cases := []struct {
		name        string
		tagType     TagType
		expected    Tag
		expectedErr error
	}{
		{
			name:        `positive case: TagTypeEnd`,
			tagType:     TagTypeEnd,
			expected:    NewEndTag(),
			expectedErr: nil,
		},
		{
			name:        `positive case: TagTypeByte`,
			tagType:     TagTypeByte,
			expected:    NewByteTag(new(TagName), new(BytePayload)),
			expectedErr: nil,
		},
		{
			name:        `positive case: TagTypeShort`,
			tagType:     TagTypeShort,
			expected:    NewShortTag(new(TagName), new(ShortPayload)),
			expectedErr: nil,
		},
		{
			name:        `positive case: TagTypeInt`,
			tagType:     TagTypeInt,
			expected:    NewIntTag(new(TagName), new(IntPayload)),
			expectedErr: nil,
		},
		{
			name:        `positive case: TagTypeLong`,
			tagType:     TagTypeLong,
			expected:    NewLongTag(new(TagName), new(LongPayload)),
			expectedErr: nil,
		},
		{
			name:        `positive case: TagTypeFloat`,
			tagType:     TagTypeFloat,
			expected:    NewFloatTag(new(TagName), new(FloatPayload)),
			expectedErr: nil,
		},
		{
			name:        `positive case: TagTypeDouble`,
			tagType:     TagTypeDouble,
			expected:    NewDoubleTag(new(TagName), new(DoublePayload)),
			expectedErr: nil,
		},
		{
			name:        `positive case: TagTypeByteArray`,
			tagType:     TagTypeByteArray,
			expected:    NewByteArrayTag(new(TagName), new(ByteArrayPayload)),
			expectedErr: nil,
		},
		{
			name:        `positive case: TagTypeString`,
			tagType:     TagTypeString,
			expected:    NewStringTag(new(TagName), new(StringPayload)),
			expectedErr: nil,
		},
		{
			name:        `positive case: TagTypeList`,
			tagType:     TagTypeList,
			expected:    NewListTag(new(TagName), new(ListPayload)),
			expectedErr: nil,
		},
		{
			name:        `positive case: TagTypeCompound`,
			tagType:     TagTypeCompound,
			expected:    NewCompoundTag(new(TagName), new(CompoundPayload)),
			expectedErr: nil,
		},
		{
			name:        `positive case: TagTypeIntArray`,
			tagType:     TagTypeIntArray,
			expected:    NewIntArrayTag(new(TagName), new(IntArrayPayload)),
			expectedErr: nil,
		},
		{
			name:        `positive case: TagTypeLongArray`,
			tagType:     TagTypeLongArray,
			expected:    NewLongArrayTag(new(TagName), new(LongArrayPayload)),
			expectedErr: nil,
		},
		{
			name:        `negative case: out of range`,
			tagType:     TagType(0x0D),
			expected:    nil,
			expectedErr: &NbtError{Op: "new", Err: ErrInvalidTagType},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tag, err := NewTag(tt.tagType)

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

func TestEncode(t *testing.T) {
	type Case struct {
		name        string
		nbt         Tag
		expected    []byte
		expectedErr error
	}

	cases := []Case{}

	for _, c := range nbtCases {
		cases = append(cases, Case{
			name:        c.name,
			nbt:         c.nbt,
			expected:    c.raw,
			expectedErr: nil,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			err := Encode(buf, tt.nbt)

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

func TestDecode(t *testing.T) {
	type Case struct {
		name        string
		raw         []byte
		expected    Tag
		expectedErr error
	}

	cases := []Case{}

	for _, c := range nbtCases {
		cases = append(cases, Case{
			name:        c.name,
			raw:         c.raw,
			expected:    c.nbt,
			expectedErr: nil,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer(tt.raw)
			actual, err := Decode(buf)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, actual)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr, err)
			}
		})
	}
}

func TestStringify(t *testing.T) {
	type Case struct {
		name     string
		nbt      Tag
		expected string
	}

	cases := []Case{}

	for _, c := range nbtCases {
		cases = append(cases, Case{
			name:     c.name,
			nbt:      c.nbt,
			expected: c.snbt.typeDefault,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := Stringify(tt.nbt)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestCompactStringify(t *testing.T) {
	type Case struct {
		name     string
		nbt      Tag
		expected string
	}

	cases := []Case{}

	for _, c := range nbtCases {
		cases = append(cases, Case{
			name:     c.name,
			nbt:      c.nbt,
			expected: c.snbt.typeCompact,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := CompactStringify(tt.nbt)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestPrettyStringify(t *testing.T) {
	type Case struct {
		name     string
		nbt      Tag
		expected string
	}

	cases := []Case{}

	for _, c := range nbtCases {
		cases = append(cases, Case{
			name:     c.name,
			nbt:      c.nbt,
			expected: c.snbt.typePretty,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := PrettyStringify(tt.nbt, "  ")
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestParse_default(t *testing.T) {
	type Case struct {
		name        string
		snbt        string
		expected    Tag
		expectedErr error
	}

	cases := []Case{}

	for _, c := range nbtCases {
		cases = append(cases, Case{
			name:        c.name,
			snbt:        c.snbt.typeDefault,
			expected:    c.nbt,
			expectedErr: nil,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := Parse(tt.snbt)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, Stringify(tt.expected), Stringify(actual))
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr, err)
			}
		})
	}
}

func TestParse_compact(t *testing.T) {
	type Case struct {
		name        string
		snbt        string
		expected    Tag
		expectedErr error
	}

	cases := []Case{}

	for _, c := range nbtCases {
		cases = append(cases, Case{
			name:        c.name,
			snbt:        c.snbt.typeCompact,
			expected:    c.nbt,
			expectedErr: nil,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := Parse(tt.snbt)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, Stringify(tt.expected), Stringify(actual))
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr, err)
			}
		})
	}
}

func TestParse_pretty(t *testing.T) {
	type Case struct {
		name        string
		snbt        string
		expected    Tag
		expectedErr error
	}

	cases := []Case{}

	for _, c := range nbtCases {
		cases = append(cases, Case{
			name:        c.name,
			snbt:        c.snbt.typePretty,
			expected:    c.nbt,
			expectedErr: nil,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := Parse(tt.snbt)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, Stringify(tt.expected), Stringify(actual))
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr, err)
			}
		})
	}
}

func TestJson(t *testing.T) {
	type Case struct {
		name     string
		nbt      Tag
		expected string
	}

	cases := []Case{}

	for _, c := range nbtCases {
		cases = append(cases, Case{
			name:     c.name,
			nbt:      c.nbt,
			expected: c.json.typeDefault,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := Json(tt.nbt)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestCompactJson(t *testing.T) {
	type Case struct {
		name     string
		nbt      Tag
		expected string
	}

	cases := []Case{}

	for _, c := range nbtCases {
		cases = append(cases, Case{
			name:     c.name,
			nbt:      c.nbt,
			expected: c.json.typeCompact,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := CompactJson(tt.nbt)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestPrettyJson(t *testing.T) {
	type Case struct {
		name     string
		nbt      Tag
		expected string
	}

	cases := []Case{}

	for _, c := range nbtCases {
		cases = append(cases, Case{
			name:     c.name,
			nbt:      c.nbt,
			expected: c.json.typePretty,
		})
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := PrettyJson(tt.nbt, "  ")
			assert.Equal(t, tt.expected, actual)
		})
	}
}
