package gonbt

import (
	"errors"
	"strings"
)

type Token struct {
	Char  rune
	Index int
}

type SnbtTokenBitmaps struct {
	Raw              []byte
	StringMaskBitmap []uint64
	ValueMaskBitmap  []uint64
	TokenBitmaps     map[rune][]uint64
	PrevToken        *Token
	CurrToken        *Token
}

func NewSnbtTokenBitmaps(snbt string) SnbtTokenBitmaps {
	bmLen := len(snbt)/64 + 1
	bm := newSnbtTokenBitmapsWithoutRaw(bmLen)
	bm.Raw = []byte(snbt)
	return bm
}

func NewSnbtTokenBitmapsWithCapacity(snbtLen int) SnbtTokenBitmaps {
	bmLen := snbtLen/64 + 1
	bm := newSnbtTokenBitmapsWithoutRaw(bmLen)
	bm.Raw = make([]byte, 0, snbtLen)
	return bm
}

func newSnbtTokenBitmapsWithoutRaw(bitmapLen int) SnbtTokenBitmaps {
	return SnbtTokenBitmaps{
		StringMaskBitmap: make([]uint64, bitmapLen),
		ValueMaskBitmap:  make([]uint64, bitmapLen),
		TokenBitmaps: map[rune][]uint64{
			'"': make([]uint64, bitmapLen),
			' ': make([]uint64, bitmapLen),
			'{': make([]uint64, bitmapLen),
			'}': make([]uint64, bitmapLen),
			'[': make([]uint64, bitmapLen),
			']': make([]uint64, bitmapLen),
			',': make([]uint64, bitmapLen),
			':': make([]uint64, bitmapLen),
			';': make([]uint64, bitmapLen),
		},
		CurrToken: &Token{Char: ':', Index: 0},
	}
}

func (bm *SnbtTokenBitmaps) SetTokenBitmaps() {
	isQuoted := map[string]bool{"single": false, "double": false}
	for i := range bm.Raw {
		isEscaped := false
		if i > 0 && bm.Raw[i-1] == '\\' {
			isEscaped = true
		}

		j, k := i/64, i%64
		switch bm.Raw[i] {
		case '\'':
			if isEscaped || isQuoted["double"] {
				continue
			}

			isQuoted["single"] = !isQuoted["single"]
			bm.TokenBitmaps['"'][j] |= 1 << k
		case '"':
			if isEscaped || isQuoted["single"] {
				continue
			}

			isQuoted["double"] = !isQuoted["double"]
			bm.TokenBitmaps['"'][j] |= 1 << k
		case ' ', '\t', '\n', '\r', '\f':
			bm.TokenBitmaps[' '][j] |= 1 << k
		case '{':
			bm.TokenBitmaps['{'][j] |= 1 << k
		case '}':
			bm.TokenBitmaps['}'][j] |= 1 << k
		case '[':
			bm.TokenBitmaps['['][j] |= 1 << k
		case ']':
			bm.TokenBitmaps[']'][j] |= 1 << k
		case ',':
			bm.TokenBitmaps[','][j] |= 1 << k
		case ':':
			bm.TokenBitmaps[':'][j] |= 1 << k
		case ';':
			bm.TokenBitmaps[';'][j] |= 1 << k
		}
	}
}

func (bm *SnbtTokenBitmaps) SetMaskBitmaps() {
	l := len(bm.TokenBitmaps['"'])
	quoteBitmap := make([]uint64, l)
	copy(quoteBitmap, bm.TokenBitmaps['"'])

	quoteCnt := 0
	for i := 0; i < l; i++ {
		for quoteBitmap[i] != 0 {
			bm.StringMaskBitmap[i] ^= smearRightmost(quoteBitmap[i])
			quoteBitmap[i] = removeRightmost(quoteBitmap[i])
			quoteCnt++
		}

		if quoteCnt%2 == 1 {
			bm.StringMaskBitmap[i] = ^bm.StringMaskBitmap[i]
		}

		for k := range bm.TokenBitmaps {
			if k == '"' {
				continue
			}

			bm.TokenBitmaps[k][i] &= ^bm.StringMaskBitmap[i]
		}

		bm.ValueMaskBitmap[i] = ^bm.TokenBitmaps[' '][i]
	}
}

func (bm *SnbtTokenBitmaps) NextToken(allow, deny string) error {
	l := len(bm.Raw)

	var token rune
	index := l

	for k, v := range bm.TokenBitmaps {
		if (allow != `` && !strings.ContainsRune(allow, k)) || (deny != `` && strings.ContainsRune(deny, k)) {
			continue
		}

		for i, b := range v {
			if b == 0 {
				continue
			}

			if j := i*64 + rightmostIndex(b); j < index {
				token = k
				index = j
			}

			break
		}
	}

	if index == l {
		bm.PrevToken = bm.CurrToken
		bm.CurrToken = nil
		return errors.New("stop iteration")
	}

	i := index / 64
	bm.TokenBitmaps[token][i] = removeRightmost(bm.TokenBitmaps[token][i])

	bm.PrevToken = bm.CurrToken
	bm.CurrToken = &Token{Char: token, Index: index}
	return nil
}

func (bm *SnbtTokenBitmaps) Compact() SnbtTokenBitmaps {
	err := bm.NextToken(``, ` `)

	ci := 0
	cl := len(bm.Raw) - bitmapCount(bm.TokenBitmaps[' '])
	cbm := NewSnbtTokenBitmapsWithCapacity(cl)

	for i := range bm.ValueMaskBitmap {
		for bm.ValueMaskBitmap[i] != 0 {
			if ci == cl {
				break
			}

			j := i*64 + rightmostIndex(bm.ValueMaskBitmap[i])
			bm.ValueMaskBitmap[i] = removeRightmost(bm.ValueMaskBitmap[i])

			if j == bm.CurrToken.Index && err == nil {
				cbm.TokenBitmaps[bm.CurrToken.Char][ci/64] |= 1 << (ci % 64)
				err = bm.NextToken(``, ` `)
			}

			ci++
			cbm.Raw = append(cbm.Raw, bm.Raw[j])
		}
	}

	cbm.SetMaskBitmaps()

	return cbm
}

// Bitwise Operations
var (
	tarzHash  = uint64(0x03F566ED27179461)
	tarzTable = [64]int{
		0, 1, 59, 2, 60, 40, 54, 3, 61, 32, 49, 41, 55, 19, 35, 4,
		62, 52, 30, 33, 50, 12, 14, 42, 56, 16, 27, 20, 36, 23, 44, 5,
		63, 58, 39, 53, 31, 48, 18, 34, 51, 29, 11, 13, 15, 26, 22, 43,
		57, 38, 47, 17, 28, 10, 25, 21, 37, 46, 9, 24, 45, 8, 7, 6,
	}
)

func removeRightmost(x uint64) uint64 {
	return x & (x - 1)
}

func extractRightmost(x uint64) uint64 {
	return x & -x
}

func smearRightmost(x uint64) uint64 {
	return x ^ (x - 1)
}

func rightmostIndex(x uint64) int {
	r := extractRightmost(x)
	return tarzTable[(r*tarzHash)>>58]
}

func popCount(x uint64) int {
	x = (x & 0x5555555555555555) + ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x & 0x0F0F0F0F0F0F0F0F) + ((x >> 4) & 0x0F0F0F0F0F0F0F0F)
	x += x >> 8
	x += x >> 16
	x += x >> 32
	return int(x & 0x7f)
}

func bitmapCount(bm []uint64) int {
	c := 0
	for _, x := range bm {
		c += popCount(x)
	}

	return c
}
