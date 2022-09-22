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
	"io"
)

type StringTag struct {
	tagName TagName
	payload StringPayload
}

func NewStringTag() Tag {
	return new(StringTag)
}

func (t *StringTag) TypeId() TagType {
	return t.Payload().TypeId()
}

func (t *StringTag) TagName() *TagName {
	return &t.tagName
}

func (t *StringTag) Payload() Payload {
	return &t.payload
}

func (t *StringTag) Encode(w io.Writer) error {
	return encodeTag(w, t)
}

func (t *StringTag) Decode(r io.Reader) error {
	tag, err := decodeTag(r)
	if err != nil {
		return err
	}

	v, ok := tag.(*StringTag)
	if !ok {
		return errors.New("decode failed")
	}

	*t = *v

	return nil
}

type StringPayload string

func NewStringPayload() Payload {
	return new(StringPayload)
}

func (p *StringPayload) TypeId() TagType {
	return StringType
}

func (p *StringPayload) Encode(w io.Writer) error {
	l := uint16(len(*p))
	if err := binary.Write(w, binary.BigEndian, &l); err != nil {
		return err
	}

	b := []byte(*p)
	if err := binary.Write(w, binary.BigEndian, b); err != nil {
		return err
	}

	return nil
}

func (p *StringPayload) Decode(r io.Reader) error {
	var l uint16
	if err := binary.Read(r, binary.BigEndian, &l); err != nil {
		return err
	}

	b := make([]byte, l)
	if err := binary.Read(r, binary.BigEndian, b); err != nil {
		return err
	}
	*p = StringPayload(b)

	return nil
}