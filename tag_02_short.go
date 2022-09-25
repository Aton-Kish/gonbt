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
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/Aton-Kish/gonbt/pointer"
	"github.com/Aton-Kish/gonbt/snbt"
)

type ShortTag struct {
	tagName *TagName
	payload *ShortPayload
}

func NewShortTag(tagName *TagName, payload *ShortPayload) *ShortTag {
	return &ShortTag{
		tagName: tagName,
		payload: payload,
	}
}

func (t *ShortTag) TypeId() TagType {
	return t.Payload().TypeId()
}

func (t *ShortTag) TagName() *TagName {
	return t.tagName
}

func (t *ShortTag) Payload() Payload {
	return t.payload
}

func (t *ShortTag) encode(w io.Writer) error {
	return Encode(w, t)
}

func (t *ShortTag) decode(r io.Reader) error {
	tag, err := Decode(r)
	if err != nil {
		return err
	}

	v, ok := tag.(*ShortTag)
	if !ok {
		return errors.New("decode failed")
	}

	*t = *v

	return nil
}

func (t *ShortTag) stringify(space string, indent string, depth int) string {
	return stringifyTag(t, space, indent, depth)
}

func (t *ShortTag) parse(parser *snbt.Parser) error {
	return parseTag(t, parser)
}

func (t *ShortTag) json(space string, indent string, depth int) string {
	return jsonTag(t, space, indent, depth)
}

type ShortPayload int16

func NewShortPayload(value int16) *ShortPayload {
	return pointer.Pointer(ShortPayload(value))
}

func (p *ShortPayload) TypeId() TagType {
	return ShortType
}

func (p *ShortPayload) encode(w io.Writer) error {
	return encodeNumericPayload(w, p)
}

func (p *ShortPayload) decode(r io.Reader) error {
	payload, err := decodeNumericPayload[ShortPayload](r)
	if err != nil {
		return err
	}

	*p = *payload

	return nil
}

func (p *ShortPayload) stringify(space string, indent string, depth int) string {
	return fmt.Sprintf("%ds", *p)
}

func (p *ShortPayload) parse(parser *snbt.Parser) error {
	b, err := parser.Slice(parser.PrevToken().Index()+1, parser.CurrToken().Index())
	if err != nil {
		return err
	}

	g := shortPattern.FindSubmatch(b)
	if len(g) < 2 {
		return errors.New("invalid snbt format")
	}

	i, err := strconv.ParseInt(string(g[1]), 10, 16)
	if err != nil {
		return err
	}

	*p = *NewShortPayload(int16(i))

	return nil
}

func (p *ShortPayload) json(space string, indent string, depth int) string {
	return fmt.Sprintf("%d", *p)
}
