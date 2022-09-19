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
	"io"
)

type ByteArrayTag struct {
	TagName
	ByteArrayPayload
}

func NewByteArrayTag() Tag {
	return new(ByteArrayTag)
}

func (t *ByteArrayTag) TypeId() TagType {
	return t.ByteArrayPayload.TypeId()
}

func (t *ByteArrayTag) Encode(w io.Writer) error {
	typ := t.TypeId()
	if err := binary.Write(w, binary.BigEndian, &typ); err != nil {
		return err
	}

	if err := t.TagName.Encode(w); err != nil {
		return err
	}

	if err := t.ByteArrayPayload.Encode(w); err != nil {
		return err
	}

	return nil
}

type ByteArrayPayload []int8

func NewByteArrayPayload() Payload {
	return new(ByteArrayPayload)
}

func (p *ByteArrayPayload) TypeId() TagType {
	return ByteArrayType
}

func (p *ByteArrayPayload) Encode(w io.Writer) error {
	l := int32(len(*p))
	if err := binary.Write(w, binary.BigEndian, &l); err != nil {
		return err
	}

	if err := binary.Write(w, binary.BigEndian, p); err != nil {
		return err
	}

	return nil
}
