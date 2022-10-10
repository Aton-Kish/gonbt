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
	"regexp"
	"strconv"
	"strings"

	"github.com/Aton-Kish/gonbt/pointer"
	"github.com/Aton-Kish/gonbt/snbt"
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

// Tag Name

var quoteRequiredCharacters = regexp.MustCompile(`[ !"#$%&'()*,/:;<=>?@[\\\]^\x60{|}~]`)

type TagName string

func NewTagName(value string) *TagName {
	return pointer.Pointer(TagName(value))
}

func (n *TagName) encode(w io.Writer) error {
	l := uint16(len(*n))
	if err := binary.Write(w, binary.BigEndian, &l); err != nil {
		err = &NbtError{Op: "encode", Err: err}
		return err
	}

	b := []byte(*n)
	if err := binary.Write(w, binary.BigEndian, b); err != nil {
		err = &NbtError{Op: "encode", Err: err}
		return err
	}

	return nil
}

func (n *TagName) decode(r io.Reader) error {
	var l uint16
	if err := binary.Read(r, binary.BigEndian, &l); err != nil {
		err = &NbtError{Op: "decode", Err: err}
		return err
	}

	b := make([]byte, l)
	if err := binary.Read(r, binary.BigEndian, b); err != nil {
		err = &NbtError{Op: "decode", Err: err}
		return err
	}
	*n = TagName(b)

	return nil
}

func (n *TagName) stringify() string {
	s := string(*n)
	if !quoteRequiredCharacters.MatchString(s) {
		return s
	}

	qs := strconv.Quote(s)
	if strings.Contains(s, "\"") && !strings.Contains(s, "'") {
		qs = fmt.Sprintf("'%s'", qs[1:len(qs)-1])
		qs = strings.ReplaceAll(qs, "\\\"", "\"")
		return qs
	}

	return qs
}

func (n *TagName) parse(parser *snbt.Parser) error {
	b, err := parser.Slice(parser.PrevToken().Index()+1, parser.CurrToken().Index())
	if err != nil {
		err = &NbtError{Op: "parse", Err: err}
		return err
	}

	if s, err := strconv.Unquote(string(b)); err != nil {
		*n = TagName(b)
	} else {
		*n = TagName(s)
	}

	if err := parser.Next(); err != nil {
		err = &NbtError{Op: "parse", Err: err}
		return err
	}

	return nil
}

func (n *TagName) json() string {
	s := string(*n)
	return strconv.Quote(s)
}

// Tag

type Tag interface {
	TypeId() TagType
	TagName() *TagName
	Payload() Payload
	encode(w io.Writer) error
	decode(r io.Reader) error
	stringify(space string, indent string, depth int) string
	parse(parser *snbt.Parser) error
	json(space string, indent string, depth int) string
}

func NewTag(typ TagType) (Tag, error) {
	switch typ {
	case EndType:
		return NewEndTag(), nil
	case ByteType:
		return NewByteTag(new(TagName), new(BytePayload)), nil
	case ShortType:
		return NewShortTag(new(TagName), new(ShortPayload)), nil
	case IntType:
		return NewIntTag(new(TagName), new(IntPayload)), nil
	case LongType:
		return NewLongTag(new(TagName), new(LongPayload)), nil
	case FloatType:
		return NewFloatTag(new(TagName), new(FloatPayload)), nil
	case DoubleType:
		return NewDoubleTag(new(TagName), new(DoublePayload)), nil
	case ByteArrayType:
		return NewByteArrayTag(new(TagName), new(ByteArrayPayload)), nil
	case StringType:
		return NewStringTag(new(TagName), new(StringPayload)), nil
	case ListType:
		return NewListTag(new(TagName), new(ListPayload)), nil
	case CompoundType:
		return NewCompoundTag(new(TagName), new(CompoundPayload)), nil
	case IntArrayType:
		return NewIntArrayTag(new(TagName), new(IntArrayPayload)), nil
	case LongArrayType:
		return NewLongArrayTag(new(TagName), new(LongArrayPayload)), nil
	default:
		err := &NbtError{Op: "new", Err: invalidTagTypeError}
		return nil, err
	}
}

