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

type ByteArrayTag struct {
	tagName *TagName
	payload *ByteArrayPayload
}

func NewByteArrayTag(tagName *TagName, payload *ByteArrayPayload) *ByteArrayTag {
	return &ByteArrayTag{
		tagName: tagName,
		payload: payload,
	}
}

func (t ByteArrayTag) String() string {
	return t.stringify(" ", "", 0)
}

func (t *ByteArrayTag) TypeId() TagType {
	return t.Payload().TypeId()
}

func (t *ByteArrayTag) TagName() *TagName {
	return t.tagName
}

func (t *ByteArrayTag) Payload() Payload {
	return t.payload
}

func (t *ByteArrayTag) encode(w io.Writer) error {
	if err := Encode(w, t); err != nil {
		logger.Printf("(*ByteArrayTag).encode; tag: %s; error: %s", t, err)
		return err
	}

	return nil
}

func (t *ByteArrayTag) decode(r io.Reader) error {
	tag, err := Decode(r)
	if err != nil {
		logger.Printf("(*ByteArrayTag).decode; tag: %s; error: %s", t, err)
		return err
	}

	v, ok := tag.(*ByteArrayTag)
	if !ok {
		err = &NbtError{Op: "decode", Err: ErrDecode}
		logger.Printf("(*ByteArrayTag).decode; tag: %s; error: %s", t, err)
		return err
	}

	*t = *v

	return nil
}

func (t *ByteArrayTag) stringify(space string, indent string, depth int) string {
	return stringifyTag(t, space, indent, depth)
}

func (t *ByteArrayTag) parse(parser *snbt.Parser) error {
	if err := parseTag(t, parser); err != nil {
		logger.Printf("(*ByteArrayTag).parse; tag: %s; error: %s", t, err)
		return err
	}

	return nil
}

func (t *ByteArrayTag) json(space string, indent string, depth int) string {
	return jsonTag(t, space, indent, depth)
}
