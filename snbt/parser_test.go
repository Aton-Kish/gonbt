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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewParser(t *testing.T) {
	cases := []struct {
		name     string
		snbt     string
		expected *Parser
	}{
		{
			name: `positive case: Simple - default`,
			snbt: `{"Hello World": {Name: "Steve"}}`,
			expected: &Parser{
				raw: []byte(`{"Hello World": {Name: "Steve"}}`),
				//                                                            }}"evetS" :emaN{ :"dlroW olleH"{
				stringMask:        []uint64{0b0000000000000000000000000000000000111111000000000011111111111100},
				quoteToken:        []uint64{0b0000000000000000000000000000000000100000100000000010000000000010},
				spaceToken:        []uint64{0b0000000000000000000000000000000000000000010000001000000000000000},
				leftBraceToken:    []uint64{0b0000000000000000000000000000000000000000000000010000000000000000},
				rightBraceToken:   []uint64{0b0000000000000000000000000000000011000000000000000000000000000000},
				leftBracketToken:  []uint64{0b0000000000000000000000000000000000000000000000000000000000000000},
				rightBracketToken: []uint64{0b0000000000000000000000000000000000000000000000000000000000000000},
				commaToken:        []uint64{0b0000000000000000000000000000000000000000000000000000000000000000},
				colonToken:        []uint64{0b0000000000000000000000000000000000000000001000000100000000000000},
				semicolonToken:    []uint64{0b0000000000000000000000000000000000000000000000000000000000000000},
				prev:              Token{index: -1},
				curr:              Token{index: 0, char: '{'},
			},
		},
		{
			name: `positive case: Simple - compact`,
			snbt: `{"Hello World":{Name:"Steve"}}`,
			expected: &Parser{
				raw: []byte(`{"Hello World":{Name:"Steve"}}`),
				//                                                              }}"evetS":emaN{:"dlroW olleH"{
				stringMask:        []uint64{0b0000000000000000000000000000000000001111110000000011111111111100},
				quoteToken:        []uint64{0b0000000000000000000000000000000000001000001000000010000000000010},
				spaceToken:        []uint64{0b0000000000000000000000000000000000000000000000000000000000000000},
				leftBraceToken:    []uint64{0b0000000000000000000000000000000000000000000000001000000000000000},
				rightBraceToken:   []uint64{0b0000000000000000000000000000000000110000000000000000000000000000},
				leftBracketToken:  []uint64{0b0000000000000000000000000000000000000000000000000000000000000000},
				rightBracketToken: []uint64{0b0000000000000000000000000000000000000000000000000000000000000000},
				commaToken:        []uint64{0b0000000000000000000000000000000000000000000000000000000000000000},
				colonToken:        []uint64{0b0000000000000000000000000000000000000000000100000100000000000000},
				semicolonToken:    []uint64{0b0000000000000000000000000000000000000000000000000000000000000000},
				prev:              Token{index: -1},
				curr:              Token{index: 0, char: '{'},
			},
		},
		{
			name: `positive case: Simple - pretty`,
			snbt: `{
  "Hello World": {
    Name: "Steve"
  }
}`,
			expected: &Parser{
				raw: []byte(`{
  "Hello World": {
    Name: "Steve"
  }
}`),
				//                                                }\n}  \n"evetS" :emaN    \n{ :"dlroW olleH"  \n{
				stringMask:        []uint64{0b000000000000000000000_0000_011111100000000000_0000111111111111000_00},
				quoteToken:        []uint64{0b000000000000000000000_0000_010000010000000000_0000100000000000100_00},
				spaceToken:        []uint64{0b000000000000000000000_1011_100000001000001111_1010000000000000011_10},
				leftBraceToken:    []uint64{0b000000000000000000000_0000_000000000000000000_0100000000000000000_00},
				rightBraceToken:   []uint64{0b000000000000000000001_0100_000000000000000000_0000000000000000000_00},
				leftBracketToken:  []uint64{0b000000000000000000000_0000_000000000000000000_0000000000000000000_00},
				rightBracketToken: []uint64{0b000000000000000000000_0000_000000000000000000_0000000000000000000_00},
				commaToken:        []uint64{0b000000000000000000000_0000_000000000000000000_0000000000000000000_00},
				colonToken:        []uint64{0b000000000000000000000_0000_000000000100000000_0001000000000000000_00},
				semicolonToken:    []uint64{0b000000000000000000000_0000_000000000000000000_0000000000000000000_00},
				prev:              Token{index: -1},
				curr:              Token{index: 0, char: '{'},
			},
		},
		{
			name: `positive case: Tag Check - default`,
			snbt: `{Compound: {ByteArray: [B; 0b, 1b], Compound: {String: "World"}, List: [123b], Short: 12345s, String: "Hello"}}`,
			expected: &Parser{
				raw: []byte(`{Compound: {ByteArray: [B; 0b, 1b], Compound: {String: "World"}, List: [123b], Short: 12345s, String: "Hello"}}`),
				//                            ,}"dlroW" :gnirtS{ :dnuopmoC ,]b1 ,b0 ;B[ :yarrAetyB{ :dnuopmoC{                     }}"olleH" :gnirtS ,s54321 :trohS ,]b321[ :tsiL
				stringMask:        []uint64{0b0011111100000000000000000000000000000000000000000000000000000000, 0b0000000000000000000111111000000000000000000000000000000000000000},
				quoteToken:        []uint64{0b0010000010000000000000000000000000000000000000000000000000000000, 0b0000000000000000000100000100000000000000000000000000000000000000},
				spaceToken:        []uint64{0b0000000001000000001000000000100001000100010000000000010000000000, 0b0000000000000000000000000010000000100000001000000100000001000001},
				leftBraceToken:    []uint64{0b0000000000000000010000000000000000000000000000000000100000000000, 0b0000000000000000000000000000000000000000000000000000000000000000},
				rightBraceToken:   []uint64{0b0100000000000000000000000000000000000000000000000000000000000000, 0b0000000000000000011000000000000000000000000000000000000000000000},
				leftBracketToken:  []uint64{0b0000000000000000000000000000000000000000100000000000000000000000, 0b0000000000000000000000000000000000000000000000000000000010000000},
				rightBracketToken: []uint64{0b0000000000000000000000000000001000000000000000000000000000000000, 0b0000000000000000000000000000000000000000000000000001000000000000},
				commaToken:        []uint64{0b1000000000000000000000000000010000100000000000000000000000000000, 0b0000000000000000000000000000000000010000000000000010000000000000},
				colonToken:        []uint64{0b0000000000100000000100000000000000000000001000000000001000000000, 0b0000000000000000000000000001000000000000000100000000000000100000},
				semicolonToken:    []uint64{0b0000000000000000000000000000000000000010000000000000000000000000, 0b0000000000000000000000000000000000000000000000000000000000000000},
				prev:              Token{index: -1},
				curr:              Token{index: 0, char: '{'},
			},
		},
		{
			name: `positive case: Tag Check - compact`,
			snbt: `{Compound:{ByteArray:[B;0b,1b],Compound:{String:"World"},List:[123b],Short:12345s,String:"Hello"}}`,
			expected: &Parser{
				raw: []byte(`{Compound:{ByteArray:[B;0b,1b],Compound:{String:"World"},List:[123b],Short:12345s,String:"Hello"}}`),
				//                            1[:tsiL,}"dlroW":gnirtS{:dnuopmoC,]b1,b0;B[:yarrAetyB{:dnuopmoC{                                  }}"olleH":gnirtS,s54321:trohS,]b32
				stringMask:        []uint64{0b0000000001111110000000000000000000000000000000000000000000000000, 0b0000000000000000000000000000000011111100000000000000000000000000},
				quoteToken:        []uint64{0b0000000001000001000000000000000000000000000000000000000000000000, 0b0000000000000000000000000000000010000010000000000000000000000000},
				spaceToken:        []uint64{0b0000000000000000000000000000000000000000000000000000000000000000, 0b0000000000000000000000000000000000000000000000000000000000000000},
				leftBraceToken:    []uint64{0b0000000000000000000000010000000000000000000000000000010000000000, 0b0000000000000000000000000000000000000000000000000000000000000000},
				rightBraceToken:   []uint64{0b0000000010000000000000000000000000000000000000000000000000000000, 0b0000000000000000000000000000001100000000000000000000000000000000},
				leftBracketToken:  []uint64{0b0100000000000000000000000000000000000000001000000000000000000000, 0b0000000000000000000000000000000000000000000000000000000000000000},
				rightBracketToken: []uint64{0b0000000000000000000000000000000000100000000000000000000000000000, 0b0000000000000000000000000000000000000000000000000000000000001000},
				commaToken:        []uint64{0b0000000100000000000000000000000001000100000000000000000000000000, 0b0000000000000000000000000000000000000000000000100000000000010000},
				colonToken:        []uint64{0b0010000000000000100000001000000000000000000100000000001000000000, 0b0000000000000000000000000000000000000001000000000000010000000000},
				semicolonToken:    []uint64{0b0000000000000000000000000000000000000000100000000000000000000000, 0b0000000000000000000000000000000000000000000000000000000000000000},
				prev:              Token{index: -1},
				curr:              Token{index: 0, char: '{'},
			},
		},
		{
			name: `positive case: Tag Check - pretty`,
			snbt: `{
  Compound: {
    ByteArray: [B; 0b, 1b],
    Compound: {
      String: "World"
    },
    List: [
      123b
    ],
    Short: 12345s,
    String: "Hello"
  }
}`,
			expected: &Parser{
				raw: []byte(`{
  Compound: {
    ByteArray: [B; 0b, 1b],
    Compound: {
      String: "World"
    },
    List: [
      123b
    ],
    Short: 12345s,
    String: "Hello"
  }
}`),
				//                                \n{ :dnuopmoC    \n,]b1 ,b0 ;B[ :yarrAetyB    \n{ :dnuopmoC  \n{    trohS    \n,]    \nb321      \n[ :tsiL    \n,}    \n"dlroW" :gnirtS                                   }\n}  \n"olleH" :gnirtS    \n,s54321 :
				stringMask:        []uint64{0b0000_0000000000000000_0000000000000000000000000000_00000000000000_00, 0b000000000_0000000_00000000000_000000000000_0000000_011111100000000000, 0b000000000000000000000000000000_0000_01111110000000000000_0000000000},
				quoteToken:        []uint64{0b0000_0000000000000000_0000000000000000000000000000_00000000000000_00, 0b000000000_0000000_00000000000_000000000000_0000000_010000010000000000, 0b000000000000000000000000000000_0000_01000001000000000000_0000000000},
				spaceToken:        []uint64{0b1111_1010000000001111_1000010001000100000000001111_10100000000011_10, 0b000001111_1001111_10000111111_101000001111_1001111_100000001000000011, 0b000000000000000000000000000000_1011_10000000100000001111_1000000010},
				leftBraceToken:    []uint64{0b0000_0100000000000000_0000000000000000000000000000_01000000000000_00, 0b000000000_0000000_00000000000_000000000000_0000000_000000000000000000, 0b000000000000000000000000000000_0000_00000000000000000000_0000000000},
				rightBraceToken:   []uint64{0b0000_0000000000000000_0000000000000000000000000000_00000000000000_00, 0b000000000_0000000_00000000000_000000000000_0010000_000000000000000000, 0b000000000000000000000000000001_0100_00000000000000000000_0000000000},
				leftBracketToken:  []uint64{0b0000_0000000000000000_0000000000001000000000000000_00000000000000_00, 0b000000000_0000000_00000000000_010000000000_0000000_000000000000000000, 0b000000000000000000000000000000_0000_00000000000000000000_0000000000},
				rightBracketToken: []uint64{0b0000_0000000000000000_0010000000000000000000000000_00000000000000_00, 0b000000000_0010000_00000000000_000000000000_0000000_000000000000000000, 0b000000000000000000000000000000_0000_00000000000000000000_0000000000},
				commaToken:        []uint64{0b0000_0000000000000000_0100001000000000000000000000_00000000000000_00, 0b000000000_0100000_00000000000_000000000000_0100000_000000000000000000, 0b000000000000000000000000000000_0000_00000000000000000000_0100000000},
				colonToken:        []uint64{0b0000_0001000000000000_0000000000000010000000000000_00010000000000_00, 0b000000000_0000000_00000000000_000100000000_0000000_000000000100000000, 0b000000000000000000000000000000_0000_00000000010000000000_0000000001},
				semicolonToken:    []uint64{0b0000_0000000000000000_0000000000100000000000000000_00000000000000_00, 0b000000000_0000000_00000000000_000000000000_0000000_000000000000000000, 0b000000000000000000000000000000_0000_00000000000000000000_0000000000},
				prev:              Token{index: -1},
				curr:              Token{index: 0, char: '{'},
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewParser(tt.snbt)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestParser_Char(t *testing.T) {
	cases := []struct {
		name        string
		snbt        string
		index       int
		expected    rune
		expectedErr error
	}{
		{
			name:     `positive case: Simple - 0`,
			snbt:     `{"Hello World": {Name: "Steve"}}`,
			index:    0,
			expected: '{',
		},
		{
			name:     `positive case: Simple - 10`,
			snbt:     `{"Hello World": {Name: "Steve"}}`,
			index:    10,
			expected: 'r',
		},
		{
			name:        `negative case: Simple - -1`,
			snbt:        `{"Hello World": {Name: "Steve"}}`,
			index:       -1,
			expectedErr: &SnbtError{Op: "char", Err: ErrOutOfRange},
		},
		{
			name:        `negative case: Simple - 32`,
			snbt:        `{"Hello World": {Name: "Steve"}}`,
			index:       32,
			expectedErr: &SnbtError{Op: "char", Err: ErrOutOfRange},
		},
		{
			name:     `positive case: Tag Check - 0`,
			snbt:     `{Compound: {ByteArray: [B; 0b, 1b], Compound: {String: "World"}, List: [123b], Short: 12345s, String: "Hello"}}`,
			index:    0,
			expected: '{',
		},
		{
			name:     `positive case: Tag Check - 10`,
			snbt:     `{Compound: {ByteArray: [B; 0b, 1b], Compound: {String: "World"}, List: [123b], Short: 12345s, String: "Hello"}}`,
			index:    10,
			expected: ' ',
		},
		{
			name:        `positive case: Tag Check - -1`,
			snbt:        `{Compound: {ByteArray: [B; 0b, 1b], Compound: {String: "World"}, List: [123b], Short: 12345s, String: "Hello"}}`,
			index:       -1,
			expectedErr: &SnbtError{Op: "char", Err: ErrOutOfRange},
		},
		{
			name:        `positive case: Tag Check - 111`,
			snbt:        `{Compound: {ByteArray: [B; 0b, 1b], Compound: {String: "World"}, List: [123b], Short: 12345s, String: "Hello"}}`,
			index:       111,
			expectedErr: &SnbtError{Op: "char", Err: ErrOutOfRange},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			p := NewParser(tt.snbt)
			actual, err := p.Char(tt.index)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, actual)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr, err)
			}
		})
	}
}

