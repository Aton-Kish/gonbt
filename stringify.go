package gonbt

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

var quotedPattern = regexp.MustCompile(`[ !"#$%&'()*,/:;<=>?@[\\\]^\x60{|}~]`)

func Stringify(tag *Tag) string {
	return PrettyStringify(tag, " ", "")
}

func CompactStringify(tag *Tag) string {
	return PrettyStringify(tag, "", "")
}

func PrettyStringify(tag *Tag, space string, indent string) string {
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
		snbt := (*tag).Stringify(space, indent, 0)
		return strings.TrimLeft(snbt[1:], " ")
	}

	if indent == "" {
		return "{" + indent + (*tag).Stringify(space, indent, 1) + "}"
	}

	return "{\n" + indent + (*tag).Stringify(space, indent, 1) + "\n}"
}

// Tag Name
func (t *TagName) Stringify() string {
	s := string(*t)
	if !quotedPattern.MatchString(s) {
		return s
	}

	qs := fmt.Sprintf("%#v", s)
	if strings.Contains(s, `"`) && !strings.Contains(s, `'`) {
		qs = `'` + qs[1:len(qs)-1] + `'`
		qs = strings.ReplaceAll(qs, `\"`, `"`)
		return qs
	}

	return qs
}

// Payload
func (p *BytePayload) Stringify(space string, indent string, depth int) string {
	return fmt.Sprintf("%db", *p)
}

func (p *ShortPayload) Stringify(space string, indent string, depth int) string {
	return fmt.Sprintf("%ds", *p)
}

func (p *IntPayload) Stringify(space string, indent string, depth int) string {
	return fmt.Sprintf("%d", *p)
}

func (p *LongPayload) Stringify(space string, indent string, depth int) string {
	return fmt.Sprintf("%dL", *p)
}

func (p *FloatPayload) Stringify(space string, indent string, depth int) string {
	return fmt.Sprintf("%gf", *p)
}

func (p *DoublePayload) Stringify(space string, indent string, depth int) string {
	return fmt.Sprintf("%gd", *p)
}

func (p *ByteArrayPayload) Stringify(space string, indent string, depth int) string {
	sl := make([]string, 0, len(*p))
	for _, v := range *p {
		sl = append(sl, fmt.Sprintf("%db", v))
	}

	return "[B;" + space + strings.Join(sl, ","+space) + "]"
}

func (p *StringPayload) Stringify(space string, indent string, depth int) string {
	s := string(*p)
	qs := fmt.Sprintf("%#v", s)
	if strings.Contains(s, `"`) && !strings.Contains(s, `'`) {
		qs = `'` + qs[1:len(qs)-1] + `'`
		qs = strings.ReplaceAll(qs, `\"`, `"`)
		return qs
	}

	return qs
}

func (p *ListPayload) Stringify(space string, indent string, depth int) string {
	l := len(p.Payloads)
	sl := make([]string, 0, l)
	for _, payload := range p.Payloads {
		sl = append(sl, payload.Stringify(space, indent, depth+1))
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

func (p *CompoundPayload) Stringify(space string, indent string, depth int) string {
	sl := make([]string, 0, len(*p))
	for _, tag := range *p {
		if tag.TypeId() == TagEnd {
			break
		}

		sl = append(sl, tag.Stringify(space, indent, depth+1))
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

func (p *IntArrayPayload) Stringify(space string, indent string, depth int) string {
	sl := make([]string, 0, len(*p))
	for _, v := range *p {
		sl = append(sl, fmt.Sprintf("%d", v))
	}

	return "[I;" + space + strings.Join(sl, ","+space) + "]"
}

func (p *LongArrayPayload) Stringify(space string, indent string, depth int) string {
	sl := make([]string, 0, len(*p))
	for _, v := range *p {
		sl = append(sl, fmt.Sprintf("%dL", v))
	}

	return "[L;" + space + strings.Join(sl, ","+space) + "]"
}

// Tag
func (t *EndTag) Stringify(space string, indent string, depth int) string {
	return ""
}

func (t *ByteTag) Stringify(space string, indent string, depth int) string {
	return t.TagName.Stringify() + ":" + space + t.BytePayload.Stringify(space, indent, depth)
}

func (t *ShortTag) Stringify(space string, indent string, depth int) string {
	return t.TagName.Stringify() + ":" + space + t.ShortPayload.Stringify(space, indent, depth)
}

func (t *IntTag) Stringify(space string, indent string, depth int) string {
	return t.TagName.Stringify() + ":" + space + t.IntPayload.Stringify(space, indent, depth)
}

func (t *LongTag) Stringify(space string, indent string, depth int) string {
	return t.TagName.Stringify() + ":" + space + t.LongPayload.Stringify(space, indent, depth)
}

func (t *FloatTag) Stringify(space string, indent string, depth int) string {
	return t.TagName.Stringify() + ":" + space + t.FloatPayload.Stringify(space, indent, depth)
}

func (t *DoubleTag) Stringify(space string, indent string, depth int) string {
	return t.TagName.Stringify() + ":" + space + t.DoublePayload.Stringify(space, indent, depth)
}

func (t *ByteArrayTag) Stringify(space string, indent string, depth int) string {
	return t.TagName.Stringify() + ":" + space + t.ByteArrayPayload.Stringify(space, indent, depth)
}

func (t *StringTag) Stringify(space string, indent string, depth int) string {
	return t.TagName.Stringify() + ":" + space + t.StringPayload.Stringify(space, indent, depth)
}

func (t *ListTag) Stringify(space string, indent string, depth int) string {
	return t.TagName.Stringify() + ":" + space + t.ListPayload.Stringify(space, indent, depth)
}

func (t *CompoundTag) Stringify(space string, indent string, depth int) string {
	return t.TagName.Stringify() + ":" + space + t.CompoundPayload.Stringify(space, indent, depth)
}

func (t *IntArrayTag) Stringify(space string, indent string, depth int) string {
	return t.TagName.Stringify() + ":" + space + t.IntArrayPayload.Stringify(space, indent, depth)
}

func (t *LongArrayTag) Stringify(space string, indent string, depth int) string {
	return t.TagName.Stringify() + ":" + space + t.LongArrayPayload.Stringify(space, indent, depth)
}
