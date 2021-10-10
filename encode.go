package gonbt

import (
	"encoding/binary"
	"io"
)

func EncodeNBT(w io.Writer, tag *Tag) error {
	if err := (*tag).Encode(w); err != nil {
		return err
	}

	return nil
}

// Tag Name
func (n *TagName) Encode(w io.Writer) error {
	l := uint16(len(*n))
	if err := binary.Write(w, binary.BigEndian, &l); err != nil {
		return err
	}

	b := []byte(*n)
	if err := binary.Write(w, binary.BigEndian, b); err != nil {
		return err
	}

	return nil
}

// Payload
func (p *BytePayload) Encode(w io.Writer) error {
	if err := binary.Write(w, binary.BigEndian, p); err != nil {
		return err
	}

	return nil
}

func (p *ShortPayload) Encode(w io.Writer) error {
	if err := binary.Write(w, binary.BigEndian, p); err != nil {
		return err
	}

	return nil
}

func (p *IntPayload) Encode(w io.Writer) error {
	if err := binary.Write(w, binary.BigEndian, p); err != nil {
		return err
	}

	return nil
}

func (p *LongPayload) Encode(w io.Writer) error {
	if err := binary.Write(w, binary.BigEndian, p); err != nil {
		return err
	}

	return nil
}

func (p *FloatPayload) Encode(w io.Writer) error {
	if err := binary.Write(w, binary.BigEndian, p); err != nil {
		return err
	}

	return nil
}

func (p *DoublePayload) Encode(w io.Writer) error {
	if err := binary.Write(w, binary.BigEndian, p); err != nil {
		return err
	}

	return nil
}

func (p *ByteArrayPayload) Encode(w io.Writer) error {
	l := int32(len(*p))
	if err := binary.Write(w, binary.BigEndian, &l); err != nil {
		return err
	}

	if err := binary.Write(w, binary.BigEndian, p); err != nil {
		return err
	}

	return nil
}

func (p *StringPayload) Encode(w io.Writer) error {
	l := uint16(len(*p))
	if err := binary.Write(w, binary.BigEndian, &l); err != nil {
		return err
	}

	b := []byte(*p)
	if err := binary.Write(w, binary.BigEndian, b); err != nil {
		return err
	}

	return nil
}

func (p *ListPayload) Encode(w io.Writer) error {
	if err := binary.Write(w, binary.BigEndian, &p.PayloadType); err != nil {
		return err
	}

	l := int32(len(p.Payloads))
	if err := binary.Write(w, binary.BigEndian, &l); err != nil {
		return err
	}

	for _, payload := range p.Payloads {
		if err := payload.Encode(w); err != nil {
			return err
		}
	}

	return nil
}

func (p *CompoundPayload) Encode(w io.Writer) error {
	for _, tag := range *p {
		if err := tag.Encode(w); err != nil {
			return err
		}
	}

	return nil
}

func (p *IntArrayPayload) Encode(w io.Writer) error {
	l := int32(len(*p))
	if err := binary.Write(w, binary.BigEndian, &l); err != nil {
		return err
	}

	if err := binary.Write(w, binary.BigEndian, p); err != nil {
		return err
	}

	return nil
}

func (p *LongArrayPayload) Encode(w io.Writer) error {
	l := int32(len(*p))
	if err := binary.Write(w, binary.BigEndian, &l); err != nil {
		return err
	}

	if err := binary.Write(w, binary.BigEndian, p); err != nil {
		return err
	}

	return nil
}

// Tag
func (t *EndTag) Encode(w io.Writer) error {
	typ := t.TypeId()
	if err := binary.Write(w, binary.BigEndian, &typ); err != nil {
		return err
	}

	return nil
}

func (t *ByteTag) Encode(w io.Writer) error {
	typ := t.TypeId()
	if err := binary.Write(w, binary.BigEndian, &typ); err != nil {
		return err
	}

	if err := t.TagName.Encode(w); err != nil {
		return err
	}

	if err := t.BytePayload.Encode(w); err != nil {
		return err
	}

	return nil
}

