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
	"strings"

	"github.com/Aton-Kish/gonbt/pointer"
	"github.com/Aton-Kish/gonbt/snbt"
)

type ListTag struct {
	tagName *TagName
	payload *ListPayload
}

func NewListTag(tagName *TagName, payload *ListPayload) *ListTag {
	return &ListTag{
		tagName: tagName,
		payload: payload,
	}
}

func (t *ListTag) TypeId() TagType {
	return t.Payload().TypeId()
}

func (t *ListTag) TagName() *TagName {
	return t.tagName
}

func (t *ListTag) Payload() Payload {
	return t.payload
}

func (t *ListTag) encode(w io.Writer) error {
	return Encode(w, t)
}

func (t *ListTag) decode(r io.Reader) error {
	tag, err := Decode(r)
	if err != nil {
		return err
	}

	v, ok := tag.(*ListTag)
	if !ok {
		return errors.New("decode failed")
	}

	*t = *v

	return nil
}

func (t *ListTag) stringify(space string, indent string, depth int) string {
	return stringifyTag(t, space, indent, depth)
}

func (t *ListTag) parse(parser *snbt.Parser) error {
	return parseTag(t, parser)
}

func (t *ListTag) json(space string, indent string, depth int) string {
	return jsonTag(t, space, indent, depth)
}

type ListPayload []Payload

func NewListPayload(values ...Payload) *ListPayload {
	if values == nil {
		values = []Payload{}
	}

	return pointer.Pointer(ListPayload(values))
}

func (p *ListPayload) TypeId() TagType {
	return ListType
}

func (p *ListPayload) encode(w io.Writer) error {
	l := int32(len(*p))

	typ := EndType
	if l > 0 {
		typ = []Payload(*p)[0].TypeId()
	}

	if err := binary.Write(w, binary.BigEndian, &typ); err != nil {
		return err
	}

	if err := binary.Write(w, binary.BigEndian, &l); err != nil {
		return err
	}

	for _, payload := range *p {
		if err := payload.encode(w); err != nil {
			return err
		}
	}

	return nil
}

func (p *ListPayload) decode(r io.Reader) error {
	var typ TagType
	if err := binary.Read(r, binary.BigEndian, &typ); err != nil {
		return err
	}

	var l int32
	if err := binary.Read(r, binary.BigEndian, &l); err != nil {
		return err
	}

	*p = make([]Payload, 0, l)
	for i := 0; i < int(l); i++ {
		payload, err := NewPayload(typ)
		if err != nil {
			return err
		}

		if err := payload.decode(r); err != nil {
			return err
		}

		*p = append(*p, payload)
	}

	return nil
}

func (p *ListPayload) stringify(space string, indent string, depth int) string {
	l := len(*p)
	strs := make([]string, 0, l)
	for _, payload := range *p {
		strs = append(strs, payload.stringify(space, indent, depth+1))
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

func (p *ListPayload) parse(parser *snbt.Parser) error {
	for parser.CurrToken().Char() != ']' {
		if err := parser.Next(); err != nil {
			return err
		}

		payload, err := newPayloadFromSnbt(parser)
		if err != nil {
			return err
		}

		if err := payload.parse(parser); err != nil {
			return err
		}

		*p = append(*p, payload)
	}

	if err := parser.Next(); err != nil {
		return err
	}

	return nil
}

func (p *ListPayload) json(space string, indent string, depth int) string {
	l := len(*p)
	strs := make([]string, 0, l)
	for _, payload := range *p {
		strs = append(strs, payload.json(space, indent, depth+1))
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
