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
	"strconv"
	"strings"

	"github.com/Aton-Kish/gonbt/pointer"
	"github.com/Aton-Kish/gonbt/snbt"
)

type ByteArrayTag struct {
	tagName *TagName
	payload *ByteArrayPayload
}

func NewByteArrayTag(tagName *TagName, payload *ByteArrayPayload) *ByteArrayTag {
	return &ByteArrayTag{
		tagName: tagName,
		payload: payload,
	}
}

func (t *ByteArrayTag) TypeId() TagType {
	return t.Payload().TypeId()
}

func (t *ByteArrayTag) TagName() *TagName {
	return t.tagName
}

func (t *ByteArrayTag) Payload() Payload {
	return t.payload
}

func (t *ByteArrayTag) encode(w io.Writer) error {
	return Encode(w, t)
}

func (t *ByteArrayTag) decode(r io.Reader) error {
	tag, err := Decode(r)
	if err != nil {
		return err
	}

	v, ok := tag.(*ByteArrayTag)
	if !ok {
		err = &NbtError{Op: "decode", Err: decodeError}
		return err
	}

	*t = *v

	return nil
}

func (t *ByteArrayTag) stringify(space string, indent string, depth int) string {
	return stringifyTag(t, space, indent, depth)
}

func (t *ByteArrayTag) parse(parser *snbt.Parser) error {
	return parseTag(t, parser)
}

func (t *ByteArrayTag) json(space string, indent string, depth int) string {
	return jsonTag(t, space, indent, depth)
}

type ByteArrayPayload []int8

func NewByteArrayPayload(values ...int8) *ByteArrayPayload {
	if values == nil {
		values = []int8{}
	}

	return pointer.Pointer(ByteArrayPayload(values))
}

func (p *ByteArrayPayload) TypeId() TagType {
	return ByteArrayType
}

func (p *ByteArrayPayload) encode(w io.Writer) error {
	return encodeArrayPayload(w, p)
}

func (p *ByteArrayPayload) decode(r io.Reader) error {
	var l int32
	if err := binary.Read(r, binary.BigEndian, &l); err != nil {
		err = &NbtError{Op: "decode", Err: err}
		return err
	}

	*p = make(ByteArrayPayload, l)
	if err := binary.Read(r, binary.BigEndian, p); err != nil {
		err = &NbtError{Op: "decode", Err: err}
		return err
	}

	return nil
}

func (p *ByteArrayPayload) stringify(space string, indent string, depth int) string {
	strs := make([]string, 0, len(*p))
	for _, v := range *p {
		strs = append(strs, fmt.Sprintf("%db", v))
	}

	return fmt.Sprintf("[B;%s%s]", space, strings.Join(strs, fmt.Sprintf(",%s", space)))
}

func (p *ByteArrayPayload) parse(parser *snbt.Parser) error {
	if err := parser.Next(); err != nil {
		err = &NbtError{Op: "parse", Err: err}
		return err
	}

	if parser.CurrToken().Char() != ';' {
		err := &NbtError{Op: "parse", Err: invalidSnbtFormatError}
		return err
	}

	for parser.CurrToken().Char() != ']' {
		if err := parser.Next(); err != nil {
			err = &NbtError{Op: "parse", Err: err}
			return err
		}

		// NOTE: break if empty
		if parser.CurrToken().Char() == ']' && parser.PrevToken().Char() == ';' &&
			parser.CurrToken().Index()-parser.PrevToken().Index() == 1 {
			break
		}

		b, err := parser.Slice(parser.PrevToken().Index()+1, parser.CurrToken().Index())
		if err != nil {
			err = &NbtError{Op: "parse", Err: err}
			return err
		}

		g := bytePattern.FindSubmatch(b)
		if len(g) < 2 {
			err = &NbtError{Op: "parse", Err: invalidSnbtFormatError}
			return err
		}

		i, err := strconv.ParseInt(string(g[1]), 10, 8)
		if err != nil {
			err = &NbtError{Op: "parse", Err: err}
			return err
		}

		*p = append(*p, int8(i))
	}

	if err := parser.Next(); err != nil {
		err = &NbtError{Op: "parse", Err: err}
		return err
	}

	return nil
}

func (p *ByteArrayPayload) json(space string, indent string, depth int) string {
	l := len(*p)
	strs := make([]string, 0, l)
	for _, v := range *p {
		strs = append(strs, fmt.Sprintf("%d", v))
	}

	if indent == "" || l == 0 {
		return fmt.Sprintf("[%s]", strings.Join(strs, fmt.Sprintf(",%s", space)))
	}

	indents := ""
	for i := 0; i < depth; i++ {
		indents += indent
	}

	return fmt.Sprintf("[\n%s%s%s\n%s]", indents, indent, strings.Join(strs, fmt.Sprintf(",\n%s%s", indents, indent)), indents)
}
