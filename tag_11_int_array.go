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
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/Aton-Kish/gonbt/pointer"
	"github.com/Aton-Kish/gonbt/snbt"
)

type IntArrayTag struct {
	tagName *TagName
	payload *IntArrayPayload
}

func NewIntArrayTag(tagName *TagName, payload *IntArrayPayload) *IntArrayTag {
	return &IntArrayTag{
		tagName: tagName,
		payload: payload,
	}
}

func (t *IntArrayTag) TypeId() TagType {
	return t.Payload().TypeId()
}

func (t *IntArrayTag) TagName() *TagName {
	return t.tagName
}

func (t *IntArrayTag) Payload() Payload {
	return t.payload
}

func (t *IntArrayTag) encode(w io.Writer) error {
	return Encode(w, t)
}

func (t *IntArrayTag) decode(r io.Reader) error {
	tag, err := Decode(r)
	if err != nil {
		return err
	}

	v, ok := tag.(*IntArrayTag)
	if !ok {
		return errors.New("decode failed")
	}

	*t = *v

	return nil
}

func (t *IntArrayTag) stringify(space string, indent string, depth int) string {
	return stringifyTag(t, space, indent, depth)
}

func (t *IntArrayTag) parse(parser *snbt.Parser) error {
	return parseTag(t, parser)
}

func (t *IntArrayTag) json(space string, indent string, depth int) string {
	return jsonTag(t, space, indent, depth)
}

type IntArrayPayload []int32

func NewIntArrayPayload(values ...int32) *IntArrayPayload {
	if values == nil {
		values = []int32{}
	}

	return pointer.Pointer(IntArrayPayload(values))
}

func (p *IntArrayPayload) TypeId() TagType {
	return IntArrayType
}

func (p *IntArrayPayload) encode(w io.Writer) error {
	return encodeArrayPayload(w, p)
}

func (p *IntArrayPayload) decode(r io.Reader) error {
	var l int32
	if err := binary.Read(r, binary.BigEndian, &l); err != nil {
		return err
	}

	*p = make(IntArrayPayload, l)
	if err := binary.Read(r, binary.BigEndian, p); err != nil {
		return err
	}

	return nil
}

func (p *IntArrayPayload) stringify(space string, indent string, depth int) string {
	strs := make([]string, 0, len(*p))
	for _, v := range *p {
		strs = append(strs, fmt.Sprintf("%d", v))
	}

	return fmt.Sprintf("[I;%s%s]", space, strings.Join(strs, fmt.Sprintf(",%s", space)))
}

func (p *IntArrayPayload) parse(parser *snbt.Parser) error {
	if err := parser.Next(); err != nil {
		return err
	}

	if parser.CurrToken().Char() != ';' {
		return errors.New("invalid snbt format")
	}

	for parser.CurrToken().Char() != ']' {
		if err := parser.Next(); err != nil {
			return err
		}

		// NOTE: break if empty
		if parser.CurrToken().Char() == ']' && parser.PrevToken().Char() == ';' &&
			parser.CurrToken().Index()-parser.PrevToken().Index() == 1 {
			break
		}

		b, err := parser.Slice(parser.PrevToken().Index()+1, parser.CurrToken().Index())
		if err != nil {
			return err
		}

		i, err := strconv.ParseInt(string(b), 10, 32)
		if err != nil {
			return err
		}

		*p = append(*p, int32(i))
	}

	if err := parser.Next(); err != nil {
		return err
	}

	return nil
}

func (p *IntArrayPayload) json(space string, indent string, depth int) string {
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
