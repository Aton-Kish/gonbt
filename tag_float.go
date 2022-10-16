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

func (t FloatTag) String() string {
	return t.stringify(" ", "", 0)
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
	if err := Encode(w, t); err != nil {
		logger.Println("failed to encode", "func", getFuncName(), "tag", t, "error", err)
		return err
	}

	return nil
}

func (t *FloatTag) decode(r io.Reader) error {
	tag, err := Decode(r)
	if err != nil {
		logger.Println("failed to decode", "func", getFuncName(), "tag", t, "error", err)
		return err
	}

	v, ok := tag.(*FloatTag)
	if !ok {
		err = &NbtError{Op: "decode", Err: ErrDecode}
		logger.Println("failed to decode", "func", getFuncName(), "tag", t, "error", err)
		return err
	}

	*t = *v

	return nil
}

func (t *FloatTag) stringify(space string, indent string, depth int) string {
	return stringifyTag(t, space, indent, depth)
}

func (t *FloatTag) parse(parser *snbt.Parser) error {
	if err := parseTag(t, parser); err != nil {
		logger.Println("failed to parse", "func", getFuncName(), "tag", t, "error", err)
		return err
	}

	return nil
}

func (t *FloatTag) json(space string, indent string, depth int) string {
	return jsonTag(t, space, indent, depth)
}