func (t *ShortTag) Encode(w io.Writer) error {
	typ := t.TypeId()
	if err := binary.Write(w, binary.BigEndian, &typ); err != nil {
		return err
	}

	if err := t.TagName.Encode(w); err != nil {
		return err
	}

	if err := t.ShortPayload.Encode(w); err != nil {
		return err
	}

	return nil
}

func (t *IntTag) Encode(w io.Writer) error {
	typ := t.TypeId()
	if err := binary.Write(w, binary.BigEndian, &typ); err != nil {
		return err
	}

	if err := t.TagName.Encode(w); err != nil {
		return err
	}

	if err := t.IntPayload.Encode(w); err != nil {
		return err
	}

	return nil
}

func (t *LongTag) Encode(w io.Writer) error {
	typ := t.TypeId()
	if err := binary.Write(w, binary.BigEndian, &typ); err != nil {
		return err
	}

	if err := t.TagName.Encode(w); err != nil {
		return err
	}

	if err := t.LongPayload.Encode(w); err != nil {
		return err
	}

	return nil
}

func (t *FloatTag) Encode(w io.Writer) error {
	typ := t.TypeId()
	if err := binary.Write(w, binary.BigEndian, &typ); err != nil {
		return err
	}

	if err := t.TagName.Encode(w); err != nil {
		return err
	}

	if err := t.FloatPayload.Encode(w); err != nil {
		return err
	}

	return nil
}

func (t *DoubleTag) Encode(w io.Writer) error {
	typ := t.TypeId()
	if err := binary.Write(w, binary.BigEndian, &typ); err != nil {
		return err
	}

	if err := t.TagName.Encode(w); err != nil {
		return err
	}

	if err := t.DoublePayload.Encode(w); err != nil {
		return err
	}

	return nil
}

func (t *ByteArrayTag) Encode(w io.Writer) error {
	typ := t.TypeId()
	if err := binary.Write(w, binary.BigEndian, &typ); err != nil {
		return err
	}

	if err := t.TagName.Encode(w); err != nil {
		return err
	}

	if err := t.ByteArrayPayload.Encode(w); err != nil {
		return err
	}

	return nil
}

func (t *StringTag) Encode(w io.Writer) error {
	typ := t.TypeId()
	if err := binary.Write(w, binary.BigEndian, &typ); err != nil {
		return err
	}

	if err := t.TagName.Encode(w); err != nil {
		return err
	}

	if err := t.StringPayload.Encode(w); err != nil {
		return err
	}

	return nil
}

func (t *ListTag) Encode(w io.Writer) error {
	typ := t.TypeId()
	if err := binary.Write(w, binary.BigEndian, &typ); err != nil {
		return err
	}

	if err := t.TagName.Encode(w); err != nil {
		return err
	}

	if err := t.ListPayload.Encode(w); err != nil {
		return err
	}

	return nil
}

func (t *CompoundTag) Encode(w io.Writer) error {
	typ := t.TypeId()
	if err := binary.Write(w, binary.BigEndian, &typ); err != nil {
		return err
	}

	if err := t.TagName.Encode(w); err != nil {
		return err
	}

	if err := t.CompoundPayload.Encode(w); err != nil {
		return err
	}

	return nil
}

func (t *IntArrayTag) Encode(w io.Writer) error {
	typ := t.TypeId()
	if err := binary.Write(w, binary.BigEndian, &typ); err != nil {
		return err
	}

	if err := t.TagName.Encode(w); err != nil {
		return err
	}

	if err := t.IntArrayPayload.Encode(w); err != nil {
		return err
	}

	return nil
}

func (t *LongArrayTag) Encode(w io.Writer) error {
	typ := t.TypeId()
	if err := binary.Write(w, binary.BigEndian, &typ); err != nil {
		return err
	}

	if err := t.TagName.Encode(w); err != nil {
		return err
	}

	if err := t.LongArrayPayload.Encode(w); err != nil {
		return err
	}

	return nil
}
