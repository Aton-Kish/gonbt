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

	"github.com/Aton-Kish/gonbt/pointer"
)

type LongArrayTag struct {
	tagName *TagName
	payload *LongArrayPayload
}

func NewLongArrayTag(tagName *TagName, payload *LongArrayPayload) *LongArrayTag {
	return &LongArrayTag{
		tagName: tagName,
		payload: payload,
	}
}

func (t *LongArrayTag) TypeId() TagType {
	return t.Payload().TypeId()
}

func (t *LongArrayTag) TagName() *TagName {
	return t.tagName
}

func (t *LongArrayTag) Payload() Payload {
	return t.payload
}

func (t *LongArrayTag) encode(w io.Writer) error {
	return Encode(w, t)
}

func (t *LongArrayTag) decode(r io.Reader) error {
	tag, err := Decode(r)
	if err != nil {
		return err
	}

	v, ok := tag.(*LongArrayTag)
	if !ok {
		return errors.New("decode failed")
	}

	*t = *v

	return nil
}

type LongArrayPayload []int64

func NewLongArrayPayload(values ...int64) *LongArrayPayload {
	if values == nil {
		values = []int64{}
	}

	return pointer.Pointer(LongArrayPayload(values))
}

func (p *LongArrayPayload) TypeId() TagType {
	return LongArrayType
}

func (p *LongArrayPayload) encode(w io.Writer) error {
	return encodeArrayPayload(w, p)
}

func (p *LongArrayPayload) decode(r io.Reader) error {
	var l int32
	if err := binary.Read(r, binary.BigEndian, &l); err != nil {
		return err
	}

	*p = make(LongArrayPayload, l)
	if err := binary.Read(r, binary.BigEndian, p); err != nil {
		return err
	}

	return nil
}
