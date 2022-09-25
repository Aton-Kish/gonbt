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
	"sort"
	"strings"

	"github.com/Aton-Kish/gonbt/pointer"
	"github.com/Aton-Kish/gonbt/snbt"
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

func (t *CompoundTag) stringify(space string, indent string, depth int) string {
	return stringifyTag(t, space, indent, depth)
}

func (t *CompoundTag) parse(parser *snbt.Parser) error {
	return parseTag(t, parser)
}

func (t *CompoundTag) json(space string, indent string, depth int) string {
	return jsonTag(t, space, indent, depth)
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

func (p *CompoundPayload) stringify(space string, indent string, depth int) string {
	strs := make([]string, 0, len(*p))
	for _, tag := range *p {
		if tag.TypeId() == EndType {
			break
		}

		strs = append(strs, tag.stringify(space, indent, depth+1))
	}

	l := len(strs)
	sort.SliceStable(strs, func(i, j int) bool { return strs[i] < strs[j] })

	if indent == "" || l == 0 {
		return fmt.Sprintf("{%s}", strings.Join(strs, fmt.Sprintf(",%s", space)))
	}

	indents := ""
	for i := 0; i < depth; i++ {
		indents += indent
	}

	return fmt.Sprintf("{\n%s%s%s\n%s}", indents, indent, strings.Join(strs, fmt.Sprintf(",\n%s%s", indents, indent)), indents)
}

func (p *CompoundPayload) parse(parser *snbt.Parser) error {
	for parser.CurrToken().Char() != '}' {
		if err := parser.Next(); err != nil {
			return err
		}

		tag, err := newTagFromSnbt(parser)
		if err != nil {
			return err
		}

		if err := tag.parse(parser); err != nil {
			return err
		}

		*p = append(*p, tag)
	}

	*p = append(*p, &EndTag{})

	// NOTE: ignore stop iteration error
	if err := parser.Next(); err != nil && err.Error() != "stop iteration" {
		return err
	}

	return nil
}

func (p *CompoundPayload) json(space string, indent string, depth int) string {
	strs := make([]string, 0, len(*p))
	for _, tag := range *p {
		if tag.TypeId() == EndType {
			break
		}

		strs = append(strs, tag.json(space, indent, depth+1))
	}

	l := len(strs)
	sort.SliceStable(strs, func(i, j int) bool { return strs[i] < strs[j] })

	if indent == "" || l == 0 {
		return fmt.Sprintf("{%s}", strings.Join(strs, fmt.Sprintf(",%s", space)))
	}

	indents := ""
	for i := 0; i < depth; i++ {
		indents += indent
	}

	return fmt.Sprintf("{\n%s%s%s\n%s}", indents, indent, strings.Join(strs, fmt.Sprintf(",\n%s%s", indents, indent)), indents)
}
