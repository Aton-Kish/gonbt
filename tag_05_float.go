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
	"fmt"
	"io"

	"github.com/Aton-Kish/gonbt/pointer"
)

type FloatTag struct {
	tagName *TagName
	payload *FloatPayload
}

func NewFloatTag(tagName *TagName, payload *FloatPayload) *FloatTag {
	return &FloatTag{
		tagName: tagName,
		payload: payload,
	}
}
func (t *FloatTag) TypeId() TagType {
	return t.Payload().TypeId()
}

func (t *FloatTag) TagName() *TagName {
	return t.tagName
}

func (t *FloatTag) Payload() Payload {
	return t.payload
}

func (t *FloatTag) encode(w io.Writer) error {
	return Encode(w, t)
}

func (t *FloatTag) decode(r io.Reader) error {
	tag, err := Decode(r)
	if err != nil {
		return err
	}

	v, ok := tag.(*FloatTag)
	if !ok {
		return errors.New("decode failed")
	}

	*t = *v

	return nil
}

func (t *FloatTag) stringify(space string, indent string, depth int) string {
	return stringifyTag(t, space, indent, depth)
}

func (t *FloatTag) json(space string, indent string, depth int) string {
	return jsonTag(t, space, indent, depth)
}

type FloatPayload float32

func NewFloatPayload(value float32) *FloatPayload {
	return pointer.Pointer(FloatPayload(value))
}

func (p *FloatPayload) TypeId() TagType {
	return FloatType
}

func (p *FloatPayload) encode(w io.Writer) error {
	return encodeNumericPayload(w, p)
}

func (p *FloatPayload) decode(r io.Reader) error {
	payload, err := decodeNumericPayload[FloatPayload](r)
	if err != nil {
		return err
	}

	*p = *payload

	return nil
}

func (p *FloatPayload) stringify(space string, indent string, depth int) string {
	return fmt.Sprintf("%gf", *p)
}

func (p *FloatPayload) json(space string, indent string, depth int) string {
	return fmt.Sprintf("%g", *p)
}
