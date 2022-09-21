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

type ShortTag struct {
	tagName TagName
	payload ShortPayload
}

func NewShortTag() Tag {
	return new(ShortTag)
}

func (t *ShortTag) TypeId() TagType {
	return t.Payload().TypeId()
}

func (t *ShortTag) TagName() *TagName {
	return &t.tagName
}

func (t *ShortTag) Payload() Payload {
	return &t.payload
}

func (t *ShortTag) Encode(w io.Writer) error {
	return encodeTagExcludeEndTag(w, t)
}

func (t *ShortTag) Decode(r io.Reader) error {
	return decodeTagExcludeEndTag(r, t)
}

type ShortPayload int16

func NewShortPayload() Payload {
	return new(ShortPayload)
}

func (p *ShortPayload) TypeId() TagType {
	return ShortType
}

func (p *ShortPayload) Encode(w io.Writer) error {
	return encodeNumericPayload(w, p)
}

func (p *ShortPayload) Decode(r io.Reader) error {
	return decodeNumericPayload(r, p)
}
