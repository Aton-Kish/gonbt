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

package main

import (
	nbt "github.com/Aton-Kish/gonbt"
	"github.com/iancoleman/strcase"
)

func public(typ nbt.TagType) string {
	return strcase.ToCamel(typ.String())
}

func private(typ nbt.TagType) string {
	return strcase.ToLowerCamel(typ.String())
}

func typeof(typ nbt.TagType) string {
	switch typ {
	case nbt.ByteType:
		return "int8"
	case nbt.ShortType:
		return "int16"
	case nbt.IntType:
		return "int32"
	case nbt.LongType:
		return "int64"
	case nbt.FloatType:
		return "float32"
	case nbt.DoubleType:
		return "float64"
	case nbt.ByteArrayType:
		return "[]int8"
	case nbt.StringType:
		return "string"
	case nbt.ListType:
		return "[]Payload"
	case nbt.CompoundType:
		return "[]Tag"
	case nbt.IntArrayType:
		return "[]int32"
	case nbt.LongArrayType:
		return "[]int64"
	default:
		return ""
	}
}
