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
	"golang.org/x/exp/constraints"
)

var (
	hashTarZ  = uint64(0x03F566ED27179461)
	tableTarZ = [64]int{
		0, 1, 59, 2, 60, 40, 54, 3, 61, 32, 49, 41, 55, 19, 35, 4,
		62, 52, 30, 33, 50, 12, 14, 42, 56, 16, 27, 20, 36, 23, 44, 5,
		63, 58, 39, 53, 31, 48, 18, 34, 51, 29, 11, 13, 15, 26, 22, 43,
		57, 38, 47, 17, 28, 10, 25, 21, 37, 46, 9, 24, 45, 8, 7, 6,
	}
)

func removeRightmost[T constraints.Unsigned](bitmap T) T {
	return bitmap & (bitmap - 1)
}

func extractRightmost[T constraints.Unsigned](bitmap T) T {
	return bitmap & -bitmap
}

func smearRightmost[T constraints.Unsigned](bitmap T) T {
	return bitmap ^ (bitmap - 1)
}

func rightmostIndex[T constraints.Unsigned](bitmap T) int {
	r := extractRightmost(bitmap)
	x := uint64(r)
	return tableTarZ[(x*hashTarZ)>>58]
}

func popCount[T constraints.Unsigned](bitmaps ...T) int {
	c := 0
	for _, bitmap := range bitmaps {
		x := uint64(bitmap)
		x = (x & 0x5555555555555555) + ((x >> 1) & 0x5555555555555555)
		x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
		x = (x & 0x0F0F0F0F0F0F0F0F) + ((x >> 4) & 0x0F0F0F0F0F0F0F0F)
		x += x >> 8
		x += x >> 16
		x += x >> 32
		c += int(x & 0x7f)
	}

	return c
}