func TestParser_Slice(t *testing.T) {
	cases := []struct {
		name        string
		snbt        string
		start       int
		end         int
		expected    []byte
		expectedErr error
	}{
		{
			name:     `positive case: Simple - [0:0]`,
			snbt:     `{"Hello World": {Name: "Steve"}}`,
			start:    0,
			end:      0,
			expected: []byte{},
		},
		{
			name:     `positive case: Simple - [0:10]`,
			snbt:     `{"Hello World": {Name: "Steve"}}`,
			start:    0,
			end:      10,
			expected: []byte(`{"Hello Wo`),
		},
		{
			name:        `negative case: Simple - [-1:0]`,
			snbt:        `{"Hello World": {Name: "Steve"}}`,
			start:       -1,
			end:         0,
			expectedErr: &SnbtError{Op: "slice", Err: ErrOutOfRange},
		},
		{
			name:        `negative case: Simple - [0:33]`,
			snbt:        `{"Hello World": {Name: "Steve"}}`,
			start:       0,
			end:         33,
			expectedErr: &SnbtError{Op: "slice", Err: ErrOutOfRange},
		},
		{
			name:        `negative case: Simple - [10:0]`,
			snbt:        `{"Hello World": {Name: "Steve"}}`,
			start:       10,
			end:         0,
			expectedErr: &SnbtError{Op: "slice", Err: ErrOutOfRange},
		},
		{
			name:     `positive case: Tag Check - [0:0]`,
			snbt:     `{Compound: {ByteArray: [B; 0b, 1b], Compound: {String: "World"}, List: [123b], Short: 12345s, String: "Hello"}}`,
			start:    0,
			end:      0,
			expected: []byte{},
		},
		{
			name:     `positive case: Tag Check - [0:10]`,
			snbt:     `{Compound: {ByteArray: [B; 0b, 1b], Compound: {String: "World"}, List: [123b], Short: 12345s, String: "Hello"}}`,
			start:    0,
			end:      10,
			expected: []byte(`{Compound:`),
		},
		{
			name:        `negative case: Tag Check - [-1:0]`,
			snbt:        `{Compound: {ByteArray: [B; 0b, 1b], Compound: {String: "World"}, List: [123b], Short: 12345s, String: "Hello"}}`,
			start:       -1,
			end:         0,
			expectedErr: &SnbtError{Op: "slice", Err: ErrOutOfRange},
		},
		{
			name:        `negative case: Tag Check - [0:112]`,
			snbt:        `{Compound: {ByteArray: [B; 0b, 1b], Compound: {String: "World"}, List: [123b], Short: 12345s, String: "Hello"}}`,
			start:       0,
			end:         112,
			expectedErr: &SnbtError{Op: "slice", Err: ErrOutOfRange},
		},
		{
			name:        `negative case: Tag Check - [10:0]`,
			snbt:        `{Compound: {ByteArray: [B; 0b, 1b], Compound: {String: "World"}, List: [123b], Short: 12345s, String: "Hello"}}`,
			start:       10,
			end:         0,
			expectedErr: &SnbtError{Op: "slice", Err: ErrOutOfRange},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			p := NewParser(tt.snbt)
			actual, err := p.Slice(tt.start, tt.end)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, actual)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr, err)
			}
		})
	}
}

