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

package snbt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeRightmost_uint8(t *testing.T) {
	cases := []struct {
		name     string
		bitmap   uint8
		expected uint8
	}{
		{
			name:     "positive case: 0b_00000000",
			bitmap:   0b_00000000,
			expected: 0b_00000000,
		},
		{
			name:     "positive case: 0b_00000001",
			bitmap:   0b_00000001,
			expected: 0b_00000000,
		},
		{
			name:     "positive case: 0b_00000010",
			bitmap:   0b_00000010,
			expected: 0b_00000000,
		},
		{
			name:     "positive case: 0b_00000110",
			bitmap:   0b_00000110,
			expected: 0b_00000100,
		},
		{
			name:     "positive case: 0b_11111111",
			bitmap:   0b_11111111,
			expected: 0b_11111110,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := removeRightmost(tt.bitmap)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func Test_removeRightmost_uint64(t *testing.T) {
	cases := []struct {
		name     string
		bitmap   uint64
		expected uint64
	}{
		{
			name:     "positive case: 0b_00000000_00000000_00000000_00000000_00000000_00000000_00000000_00000000",
			bitmap:   0b_00000000_00000000_00000000_00000000_00000000_00000000_00000000_00000000,
			expected: 0b_00000000_00000000_00000000_00000000_00000000_00000000_00000000_00000000,
		},
		{
			name:     "positive case: 0b_00000000_00000000_00000000_00000000_00000000_00000000_00000000_00000001",
			bitmap:   0b_00000000_00000000_00000000_00000000_00000000_00000000_00000000_00000001,
			expected: 0b_00000000_00000000_00000000_00000000_00000000_00000000_00000000_00000000,
		},
		{
			name:     "positive case: 0b_00000000_00000000_00000000_00000000_00000000_00000000_00000001_00000000",
			bitmap:   0b_00000000_00000000_00000000_00000000_00000000_00000000_00000001_00000000,
			expected: 0b_00000000_00000000_00000000_00000000_00000000_00000000_00000000_00000000,
		},
		{
			name:     "positive case: 0b_00000000_00000000_00000000_00000000_00000000_00000001_00000001_00000000",
			bitmap:   0b_00000000_00000000_00000000_00000000_00000000_00000001_00000001_00000000,
			expected: 0b_00000000_00000000_00000000_00000000_00000000_00000001_00000000_00000000,
		},
		{
			name:     "positive case: 0b_11111111_11111111_11111111_11111111_11111111_11111111_11111111_11111111",
			bitmap:   0b_11111111_11111111_11111111_11111111_11111111_11111111_11111111_11111111,
			expected: 0b_11111111_11111111_11111111_11111111_11111111_11111111_11111111_11111110,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := removeRightmost(tt.bitmap)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func Test_extractRightmost_uint8(t *testing.T) {
	cases := []struct {
		name     string
		bitmap   uint8
		expected uint8
	}{
		{
			name:     "positive case: 0b_00000000",
			bitmap:   0b_00000000,
			expected: 0b_00000000,
		},
		{
			name:     "positive case: 0b_00000001",
			bitmap:   0b_00000001,
			expected: 0b_00000001,
		},
		{
			name:     "positive case: 0b_00000010",
			bitmap:   0b_00000010,
			expected: 0b_00000010,
		},
		{
			name:     "positive case: 0b_00000110",
			bitmap:   0b_00000110,
			expected: 0b_00000010,
		},
		{
			name:     "positive case: 0b_11111111",
			bitmap:   0b_11111111,
			expected: 0b_00000001,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := extractRightmost(tt.bitmap)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func Test_extractRightmost_uint64(t *testing.T) {
	cases := []struct {
		name     string
		bitmap   uint64
		expected uint64
	}{
		{
			name:     "positive case: 0b_00000000_00000000_00000000_00000000_00000000_00000000_00000000_00000000",
			bitmap:   0b_00000000_00000000_00000000_00000000_00000000_00000000_00000000_00000000,
			expected: 0b_00000000_00000000_00000000_00000000_00000000_00000000_00000000_00000000,
		},
		{
			name:     "positive case: 0b_00000000_00000000_00000000_00000000_00000000_00000000_00000000_00000001",
			bitmap:   0b_00000000_00000000_00000000_00000000_00000000_00000000_00000000_00000001,
			expected: 0b_00000000_00000000_00000000_00000000_00000000_00000000_00000000_00000001,
		},
		{
			name:     "positive case: 0b_00000000_00000000_00000000_00000000_00000000_00000000_00000001_00000000",
			bitmap:   0b_00000000_00000000_00000000_00000000_00000000_00000000_00000001_00000000,
			expected: 0b_00000000_00000000_00000000_00000000_00000000_00000000_00000001_00000000,
		},
		{
			name:     "positive case: 0b_00000000_00000000_00000000_00000000_00000000_00000001_00000001_00000000",
			bitmap:   0b_00000000_00000000_00000000_00000000_00000000_00000001_00000001_00000000,
			expected: 0b_00000000_00000000_00000000_00000000_00000000_00000000_00000001_00000000,
		},
		{
			name:     "positive case: 0b_11111111_11111111_11111111_11111111_11111111_11111111_11111111_11111111",
			bitmap:   0b_11111111_11111111_11111111_11111111_11111111_11111111_11111111_11111111,
			expected: 0b_00000000_00000000_00000000_00000000_00000000_00000000_00000000_00000001,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := extractRightmost(tt.bitmap)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func Test_smearRightmost_uint8(t *testing.T) {
	cases := []struct {
		name     string
		bitmap   uint8
		expected uint8
	}{
		{
			name:     "positive case: 0b_00000000",
			bitmap:   0b_00000000,
			expected: 0b_11111111,
		},
		{
			name:     "positive case: 0b_00000001",
			bitmap:   0b_00000001,
			expected: 0b_00000001,
		},
		{
			name:     "positive case: 0b_00000010",
			bitmap:   0b_00000010,
			expected: 0b_00000011,
		},
		{
			name:     "positive case: 0b_00000110",
			bitmap:   0b_00000110,
			expected: 0b_00000011,
		},
		{
			name:     "positive case: 0b_11111111",
			bitmap:   0b_11111111,
			expected: 0b_00000001,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := smearRightmost(tt.bitmap)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func Test_smearRightmost_uint64(t *testing.T) {
	cases := []struct {
		name     string
		bitmap   uint64
		expected uint64
	}{
		{
			name:     "positive case: 0b_00000000_00000000_00000000_00000000_00000000_00000000_00000000_00000000",
			bitmap:   0b_00000000_00000000_00000000_00000000_00000000_00000000_00000000_00000000,
			expected: 0b_11111111_11111111_11111111_11111111_11111111_11111111_11111111_11111111,
		},
		{
			name:     "positive case: 0b_00000000_00000000_00000000_00000000_00000000_00000000_00000000_00000001",
			bitmap:   0b_00000000_00000000_00000000_00000000_00000000_00000000_00000000_00000001,
			expected: 0b_00000000_00000000_00000000_00000000_00000000_00000000_00000000_00000001,
		},
		{
			name:     "positive case: 0b_00000000_00000000_00000000_00000000_00000000_00000000_00000001_00000000",
			bitmap:   0b_00000000_00000000_00000000_00000000_00000000_00000000_00000001_00000000,
			expected: 0b_00000000_00000000_00000000_00000000_00000000_00000000_00000001_11111111,
		},
		{
			name:     "positive case: 0b_00000000_00000000_00000000_00000000_00000000_00000001_00000001_00000000",
			bitmap:   0b_00000000_00000000_00000000_00000000_00000000_00000001_00000001_00000000,
			expected: 0b_00000000_00000000_00000000_00000000_00000000_00000000_00000001_11111111,
		},
		{
			name:     "positive case: 0b_11111111_11111111_11111111_11111111_11111111_11111111_11111111_11111111",
			bitmap:   0b_11111111_11111111_11111111_11111111_11111111_11111111_11111111_11111111,
			expected: 0b_00000000_00000000_00000000_00000000_00000000_00000000_00000000_00000001,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := smearRightmost(tt.bitmap)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func Test_rightmostIndex_uint8(t *testing.T) {
	cases := []struct {
		name     string
		bitmap   uint8
		expected int
	}{
		{
			name:     "positive case: 0b_00000000",
			bitmap:   0b_00000000,
			expected: 0,
		},
		{
			name:     "positive case: 0b_00000001",
			bitmap:   0b_00000001,
			expected: 0,
		},
		{
			name:     "positive case: 0b_00000010",
			bitmap:   0b_00000010,
			expected: 1,
		},
		{
			name:     "positive case: 0b_00000110",
			bitmap:   0b_00000110,
			expected: 1,
		},
		{
			name:     "positive case: 0b_11111111",
			bitmap:   0b_11111111,
			expected: 0,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := rightmostIndex(tt.bitmap)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func Test_rightmostIndex_uint64(t *testing.T) {
	cases := []struct {
		name     string
		bitmap   uint64
		expected int
	}{
		{
			name:     "positive case: 0b_00000000_00000000_00000000_00000000_00000000_00000000_00000000_00000000",
			bitmap:   0b_00000000_00000000_00000000_00000000_00000000_00000000_00000000_00000000,
			expected: 0,
		},
		{
			name:     "positive case: 0b_00000000_00000000_00000000_00000000_00000000_00000000_00000000_00000001",
			bitmap:   0b_00000000_00000000_00000000_00000000_00000000_00000000_00000000_00000001,
			expected: 0,
		},
		{
			name:     "positive case: 0b_00000000_00000000_00000000_00000000_00000000_00000000_00000001_00000000",
			bitmap:   0b_00000000_00000000_00000000_00000000_00000000_00000000_00000001_00000000,
			expected: 8,
		},
		{
			name:     "positive case: 0b_00000000_00000000_00000000_00000000_00000000_00000001_00000001_00000000",
			bitmap:   0b_00000000_00000000_00000000_00000000_00000000_00000001_00000001_00000000,
			expected: 8,
		},
		{
			name:     "positive case: 0b_11111111_11111111_11111111_11111111_11111111_11111111_11111111_11111111",
			bitmap:   0b_11111111_11111111_11111111_11111111_11111111_11111111_11111111_11111111,
			expected: 0,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := rightmostIndex(tt.bitmap)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func Test_popCount_uint8(t *testing.T) {
	cases := []struct {
		name     string
		bitmaps  []uint8
		expected int
	}{
		{
			name:     "positive case: 0x_00",
			bitmaps:  []uint8{0x_00},
			expected: 0,
		},
		{
			name:     "positive case: 0x_01",
			bitmaps:  []uint8{0x_01},
			expected: 1,
		},
		{
			name:     "positive case: 0x_02",
			bitmaps:  []uint8{0x_02},
			expected: 1,
		},
		{
			name:     "positive case: 0x_06",
			bitmaps:  []uint8{0x_06},
			expected: 2,
		},
		{
			name:     "positive case: 0x_FF",
			bitmaps:  []uint8{0x_FF},
			expected: 8,
		},
		{
			name:     "positive case: 0x_00 0x_01",
			bitmaps:  []uint8{0x_00, 0x_01},
			expected: 1,
		},
		{
			name:     "positive case: 0x_01 0x_02",
			bitmaps:  []uint8{0x_01, 0x_02},
			expected: 2,
		},
		{
			name:     "positive case: 0x_02 0x_06",
			bitmaps:  []uint8{0x_02, 0x_06},
			expected: 3,
		},
		{
			name:     "positive case: 0x_06 0x_FF",
			bitmaps:  []uint8{0x_06, 0x_FF},
			expected: 10,
		},
		{
			name:     "positive case: 0x_FF 0x_00",
			bitmaps:  []uint8{0x_FF, 0x_00},
			expected: 8,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := popCount(tt.bitmaps...)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func Test_popCount_uint64(t *testing.T) {
	cases := []struct {
		name     string
		bitmaps  []uint64
		expected int
	}{
		{
			name:     "positive case: 0x_00_00_00_00_00_00_00_00",
			bitmaps:  []uint64{0x_00_00_00_00_00_00_00_00},
			expected: 0,
		},
		{
			name:     "positive case: 0x_00_00_00_00_00_00_00_01",
			bitmaps:  []uint64{0x_00_00_00_00_00_00_00_01},
			expected: 1,
		},
		{
			name:     "positive case: 0x_00_00_00_00_00_00_01_00",
			bitmaps:  []uint64{0x_00_00_00_00_00_00_01_00},
			expected: 1,
		},
		{
			name:     "positive case: 0x_00_00_00_00_00_01_01_00",
			bitmaps:  []uint64{0x_00_00_00_00_00_01_01_00},
			expected: 2,
		},
		{
			name:     "positive case: 0x_FF_FF_FF_FF_FF_FF_FF_FF",
			bitmaps:  []uint64{0x_FF_FF_FF_FF_FF_FF_FF_FF},
			expected: 64,
		},
		{
			name:     "positive case: 0x_00_00_00_00_00_00_00_00 0x_00_00_00_00_00_00_00_01",
			bitmaps:  []uint64{0x_00_00_00_00_00_00_00_00, 0x_00_00_00_00_00_00_00_01},
			expected: 1,
		},
		{
			name:     "positive case: 0x_00_00_00_00_00_00_00_01 0x_00_00_00_00_00_00_01_00",
			bitmaps:  []uint64{0x_00_00_00_00_00_00_00_01, 0x_00_00_00_00_00_00_01_00},
			expected: 2,
		},
		{
			name:     "positive case: 0x_00_00_00_00_00_00_01_00 0x_00_00_00_00_00_01_01_00",
			bitmaps:  []uint64{0x_00_00_00_00_00_00_01_00, 0x_00_00_00_00_00_01_01_00},
			expected: 3,
		},
		{
			name:     "positive case: 0x_00_00_00_00_00_01_01_00 0x_FF_FF_FF_FF_FF_FF_FF_FF",
			bitmaps:  []uint64{0x_00_00_00_00_00_01_01_00, 0x_FF_FF_FF_FF_FF_FF_FF_FF},
			expected: 66,
		},
		{
			name:     "positive case: 0x_FF_FF_FF_FF_FF_FF_FF_FF 0x_00_00_00_00_00_00_00_00",
			bitmaps:  []uint64{0x_FF_FF_FF_FF_FF_FF_FF_FF, 0x_00_00_00_00_00_00_00_00},
			expected: 64,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := popCount(tt.bitmaps...)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
