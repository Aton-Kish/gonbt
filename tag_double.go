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

type DoubleTag struct {
	tagName *TagName
	payload *DoublePayload
}

func NewDoubleTag(tagName *TagName, payload *DoublePayload) *DoubleTag {
	return &DoubleTag{
		tagName: tagName,
		payload: payload,
	}
}

func (t DoubleTag) String() string {
	return t.stringify(" ", "", 0)
}

func (t *DoubleTag) TypeId() TagType {
	return t.Payload().TypeId()
}

func (t *DoubleTag) TagName() *TagName {
	return t.tagName
}

func (t *DoubleTag) Payload() Payload {
	return t.payload
}

func (t *DoubleTag) encode(w io.Writer) error {
	if err := Encode(w, t); err != nil {
		return err
	}

	return nil
}

func (t *DoubleTag) decode(r io.Reader) error {
	tag, err := Decode(r)
	if err != nil {
		return err
	}

	v, ok := tag.(*DoubleTag)
	if !ok {
		err = &NbtError{Op: "decode", Err: ErrDecode}
		return err
	}

	*t = *v

	return nil
}

func (t *DoubleTag) stringify(space string, indent string, depth int) string {
	return stringifyTag(t, space, indent, depth)
}

func (t *DoubleTag) parse(parser *snbt.Parser) error {
	if err := parseTag(t, parser); err != nil {
		return err
	}

	return nil
}

func (t *DoubleTag) json(space string, indent string, depth int) string {
	return jsonTag(t, space, indent, depth)
}
