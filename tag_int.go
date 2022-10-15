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

	"github.com/Aton-Kish/gonbt/snbt"
)

type IntTag struct {
	tagName *TagName
	payload *IntPayload
}

func NewIntTag(tagName *TagName, payload *IntPayload) *IntTag {
	return &IntTag{
		tagName: tagName,
		payload: payload,
	}
}

func (t *IntTag) TypeId() TagType {
	return t.Payload().TypeId()
}

func (t *IntTag) TagName() *TagName {
	return t.tagName
}

func (t *IntTag) Payload() Payload {
	return t.payload
}

func (t *IntTag) encode(w io.Writer) error {
	return Encode(w, t)
}

func (t *IntTag) decode(r io.Reader) error {
	tag, err := Decode(r)
	if err != nil {
		return err
	}

	v, ok := tag.(*IntTag)
	if !ok {
		err = &NbtError{Op: "decode", Err: ErrDecode}
		return err
	}

	*t = *v

	return nil
}

func (t *IntTag) stringify(space string, indent string, depth int) string {
	return stringifyTag(t, space, indent, depth)
}

func (t *IntTag) parse(parser *snbt.Parser) error {
	return parseTag(t, parser)
}

func (t *IntTag) json(space string, indent string, depth int) string {
	return jsonTag(t, space, indent, depth)
}
