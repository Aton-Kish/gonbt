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
	"errors"
	"io"

	"github.com/Aton-Kish/gonbt/pointer"
)

type DoubleTag struct {
	tagName TagName
	payload DoublePayload
}

func NewDoubleTag(tagName TagName, payload DoublePayload) Tag {
	return &DoubleTag{
		tagName: tagName,
		payload: payload,
	}
}

func (t *DoubleTag) TypeId() TagType {
	return t.Payload().TypeId()
}

func (t *DoubleTag) TagName() *TagName {
	return &t.tagName
}

func (t *DoubleTag) Payload() Payload {
	return &t.payload
}

func (t *DoubleTag) Encode(w io.Writer) error {
	return Encode(w, t)
}

func (t *DoubleTag) Decode(r io.Reader) error {
	tag, err := Decode(r)
	if err != nil {
		return err
	}

	v, ok := tag.(*DoubleTag)
	if !ok {
		return errors.New("decode failed")
	}

	*t = *v

	return nil
}

type DoublePayload float64

func NewDoublePayload(value float64) Payload {
	return pointer.Pointer(DoublePayload(value))
}

func (p *DoublePayload) TypeId() TagType {
	return DoubleType
}

func (p *DoublePayload) Encode(w io.Writer) error {
	return encodeNumericPayload(w, p)
}

func (p *DoublePayload) Decode(r io.Reader) error {
	payload, err := decodeNumericPayload[DoublePayload](r)
	if err != nil {
		return err
	}

	*p = *payload

	return nil
}
