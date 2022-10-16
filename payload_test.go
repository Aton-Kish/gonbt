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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPayload(t *testing.T) {
	cases := []struct {
		name        string
		tagType     TagType
		expected    Payload
		expectedErr error
	}{
		{
			name:        `positive case: TagTypeByte`,
			tagType:     TagTypeByte,
			expected:    new(BytePayload),
			expectedErr: nil,
		},
		{
			name:        `positive case: TagTypeShort`,
			tagType:     TagTypeShort,
			expected:    new(ShortPayload),
			expectedErr: nil,
		},
		{
			name:        `positive case: TagTypeInt`,
			tagType:     TagTypeInt,
			expected:    new(IntPayload),
			expectedErr: nil,
		},
		{
			name:        `positive case: TagTypeLong`,
			tagType:     TagTypeLong,
			expected:    new(LongPayload),
			expectedErr: nil,
		},
		{
			name:        `positive case: TagTypeFloat`,
			tagType:     TagTypeFloat,
			expected:    new(FloatPayload),
			expectedErr: nil,
		},
		{
			name:        `positive case: TagTypeDouble`,
			tagType:     TagTypeDouble,
			expected:    new(DoublePayload),
			expectedErr: nil,
		},
		{
			name:        `positive case: TagTypeByteArray`,
			tagType:     TagTypeByteArray,
			expected:    new(ByteArrayPayload),
			expectedErr: nil,
		},
		{
			name:        `positive case: TagTypeString`,
			tagType:     TagTypeString,
			expected:    new(StringPayload),
			expectedErr: nil,
		},
		{
			name:        `positive case: TagTypeList`,
			tagType:     TagTypeList,
			expected:    new(ListPayload),
			expectedErr: nil,
		},
		{
			name:        `positive case: TagTypeCompound`,
			tagType:     TagTypeCompound,
			expected:    new(CompoundPayload),
			expectedErr: nil,
		},
		{
			name:        `positive case: TagTypeIntArray`,
			tagType:     TagTypeIntArray,
			expected:    new(IntArrayPayload),
			expectedErr: nil,
		},
		{
			name:        `positive case: TagTypeLongArray`,
			tagType:     TagTypeLongArray,
			expected:    new(LongArrayPayload),
			expectedErr: nil,
		},
		{
			name:        `negative case`,
			tagType:     TagType(0x0D),
			expected:    nil,
			expectedErr: &NbtError{Op: "new", Err: ErrInvalidTagType},
		},
		{
			name:        `negative case: TagTypeEnd`,
			tagType:     TagTypeEnd,
			expected:    nil,
			expectedErr: &NbtError{Op: "new", Err: ErrInvalidTagType},
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
			payload, err := NewPayload(tt.tagType)

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
