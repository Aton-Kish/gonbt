package gonbt

import (
	"errors"
	"fmt"
	"io"
	"strconv"
)

// Tag Name
type TagName string

// Payload
type (
	Payload interface {
		TypeId() TagType
		Decode(r io.Reader) error
		Encode(w io.Writer) error
		Stringify(string, string, int) string
		Parse(*SnbtTokenBitmaps) error
	}

	BytePayload int8

	ShortPayload int16

	IntPayload int32

	LongPayload int64

	FloatPayload float32

	DoublePayload float64

	ByteArrayPayload []int8

	StringPayload string

	ListPayload []Payload

	CompoundPayload []Tag

	IntArrayPayload []int32

	LongArrayPayload []int64
)

func NewPayload(typ TagType) (Payload, error) {
	switch typ {
	case TagByte:
		return new(BytePayload), nil
	case TagShort:
		return new(ShortPayload), nil
	case TagInt:
		return new(IntPayload), nil
	case TagLong:
		return new(LongPayload), nil
	case TagFloat:
		return new(FloatPayload), nil
	case TagDouble:
		return new(DoublePayload), nil
	case TagByteArray:
		return new(ByteArrayPayload), nil
	case TagString:
		return new(StringPayload), nil
	case TagList:
		return new(ListPayload), nil
	case TagCompound:
		return new(CompoundPayload), nil
	case TagIntArray:
		return new(IntArrayPayload), nil
	case TagLongArray:
		return new(LongArrayPayload), nil
	}

	return nil, fmt.Errorf(`invalid tag type %d`, typ)
}

func NewPayloadFromSnbt(bm *SnbtTokenBitmaps) (Payload, error) {
	switch bm.CurrToken.Char {
	case '{':
		return new(CompoundPayload), nil
	case '[':
		typ := bm.Raw[bm.CurrToken.Index+1]
		switch typ {
		case 'B', 'I', 'L':
			bm.NextToken(``, `" `)

			if bm.CurrToken.Char != ';' {
				return nil, errors.New("invalid snbt format")
			}

			switch typ {
			case 'B':
				return new(ByteArrayPayload), nil
			case 'I':
				return new(IntArrayPayload), nil
			case 'L':
				return new(LongArrayPayload), nil
			}
		}

		return new(ListPayload), nil
	case '"', ' ', ':', ';':
		return nil, errors.New("invalid snbt format")
	}

	var b []byte
	if bm.CurrToken.Index < 0 {
		b = bm.Raw[bm.PrevToken.Index+1:]
	} else {
		b = bm.Raw[bm.PrevToken.Index+1 : bm.CurrToken.Index]
	}

	if bytePattern.Match(b) {
		return new(BytePayload), nil
	}

	if shortPattern.Match(b) {
		return new(ShortPayload), nil
	}

	if intPattern.Match(b) {
		return new(IntPayload), nil
	}

	if longPattern.Match(b) {
		return new(LongPayload), nil
	}

	if floatPattern.Match(b) {
		return new(FloatPayload), nil
	}

	if doublePattern.Match(b) {
		return new(DoublePayload), nil
	}

	return new(StringPayload), nil
}

// Tag
type (
	Tag interface {
		TypeId() TagType
		Decode(r io.Reader) error
		Encode(w io.Writer) error
		Stringify(string, string, int) string
		Parse(*SnbtTokenBitmaps) error
	}

	EndTag struct {
	}

	ByteTag struct {
		TagName
		BytePayload
	}

	ShortTag struct {
		TagName
		ShortPayload
	}

	IntTag struct {
		TagName
		IntPayload
	}

	LongTag struct {
		TagName
		LongPayload
	}

	FloatTag struct {
		TagName
		FloatPayload
	}

	DoubleTag struct {
		TagName
		DoublePayload
	}

	ByteArrayTag struct {
		TagName
		ByteArrayPayload
	}

	StringTag struct {
		TagName
		StringPayload
	}

	ListTag struct {
		TagName
		ListPayload
	}

	CompoundTag struct {
		TagName
		CompoundPayload
	}

	IntArrayTag struct {
		TagName
		IntArrayPayload
	}

	LongArrayTag struct {
		TagName
		LongArrayPayload
	}
)

func NewTag(typ TagType) (Tag, error) {
	switch typ {
	case TagEnd:
		return new(EndTag), nil
	case TagByte:
		return new(ByteTag), nil
	case TagShort:
		return new(ShortTag), nil
	case TagInt:
		return new(IntTag), nil
	case TagLong:
		return new(LongTag), nil
	case TagFloat:
		return new(FloatTag), nil
	case TagDouble:
		return new(DoubleTag), nil
	case TagByteArray:
		return new(ByteArrayTag), nil
	case TagString:
		return new(StringTag), nil
	case TagList:
		return new(ListTag), nil
	case TagCompound:
		return new(CompoundTag), nil
	case TagIntArray:
		return new(IntArrayTag), nil
	case TagLongArray:
		return new(LongArrayTag), nil
	}

	return nil, fmt.Errorf(`invalid tag type %d`, typ)
}

func NewTagFromSnbt(bm *SnbtTokenBitmaps) (Tag, error) {
	bm.NextToken(``, `" `)

	var n TagName
	if bm.CurrToken.Index > 0 {
		b := bm.Raw[bm.PrevToken.Index+1 : bm.CurrToken.Index]

		if s, err := strconv.Unquote(string(b)); err != nil {
			n = TagName(b)
		} else {
			n = TagName(s)
		}

		bm.NextToken(``, `" `)
	}

	p, err := NewPayloadFromSnbt(bm)
	if err != nil {
		return nil, err
	}

	switch cp := p.(type) {
	case *BytePayload:
		return &ByteTag{TagName: n, BytePayload: *cp}, nil
	case *ShortPayload:
		return &ShortTag{TagName: n, ShortPayload: *cp}, nil
	case *IntPayload:
		return &IntTag{TagName: n, IntPayload: *cp}, nil
	case *LongPayload:
		return &LongTag{TagName: n, LongPayload: *cp}, nil
	case *FloatPayload:
		return &FloatTag{TagName: n, FloatPayload: *cp}, nil
	case *DoublePayload:
		return &DoubleTag{TagName: n, DoublePayload: *cp}, nil
	case *ByteArrayPayload:
		return &ByteArrayTag{TagName: n, ByteArrayPayload: *cp}, nil
	case *StringPayload:
		return &StringTag{TagName: n, StringPayload: *cp}, nil
	case *ListPayload:
		return &ListTag{TagName: n, ListPayload: *cp}, nil
	case *CompoundPayload:
		return &CompoundTag{TagName: n, CompoundPayload: *cp}, nil
	case *IntArrayPayload:
		return &IntArrayTag{TagName: n, IntArrayPayload: *cp}, nil
	case *LongArrayPayload:
		return &LongArrayTag{TagName: n, LongArrayPayload: *cp}, nil
	}

	return nil, errors.New("invalid snbt format")
}
