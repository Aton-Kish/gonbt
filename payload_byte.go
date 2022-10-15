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

type BytePayload int8

func NewBytePayload(value int8) *BytePayload {
	return pointer.Pointer(BytePayload(value))
}

func (p BytePayload) String() string {
	return p.stringify(" ", "", 0)
}

func (p *BytePayload) TypeId() TagType {
	return TagTypeByte
}

func (p *BytePayload) encode(w io.Writer) error {
	if err := binary.Write(w, binary.BigEndian, p); err != nil {
		err = &NbtError{Op: "encode", Err: err}
		return err
	}

	return nil
}

func (p *BytePayload) decode(r io.Reader) error {
	payload := new(BytePayload)
	if err := binary.Read(r, binary.BigEndian, payload); err != nil {
		err = &NbtError{Op: "decode", Err: err}
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
		err = &NbtError{Op: "parse", Err: ErrInvalidSnbtFormat}
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
