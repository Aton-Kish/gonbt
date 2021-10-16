package gonbt

import (
	"encoding/binary"
	"io"
)

func Decode(r io.Reader, tag *Tag) error {
	var typ TagType
	if err := binary.Read(r, binary.BigEndian, &typ); err != nil {
		return err
	}

	if t, err := NewTag(typ); err != nil {
		return err
	} else {
		*tag = t
	}

	if err := (*tag).Decode(r); err != nil {
		return err
	}

	return nil
}

// Tag Name
func (n *TagName) Decode(r io.Reader) error {
	var l uint16
	if err := binary.Read(r, binary.BigEndian, &l); err != nil {
		return err
	}

	b := make([]byte, l)
	if err := binary.Read(r, binary.BigEndian, b); err != nil {
		return err
	}
	*n = TagName(b)

	return nil
}

// Payload
func (p *BytePayload) Decode(r io.Reader) error {
	if err := binary.Read(r, binary.BigEndian, p); err != nil {
		return err
	}

	return nil
}

func (p *ShortPayload) Decode(r io.Reader) error {
	if err := binary.Read(r, binary.BigEndian, p); err != nil {
		return err
	}

	return nil
}

func (p *IntPayload) Decode(r io.Reader) error {
	if err := binary.Read(r, binary.BigEndian, p); err != nil {
		return err
	}

	return nil
}

func (p *LongPayload) Decode(r io.Reader) error {
	if err := binary.Read(r, binary.BigEndian, p); err != nil {
		return err
	}

	return nil
}

func (p *FloatPayload) Decode(r io.Reader) error {
	if err := binary.Read(r, binary.BigEndian, p); err != nil {
		return err
	}

	return nil
}

func (p *DoublePayload) Decode(r io.Reader) error {
	if err := binary.Read(r, binary.BigEndian, p); err != nil {
		return err
	}

	return nil
}

func (p *ByteArrayPayload) Decode(r io.Reader) error {
	var l int32
	if err := binary.Read(r, binary.BigEndian, &l); err != nil {
		return err
	}

	*p = make(ByteArrayPayload, l)
	if err := binary.Read(r, binary.BigEndian, p); err != nil {
		return err
	}

	return nil
}

func (p *StringPayload) Decode(r io.Reader) error {
	var l uint16
	if err := binary.Read(r, binary.BigEndian, &l); err != nil {
		return err
	}

	b := make([]byte, l)
	if err := binary.Read(r, binary.BigEndian, b); err != nil {
		return err
	}
	*p = StringPayload(b)

	return nil
}

func (p *ListPayload) Decode(r io.Reader) error {
	if err := binary.Read(r, binary.BigEndian, &p.PayloadType); err != nil {
		return err
	}

	var l int32
	if err := binary.Read(r, binary.BigEndian, &l); err != nil {
		return err
	}

	p.Payloads = make([]Payload, 0, l)
	for i := 0; i < int(l); i++ {
		payload, err := NewPayload(p.PayloadType)
		if err != nil {
			return err
		}

		if err := payload.Decode(r); err != nil {
			return err
		}

		p.Payloads = append(p.Payloads, payload)
	}

	return nil
}

func (p *CompoundPayload) Decode(r io.Reader) error {
	for {
		tag := new(Tag)
		if err := Decode(r, tag); err != nil {
			return err
		}

		*p = append(*p, *tag)

		if (*tag).TypeId() == TagEnd {
			break
		}
	}

	return nil
}

func (p *IntArrayPayload) Decode(r io.Reader) error {
	var l int32
	if err := binary.Read(r, binary.BigEndian, &l); err != nil {
		return err
	}

	*p = make(IntArrayPayload, l)
	if err := binary.Read(r, binary.BigEndian, p); err != nil {
		return err
	}

	return nil
}

func (p *LongArrayPayload) Decode(r io.Reader) error {
	var l int32
	if err := binary.Read(r, binary.BigEndian, &l); err != nil {
		return err
	}

	*p = make(LongArrayPayload, l)
	if err := binary.Read(r, binary.BigEndian, p); err != nil {
		return err
	}

	return nil
}

// Tag
func (t *EndTag) Decode(r io.Reader) error {
	return nil
}

func (t *ByteTag) Decode(r io.Reader) error {
	if err := t.TagName.Decode(r); err != nil {
		return err
	}

	if err := t.BytePayload.Decode(r); err != nil {
		return err
	}

	return nil
}

func (t *ShortTag) Decode(r io.Reader) error {
	if err := t.TagName.Decode(r); err != nil {
		return err
	}

	if err := t.ShortPayload.Decode(r); err != nil {
		return err
	}

	return nil
}

func (t *IntTag) Decode(r io.Reader) error {
	if err := t.TagName.Decode(r); err != nil {
		return err
	}

	if err := t.IntPayload.Decode(r); err != nil {
		return err
	}

	return nil
}

func (t *LongTag) Decode(r io.Reader) error {
	if err := t.TagName.Decode(r); err != nil {
		return err
	}

	if err := t.LongPayload.Decode(r); err != nil {
		return err
	}

	return nil
}

func (t *FloatTag) Decode(r io.Reader) error {
	if err := t.TagName.Decode(r); err != nil {
		return err
	}

	if err := t.FloatPayload.Decode(r); err != nil {
		return err
	}

	return nil
}

func (t *DoubleTag) Decode(r io.Reader) error {
	if err := t.TagName.Decode(r); err != nil {
		return err
	}

	if err := t.DoublePayload.Decode(r); err != nil {
		return err
	}

	return nil
}

func (t *ByteArrayTag) Decode(r io.Reader) error {
	if err := t.TagName.Decode(r); err != nil {
		return err
	}

	if err := t.ByteArrayPayload.Decode(r); err != nil {
		return err
	}

	return nil
}

func (t *StringTag) Decode(r io.Reader) error {
	if err := t.TagName.Decode(r); err != nil {
		return err
	}

	if err := t.StringPayload.Decode(r); err != nil {
		return err
	}

	return nil
}

func (t *ListTag) Decode(r io.Reader) error {
	if err := t.TagName.Decode(r); err != nil {
		return err
	}

	if err := t.ListPayload.Decode(r); err != nil {
		return err
	}

	return nil
}

func (t *CompoundTag) Decode(r io.Reader) error {
	if err := t.TagName.Decode(r); err != nil {
		return err
	}

	if err := t.CompoundPayload.Decode(r); err != nil {
		return err
	}

	return nil
}

func (t *IntArrayTag) Decode(r io.Reader) error {
	if err := t.TagName.Decode(r); err != nil {
		return err
	}

	if err := t.IntArrayPayload.Decode(r); err != nil {
		return err
	}

	return nil
}

func (t *LongArrayTag) Decode(r io.Reader) error {
	if err := t.TagName.Decode(r); err != nil {
		return err
	}

	if err := t.LongArrayPayload.Decode(r); err != nil {
		return err
	}

	return nil
}
