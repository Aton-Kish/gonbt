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

type ListTag struct {
	tagName TagName
	payload ListPayload
}

func NewListTag() Tag {
	return new(ListTag)
}

func (t *ListTag) TypeId() TagType {
	return t.Payload().TypeId()
}

func (t *ListTag) TagName() *TagName {
	return &t.tagName
}

func (t *ListTag) Payload() Payload {
	return &t.payload
}

func (t *ListTag) Encode(w io.Writer) error {
	return Encode(w, t)
}

func (t *ListTag) Decode(r io.Reader) error {
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

type ListPayload []Payload

func NewListPayload() Payload {
	return new(ListPayload)
}

func (p *ListPayload) TypeId() TagType {
	return ListType
}

func (p *ListPayload) Encode(w io.Writer) error {
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
		if err := payload.Encode(w); err != nil {
			return err
		}
	}

	return nil
}

func (p *ListPayload) Decode(r io.Reader) error {
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

		if err := payload.Decode(r); err != nil {
			return err
		}

		*p = append(*p, payload)
	}

	return nil
}
