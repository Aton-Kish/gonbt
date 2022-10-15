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

type IntArrayTag struct {
	tagName *TagName
	payload *IntArrayPayload
}

func NewIntArrayTag(tagName *TagName, payload *IntArrayPayload) *IntArrayTag {
	return &IntArrayTag{
		tagName: tagName,
		payload: payload,
	}
}

func (t IntArrayTag) String() string {
	return t.stringify(" ", "", 0)
}

func (t *IntArrayTag) TypeId() TagType {
	return t.Payload().TypeId()
}

func (t *IntArrayTag) TagName() *TagName {
	return t.tagName
}

func (t *IntArrayTag) Payload() Payload {
	return t.payload
}

func (t *IntArrayTag) encode(w io.Writer) error {
	if err := Encode(w, t); err != nil {
		return err
	}

	return nil
}

func (t *IntArrayTag) decode(r io.Reader) error {
	tag, err := Decode(r)
	if err != nil {
		return err
	}

	v, ok := tag.(*IntArrayTag)
	if !ok {
		err = &NbtError{Op: "decode", Err: ErrDecode}
		return err
	}

	*t = *v

	return nil
}

func (t *IntArrayTag) stringify(space string, indent string, depth int) string {
	return stringifyTag(t, space, indent, depth)
}

func (t *IntArrayTag) parse(parser *snbt.Parser) error {
	if err := parseTag(t, parser); err != nil {
		return err
	}

	return nil
}

func (t *IntArrayTag) json(space string, indent string, depth int) string {
	return jsonTag(t, space, indent, depth)
}
