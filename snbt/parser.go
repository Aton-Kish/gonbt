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

package snbt

var (
	bitmapSize = 64
)

type Bitmaps []uint64

type Parser struct {
	raw               []byte
	stringMask        Bitmaps
	quoteToken        Bitmaps
	spaceToken        Bitmaps
	leftBraceToken    Bitmaps
	rightBraceToken   Bitmaps
	leftBracketToken  Bitmaps
	rightBracketToken Bitmaps
	commaToken        Bitmaps
	colonToken        Bitmaps
	semicolonToken    Bitmaps
}

func NewParser(snbt string) *Parser {
	l := len(snbt)/bitmapSize + 1

	p := &Parser{
		raw:               []byte(snbt),
		stringMask:        make(Bitmaps, l),
		quoteToken:        make(Bitmaps, l),
		spaceToken:        make(Bitmaps, l),
		leftBraceToken:    make(Bitmaps, l),
		rightBraceToken:   make(Bitmaps, l),
		leftBracketToken:  make(Bitmaps, l),
		rightBracketToken: make(Bitmaps, l),
		commaToken:        make(Bitmaps, l),
		colonToken:        make(Bitmaps, l),
		semicolonToken:    make(Bitmaps, l),
	}

	p.parseToken()
	p.parseMask()

	return p
}

func (p *Parser) parseToken() {
	isSingleQuoted := false
	isDoubleQuoted := false
	isEscaped := false

	for i, c := range p.raw {
		idx, pos := i/bitmapSize, i%bitmapSize
		switch c {
		case '\\':
			if !isEscaped {
				isEscaped = true
				continue
			}
		case '\'':
			if isEscaped || isDoubleQuoted {
				break
			}

			isSingleQuoted = !isSingleQuoted
			p.quoteToken[idx] |= 1 << pos
		case '"':
			if isEscaped || isSingleQuoted {
				break
			}

			isDoubleQuoted = !isDoubleQuoted
			p.quoteToken[idx] |= 1 << pos
		case ' ', '\t', '\n', '\r', '\f':
			p.spaceToken[idx] |= 1 << pos
		case '{':
			p.leftBraceToken[idx] |= 1 << pos
		case '}':
			p.rightBraceToken[idx] |= 1 << pos
		case '[':
			p.leftBracketToken[idx] |= 1 << pos
		case ']':
			p.rightBracketToken[idx] |= 1 << pos
		case ',':
			p.commaToken[idx] |= 1 << pos
		case ':':
			p.colonToken[idx] |= 1 << pos
		case ';':
			p.semicolonToken[idx] |= 1 << pos
		}

		isEscaped = false
	}
}

func (p *Parser) parseMask() {
	l := len(p.quoteToken)
	quoteBitmaps := make(Bitmaps, l)
	copy(quoteBitmaps, p.quoteToken)

	isQuoted := false
	for idx := 0; idx < l; idx++ {
		for quoteBitmaps[idx] != 0 {
			p.stringMask[idx] ^= smearRightmost(quoteBitmaps[idx])
			quoteBitmaps[idx] = removeRightmost(quoteBitmaps[idx])
			isQuoted = !isQuoted
		}

		if isQuoted {
			p.stringMask[idx] = ^p.stringMask[idx]
		}

		p.spaceToken[idx] &= ^p.stringMask[idx]
		p.leftBraceToken[idx] &= ^p.stringMask[idx]
		p.rightBraceToken[idx] &= ^p.stringMask[idx]
		p.leftBracketToken[idx] &= ^p.stringMask[idx]
		p.rightBracketToken[idx] &= ^p.stringMask[idx]
		p.commaToken[idx] &= ^p.stringMask[idx]
		p.colonToken[idx] &= ^p.stringMask[idx]
		p.semicolonToken[idx] &= ^p.stringMask[idx]
	}
}
