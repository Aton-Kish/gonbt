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
	"fmt"
	"io"
)

// Tag Type

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

func (t *TagType) Encode(w io.Writer) error {
	if err := binary.Write(w, binary.BigEndian, t); err != nil {
		return err
	}

	return nil
}

func (t *TagType) Decode(r io.Reader) error {
	if err := binary.Read(r, binary.BigEndian, t); err != nil {
		return err
	}

	return nil
}

// Tag Name

type TagName string

func (n *TagName) Encode(w io.Writer) error {
	l := uint16(len(*n))
	if err := binary.Write(w, binary.BigEndian, &l); err != nil {
		return err
	}

	b := []byte(*n)
	if err := binary.Write(w, binary.BigEndian, b); err != nil {
		return err
	}

	return nil
}

func (n *TagName) Decode(r io.Reader) error {
	var l uint16
	if err := binary.Read(r, binary.BigEndian, &l); err != nil {
		return err
	}

	b := make([]byte, l)
	if err := binary.Read(r, binary.BigEndian, b); err != nil {
		return err
	}
	*n = TagName(b)

	return nil
}

// Tag

type Tag interface {
	TypeId() TagType
	TagName() *TagName
	Payload() Payload
	Encode(w io.Writer) error
	Decode(r io.Reader) error
}

func NewTag(typ TagType) (Tag, error) {
	switch typ {
	case EndType:
		return new(EndTag), nil
	case ByteType:
		return new(ByteTag), nil
	case ShortType:
		return new(ShortTag), nil
	case IntType:
		return new(IntTag), nil
	case LongType:
		return new(LongTag), nil
	case FloatType:
		return new(FloatTag), nil
	case DoubleType:
		return new(DoubleTag), nil
	case ByteArrayType:
		return new(ByteArrayTag), nil
	case StringType:
		return new(StringTag), nil
	case ListType:
		return new(ListTag), nil
	case CompoundType:
		return new(CompoundTag), nil
	case IntArrayType:
		return new(IntArrayTag), nil
	case LongArrayType:
		return new(LongArrayTag), nil
	default:
		return nil, fmt.Errorf("invalid tag type id %d", typ)
	}
}

func encodeTag(w io.Writer, tag Tag) error {
	typ := tag.TypeId()
	if err := typ.Encode(w); err != nil {
		return err
	}

	if typ == EndType {
		return nil
	}

	if err := tag.TagName().Encode(w); err != nil {
		return err
	}

	if err := tag.Payload().Encode(w); err != nil {
		return err
	}

	return nil
}

func decodeTagExcludeEndTag(r io.Reader, tag Tag) error {
	typ := tag.TypeId()
	if err := typ.Decode(r); err != nil {
		return err
	}

	if err := tag.TagName().Decode(r); err != nil {
		return err
	}

	if err := tag.Payload().Decode(r); err != nil {
		return err
	}

	return nil
}

// Payload

type Payload interface {
	TypeId() TagType
	Encode(w io.Writer) error
	Decode(r io.Reader) error
}

type NumericPayload interface {
	BytePayload | ShortPayload | IntPayload | LongPayload | FloatPayload | DoublePayload
}

type PrimitivePayload interface {
	NumericPayload | StringPayload
}

type ArrayPayload interface {
	ByteArrayPayload | IntArrayPayload | LongArrayPayload
}

func NewPayload(typ TagType) (Payload, error) {
	switch typ {
	case ByteType:
		return new(BytePayload), nil
	case ShortType:
		return new(ShortPayload), nil
	case IntType:
		return new(IntPayload), nil
	case LongType:
		return new(LongPayload), nil
	case FloatType:
		return new(FloatPayload), nil
	case DoubleType:
		return new(DoublePayload), nil
	case ByteArrayType:
		return new(ByteArrayPayload), nil
	case StringType:
		return new(StringPayload), nil
	case ListType:
		return new(ListPayload), nil
	case CompoundType:
		return new(CompoundPayload), nil
	case IntArrayType:
		return new(IntArrayPayload), nil
	case LongArrayType:
		return new(LongArrayPayload), nil
	default:
		return nil, fmt.Errorf("invalid tag type id %d", typ)
	}
}

func PrimitivePayloadPointer[T PrimitivePayload](x T) *T {
	return &x
}

func encodeNumericPayload[T NumericPayload](w io.Writer, payload *T) error {
	if err := binary.Write(w, binary.BigEndian, payload); err != nil {
		return err
	}

	return nil
}

func encodeArrayPayload[T ArrayPayload](w io.Writer, payload *T) error {
	l := int32(len(*payload))
	if err := binary.Write(w, binary.BigEndian, &l); err != nil {
		return err
	}

	if err := binary.Write(w, binary.BigEndian, payload); err != nil {
		return err
	}

	return nil
}

func decodeNumericPayload[T NumericPayload](r io.Reader, payload *T) error {
	if err := binary.Read(r, binary.BigEndian, payload); err != nil {
		return err
	}

	return nil
}
