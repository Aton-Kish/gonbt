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
	"encoding/binary"
	"io"
)

type TagType byte

const (
	EndType TagType = iota
	ByteType
	ShortType
	IntType
	LongType
	FloatType
	DoubleType
	ByteArrayType
	StringType
	ListType
	CompoundType
	IntArrayType
	LongArrayType
)

var TagTypes []TagType = []TagType{
	EndType, ByteType, ShortType, IntType, LongType, FloatType, DoubleType,
	ByteArrayType, StringType, ListType, CompoundType, IntArrayType, LongArrayType,
}

func (t TagType) String() string {
	switch t {
	case EndType:
		return "End"
	case ByteType:
		return "Byte"
	case ShortType:
		return "Short"
	case IntType:
		return "Int"
	case LongType:
		return "Long"
	case FloatType:
		return "Float"
	case DoubleType:
		return "Double"
	case ByteArrayType:
		return "ByteArray"
	case StringType:
		return "String"
	case ListType:
		return "List"
	case CompoundType:
		return "Compound"
	case IntArrayType:
		return "IntArray"
	case LongArrayType:
		return "LongArray"
	default:
		return ""
	}
}

func (t *TagType) encode(w io.Writer) error {
	if err := binary.Write(w, binary.BigEndian, t); err != nil {
		err = &NbtError{Op: "encode", Err: err}
		return err
	}

	return nil
}

func (t *TagType) decode(r io.Reader) error {
	if err := binary.Read(r, binary.BigEndian, t); err != nil {
		err = &NbtError{Op: "decode", Err: err}
		return err
	}

	return nil
}
