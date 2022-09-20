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
	"io"
)

type CompoundTag struct {
	tagName TagName
	payload CompoundPayload
}

func NewCompoundTag() Tag {
	return new(CompoundTag)
}

func (t *CompoundTag) TypeId() TagType {
	return t.Payload().TypeId()
}

func (t *CompoundTag) TagName() *TagName {
	return &t.tagName
}

func (t *CompoundTag) Payload() Payload {
	return &t.payload
}

func (t *CompoundTag) Encode(w io.Writer) error {
	return encodeTagExcludeEndTag(w, t)
}

type CompoundPayload []Tag

func NewCompoundPayload() Payload {
	return new(CompoundPayload)
}

func (p *CompoundPayload) TypeId() TagType {
	return CompoundType
}

func (p *CompoundPayload) Encode(w io.Writer) error {
	for _, tag := range *p {
		if err := tag.Encode(w); err != nil {
			return err
		}
	}

	return nil
}
