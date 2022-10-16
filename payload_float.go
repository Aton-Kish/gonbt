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

	"github.com/Aton-Kish/gonbt/pointer"
	"github.com/Aton-Kish/gonbt/snbt"
)

type FloatPayload float32

func NewFloatPayload(value float32) *FloatPayload {
	return pointer.Pointer(FloatPayload(value))
}

func (p FloatPayload) String() string {
	return p.stringify(" ", "", 0)
}

func (p *FloatPayload) TypeId() TagType {
	return TagTypeFloat
}

func (p *FloatPayload) encode(w io.Writer) error {
	if err := binary.Write(w, binary.BigEndian, p); err != nil {
		err = &NbtError{Op: "encode", Err: err}
		logger.Printf("(*FloatPayload).encode; payload: %s; error: %s", p, err)
		return err
	}

	return nil
}

func (p *FloatPayload) decode(r io.Reader) error {
	payload := new(FloatPayload)
	if err := binary.Read(r, binary.BigEndian, payload); err != nil {
		err = &NbtError{Op: "decode", Err: err}
		logger.Printf("(*FloatPayload).decode; payload: %s; error: %s", p, err)
		return err
	}

	*p = *payload

	return nil
}

func (p *FloatPayload) stringify(space string, indent string, depth int) string {
	return fmt.Sprintf("%gf", *p)
}

func (p *FloatPayload) parse(parser *snbt.Parser) error {
	b, err := parser.Slice(parser.PrevToken().Index()+1, parser.CurrToken().Index())
	if err != nil {
		err = &NbtError{Op: "parse", Err: err}
		logger.Printf("(*FloatPayload).parse; payload: %s; error: %s", p, err)
		return err
	}

	g := floatPattern.FindSubmatch(b)
	if len(g) < 2 {
		err = &NbtError{Op: "parse", Err: ErrInvalidSnbtFormat}
		logger.Printf("(*FloatPayload).parse; payload: %s; error: %s", p, err)
		return err
	}

	f, err := strconv.ParseFloat(string(g[1]), 32)
	if err != nil {
		err = &NbtError{Op: "parse", Err: err}
		logger.Printf("(*FloatPayload).parse; payload: %s; error: %s", p, err)
		return err
	}

	*p = *NewFloatPayload(float32(f))

	return nil
}

func (p *FloatPayload) json(space string, indent string, depth int) string {
	return fmt.Sprintf("%g", *p)
}
