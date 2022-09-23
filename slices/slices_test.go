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

package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcat_int(t *testing.T) {
	cases := []struct {
		name     string
		slices   [][]int
		expected []int
	}{
		{
			name:     "positive case: empty",
			slices:   [][]int{},
			expected: []int{},
		},
		{
			name:     "positive case: []int",
			slices:   [][]int{{0, 1}},
			expected: []int{0, 1},
		},
		{
			name:     "positive case: []int []int",
			slices:   [][]int{{0, 1}, {2, 3}},
			expected: []int{0, 1, 2, 3},
		},
		{
			name:     "positive case: []int []int []int",
			slices:   [][]int{{0, 1}, {2, 3}, {4, 5}},
			expected: []int{0, 1, 2, 3, 4, 5},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := Concat(tt.slices...)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestConcat_string(t *testing.T) {
	cases := []struct {
		name     string
		slices   [][]string
		expected []string
	}{
		{
			name:     "positive case: empty",
			slices:   [][]string{},
			expected: []string{},
		},
		{
			name:     "positive case: []string",
			slices:   [][]string{{"Hello", "World"}},
			expected: []string{"Hello", "World"},
		},
		{
			name:     "positive case: []string []string",
			slices:   [][]string{{"Hello", "World"}, {"Hello", "World"}},
			expected: []string{"Hello", "World", "Hello", "World"},
		},
		{
			name:     "positive case: []string []string []string",
			slices:   [][]string{{"Hello", "World"}, {"Hello", "World"}, {"Hello", "World"}},
			expected: []string{"Hello", "World", "Hello", "World", "Hello", "World"},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := Concat(tt.slices...)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
