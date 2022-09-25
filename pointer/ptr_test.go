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

package pointer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointer_int8(t *testing.T) {
	value := int8(123)
	expected := &value
	actual := Pointer(value)
	assert.Equal(t, expected, actual)
}

func TestPointer_int16(t *testing.T) {
	value := int16(12345)
	expected := &value
	actual := Pointer(value)
	assert.Equal(t, expected, actual)
}

func TestPointer_int32(t *testing.T) {
	value := int32(123456789)
	expected := &value
	actual := Pointer(value)
	assert.Equal(t, expected, actual)
}

func TestPointer_int64(t *testing.T) {
	value := int64(123456789123456789)
	expected := &value
	actual := Pointer(value)
	assert.Equal(t, expected, actual)
}

func TestPointer_float32(t *testing.T) {
	value := float32(0.12345678)
	expected := &value
	actual := Pointer(value)
	assert.Equal(t, expected, actual)
}

func TestPointer_float64(t *testing.T) {
	value := float64(0.123456789)
	expected := &value
	actual := Pointer(value)
	assert.Equal(t, expected, actual)
}

func TestPointer_string(t *testing.T) {
	value := string("Test")
	expected := &value
	actual := Pointer(value)
	assert.Equal(t, expected, actual)
}
