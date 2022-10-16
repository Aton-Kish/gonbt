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

type LongArrayPayload []int64

func NewLongArrayPayload(values ...int64) *LongArrayPayload {
	if values == nil {
		values = []int64{}
	}

	return pointer.Pointer(LongArrayPayload(values))
}

func (p LongArrayPayload) String() string {
	return p.stringify(" ", "", 0)
}

func (p *LongArrayPayload) TypeId() TagType {
	return TagTypeLongArray
}

func (p *LongArrayPayload) encode(w io.Writer) error {
	l := int32(len(*p))
	if err := binary.Write(w, binary.BigEndian, &l); err != nil {
		err = &NbtError{Op: "encode", Err: err}
		logger.Printf("(*LongArrayPayload).encode; payload: %s; error: %s", p, err)
		return err
	}

	if err := binary.Write(w, binary.BigEndian, p); err != nil {
		err = &NbtError{Op: "encode", Err: err}
		logger.Printf("(*LongArrayPayload).encode; payload: %s; error: %s", p, err)
		return err
	}

	return nil
}

func (p *LongArrayPayload) decode(r io.Reader) error {
	var l int32
	if err := binary.Read(r, binary.BigEndian, &l); err != nil {
		err = &NbtError{Op: "decode", Err: err}
		logger.Printf("(*LongArrayPayload).decode; payload: %s; error: %s", p, err)
		return err
	}

	*p = make(LongArrayPayload, l)
	if err := binary.Read(r, binary.BigEndian, p); err != nil {
		err = &NbtError{Op: "decode", Err: err}
		logger.Printf("(*LongArrayPayload).decode; payload: %s; error: %s", p, err)
		return err
	}

	return nil
}

func (p *LongArrayPayload) stringify(space string, indent string, depth int) string {
	strs := make([]string, 0, len(*p))
	for _, v := range *p {
		strs = append(strs, fmt.Sprintf("%dL", v))
	}

	return fmt.Sprintf("[L;%s%s]", space, strings.Join(strs, fmt.Sprintf(",%s", space)))
}

func (p *LongArrayPayload) parse(parser *snbt.Parser) error {
	if err := parser.Next(); err != nil {
		err = &NbtError{Op: "parse", Err: err}
		logger.Printf("(*LongArrayPayload).parse; payload: %s; error: %s", p, err)
		return err
	}

	if parser.CurrToken().Char() != ';' {
		err := &NbtError{Op: "parse", Err: ErrInvalidSnbtFormat}
		logger.Printf("(*LongArrayPayload).parse; payload: %s; error: %s", p, err)
		return err
	}

	for parser.CurrToken().Char() != ']' {
		if err := parser.Next(); err != nil {
			err = &NbtError{Op: "parse", Err: err}
			logger.Printf("(*LongArrayPayload).parse; payload: %s; error: %s", p, err)
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
			logger.Printf("(*LongArrayPayload).parse; payload: %s; error: %s", p, err)
			return err
		}

		g := longPattern.FindSubmatch(b)
		if len(g) < 2 {
			err = &NbtError{Op: "parse", Err: ErrInvalidSnbtFormat}
			logger.Printf("(*LongArrayPayload).parse; payload: %s; error: %s", p, err)
			return err
		}

		i, err := strconv.ParseInt(string(g[1]), 10, 8)
		if err != nil {
			err = &NbtError{Op: "parse", Err: err}
			logger.Printf("(*LongArrayPayload).parse; payload: %s; error: %s", p, err)
			return err
		}

		*p = append(*p, int64(i))
	}

	if err := parser.Next(); err != nil {
		err = &NbtError{Op: "parse", Err: err}
		logger.Printf("(*LongArrayPayload).parse; payload: %s; error: %s", p, err)
		return err
	}

	return nil
}

func (p *LongArrayPayload) json(space string, indent string, depth int) string {
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
