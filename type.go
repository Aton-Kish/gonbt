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
	TagTypeEnd TagType = iota
	TagTypeByte
	TagTypeShort
	TagTypeInt
	TagTypeLong
	TagTypeFloat
	TagTypeDouble
	TagTypeByteArray
	TagTypeString
	TagTypeList
	TagTypeCompound
	TagTypeIntArray
	TagTypeLongArray
)

var TagTypes []TagType = []TagType{
	TagTypeEnd,
	TagTypeByte,
	TagTypeShort,
	TagTypeInt,
	TagTypeLong,
	TagTypeFloat,
	TagTypeDouble,
	TagTypeByteArray,
	TagTypeString,
	TagTypeList,
	TagTypeCompound,
	TagTypeIntArray,
	TagTypeLongArray,
}

func (t TagType) String() string {
	switch t {
	case TagTypeEnd:
		return "End"
	case TagTypeByte:
		return "Byte"
	case TagTypeShort:
		return "Short"
	case TagTypeInt:
		return "Int"
	case TagTypeLong:
		return "Long"
	case TagTypeFloat:
		return "Float"
	case TagTypeDouble:
		return "Double"
	case TagTypeByteArray:
		return "ByteArray"
	case TagTypeString:
		return "String"
	case TagTypeList:
		return "List"
	case TagTypeCompound:
		return "Compound"
	case TagTypeIntArray:
		return "IntArray"
	case TagTypeLongArray:
		return "LongArray"
	default:
		return ""
	}
}

func (t *TagType) encode(w io.Writer) error {
	if err := binary.Write(w, binary.BigEndian, t); err != nil {
		err = &NbtError{Op: "encode", Err: err}
		logger.Println("failed to encode", "func", getFuncName(), "type", t, "error", err)
		return err
	}

	return nil
}

func (t *TagType) decode(r io.Reader) error {
	if err := binary.Read(r, binary.BigEndian, t); err != nil {
		err = &NbtError{Op: "decode", Err: err}
		logger.Println("failed to decode", "func", getFuncName(), "type", t, "error", err)
		return err
	}

	return nil
}
