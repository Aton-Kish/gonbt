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
	"regexp"

	"github.com/Aton-Kish/gonbt/snbt"
)

var (
	bytePattern   = regexp.MustCompile(`^(-?\d+)[Bb]$`)
	shortPattern  = regexp.MustCompile(`^(-?\d+)[Ss]$`)
	intPattern    = regexp.MustCompile(`^-?\d+$`)
	longPattern   = regexp.MustCompile(`^(-?\d+)[Ll]$`)
	floatPattern  = regexp.MustCompile(`^(-?\d+(\.\d+)?([Ee][+-]?\d+)?)[Ff]$`)
	doublePattern = regexp.MustCompile(`^(-?\d+(\.\d+)?([Ee][+-]?\d+)?)[Dd]?$`)
)

type Payload interface {
	TypeId() TagType
	encode(w io.Writer) error
	decode(r io.Reader) error
	stringify(space string, indent string, depth int) string
	parse(parser *snbt.Parser) error
	json(space string, indent string, depth int) string
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
		err := &NbtError{Op: "new", Err: ErrInvalidTagType}
		return nil, err
	}
}

func newPayloadFromSnbt(parser *snbt.Parser) (Payload, error) {
	switch parser.CurrToken().Char() {
	case '{':
		return new(CompoundPayload), nil
	case '[':
		typ, err := parser.Char(parser.CurrToken().Index() + 1)
		if err != nil {
			err = &NbtError{Op: "new", Err: err}
			return nil, err
		}

		switch typ {
		case 'B':
			return new(ByteArrayPayload), nil
		case 'I':
			return new(IntArrayPayload), nil
		case 'L':
			return new(LongArrayPayload), nil
		}

		return new(ListPayload), nil
	case *new(rune), '"', ' ', ':', ';':
		err := &NbtError{Op: "new", Err: ErrInvalidSnbtFormat}
		return nil, err
	}

	b, err := parser.Slice(parser.PrevToken().Index()+1, parser.CurrToken().Index())
	if err != nil {
		err = &NbtError{Op: "new", Err: err}
		return nil, err
	}

	if bytePattern.Match(b) {
		return new(BytePayload), nil
	}

	if shortPattern.Match(b) {
		return new(ShortPayload), nil
	}

	if intPattern.Match(b) {
		return new(IntPayload), nil
	}

	if longPattern.Match(b) {
		return new(LongPayload), nil
	}

	if floatPattern.Match(b) {
		return new(FloatPayload), nil
	}

	if doublePattern.Match(b) {
		return new(DoublePayload), nil
	}

	return new(StringPayload), nil
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

func encodeNumericPayload[T NumericPayload](w io.Writer, payload *T) error {
	if err := binary.Write(w, binary.BigEndian, payload); err != nil {
		err = &NbtError{Op: "encode", Err: err}
		return err
	}

	return nil
}

func encodeArrayPayload[T ArrayPayload](w io.Writer, payload *T) error {
	l := int32(len(*payload))
	if err := binary.Write(w, binary.BigEndian, &l); err != nil {
		err = &NbtError{Op: "encode", Err: err}
		return err
	}

	if err := binary.Write(w, binary.BigEndian, payload); err != nil {
		err = &NbtError{Op: "encode", Err: err}
		return err
	}

	return nil
}

func decodeNumericPayload[T NumericPayload](r io.Reader) (*T, error) {
	var payload T
	if err := binary.Read(r, binary.BigEndian, &payload); err != nil {
		err = &NbtError{Op: "decode", Err: err}
		return nil, err
	}

	return &payload, nil
}
