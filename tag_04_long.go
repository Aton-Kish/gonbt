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

type LongTag struct {
	tagName TagName
	payload LongPayload
}

func NewLongTag() Tag {
	return new(LongTag)
}

func (t *LongTag) TypeId() TagType {
	return t.Payload().TypeId()
}

func (t *LongTag) TagName() *TagName {
	return &t.tagName
}

func (t *LongTag) Payload() Payload {
	return &t.payload
}

func (t *LongTag) Encode(w io.Writer) error {
	return encodeTagExcludeEndTag(w, t)
}

type LongPayload int64

func NewLongPayload() Payload {
	return new(LongPayload)
}

func (p *LongPayload) TypeId() TagType {
	return LongType
}

func (p *LongPayload) Encode(w io.Writer) error {
	return encodeNumericPayload(w, p)
}

func (p *LongPayload) Decode(r io.Reader) error {
	return decodeNumericPayload(r, p)
}
