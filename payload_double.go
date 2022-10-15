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

type DoublePayload float64

func NewDoublePayload(value float64) *DoublePayload {
	return pointer.Pointer(DoublePayload(value))
}

func (p *DoublePayload) TypeId() TagType {
	return TagTypeDouble
}

func (p *DoublePayload) encode(w io.Writer) error {
	return encodeNumericPayload(w, p)
}

func (p *DoublePayload) decode(r io.Reader) error {
	payload, err := decodeNumericPayload[DoublePayload](r)
	if err != nil {
		return err
	}

	*p = *payload

	return nil
}

func (p *DoublePayload) stringify(space string, indent string, depth int) string {
	return fmt.Sprintf("%gd", *p)
}

func (p *DoublePayload) parse(parser *snbt.Parser) error {
	b, err := parser.Slice(parser.PrevToken().Index()+1, parser.CurrToken().Index())
	if err != nil {
		err = &NbtError{Op: "parse", Err: err}
		return err
	}

	g := doublePattern.FindSubmatch(b)
	if len(g) < 2 {
		err = &NbtError{Op: "parse", Err: ErrInvalidSnbtFormat}
		return err
	}

	f, err := strconv.ParseFloat(string(g[1]), 64)
	if err != nil {
		err = &NbtError{Op: "parse", Err: err}
		return err
	}

	*p = *NewDoublePayload(f)

	return nil
}

func (p *DoublePayload) json(space string, indent string, depth int) string {
	return fmt.Sprintf("%g", *p)
}
