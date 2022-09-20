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

func TestPayload_BytePayload(t *testing.T) {
	value := BytePayload(123)
	expected := &value
	actual := PrimitivePayloadPointer(value)
	assert.Equal(t, expected, actual)
}

func TestPayload_ShortPayload(t *testing.T) {
	value := ShortPayload(12345)
	expected := &value
	actual := PrimitivePayloadPointer(value)
	assert.Equal(t, expected, actual)
}

func TestPayload_IntPayload(t *testing.T) {
	value := IntPayload(123456789)
	expected := &value
	actual := PrimitivePayloadPointer(value)
	assert.Equal(t, expected, actual)
}

func TestPayload_LongPayload(t *testing.T) {
	value := LongPayload(123456789123456789)
	expected := &value
	actual := PrimitivePayloadPointer(value)
	assert.Equal(t, expected, actual)
}

func TestPayload_FloatPayload(t *testing.T) {
	value := FloatPayload(0.12345678)
	expected := &value
	actual := PrimitivePayloadPointer(value)
	assert.Equal(t, expected, actual)
}

func TestPayload_DoublePayload(t *testing.T) {
	value := DoublePayload(0.123456789)
	expected := &value
	actual := PrimitivePayloadPointer(value)
	assert.Equal(t, expected, actual)
}

func TestPayload_StringPayload(t *testing.T) {
	value := StringPayload("Test")
	expected := &value
	actual := PrimitivePayloadPointer(value)
	assert.Equal(t, expected, actual)
}