func newTagFromSnbt(parser *snbt.Parser) (Tag, error) {
	var name TagName
	// NOTE: skip if nameless root
	if parser.CurrToken().Index() > 0 {
		if err := name.parse(parser); err != nil {
			err = &NbtError{Op: "new", Err: err}
			return nil, err
		}
	}

	p, err := newPayloadFromSnbt(parser)
	if err != nil {
		return nil, err
	}

	switch payload := p.(type) {
	case *BytePayload:
		return NewByteTag(&name, payload), nil
	case *ShortPayload:
		return NewShortTag(&name, payload), nil
	case *IntPayload:
		return NewIntTag(&name, payload), nil
	case *LongPayload:
		return NewLongTag(&name, payload), nil
	case *FloatPayload:
		return NewFloatTag(&name, payload), nil
	case *DoublePayload:
		return NewDoubleTag(&name, payload), nil
	case *ByteArrayPayload:
		return NewByteArrayTag(&name, payload), nil
	case *StringPayload:
		return NewStringTag(&name, payload), nil
	case *ListPayload:
		return NewListTag(&name, payload), nil
	case *CompoundPayload:
		return NewCompoundTag(&name, payload), nil
	case *IntArrayPayload:
		return NewIntArrayTag(&name, payload), nil
	case *LongArrayPayload:
		return NewLongArrayTag(&name, payload), nil
	default:
		err = &NbtError{Op: "new", Err: invalidSnbtFormatError}
		return nil, err
	}
}

func Encode(w io.Writer, tag Tag) error {
	typ := tag.TypeId()
	if err := typ.encode(w); err != nil {
		return err
	}

	if typ == EndType {
		return nil
	}

	if err := tag.TagName().encode(w); err != nil {
		return err
	}

	if err := tag.Payload().encode(w); err != nil {
		return err
	}

	return nil
}

func Decode(r io.Reader) (Tag, error) {
	var typ TagType
	if err := typ.decode(r); err != nil {
		return nil, err
	}

	tag, err := NewTag(typ)
	if err != nil {
		err = &NbtError{Op: "decode", Err: err}
		return nil, err
	}

	if typ == EndType {
		return tag, nil
	}

	if err := tag.TagName().decode(r); err != nil {
		return nil, err
	}

	if err := tag.Payload().decode(r); err != nil {
		return nil, err
	}

	return tag, nil
}

func Stringify(tag Tag) string {
	return prettyStringify(tag, " ", "")
}

func CompactStringify(tag Tag) string {
	return prettyStringify(tag, "", "")
}

func PrettyStringify(tag Tag, indent string) string {
	return prettyStringify(tag, " ", indent)
}

func prettyStringify(tag Tag, space string, indent string) string {
	rootName := ""
	if tag.TagName() != nil {
		rootName = string(*tag.TagName())
	}

	if rootName == "" {
		snbt := tag.stringify(space, indent, 0)
		return strings.TrimLeft(snbt[1:], space)
	}

	if indent == "" {
		return fmt.Sprintf("{%s%s}", indent, tag.stringify(space, indent, 1))
	}

	return fmt.Sprintf("{\n%s%s\n}", indent, tag.stringify(space, indent, 1))
}

func stringifyTag(tag Tag, space string, indent string, depth int) string {
	if tag.TypeId() == EndType {
		return ""
	}

	return fmt.Sprintf("%s:%s%s", tag.TagName().stringify(), space, tag.Payload().stringify(space, indent, depth))
}

func Parse(stringified string) (Tag, error) {
	p := snbt.NewParser(stringified)

	if err := p.Compact(); err != nil {
		err = &NbtError{Op: "parse", Err: err}
		return nil, err
	}

	tag, err := newTagFromSnbt(p)
	if err != nil {
		err = &NbtError{Op: "parse", Err: err}
		return nil, err
	}

	if err := tag.parse(p); err != nil {
		return nil, err
	}

	return tag, nil
}

func parseTag(tag Tag, parser *snbt.Parser) error {
	if tag.TypeId() == EndType {
		return nil
	}

	if err := tag.Payload().parse(parser); err != nil {
		return err
	}

	return nil
}

func Json(tag Tag) string {
	return prettyJson(tag, " ", "")
}

func CompactJson(tag Tag) string {
	return prettyJson(tag, "", "")
}

func PrettyJson(tag Tag, indent string) string {
	return prettyJson(tag, " ", indent)
}

func prettyJson(tag Tag, space string, indent string) string {
	rootName := ""
	if tag.TagName() != nil {
		rootName = string(*tag.TagName())
	}

	if rootName == "" {
		snbt := tag.json(space, indent, 0)
		return strings.TrimLeft(snbt[3:], space)
	}

	if indent == "" {
		return fmt.Sprintf("{%s%s}", indent, tag.json(space, indent, 1))
	}

	return fmt.Sprintf("{\n%s%s\n}", indent, tag.json(space, indent, 1))
}

func jsonTag(tag Tag, space string, indent string, depth int) string {
	if tag.TypeId() == EndType {
		return ""
	}

	return fmt.Sprintf("%s:%s%s", tag.TagName().json(), space, tag.Payload().json(space, indent, depth))
}

// Payload

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
		err := &NbtError{Op: "new", Err: invalidTagTypeError}
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
		err := &NbtError{Op: "new", Err: invalidSnbtFormatError}
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
