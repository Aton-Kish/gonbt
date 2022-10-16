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
	"fmt"
	"io"
	"strings"

	"github.com/Aton-Kish/gonbt/snbt"
)

type Tag interface {
	String() string
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
	case TagTypeEnd:
		return NewEndTag(), nil
	case TagTypeByte:
		return NewByteTag(new(TagName), new(BytePayload)), nil
	case TagTypeShort:
		return NewShortTag(new(TagName), new(ShortPayload)), nil
	case TagTypeInt:
		return NewIntTag(new(TagName), new(IntPayload)), nil
	case TagTypeLong:
		return NewLongTag(new(TagName), new(LongPayload)), nil
	case TagTypeFloat:
		return NewFloatTag(new(TagName), new(FloatPayload)), nil
	case TagTypeDouble:
		return NewDoubleTag(new(TagName), new(DoublePayload)), nil
	case TagTypeByteArray:
		return NewByteArrayTag(new(TagName), new(ByteArrayPayload)), nil
	case TagTypeString:
		return NewStringTag(new(TagName), new(StringPayload)), nil
	case TagTypeList:
		return NewListTag(new(TagName), new(ListPayload)), nil
	case TagTypeCompound:
		return NewCompoundTag(new(TagName), new(CompoundPayload)), nil
	case TagTypeIntArray:
		return NewIntArrayTag(new(TagName), new(IntArrayPayload)), nil
	case TagTypeLongArray:
		return NewLongArrayTag(new(TagName), new(LongArrayPayload)), nil
	default:
		err := &NbtError{Op: "new", Err: ErrInvalidTagType}
		logger.Println("failed to new", "func", getFuncName(), "error", err)
		return nil, err
	}
}

func newTagFromSnbt(parser *snbt.Parser) (Tag, error) {
	var name TagName
	// NOTE: skip if nameless root
	if parser.CurrToken().Index() > 0 {
		if err := name.parse(parser); err != nil {
			err = &NbtError{Op: "new", Err: err}
			logger.Println("failed to new", "func", getFuncName(), "error", err)
			return nil, err
		}
	}

	p, err := newPayloadFromSnbt(parser)
	if err != nil {
		logger.Println("failed to new", "func", getFuncName(), "error", err)
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
		err = &NbtError{Op: "new", Err: ErrInvalidSnbtFormat}
		logger.Println("failed to new", "func", getFuncName(), "error", err)
		return nil, err
	}
}

func Encode(w io.Writer, tag Tag) error {
	typ := tag.TypeId()
	if err := typ.encode(w); err != nil {
		logger.Println("failed to encode", "func", getFuncName(), "error", err)
		return err
	}

	if typ == TagTypeEnd {
		return nil
	}

	if err := tag.TagName().encode(w); err != nil {
		logger.Println("failed to encode", "func", getFuncName(), "error", err)
		return err
	}

	if err := tag.Payload().encode(w); err != nil {
		logger.Println("failed to encode", "func", getFuncName(), "error", err)
		return err
	}

	return nil
}

func Decode(r io.Reader) (Tag, error) {
	var typ TagType
	if err := typ.decode(r); err != nil {
		logger.Println("failed to decode", "func", getFuncName(), "error", err)
		return nil, err
	}

	tag, err := NewTag(typ)
	if err != nil {
		err = &NbtError{Op: "decode", Err: err}
		logger.Println("failed to decode", "func", getFuncName(), "error", err)
		return nil, err
	}

	if typ == TagTypeEnd {
		return tag, nil
	}

	if err := tag.TagName().decode(r); err != nil {
		logger.Println("failed to decode", "func", getFuncName(), "error", err)
		return nil, err
	}

	if err := tag.Payload().decode(r); err != nil {
		logger.Println("failed to decode", "func", getFuncName(), "error", err)
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
	if tag.TypeId() == TagTypeEnd {
		return ""
	}

	return fmt.Sprintf("%s:%s%s", tag.TagName().stringify(), space, tag.Payload().stringify(space, indent, depth))
}

func Parse(stringified string) (Tag, error) {
	p := snbt.NewParser(stringified)

	if err := p.Compact(); err != nil {
		err = &NbtError{Op: "parse", Err: err}
		logger.Println("failed to parse", "func", getFuncName(), "error", err)
		return nil, err
	}

	tag, err := newTagFromSnbt(p)
	if err != nil {
		err = &NbtError{Op: "parse", Err: err}
		logger.Println("failed to parse", "func", getFuncName(), "error", err)
		return nil, err
	}

	if err := tag.parse(p); err != nil {
		logger.Println("failed to parse", "func", getFuncName(), "error", err)
		return nil, err
	}

	return tag, nil
}

func parseTag(tag Tag, parser *snbt.Parser) error {
	if tag.TypeId() == TagTypeEnd {
		return nil
	}

	if err := tag.Payload().parse(parser); err != nil {
		logger.Println("failed to parse", "func", getFuncName(), "error", err)
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
	if tag.TypeId() == TagTypeEnd {
		return ""
	}

	return fmt.Sprintf("%s:%s%s", tag.TagName().json(), space, tag.Payload().json(space, indent, depth))
}
