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
	case nbt.TagTypeByte:
		return "int8"
	case nbt.TagTypeShort:
		return "int16"
	case nbt.TagTypeInt:
		return "int32"
	case nbt.TagTypeLong:
		return "int64"
	case nbt.TagTypeFloat:
		return "float32"
	case nbt.TagTypeDouble:
		return "float64"
	case nbt.TagTypeByteArray:
		return "[]int8"
	case nbt.TagTypeString:
		return "string"
	case nbt.TagTypeList:
		return "[]Payload"
	case nbt.TagTypeCompound:
		return "[]Tag"
	case nbt.TagTypeIntArray:
		return "[]int32"
	case nbt.TagTypeLongArray:
		return "[]int64"
	default:
		return ""
	}
}
