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
	"strings"

	"github.com/Aton-Kish/gonbt/pointer"
	"github.com/Aton-Kish/gonbt/snbt"
)

type ListPayload []Payload

func NewListPayload(values ...Payload) *ListPayload {
	if values == nil {
		values = []Payload{}
	}

	return pointer.Pointer(ListPayload(values))
}

func (p ListPayload) String() string {
	return p.stringify(" ", "", 0)
}

func (p *ListPayload) TypeId() TagType {
	return TagTypeList
}

func (p *ListPayload) encode(w io.Writer) error {
	l := int32(len(*p))

	typ := TagTypeEnd
	if l > 0 {
		typ = []Payload(*p)[0].TypeId()
	}

	if err := binary.Write(w, binary.BigEndian, &typ); err != nil {
		err = &NbtError{Op: "encode", Err: err}
		logger.Println("failed to encode", "func", getFuncName(), "payload", p, "error", err)
		return err
	}

	if err := binary.Write(w, binary.BigEndian, &l); err != nil {
		err = &NbtError{Op: "encode", Err: err}
		logger.Println("failed to encode", "func", getFuncName(), "payload", p, "error", err)
		return err
	}

	for _, payload := range *p {
		if err := payload.encode(w); err != nil {
			logger.Println("failed to encode", "func", getFuncName(), "payload", p, "error", err)
			return err
		}
	}

	return nil
}

func (p *ListPayload) decode(r io.Reader) error {
	var typ TagType
	if err := binary.Read(r, binary.BigEndian, &typ); err != nil {
		err = &NbtError{Op: "decode", Err: err}
		logger.Println("failed to decode", "func", getFuncName(), "payload", p, "error", err)
		return err
	}

	var l int32
	if err := binary.Read(r, binary.BigEndian, &l); err != nil {
		err = &NbtError{Op: "decode", Err: err}
		logger.Println("failed to decode", "func", getFuncName(), "payload", p, "error", err)
		return err
	}

	*p = make([]Payload, 0, l)
	for i := 0; i < int(l); i++ {
		payload, err := NewPayload(typ)
		if err != nil {
			err = &NbtError{Op: "decode", Err: err}
			logger.Println("failed to decode", "func", getFuncName(), "payload", p, "error", err)
			return err
		}

		if err := payload.decode(r); err != nil {
			logger.Println("failed to decode", "func", getFuncName(), "payload", p, "error", err)
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
			err = &NbtError{Op: "parse", Err: err}
			logger.Println("failed to parse", "func", getFuncName(), "payload", p, "error", err)
			return err
		}

		// NOTE: break if empty
		if parser.CurrToken().Char() == ']' && parser.PrevToken().Char() == '[' &&
			parser.CurrToken().Index()-parser.PrevToken().Index() == 1 {
			break
		}

		payload, err := newPayloadFromSnbt(parser)
		if err != nil {
			err = &NbtError{Op: "parse", Err: err}
			logger.Println("failed to parse", "func", getFuncName(), "payload", p, "error", err)
			return err
		}

		if err := payload.parse(parser); err != nil {
			logger.Println("failed to parse", "func", getFuncName(), "payload", p, "error", err)
			return err
		}

		*p = append(*p, payload)
	}

	if err := parser.Next(); err != nil {
		err = &NbtError{Op: "parse", Err: err}
		logger.Println("failed to parse", "func", getFuncName(), "payload", p, "error", err)
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
