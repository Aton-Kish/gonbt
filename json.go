package gonbt

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Json(tag *Tag) string {
	return PrettyJson(tag, " ", "")
}

func CompactJson(tag *Tag) string {
	return PrettyJson(tag, "", "")
}

func PrettyJson(tag *Tag, space string, indent string) string {
	var rootName string
	switch t := (*tag).(type) {
	case *EndTag:
		rootName = ""
	case *ByteTag:
		rootName = string(t.TagName)
	case *ShortTag:
		rootName = string(t.TagName)
	case *IntTag:
		rootName = string(t.TagName)
	case *LongTag:
		rootName = string(t.TagName)
	case *FloatTag:
		rootName = string(t.TagName)
	case *DoubleTag:
		rootName = string(t.TagName)
	case *ByteArrayTag:
		rootName = string(t.TagName)
	case *StringTag:
		rootName = string(t.TagName)
	case *ListTag:
		rootName = string(t.TagName)
	case *CompoundTag:
		rootName = string(t.TagName)
	case *IntArrayTag:
		rootName = string(t.TagName)
	case *LongArrayTag:
		rootName = string(t.TagName)
	}

	if rootName == "" {
		json := (*tag).Json(space, indent, 0)
		return strings.TrimLeft(json[3:], " ")
	}

	if indent == "" {
		return "{" + indent + (*tag).Json(space, indent, 1) + "}"
	}

	return "{\n" + indent + (*tag).Json(space, indent, 1) + "\n}"
}

// Tag Name
func (t *TagName) Json() string {
	s := string(*t)
	qs := strconv.Quote(s)
	return qs
}

// Payload
func (p *BytePayload) Json(space string, indent string, depth int) string {
	return fmt.Sprint(*p)
}

func (p *ShortPayload) Json(space string, indent string, depth int) string {
	return fmt.Sprint(*p)
}

func (p *IntPayload) Json(space string, indent string, depth int) string {
	return fmt.Sprint(*p)
}

func (p *LongPayload) Json(space string, indent string, depth int) string {
	return fmt.Sprint(*p)
}

func (p *FloatPayload) Json(space string, indent string, depth int) string {
	return fmt.Sprint(*p)
}

func (p *DoublePayload) Json(space string, indent string, depth int) string {
	return fmt.Sprint(*p)
}

func (p *ByteArrayPayload) Json(space string, indent string, depth int) string {
	l := len(*p)
	sl := make([]string, 0, l)
	for _, v := range *p {
		sl = append(sl, fmt.Sprint(v))
	}

	if indent == "" || l == 0 {
		return "[" + strings.Join(sl, ","+space) + "]"
	}

	var indentBuilder strings.Builder
	for i := 0; i < depth; i++ {
		indentBuilder.WriteString(indent)
	}
	indents := indentBuilder.String()

	indentBuilder.WriteString(indent)
	moreIndents := indentBuilder.String()

	return "[\n" +
		moreIndents + strings.Join(sl, ",\n"+moreIndents) + "\n" +
		indents + "]"
}

func (p *StringPayload) Json(space string, indent string, depth int) string {
	s := string(*p)
	qs := strconv.Quote(s)
	return qs
}

func (p *ListPayload) Json(space string, indent string, depth int) string {
	l := len(*p)
	sl := make([]string, 0, l)
	for _, payload := range *p {
		sl = append(sl, payload.Json(space, indent, depth+1))
	}

	if indent == "" || l == 0 {
		return "[" + strings.Join(sl, ","+space) + "]"
	}

	var indentBuilder strings.Builder
	for i := 0; i < depth; i++ {
		indentBuilder.WriteString(indent)
	}
	indents := indentBuilder.String()

	indentBuilder.WriteString(indent)
	moreIndents := indentBuilder.String()

	return "[\n" +
		moreIndents + strings.Join(sl, ",\n"+moreIndents) + "\n" +
		indents + "]"
}

