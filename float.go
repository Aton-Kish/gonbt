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

type FloatTag struct {
	TagName
	FloatPayload
}

func NewFloatTag() Tag {
	return new(FloatTag)
}

func (t *FloatTag) TypeId() TagType {
	return t.FloatPayload.TypeId()
}

func (t *FloatTag) Encode(w io.Writer) error {
	typ := t.TypeId()
	if err := binary.Write(w, binary.BigEndian, &typ); err != nil {
		return err
	}

	if err := t.TagName.Encode(w); err != nil {
		return err
	}

	if err := t.FloatPayload.Encode(w); err != nil {
		return err
	}

	return nil
}

type FloatPayload float32

func NewFloatPayload() Payload {
	return new(FloatPayload)
}

func (p *FloatPayload) TypeId() TagType {
	return FloatType
}

func (p *FloatPayload) Encode(w io.Writer) error {
	if err := binary.Write(w, binary.BigEndian, p); err != nil {
		return err
	}

	return nil
}
