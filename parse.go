package gonbt

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	bytePattern   = regexp.MustCompile(`^(-?\d+)[Bb]$`)
	shortPattern  = regexp.MustCompile(`^(-?\d+)[Ss]$`)
	intPattern    = regexp.MustCompile(`^-?\d+$`)
	longPattern   = regexp.MustCompile(`^(-?\d+)[Ll]$`)
	floatPattern  = regexp.MustCompile(`^(-?\d+\.\d+([Ee][+-]?\d+)?)[Ff]$`)
	doublePattern = regexp.MustCompile(`^(-?\d+\.\d+([Ee][+-]?\d+)?)[Dd]?$`)
)

func Parse(bm *SnbtTokenBitmaps, tag *Tag) error {
	t, err := NewTagFromSnbt(bm)
	if err != nil {
		return err
	}
	*tag = t

	if err := (*tag).Parse(bm); err != nil {
		return err
	}

	return nil
}

// Payload
func (p *BytePayload) Parse(bm *SnbtTokenBitmaps) error {
	b := bm.Raw[bm.PrevToken.Index+1 : bm.CurrToken.Index]

	g := bytePattern.FindSubmatch(b)
	if len(g) < 2 {
		return errors.New("invalid snbt format")
	}

	i, err := strconv.ParseInt(string(g[1]), 10, 8)
	if err != nil {
		return err
	}

	*p = BytePayload(i)

	return nil
}

func (p *ShortPayload) Parse(bm *SnbtTokenBitmaps) error {
	b := bm.Raw[bm.PrevToken.Index+1 : bm.CurrToken.Index]

	g := shortPattern.FindSubmatch(b)
	if len(g) < 2 {
		return errors.New("invalid snbt format")
	}

	i, err := strconv.ParseInt(string(g[1]), 10, 16)
	if err != nil {
		return err
	}

	*p = ShortPayload(i)

	return nil
}

func (p *IntPayload) Parse(bm *SnbtTokenBitmaps) error {
	b := bm.Raw[bm.PrevToken.Index+1 : bm.CurrToken.Index]

	i, err := strconv.ParseInt(string(b), 10, 32)
	if err != nil {
		return err
	}

	*p = IntPayload(i)

	return nil
}

func (p *LongPayload) Parse(bm *SnbtTokenBitmaps) error {
	b := bm.Raw[bm.PrevToken.Index+1 : bm.CurrToken.Index]

	g := longPattern.FindSubmatch(b)
	if len(g) < 2 {
		return errors.New("invalid snbt format")
	}

	i, err := strconv.ParseInt(string(g[1]), 10, 64)
	if err != nil {
		return err
	}

	*p = LongPayload(i)

	return nil
}

func (p *FloatPayload) Parse(bm *SnbtTokenBitmaps) error {
	b := bm.Raw[bm.PrevToken.Index+1 : bm.CurrToken.Index]

	g := floatPattern.FindSubmatch(b)
	if len(g) < 2 {
		return errors.New("invalid snbt format")
	}

	i, err := strconv.ParseFloat(string(g[1]), 32)
	if err != nil {
		return err
	}

	*p = FloatPayload(i)

	return nil
}

func (p *DoublePayload) Parse(bm *SnbtTokenBitmaps) error {
	b := bm.Raw[bm.PrevToken.Index+1 : bm.CurrToken.Index]

	g := doublePattern.FindSubmatch(b)
	if len(g) < 2 {
		return errors.New("invalid snbt format")
	}

	i, err := strconv.ParseFloat(string(g[1]), 64)
	if err != nil {
		return err
	}

	*p = DoublePayload(i)

	return nil
}

func (p *ByteArrayPayload) Parse(bm *SnbtTokenBitmaps) error {
	if c := bm.Raw[bm.CurrToken.Index+1]; c != ']' {
		for {
			if bm.CurrToken.Char == ']' {
				break
			}

			bm.NextToken(``, `" `)

			payload, err := NewPayloadFromSnbt(bm)
			if err != nil {
				return err
			}

			if err := payload.Parse(bm); err != nil {
				return err
			}

			cp, ok := payload.(*BytePayload)
			if !ok {
				return errors.New("invalid snbt format")
			}

			*p = append(*p, int8(*cp))
		}
	}

	bm.NextToken(``, `" `)

	return nil
}

func (p *StringPayload) Parse(bm *SnbtTokenBitmaps) error {
	b := bm.Raw[bm.PrevToken.Index+1 : bm.CurrToken.Index]
	qs := string(b)

	if qs[0] == '\'' {
		ms := qs[1 : len(qs)-1]
		ms = strings.ReplaceAll(ms, `\'`, `'`)
		ms = strings.ReplaceAll(ms, `"`, `\"`)
		qs = `"` + ms + `"`
	}

	s, err := strconv.Unquote(qs)
	if err != nil {
		return err
	}

	*p = StringPayload(s)

	return nil
}

