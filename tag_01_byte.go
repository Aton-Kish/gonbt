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
	"strconv"

	"github.com/Aton-Kish/gonbt/pointer"
	"github.com/Aton-Kish/gonbt/snbt"
)

type ByteTag struct {
	tagName *TagName
	payload *BytePayload
}

func NewByteTag(tagName *TagName, payload *BytePayload) *ByteTag {
	return &ByteTag{
		tagName: tagName,
		payload: payload,
	}
}

func (t *ByteTag) TypeId() TagType {
	return t.Payload().TypeId()
}

func (t *ByteTag) TagName() *TagName {
	return t.tagName
}

func (t *ByteTag) Payload() Payload {
	return t.payload
}

func (t *ByteTag) encode(w io.Writer) error {
	return Encode(w, t)
}

func (t *ByteTag) decode(r io.Reader) error {
	tag, err := Decode(r)
	if err != nil {
		return err
	}

	v, ok := tag.(*ByteTag)
	if !ok {
		err = &NbtError{Op: "decode", Err: decodeError}
		return err
	}

	*t = *v

	return nil
}

func (t *ByteTag) stringify(space string, indent string, depth int) string {
	return stringifyTag(t, space, indent, depth)
}

func (t *ByteTag) parse(parser *snbt.Parser) error {
	return parseTag(t, parser)
}

func (t *ByteTag) json(space string, indent string, depth int) string {
	return jsonTag(t, space, indent, depth)
}

type BytePayload int8

func NewBytePayload(value int8) *BytePayload {
	return pointer.Pointer(BytePayload(value))
}

func (p *BytePayload) TypeId() TagType {
	return ByteType
}

func (p *BytePayload) encode(w io.Writer) error {
	return encodeNumericPayload(w, p)
}

func (p *BytePayload) decode(r io.Reader) error {
	payload, err := decodeNumericPayload[BytePayload](r)
	if err != nil {
		return err
	}

	*p = *payload

	return nil
}

func (p *BytePayload) stringify(space string, indent string, depth int) string {
	return fmt.Sprintf("%db", *p)
}

func (p *BytePayload) parse(parser *snbt.Parser) error {
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

	*p = *NewBytePayload(int8(i))

	return nil
}

func (p *BytePayload) json(space string, indent string, depth int) string {
	return fmt.Sprintf("%d", *p)
}
