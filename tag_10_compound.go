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

type CompoundTag struct {
	tagName *TagName
	payload *CompoundPayload
}

func NewCompoundTag(tagName *TagName, payload *CompoundPayload) *CompoundTag {
	return &CompoundTag{
		tagName: tagName,
		payload: payload,
	}
}

func (t *CompoundTag) TypeId() TagType {
	return t.Payload().TypeId()
}

func (t *CompoundTag) TagName() *TagName {
	return t.tagName
}

func (t *CompoundTag) Payload() Payload {
	return t.payload
}

func (t *CompoundTag) encode(w io.Writer) error {
	return Encode(w, t)
}

func (t *CompoundTag) decode(r io.Reader) error {
	tag, err := Decode(r)
	if err != nil {
		return err
	}

	v, ok := tag.(*CompoundTag)
	if !ok {
		return errors.New("decode failed")
	}

	*t = *v

	return nil
}

type CompoundPayload []Tag

func NewCompoundPayload(values ...Tag) *CompoundPayload {
	if values == nil {
		values = []Tag{}
	}

	return pointer.Pointer(CompoundPayload(values))
}

func (p *CompoundPayload) TypeId() TagType {
	return CompoundType
}

func (p *CompoundPayload) encode(w io.Writer) error {
	for _, tag := range *p {
		if err := tag.encode(w); err != nil {
			return err
		}
	}

	return nil
}

func (p *CompoundPayload) decode(r io.Reader) error {
	for {
		tag, err := Decode(r)
		if err != nil {
			return err
		}

		*p = append(*p, tag)

		if tag.TypeId() == EndType {
			break
		}
	}

	return nil
}