func (p *CompoundPayload) Json(space string, indent string, depth int) string {
	sl := make([]string, 0, len(*p))
	for _, tag := range *p {
		if tag.TypeId() == TagEnd {
			break
		}

		sl = append(sl, tag.Json(space, indent, depth+1))
	}

	l := len(sl)
	sort.SliceStable(sl, func(i, j int) bool { return sl[i] < sl[j] })

	if indent == "" || l == 0 {
		return "{" + strings.Join(sl, ","+space) + "}"
	}

	var indentBuilder strings.Builder
	for i := 0; i < depth; i++ {
		indentBuilder.WriteString(indent)
	}
	indents := indentBuilder.String()

	indentBuilder.WriteString(indent)
	moreIndents := indentBuilder.String()

	return "{\n" +
		moreIndents + strings.Join(sl, ",\n"+moreIndents) + "\n" +
		indents + "}"
}

func (p *IntArrayPayload) Json(space string, indent string, depth int) string {
	l := len(*p)
	sl := make([]string, 0, l)
	for _, v := range *p {
		sl = append(sl, fmt.Sprint(v))
	}

	if indent == "" || l == 0 {
		return "[" + strings.Join(sl, ","+space) + "]"
	}

	var indentBuilder strings.Builder
	for i := 0; i < depth; i++ {
		indentBuilder.WriteString(indent)
	}
	indents := indentBuilder.String()

	indentBuilder.WriteString(indent)
	moreIndents := indentBuilder.String()

	return "[\n" +
		moreIndents + strings.Join(sl, ",\n"+moreIndents) + "\n" +
		indents + "]"
}

func (p *LongArrayPayload) Json(space string, indent string, depth int) string {
	l := len(*p)
	sl := make([]string, 0, l)
	for _, v := range *p {
		sl = append(sl, fmt.Sprint(v))
	}

	if indent == "" || l == 0 {
		return "[" + strings.Join(sl, ","+space) + "]"
	}

	var indentBuilder strings.Builder
	for i := 0; i < depth; i++ {
		indentBuilder.WriteString(indent)
	}
	indents := indentBuilder.String()

	indentBuilder.WriteString(indent)
	moreIndents := indentBuilder.String()

	return "[\n" +
		moreIndents + strings.Join(sl, ",\n"+moreIndents) + "\n" +
		indents + "]"
}

// Tag
func (t *EndTag) Json(space string, indent string, depth int) string {
	return ""
}

func (t *ByteTag) Json(space string, indent string, depth int) string {
	return t.TagName.Json() + ":" + space + t.BytePayload.Json(space, indent, depth)
}

func (t *ShortTag) Json(space string, indent string, depth int) string {
	return t.TagName.Json() + ":" + space + t.ShortPayload.Json(space, indent, depth)
}

func (t *IntTag) Json(space string, indent string, depth int) string {
	return t.TagName.Json() + ":" + space + t.IntPayload.Json(space, indent, depth)
}

func (t *LongTag) Json(space string, indent string, depth int) string {
	return t.TagName.Json() + ":" + space + t.LongPayload.Json(space, indent, depth)
}

func (t *FloatTag) Json(space string, indent string, depth int) string {
	return t.TagName.Json() + ":" + space + t.FloatPayload.Json(space, indent, depth)
}

func (t *DoubleTag) Json(space string, indent string, depth int) string {
	return t.TagName.Json() + ":" + space + t.DoublePayload.Json(space, indent, depth)
}

func (t *ByteArrayTag) Json(space string, indent string, depth int) string {
	return t.TagName.Json() + ":" + space + t.ByteArrayPayload.Json(space, indent, depth)
}

func (t *StringTag) Json(space string, indent string, depth int) string {
	return t.TagName.Json() + ":" + space + t.StringPayload.Json(space, indent, depth)
}

func (t *ListTag) Json(space string, indent string, depth int) string {
	return t.TagName.Json() + ":" + space + t.ListPayload.Json(space, indent, depth)
}

func (t *CompoundTag) Json(space string, indent string, depth int) string {
	return t.TagName.Json() + ":" + space + t.CompoundPayload.Json(space, indent, depth)
}

func (t *IntArrayTag) Json(space string, indent string, depth int) string {
	return t.TagName.Json() + ":" + space + t.IntArrayPayload.Json(space, indent, depth)
}

func (t *LongArrayTag) Json(space string, indent string, depth int) string {
	return t.TagName.Json() + ":" + space + t.LongArrayPayload.Json(space, indent, depth)
}