func TestParser_Next(t *testing.T) {
	cases := []struct {
		name        string
		snbt        string
		expected    []Token
		expectedErr error
	}{
		{
			name: `positive case: Simple`,
			snbt: `{"Hello World": {Name: "Steve"}}`,
			expected: []Token{
				{index: 0, char: '{'},
				{index: 14, char: ':'},
				{index: 16, char: '{'},
				{index: 21, char: ':'},
				{index: 30, char: '}'},
				{index: 31, char: '}'},
			},
			expectedErr: &SnbtError{Op: "next", Err: ErrStopIteration},
		},
		{
			name: `positive case: Tag Check`,
			snbt: `{Compound: {ByteArray: [B; 0b, 1b], Compound: {String: "World"}, List: [123b], Short: 12345s, String: "Hello"}}`,
			expected: []Token{
				{index: 0, char: '{'},
				{index: 9, char: ':'},
				{index: 11, char: '{'},
				{index: 21, char: ':'},
				{index: 23, char: '['},
				{index: 25, char: ';'},
				{index: 29, char: ','},
				{index: 33, char: ']'},
				{index: 34, char: ','},
				{index: 44, char: ':'},
				{index: 46, char: '{'},
				{index: 53, char: ':'},
				{index: 62, char: '}'},
				{index: 63, char: ','},
				{index: 69, char: ':'},
				{index: 71, char: '['},
				{index: 76, char: ']'},
				{index: 77, char: ','},
				{index: 84, char: ':'},
				{index: 92, char: ','},
				{index: 100, char: ':'},
				{index: 109, char: '}'},
				{index: 110, char: '}'},
			},
			expectedErr: &SnbtError{Op: "next", Err: ErrStopIteration},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			p := NewParser(tt.snbt)

			assert.Equal(t, tt.expected[0], *p.CurrToken())

			for _, expected := range tt.expected[1:] {
				err := p.Next()
				assert.NoError(t, err)
				assert.Equal(t, expected, *p.CurrToken())
			}

			err := p.Next()
			assert.Error(t, err, tt.expectedErr)
			assert.Equal(t, Token{index: len(tt.snbt)}, *p.CurrToken())
		})
	}
}