func (p *ListPayload) Parse(bm *SnbtTokenBitmaps) error {
	if c := bm.Raw[bm.CurrToken.Index+1]; c != ']' {
		for {
			if bm.CurrToken.Char == ']' {
				break
			}

			bm.NextToken(``, `" `)

			payload, err := NewPayloadFromSnbt(bm)
			if err != nil {
				return err
			}

			if err := payload.Parse(bm); err != nil {
				return err
			}

			p.Payloads = append(p.Payloads, payload)
		}

		p.PayloadType = p.Payloads[0].TypeId()
	}

	bm.NextToken(``, `" `)

	return nil
}

func (p *CompoundPayload) Parse(bm *SnbtTokenBitmaps) error {
	if c := bm.Raw[bm.CurrToken.Index+1]; c != '}' {
		for {
			if bm.CurrToken.Char == '}' {
				break
			}

			tag := new(Tag)
			if err := Parse(bm, tag); err != nil {
				return err
			}

			*p = append(*p, *tag)
		}

		*p = append(*p, &EndTag{})

		bm.NextToken(``, `" `)
	}

	return nil
}

func (p *IntArrayPayload) Parse(bm *SnbtTokenBitmaps) error {
	if c := bm.Raw[bm.CurrToken.Index+1]; c != ']' {
		for {
			if bm.CurrToken.Char == ']' {
				break
			}

			bm.NextToken(``, `" `)

			payload, err := NewPayloadFromSnbt(bm)
			if err != nil {
				return err
			}

			if err := payload.Parse(bm); err != nil {
				return err
			}

			cp, ok := payload.(*IntPayload)
			if !ok {
				return errors.New("invalid snbt format")
			}

			*p = append(*p, int32(*cp))
		}
	}

	bm.NextToken(``, `" `)

	return nil
}

func (p *LongArrayPayload) Parse(bm *SnbtTokenBitmaps) error {
	if c := bm.Raw[bm.CurrToken.Index+1]; c != ']' {
		for {
			if bm.CurrToken.Char == ']' {
				break
			}

			bm.NextToken(``, `" `)

			payload, err := NewPayloadFromSnbt(bm)
			if err != nil {
				return err
			}

			if err := payload.Parse(bm); err != nil {
				return err
			}

			cp, ok := payload.(*LongPayload)
			if !ok {
				return errors.New("invalid snbt format")
			}

			*p = append(*p, int64(*cp))
		}
	}

	bm.NextToken(``, `" `)

	return nil
}

// Tag
func (t *EndTag) Parse(bm *SnbtTokenBitmaps) error {
	return nil
}

func (t *ByteTag) Parse(bm *SnbtTokenBitmaps) error {
	if err := t.BytePayload.Parse(bm); err != nil {
		return err
	}

	return nil
}

func (t *ShortTag) Parse(bm *SnbtTokenBitmaps) error {
	if err := t.ShortPayload.Parse(bm); err != nil {
		return err
	}

	return nil
}

func (t *IntTag) Parse(bm *SnbtTokenBitmaps) error {
	if err := t.IntPayload.Parse(bm); err != nil {
		return err
	}

	return nil
}

func (t *LongTag) Parse(bm *SnbtTokenBitmaps) error {
	if err := t.LongPayload.Parse(bm); err != nil {
		return err
	}

	return nil
}

func (t *FloatTag) Parse(bm *SnbtTokenBitmaps) error {
	if err := t.FloatPayload.Parse(bm); err != nil {
		return err
	}

	return nil
}

func (t *DoubleTag) Parse(bm *SnbtTokenBitmaps) error {
	if err := t.DoublePayload.Parse(bm); err != nil {
		return err
	}

	return nil
}

func (t *ByteArrayTag) Parse(bm *SnbtTokenBitmaps) error {
	if err := t.ByteArrayPayload.Parse(bm); err != nil {
		return err
	}

	return nil
}

func (t *StringTag) Parse(bm *SnbtTokenBitmaps) error {
	if err := t.StringPayload.Parse(bm); err != nil {
		return err
	}

	return nil
}

func (t *ListTag) Parse(bm *SnbtTokenBitmaps) error {
	if err := t.ListPayload.Parse(bm); err != nil {
		return err
	}

	return nil
}

func (t *CompoundTag) Parse(bm *SnbtTokenBitmaps) error {
	if err := t.CompoundPayload.Parse(bm); err != nil {
		return err
	}

	return nil
}

func (t *IntArrayTag) Parse(bm *SnbtTokenBitmaps) error {
	if err := t.IntArrayPayload.Parse(bm); err != nil {
		return err
	}

	return nil
}

func (t *LongArrayTag) Parse(bm *SnbtTokenBitmaps) error {
	if err := t.LongArrayPayload.Parse(bm); err != nil {
		return err
	}

	return nil
}
