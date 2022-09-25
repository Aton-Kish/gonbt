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

import (
	"errors"
	"strings"
)

var (
	bitmapSize = 64
)

type bitmaps []uint64

type token struct {
	index int
	char  rune
}

type parseOptions struct {
	ignoreQuote        bool
	ignoreSpace        bool
	ignoreLeftBrace    bool
	ignoreRightBrace   bool
	ignoreLeftBracket  bool
	ignoreRightBracket bool
	ignoreComma        bool
	ignoreColon        bool
	ignoreSemicolon    bool
}

type Parser struct {
	raw               []byte
	stringMask        bitmaps
	quoteToken        bitmaps
	spaceToken        bitmaps
	leftBraceToken    bitmaps
	rightBraceToken   bitmaps
	leftBracketToken  bitmaps
	rightBracketToken bitmaps
	commaToken        bitmaps
	colonToken        bitmaps
	semicolonToken    bitmaps
	prev              token
	curr              token
}

func NewParser(snbt string) *Parser {
	l := len(snbt)/bitmapSize + 1
	p := &Parser{
		raw:               []byte(snbt),
		stringMask:        make(bitmaps, l),
		quoteToken:        make(bitmaps, l),
		spaceToken:        make(bitmaps, l),
		leftBraceToken:    make(bitmaps, l),
		rightBraceToken:   make(bitmaps, l),
		leftBracketToken:  make(bitmaps, l),
		rightBracketToken: make(bitmaps, l),
		commaToken:        make(bitmaps, l),
		colonToken:        make(bitmaps, l),
		semicolonToken:    make(bitmaps, l),
		prev:              token{index: -1},
		curr:              token{index: -1},
	}

	p.parseToken()
	p.parseMask()

	return p
}

func (p *Parser) tokenBitmaps(token rune) *bitmaps {
	switch token {
	case '"':
		return &p.quoteToken
	case ' ':
		return &p.spaceToken
	case '{':
		return &p.leftBraceToken
	case '}':
		return &p.rightBraceToken
	case '[':
		return &p.leftBracketToken
	case ']':
		return &p.rightBracketToken
	case ',':
		return &p.commaToken
	case ':':
		return &p.colonToken
	case ';':
		return &p.semicolonToken
	default:
		return nil
	}
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
	quoteBitmaps := make(bitmaps, l)
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

func (p *Parser) Rune(index int) (rune, error) {
	if index < 0 || index >= len(p.raw) {
		return *new(rune), errors.New("out of range")
	}

	return rune(p.raw[index]), nil
}

func (p *Parser) Slice(start int, end int) ([]byte, error) {
	if start < 0 || start > len(p.raw) || end < 0 || end > len(p.raw) || start > end {
		return nil, errors.New("out of range")
	}

	return p.raw[start:end], nil
}

func (p *Parser) PrevToken() (int, rune) {
	return p.prev.index, p.prev.char
}

func (p *Parser) CurrToken() (int, rune) {
	return p.curr.index, p.curr.char
}

func (p *Parser) Next() error {
	return p.next()
}

func (p *Parser) next(optFns ...func(*parseOptions) error) error {
	options := parseOptions{
		ignoreQuote: true,
		ignoreSpace: true,
	}
	for _, optFn := range optFns {
		if err := optFn(&options); err != nil {
			return err
		}
	}

	l := len(p.raw)

	var char rune
	index := l

	updateTokenIndexFn := func(c rune) {
		bitmaps := p.tokenBitmaps(c)
		if bitmaps == nil {
			return
		}

		for idx, bitmap := range *bitmaps {
			if bitmap == 0 {
				continue
			}

			if i := idx*bitmapSize + rightmostIndex(bitmap); i < index {
				index = i
				char = c
			}

			break
		}
	}

	if !options.ignoreQuote {
		updateTokenIndexFn('"')
	}

	if !options.ignoreSpace {
		updateTokenIndexFn(' ')
	}

	if !options.ignoreLeftBrace {
		updateTokenIndexFn('{')
	}

	if !options.ignoreRightBrace {
		updateTokenIndexFn('}')
	}

	if !options.ignoreLeftBracket {
		updateTokenIndexFn('[')
	}

	if !options.ignoreRightBracket {
		updateTokenIndexFn(']')
	}

	if !options.ignoreComma {
		updateTokenIndexFn(',')
	}

	if !options.ignoreColon {
		updateTokenIndexFn(':')
	}

	if !options.ignoreSemicolon {
		updateTokenIndexFn(';')
	}

	p.prev = p.curr
	p.curr = token{index: index, char: char}

	if index == l || !strings.ContainsRune(`" {}[],:;`, char) {
		return errors.New("stop iteration")
	}

	bitmaps := p.tokenBitmaps(char)
	if bitmaps == nil {
		return errors.New("unexpected error")
	}

	idx := index / bitmapSize
	(*bitmaps)[idx] = removeRightmost((*bitmaps)[idx])

	return nil
}
