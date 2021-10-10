package gonbt

import (
	"fmt"
	"io"
)

// Tag Name
type TagName string

// Payload
type (
	Payload interface {
		TypeId() TagType
		Decode(r io.Reader) error
	}

	BytePayload int8

	ShortPayload int16

	IntPayload int32

	LongPayload int64

	FloatPayload float32

	DoublePayload float64

	ByteArrayPayload []int8

	StringPayload string

	ListPayload struct {
		PayloadType TagType
		Payloads    []Payload
	}

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

// Tag
type (
	Tag interface {
		TypeId() TagType
		Decode(r io.Reader) error
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
