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
	"regexp"
	"strconv"
	"strings"

	"github.com/Aton-Kish/gonbt/pointer"
	"github.com/Aton-Kish/gonbt/snbt"
)

var quoteRequiredCharacters = regexp.MustCompile(`[ !"#$%&'()*,/:;<=>?@[\\\]^\x60{|}~]`)

type TagName string

func NewTagName(value string) *TagName {
	return pointer.Pointer(TagName(value))
}

func (n *TagName) encode(w io.Writer) error {
	l := uint16(len(*n))
	if err := binary.Write(w, binary.BigEndian, &l); err != nil {
		err = &NbtError{Op: "encode", Err: err}
		return err
	}

	b := []byte(*n)
	if err := binary.Write(w, binary.BigEndian, b); err != nil {
		err = &NbtError{Op: "encode", Err: err}
		return err
	}

	return nil
}

func (n *TagName) decode(r io.Reader) error {
	var l uint16
	if err := binary.Read(r, binary.BigEndian, &l); err != nil {
		err = &NbtError{Op: "decode", Err: err}
		return err
	}

	b := make([]byte, l)
	if err := binary.Read(r, binary.BigEndian, b); err != nil {
		err = &NbtError{Op: "decode", Err: err}
		return err
	}
	*n = TagName(b)

	return nil
}

func (n *TagName) stringify() string {
	s := string(*n)
	if !quoteRequiredCharacters.MatchString(s) {
		return s
	}

	qs := strconv.Quote(s)
	if strings.Contains(s, "\"") && !strings.Contains(s, "'") {
		qs = fmt.Sprintf("'%s'", qs[1:len(qs)-1])
		qs = strings.ReplaceAll(qs, "\\\"", "\"")
		return qs
	}

	return qs
}

func (n *TagName) parse(parser *snbt.Parser) error {
	b, err := parser.Slice(parser.PrevToken().Index()+1, parser.CurrToken().Index())
	if err != nil {
		err = &NbtError{Op: "parse", Err: err}
		return err
	}

	if s, err := strconv.Unquote(string(b)); err != nil {
		*n = TagName(b)
	} else {
		*n = TagName(s)
	}

	if err := parser.Next(); err != nil {
		err = &NbtError{Op: "parse", Err: err}
		return err
	}

	return nil
}

func (n *TagName) json() string {
	s := string(*n)
	return strconv.Quote(s)
}