func TestParser_Compact(t *testing.T) {
	cases := []struct {
		name        string
		parser      *Parser
		expected    *Parser
		expectedErr error
	}{
		{
			name:     `positive case: Simple - default`,
			parser:   NewParser(`{"Hello World": {Name: "Steve"}}`),
			expected: NewParser(`{"Hello World":{Name:"Steve"}}`),
		},
		{
			name:     `positive case: Simple - compact`,
			parser:   NewParser(`{"Hello World":{Name:"Steve"}}`),
			expected: NewParser(`{"Hello World":{Name:"Steve"}}`),
		},
		{
			name: `positive case: Simple - pretty`,
			parser: NewParser(`{
  "Hello World": {
    Name: "Steve"
  }
}`),
			expected: NewParser(`{"Hello World":{Name:"Steve"}}`),
		},
		{
			name:     `positive case: Tag Check - default`,
			parser:   NewParser(`{Compound: {ByteArray: [B; 0b, 1b], Compound: {String: "World"}, List: [123b], Short: 12345s, String: "Hello"}}`),
			expected: NewParser(`{Compound:{ByteArray:[B;0b,1b],Compound:{String:"World"},List:[123b],Short:12345s,String:"Hello"}}`),
		},
		{
			name:     `positive case: Tag Check - compact`,
			parser:   NewParser(`{Compound:{ByteArray:[B;0b,1b],Compound:{String:"World"},List:[123b],Short:12345s,String:"Hello"}}`),
			expected: NewParser(`{Compound:{ByteArray:[B;0b,1b],Compound:{String:"World"},List:[123b],Short:12345s,String:"Hello"}}`),
		},
		{
			name: `positive case: Tag Check - pretty`,
			parser: NewParser(`{
  Compound: {
    ByteArray: [B; 0b, 1b],
    Compound: {
      String: "World"
    },
    List: [
      123b
    ],
    Short: 12345s,
    String: "Hello"
  }
}`),
			expected: NewParser(`{Compound:{ByteArray:[B;0b,1b],Compound:{String:"World"},List:[123b],Short:12345s,String:"Hello"}}`),
		},
		{
			name:     `positive case: divisible by 64`,
			parser:   NewParser(`{Compound: {ByteArray: [B; 0b, 1b], List: [123b], Short: 1234s}}`),
			expected: NewParser(`{Compound:{ByteArray:[B;0b,1b],List:[123b],Short:1234s}}`),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.parser.Compact()

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, tt.parser)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr, err)
			}
		})
	}
}
