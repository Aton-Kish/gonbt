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

package nbt

import (
	"github.com/Aton-Kish/gonbt/slices"
)

type tagTestCase[T Payload] struct {
	name string
	nbt  nbtTestCase[T]
	snbt snbtTestCase
	json jsonTestCase
	raw  rawTestCase
}

type nbtTestCase[T Payload] struct {
	tagType TagType
	tagName TagName
	payload T
}

type snbtTestCase struct {
	tagName string
	payload stringifyType
}

type jsonTestCase struct {
	tagName string
	payload stringifyType
}

type stringifyType struct {
	typeDefault string
	typeCompact string
	typePretty  string
}

type rawTestCase struct {
	tagType []byte
	tagName []byte
	payload []byte
}

func interfacedTagTestCases[T Payload](cases []tagTestCase[T]) []tagTestCase[Payload] {
	interfacedCases := make([]tagTestCase[Payload], 0, len(cases))

	for _, c := range cases {
		interfacedCases = append(interfacedCases, tagTestCase[Payload]{
			name: c.name,
			nbt: nbtTestCase[Payload]{
				tagType: c.nbt.tagType,
				tagName: c.nbt.tagName,
				payload: c.nbt.payload,
			},
			snbt: c.snbt,
			json: c.json,
			raw:  c.raw,
		})
	}

	return interfacedCases
}

var nbtCases = []struct {
	name string
	nbt  Tag
	snbt stringifyType
	json stringifyType
	raw  []byte
}{
	{
		name: `positive case: Simple`,
		nbt: NewCompoundTag(NewTagName(``), NewCompoundPayload(
			NewCompoundTag(NewTagName(`Hello World`), NewCompoundPayload(
				NewStringTag(NewTagName(`Name`), NewStringPayload(`Steve`)),
				NewEndTag(),
			)),
			NewEndTag(),
		)),
		snbt: stringifyType{
			typeDefault: `{"Hello World": {Name: "Steve"}}`,
			typeCompact: `{"Hello World":{Name:"Steve"}}`,
			typePretty: `{
  "Hello World": {
    Name: "Steve"
  }
}`,
		},
		json: stringifyType{
			typeDefault: `{"Hello World": {"Name": "Steve"}}`,
			typeCompact: `{"Hello World":{"Name":"Steve"}}`,
			typePretty: `{
  "Hello World": {
    "Name": "Steve"
  }
}`,
		},
		raw: []byte{
			// CompoundTag():
			0x0A,
			0x00, 0x00,
			//   - CompoundTag(Hello World):
			0x0A,
			0x00, 0x0B,
			0x48, 0x65, 0x6C, 0x6C, 0x6F, 0x20, 0x57, 0x6F, 0x72, 0x6C, 0x64,
			//       - StringTag(Name): "Steve"
			0x08,
			0x00, 0x04,
			0x4E, 0x61, 0x6D, 0x65,
			0x00, 0x05,
			0x53, 0x74, 0x65, 0x76, 0x65,
			//       - EndTag
			0x00,
			//   - EndTag
			0x00,
		},
	},
	{
		name: `positive case: Tag Check`,
		nbt: NewCompoundTag(NewTagName(`Compound`), NewCompoundPayload(
			NewShortTag(NewTagName(`Short`), NewShortPayload(12345)),
			NewByteArrayTag(NewTagName(`ByteArray`), NewByteArrayPayload(0, 1)),
			NewStringTag(NewTagName(`String`), NewStringPayload(`Hello`)),
			NewListTag(NewTagName(`List`), NewListPayload(NewBytePayload(123))),
			NewCompoundTag(NewTagName(`Compound`), NewCompoundPayload(
				NewStringTag(NewTagName(`String`), NewStringPayload(`World`)),
				NewEndTag(),
			)),
			NewEndTag(),
		)),
		snbt: stringifyType{
			typeDefault: `{Compound: {ByteArray: [B; 0b, 1b], Compound: {String: "World"}, List: [123b], Short: 12345s, String: "Hello"}}`,
			typeCompact: `{Compound:{ByteArray:[B;0b,1b],Compound:{String:"World"},List:[123b],Short:12345s,String:"Hello"}}`,
			typePretty: `{
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
		},
		json: stringifyType{
			typeDefault: `{"Compound": {"ByteArray": [0, 1], "Compound": {"String": "World"}, "List": [123], "Short": 12345, "String": "Hello"}}`,
			typeCompact: `{"Compound":{"ByteArray":[0,1],"Compound":{"String":"World"},"List":[123],"Short":12345,"String":"Hello"}}`,
			typePretty: `{
  "Compound": {
    "ByteArray": [
      0,
      1
    ],
    "Compound": {
      "String": "World"
    },
    "List": [
      123
    ],
    "Short": 12345,
    "String": "Hello"
  }
}`,
		},
		raw: []byte{
			// CompoundTag(Compound):
			0x0A,
			0x00, 0x08,
			0x43, 0x6F, 0x6D, 0x70, 0x6F, 0x75, 0x6E, 0x64,
			//   - ShortTag(Short): 12345s
			0x02,
			0x00, 0x05,
			0x53, 0x68, 0x6f, 0x72, 0x74,
			0x30, 0x39,
			//   - ByteArrayTag(ByteArray): [B; 0b, 1b]
			0x07,
			0x00, 0x09,
			0x42, 0x79, 0x74, 0x65, 0x41, 0x72, 0x72, 0x61, 0x79,
			0x00, 0x00, 0x00, 0x02,
			0x00, 0x01,
			//   - StringTag(String): "Hello"
			0x08,
			0x00, 0x06,
			0x53, 0x74, 0x72, 0x69, 0x6E, 0x67,
			0x00, 0x05,
			0x48, 0x65, 0x6C, 0x6C, 0x6F,
			//   - ListTag(List): [123b]
			0x09,
			0x00, 0x04,
			0x4C, 0x69, 0x73, 0x74,
			0x01,
			0x00, 0x00, 0x00, 0x01,
			0x7B,
			//   - CompoundTag(Compound):
			0x0A,
			0x00, 0x08,
			0x43, 0x6F, 0x6D, 0x70, 0x6F, 0x75, 0x6E, 0x64,
			//       - StringTag(String): "World"
			0x08,
			0x00, 0x06,
			0x53, 0x74, 0x72, 0x69, 0x6E, 0x67,
			0x00, 0x05,
			0x57, 0x6F, 0x72, 0x6C, 0x64,
			//       - EndTag
			0x00,
			//   - EndTag
			0x00,
		},
	},
}

var endTagCases = []tagTestCase[Payload]{
	{
		name: `positive case: EndTag`,
		nbt: nbtTestCase[Payload]{
			tagType: EndType,
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: End(=0)
				0x00,
			},
		},
	},
}

var byteTagCases = []tagTestCase[*BytePayload]{
	{
		name: `positive case: ByteTag`,
		nbt: nbtTestCase[*BytePayload]{
			tagType: ByteType,
			tagName: `Byte`,
			payload: NewBytePayload(123),
		},
		snbt: snbtTestCase{
			tagName: `Byte`,
			payload: stringifyType{
				typeDefault: `123b`,
				typeCompact: `123b`,
				typePretty:  `123b`,
			},
		},
		json: jsonTestCase{
			tagName: `"Byte"`,
			payload: stringifyType{
				typeDefault: `123`,
				typeCompact: `123`,
				typePretty:  `123`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: Byte(=1)
				0x01,
			},
			tagName: []byte{
				// Name Length: 4
				0x00, 0x04,
				// Name: Byte
				0x42, 0x79, 0x74, 0x65,
			},
			payload: []byte{
				// Payload: 123b
				0x7B,
			},
		},
	},
}

var shortTagCases = []tagTestCase[*ShortPayload]{
	{
		name: `positive case: ShortTag`,
		nbt: nbtTestCase[*ShortPayload]{
			tagType: ShortType,
			tagName: `Short`,
			payload: NewShortPayload(12345),
		},
		snbt: snbtTestCase{
			tagName: `Short`,
			payload: stringifyType{
				typeDefault: `12345s`,
				typeCompact: `12345s`,
				typePretty:  `12345s`,
			},
		},
		json: jsonTestCase{
			tagName: `"Short"`,
			payload: stringifyType{
				typeDefault: `12345`,
				typeCompact: `12345`,
				typePretty:  `12345`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: Short(=2)
				0x02,
			},
			tagName: []byte{
				// Name Length: 5
				0x00, 0x05,
				// Name: Short
				0x53, 0x68, 0x6f, 0x72, 0x74,
			},
			payload: []byte{
				// Payload: 12345s
				0x30, 0x39,
			},
		},
	},
}

var intTagCases = []tagTestCase[*IntPayload]{
	{
		name: `positive case: IntTag`,
		nbt: nbtTestCase[*IntPayload]{
			tagType: IntType,
			tagName: `Int`,
			payload: NewIntPayload(123456789),
		},
		snbt: snbtTestCase{
			tagName: `Int`,
			payload: stringifyType{
				typeDefault: `123456789`,
				typeCompact: `123456789`,
				typePretty:  `123456789`,
			},
		},
		json: jsonTestCase{
			tagName: `"Int"`,
			payload: stringifyType{
				typeDefault: `123456789`,
				typeCompact: `123456789`,
				typePretty:  `123456789`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: Int(=3)
				0x03,
			},
			tagName: []byte{
				// Name Length: 3
				0x00, 0x03,
				// Name: Int
				0x49, 0x6E, 0x74,
			},
			payload: []byte{
				// Payload: 123456789
				0x07, 0x5B, 0xCD, 0x15,
			},
		},
	},
}

var longTagCases = []tagTestCase[*LongPayload]{
	{
		name: `positive case: LongTag`,
		nbt: nbtTestCase[*LongPayload]{
			tagType: LongType,
			tagName: `Long`,
			payload: NewLongPayload(123456789123456789),
		},
		snbt: snbtTestCase{
			tagName: `Long`,
			payload: stringifyType{
				typeDefault: `123456789123456789L`,
				typeCompact: `123456789123456789L`,
				typePretty:  `123456789123456789L`,
			},
		},
		json: jsonTestCase{
			tagName: `"Long"`,
			payload: stringifyType{
				typeDefault: `123456789123456789`,
				typeCompact: `123456789123456789`,
				typePretty:  `123456789123456789`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: Long(=4)
				0x04,
			},
			tagName: []byte{
				// Name Length: 4
				0x00, 0x04,
				// Name: Long
				0x4C, 0x6F, 0x6E, 0x67,
			},
			payload: []byte{
				// Payload: 123456789123456789L
				0x01, 0xB6, 0x9B, 0x4B, 0xAC, 0xD0, 0x5F, 0x15,
			},
		},
	},
}

var floatTagCases = []tagTestCase[*FloatPayload]{
	{
		name: `positive case: FloatTag - %f`,
		nbt: nbtTestCase[*FloatPayload]{
			tagType: FloatType,
			tagName: `Float`,
			payload: NewFloatPayload(0.12345678),
		},
		snbt: snbtTestCase{
			tagName: `Float`,
			payload: stringifyType{
				typeDefault: `0.12345678f`,
				typeCompact: `0.12345678f`,
				typePretty:  `0.12345678f`,
			},
		},
		json: jsonTestCase{
			tagName: `"Float"`,
			payload: stringifyType{
				typeDefault: `0.12345678`,
				typeCompact: `0.12345678`,
				typePretty:  `0.12345678`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: Float(=5)
				0x05,
			},
			tagName: []byte{
				// Name Length: 5
				0x00, 0x05,
				// Name: Float
				0x46, 0x6C, 0x6F, 0x61, 0x74,
			},
			payload: []byte{
				// Payload: 0.12345678f
				0x3D, 0xFC, 0xD6, 0xE9,
			},
		},
	},
	{
		name: `positive case: FloatTag - +%e`,
		nbt: nbtTestCase[*FloatPayload]{
			tagType: FloatType,
			tagName: `Float`,
			payload: NewFloatPayload(1234567.8),
		},
		snbt: snbtTestCase{
			tagName: `Float`,
			payload: stringifyType{
				typeDefault: `1.2345678e+06f`,
				typeCompact: `1.2345678e+06f`,
				typePretty:  `1.2345678e+06f`,
			},
		},
		json: jsonTestCase{
			tagName: `"Float"`,
			payload: stringifyType{
				typeDefault: `1.2345678e+06`,
				typeCompact: `1.2345678e+06`,
				typePretty:  `1.2345678e+06`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: Float(=5)
				0x05,
			},
			tagName: []byte{
				// Name Length: 5
				0x00, 0x05,
				// Name: Float
				0x46, 0x6C, 0x6F, 0x61, 0x74,
			},
			payload: []byte{
				// Payload: 1.2345678e+06f
				0x49, 0x96, 0xB4, 0x3E,
			},
		},
	},
	{
		name: `positive case: FloatTag - -%e`,
		nbt: nbtTestCase[*FloatPayload]{
			tagType: FloatType,
			tagName: `Float`,
			payload: NewFloatPayload(0.000012345678),
		},
		snbt: snbtTestCase{
			tagName: `Float`,
			payload: stringifyType{
				typeDefault: `1.2345678e-05f`,
				typeCompact: `1.2345678e-05f`,
				typePretty:  `1.2345678e-05f`,
			},
		},
		json: jsonTestCase{
			tagName: `"Float"`,
			payload: stringifyType{
				typeDefault: `1.2345678e-05`,
				typeCompact: `1.2345678e-05`,
				typePretty:  `1.2345678e-05`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: Float(=5)
				0x05,
			},
			tagName: []byte{
				// Name Length: 5
				0x00, 0x05,
				// Name: Float
				0x46, 0x6C, 0x6F, 0x61, 0x74,
			},
			payload: []byte{
				// Payload: 1.2345678e-05f
				0x37, 0x4F, 0x20, 0x49,
			},
		},
	},
}

var doubleTagCases = []tagTestCase[*DoublePayload]{
	{
		name: `positive case: DoubleTag - %f`,
		nbt: nbtTestCase[*DoublePayload]{
			tagType: DoubleType,
			tagName: `Double`,
			payload: NewDoublePayload(0.123456789),
		},
		snbt: snbtTestCase{
			tagName: `Double`,
			payload: stringifyType{
				typeDefault: `0.123456789d`,
				typeCompact: `0.123456789d`,
				typePretty:  `0.123456789d`,
			},
		},
		json: jsonTestCase{
			tagName: `"Double"`,
			payload: stringifyType{
				typeDefault: `0.123456789`,
				typeCompact: `0.123456789`,
				typePretty:  `0.123456789`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: Double(=6)
				0x06,
			},
			tagName: []byte{
				// Name Length: 6
				0x00, 0x06,
				// Name: Double
				0x44, 0x6F, 0x75, 0x62, 0x6C, 0x65,
			},
			payload: []byte{
				// Payload: 0.123456789d
				0x3F, 0xBF, 0x9A, 0xDD, 0x37, 0x39, 0x63, 0x5F,
			},
		},
	},
	{
		name: `positive case: DoubleTag - +%g`,
		nbt: nbtTestCase[*DoublePayload]{
			tagType: DoubleType,
			tagName: `Double`,
			payload: NewDoublePayload(1234567.89),
		},
		snbt: snbtTestCase{
			tagName: `Double`,
			payload: stringifyType{
				typeDefault: `1.23456789e+06d`,
				typeCompact: `1.23456789e+06d`,
				typePretty:  `1.23456789e+06d`,
			},
		},
		json: jsonTestCase{
			tagName: `"Double"`,
			payload: stringifyType{
				typeDefault: `1.23456789e+06`,
				typeCompact: `1.23456789e+06`,
				typePretty:  `1.23456789e+06`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: Double(=6)
				0x06,
			},
			tagName: []byte{
				// Name Length: 6
				0x00, 0x06,
				// Name: Double
				0x44, 0x6F, 0x75, 0x62, 0x6C, 0x65,
			},
			payload: []byte{
				// Payload: 1.23456789e+06d
				0x41, 0x32, 0xD6, 0x87, 0xE3, 0xD7, 0x0A, 0x3D,
			},
		},
	},
	{
		name: `positive case: DoubleTag - -%g`,
		nbt: nbtTestCase[*DoublePayload]{
			tagType: DoubleType,
			tagName: `Double`,
			payload: NewDoublePayload(0.0000123456789),
		},
		snbt: snbtTestCase{
			tagName: `Double`,
			payload: stringifyType{
				typeDefault: `1.23456789e-05d`,
				typeCompact: `1.23456789e-05d`,
				typePretty:  `1.23456789e-05d`,
			},
		},
		json: jsonTestCase{
			tagName: `"Double"`,
			payload: stringifyType{
				typeDefault: `1.23456789e-05`,
				typeCompact: `1.23456789e-05`,
				typePretty:  `1.23456789e-05`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: Double(=6)
				0x06,
			},
			tagName: []byte{
				// Name Length: 6
				0x00, 0x06,
				// Name: Double
				0x44, 0x6F, 0x75, 0x62, 0x6C, 0x65,
			},
			payload: []byte{
				// Payload: 1.23456789e-05d
				0x3E, 0xE9, 0xE4, 0x09, 0x30, 0x1B, 0x5A, 0x02,
			},
		},
	},
}

var byteArrayTagCases = []tagTestCase[*ByteArrayPayload]{
	{
		name: `positive case: ByteArrayTag - has items`,
		nbt: nbtTestCase[*ByteArrayPayload]{
			tagType: ByteArrayType,
			tagName: `ByteArray`,
			payload: NewByteArrayPayload(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
		},
		snbt: snbtTestCase{
			tagName: `ByteArray`,
			payload: stringifyType{
				typeDefault: `[B; 0b, 1b, 2b, 3b, 4b, 5b, 6b, 7b, 8b, 9b]`,
				typeCompact: `[B;0b,1b,2b,3b,4b,5b,6b,7b,8b,9b]`,
				typePretty:  `[B; 0b, 1b, 2b, 3b, 4b, 5b, 6b, 7b, 8b, 9b]`,
			},
		},
		json: jsonTestCase{
			tagName: `"ByteArray"`,
			payload: stringifyType{
				typeDefault: `[0, 1, 2, 3, 4, 5, 6, 7, 8, 9]`,
				typeCompact: `[0,1,2,3,4,5,6,7,8,9]`,
				typePretty: `[
  0,
  1,
  2,
  3,
  4,
  5,
  6,
  7,
  8,
  9
]`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: ByteArray(=7)
				0x07,
			},
			tagName: []byte{
				// Name Length: 9
				0x00, 0x09,
				// Name: ByteArray
				0x42, 0x79, 0x74, 0x65, 0x41, 0x72, 0x72, 0x61, 0x79,
			},
			payload: []byte{
				// Payload Length: 10
				0x00, 0x00, 0x00, 0x0A,
				// Payload: [B; 0b, 1b, 2b, 3b, 4b, 5b, 6b, 7b, 8b, 9b]
				0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09,
			},
		},
	},
	{
		name: `positive case: ByteArrayTag - empty`,
		nbt: nbtTestCase[*ByteArrayPayload]{
			tagType: ByteArrayType,
			tagName: `ByteArray`,
			payload: NewByteArrayPayload(),
		},
		snbt: snbtTestCase{
			tagName: `ByteArray`,
			payload: stringifyType{
				typeDefault: `[B; ]`,
				typeCompact: `[B;]`,
				typePretty:  `[B; ]`,
			},
		},
		json: jsonTestCase{
			tagName: `"ByteArray"`,
			payload: stringifyType{
				typeDefault: `[]`,
				typeCompact: `[]`,
				typePretty:  `[]`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: ByteArray(=7)
				0x07,
			},
			tagName: []byte{
				// Name Length: 9
				0x00, 0x09,
				// Name: ByteArray
				0x42, 0x79, 0x74, 0x65, 0x41, 0x72, 0x72, 0x61, 0x79,
			},
			payload: []byte{
				// Payload Length: 0
				0x00, 0x00, 0x00, 0x00,
				// Payload: [B; ]
			},
		},
	},
}

var stringTagCases = []tagTestCase[*StringPayload]{
	{
		name: `positive case: StringTag - "Hello World"`,
		nbt: nbtTestCase[*StringPayload]{
			tagType: StringType,
			tagName: `String`,
			payload: NewStringPayload(`Hello World`),
		},
		snbt: snbtTestCase{
			tagName: `String`,
			payload: stringifyType{
				typeDefault: `"Hello World"`,
				typeCompact: `"Hello World"`,
				typePretty:  `"Hello World"`,
			},
		},
		json: jsonTestCase{
			tagName: `"String"`,
			payload: stringifyType{
				typeDefault: `"Hello World"`,
				typeCompact: `"Hello World"`,
				typePretty:  `"Hello World"`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: String(=8)
				0x08,
			},
			tagName: []byte{
				// Name Length: 6
				0x00, 0x06,
				// Name: String
				0x53, 0x74, 0x72, 0x69, 0x6E, 0x67,
			},
			payload: []byte{
				// Payload Length: 11
				0x00, 0x0B,
				// Payload: "Hello World"
				0x48, 0x65, 0x6C, 0x6C, 0x6F, 0x20, 0x57, 0x6F, 0x72, 0x6C, 0x64,
			},
		},
	},
	{
		name: `positive case: StringTag - "Test"`,
		nbt: nbtTestCase[*StringPayload]{
			tagType: StringType,
			tagName: `String`,
			payload: NewStringPayload(`Test`),
		},
		snbt: snbtTestCase{
			tagName: `String`,
			payload: stringifyType{
				typeDefault: `"Test"`,
				typeCompact: `"Test"`,
				typePretty:  `"Test"`,
			},
		},
		json: jsonTestCase{
			tagName: `"String"`,
			payload: stringifyType{
				typeDefault: `"Test"`,
				typeCompact: `"Test"`,
				typePretty:  `"Test"`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: String(=8)
				0x08,
			},
			tagName: []byte{
				// Name Length: 6
				0x00, 0x06,
				// Name: String
				0x53, 0x74, 0x72, 0x69, 0x6E, 0x67,
			},
			payload: []byte{
				// Payload Length: 4
				0x00, 0x04,
				// Payload: "Test"
				0x54, 0x65, 0x73, 0x74,
			},
		},
	},
	{
		name: `positive case: StringTag - '"Test'`,
		nbt: nbtTestCase[*StringPayload]{
			tagType: StringType,
			tagName: `String`,
			payload: NewStringPayload(`"Test`),
		},
		snbt: snbtTestCase{
			tagName: `String`,
			payload: stringifyType{
				typeDefault: `'"Test'`,
				typeCompact: `'"Test'`,
				typePretty:  `'"Test'`,
			},
		},
		json: jsonTestCase{
			tagName: `"String"`,
			payload: stringifyType{
				typeDefault: `"\"Test"`,
				typeCompact: `"\"Test"`,
				typePretty:  `"\"Test"`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: String(=8)
				0x08,
			},
			tagName: []byte{
				// Name Length: 6
				0x00, 0x06,
				// Name: String
				0x53, 0x74, 0x72, 0x69, 0x6E, 0x67,
			},
			payload: []byte{
				// Payload Length: 5
				0x00, 0x05,
				// Payload: '"Test'
				0x22, 0x54, 0x65, 0x73, 0x74,
			},
		},
	},
	{
		name: `positive case: StringTag - "'Test"`,
		nbt: nbtTestCase[*StringPayload]{
			tagType: StringType,
			tagName: `String`,
			payload: NewStringPayload(`'Test`),
		},
		snbt: snbtTestCase{
			tagName: `String`,
			payload: stringifyType{
				typeDefault: `"'Test"`,
				typeCompact: `"'Test"`,
				typePretty:  `"'Test"`,
			},
		},
		json: jsonTestCase{
			tagName: `"String"`,
			payload: stringifyType{
				typeDefault: `"'Test"`,
				typeCompact: `"'Test"`,
				typePretty:  `"'Test"`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: String(=8)
				0x08,
			},
			tagName: []byte{
				// Name Length: 6
				0x00, 0x06,
				// Name: String
				0x53, 0x74, 0x72, 0x69, 0x6E, 0x67,
			},
			payload: []byte{
				// Payload Length: 5
				0x00, 0x05,
				// Payload: "'Test"
				0x27, 0x54, 0x65, 0x73, 0x74,
			},
		},
	},
	{
		name: `positive case: StringTag - "\"'Test"`,
		nbt: nbtTestCase[*StringPayload]{
			tagType: StringType,
			tagName: `String`,
			payload: NewStringPayload(`"'Test`),
		},
		snbt: snbtTestCase{
			tagName: `String`,
			payload: stringifyType{
				typeDefault: `"\"'Test"`,
				typeCompact: `"\"'Test"`,
				typePretty:  `"\"'Test"`,
			},
		},
		json: jsonTestCase{
			tagName: `"String"`,
			payload: stringifyType{
				typeDefault: `"\"'Test"`,
				typeCompact: `"\"'Test"`,
				typePretty:  `"\"'Test"`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: String(=8)
				0x08,
			},
			tagName: []byte{
				// Name Length: 6
				0x00, 0x06,
				// Name: String
				0x53, 0x74, 0x72, 0x69, 0x6E, 0x67,
			},
			payload: []byte{
				// Payload Length: 6
				0x00, 0x06,
				// Payload: "\"'Test"
				0x22, 0x27, 0x54, 0x65, 0x73, 0x74,
			},
		},
	},
	{
		name: `positive case: StringTag - "minecraft:the_end"`,
		nbt: nbtTestCase[*StringPayload]{
			tagType: StringType,
			tagName: `String`,
			payload: NewStringPayload(`minecraft:the_end`),
		},
		snbt: snbtTestCase{
			tagName: `String`,
			payload: stringifyType{
				typeDefault: `"minecraft:the_end"`,
				typeCompact: `"minecraft:the_end"`,
				typePretty:  `"minecraft:the_end"`,
			},
		},
		json: jsonTestCase{
			tagName: `"String"`,
			payload: stringifyType{
				typeDefault: `"minecraft:the_end"`,
				typeCompact: `"minecraft:the_end"`,
				typePretty:  `"minecraft:the_end"`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: String(=8)
				0x08,
			},
			tagName: []byte{
				// Name Length: 6
				0x00, 0x06,
				// Name: String
				0x53, 0x74, 0x72, 0x69, 0x6E, 0x67,
			},
			payload: []byte{
				// Payload Length: 17
				0x00, 0x11,
				// Payload: "minecraft:the_end"
				0x6D, 0x69, 0x6E, 0x65, 0x63, 0x72, 0x61, 0x66, 0x74, 0x3A,
				0x74, 0x68, 0x65, 0x5F, 0x65, 0x6E, 0x64,
			},
		},
	},
	{
		name: `positive case: StringTag - ""`,
		nbt: nbtTestCase[*StringPayload]{
			tagType: StringType,
			tagName: `String`,
			payload: NewStringPayload(``),
		},
		snbt: snbtTestCase{
			tagName: `String`,
			payload: stringifyType{
				typeDefault: `""`,
				typeCompact: `""`,
				typePretty:  `""`,
			},
		},
		json: jsonTestCase{
			tagName: `"String"`,
			payload: stringifyType{
				typeDefault: `""`,
				typeCompact: `""`,
				typePretty:  `""`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: String(=8)
				0x08,
			},
			tagName: []byte{
				// Name Length: 6
				0x00, 0x06,
				// Name: String
				0x53, 0x74, 0x72, 0x69, 0x6E, 0x67,
			},
			payload: []byte{
				// Payload Length: 0
				0x00, 0x00,
				// Payload: ""
			},
		},
	},
	{
		name: `positive case: StringTag - "マインクラフト"`,
		nbt: nbtTestCase[*StringPayload]{
			tagType: StringType,
			tagName: `String`,
			payload: NewStringPayload(`マインクラフト`),
		},
		snbt: snbtTestCase{
			tagName: `String`,
			payload: stringifyType{
				typeDefault: `"マインクラフト"`,
				typeCompact: `"マインクラフト"`,
				typePretty:  `"マインクラフト"`,
			},
		},
		json: jsonTestCase{
			tagName: `"String"`,
			payload: stringifyType{
				typeDefault: `"マインクラフト"`,
				typeCompact: `"マインクラフト"`,
				typePretty:  `"マインクラフト"`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: String(=8)
				0x08,
			},
			tagName: []byte{
				// Name Length: 6
				0x00, 0x06,
				// Name: String
				0x53, 0x74, 0x72, 0x69, 0x6E, 0x67,
			},
			payload: []byte{
				// Payload Length: 21
				0x00, 0x15,
				// Payload: "マインクラフト"
				0xE3, 0x83, 0x9E, 0xE3, 0x82, 0xA4, 0xE3, 0x83, 0xB3,
				0xE3, 0x82, 0xAF, 0xE3, 0x83, 0xA9, 0xE3, 0x83, 0x95,
				0xE3, 0x83, 0x88,
			},
		},
	},
}

var listTagCases = []tagTestCase[*ListPayload]{
	{
		name: `positive case: ListTag - Short`,
		nbt: nbtTestCase[*ListPayload]{
			tagType: ListType,
			tagName: `List`,
			payload: NewListPayload(NewShortPayload(12345), NewShortPayload(6789)),
		},
		snbt: snbtTestCase{
			tagName: `List`,
			payload: stringifyType{
				typeDefault: `[12345s, 6789s]`,
				typeCompact: `[12345s,6789s]`,
				typePretty: `[
  12345s,
  6789s
]`,
			},
		},
		json: jsonTestCase{
			tagName: `"List"`,
			payload: stringifyType{
				typeDefault: `[12345, 6789]`,
				typeCompact: `[12345,6789]`,
				typePretty: `[
  12345,
  6789
]`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: List(=9)
				0x09,
			},
			tagName: []byte{
				// Name Length: 4
				0x00, 0x04,
				// Name: List
				0x4C, 0x69, 0x73, 0x74,
			},
			payload: []byte{
				// Payload Type: TagShort(=2)
				0x02,
				// Payload Length: 2
				0x00, 0x00, 0x00, 0x02,
				// Payload:
				//   - 12345s
				0x30, 0x39,
				//   - 6789s
				0x1A, 0x85,
			},
		},
	},
	{
		name: `positive case: ListTag - ByteArray`,
		nbt: nbtTestCase[*ListPayload]{
			tagType: ListType,
			tagName: `List`,
			payload: NewListPayload(NewByteArrayPayload(0, 1), NewByteArrayPayload(2, 3)),
		},
		snbt: snbtTestCase{
			tagName: `List`,
			payload: stringifyType{
				typeDefault: `[[B; 0b, 1b], [B; 2b, 3b]]`,
				typeCompact: `[[B;0b,1b],[B;2b,3b]]`,
				typePretty: `[
  [B; 0b, 1b],
  [B; 2b, 3b]
]`,
			},
		},
		json: jsonTestCase{
			tagName: `"List"`,
			payload: stringifyType{
				typeDefault: `[[0, 1], [2, 3]]`,
				typeCompact: `[[0,1],[2,3]]`,
				typePretty: `[
  [
    0,
    1
  ],
  [
    2,
    3
  ]
]`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: List(=9)
				0x09,
			},
			tagName: []byte{
				// Name Length: 4
				0x00, 0x04,
				// Name: List
				0x4C, 0x69, 0x73, 0x74,
			},
			payload: []byte{
				// Payload Type: TagByteArray(=7)
				0x07,
				// Payload Length: 2
				0x00, 0x00, 0x00, 0x02,
				// Payload:
				//   - [B; 0b, 1b]
				0x00, 0x00, 0x00, 0x02,
				0x00, 0x01,
				//   - [B; 2b, 3b]
				0x00, 0x00, 0x00, 0x02,
				0x02, 0x03,
			},
		},
	},
	{
		name: `positive case: ListTag - String`,
		nbt: nbtTestCase[*ListPayload]{
			tagType: ListType,
			tagName: `List`,
			payload: NewListPayload(NewStringPayload(`Hello`), NewStringPayload(`World`)),
		},
		snbt: snbtTestCase{
			tagName: `List`,
			payload: stringifyType{
				typeDefault: `["Hello", "World"]`,
				typeCompact: `["Hello","World"]`,
				typePretty: `[
  "Hello",
  "World"
]`,
			},
		},
		json: jsonTestCase{
			tagName: `"List"`,
			payload: stringifyType{
				typeDefault: `["Hello", "World"]`,
				typeCompact: `["Hello","World"]`,
				typePretty: `[
  "Hello",
  "World"
]`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: List(=9)
				0x09,
			},
			tagName: []byte{
				// Name Length: 4
				0x00, 0x04,
				// Name: List
				0x4C, 0x69, 0x73, 0x74,
			},
			payload: []byte{
				// Payload Type: TagString(=8)
				0x08,
				// Payload Length: 2
				0x00, 0x00, 0x00, 0x02,
				// Payload:
				//   - "Hello"
				0x00, 0x05,
				0x48, 0x65, 0x6C, 0x6C, 0x6F,
				//   - "World"
				0x00, 0x05,
				0x57, 0x6F, 0x72, 0x6C, 0x64,
			},
		},
	},
	{
		name: `positive case: ListTag - List`,
		nbt: nbtTestCase[*ListPayload]{
			tagType: ListType,
			tagName: `List`,
			payload: NewListPayload(
				NewListPayload(NewBytePayload(123)),
				NewListPayload(NewStringPayload(`Test`)),
			),
		},
		snbt: snbtTestCase{
			tagName: `List`,
			payload: stringifyType{
				typeDefault: `[[123b], ["Test"]]`,
				typeCompact: `[[123b],["Test"]]`,
				typePretty: `[
  [
    123b
  ],
  [
    "Test"
  ]
]`,
			},
		},
		json: jsonTestCase{
			tagName: `"List"`,
			payload: stringifyType{
				typeDefault: `[[123], ["Test"]]`,
				typeCompact: `[[123],["Test"]]`,
				typePretty: `[
  [
    123
  ],
  [
    "Test"
  ]
]`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: List(=9)
				0x09,
			},
			tagName: []byte{
				// Name Length: 4
				0x00, 0x04,
				// Name: List
				0x4C, 0x69, 0x73, 0x74,
			},
			payload: []byte{
				// Payload Type: TagList(=9)
				0x09,
				// Payload Length: 2
				0x00, 0x00, 0x00, 0x02,
				// Payload:
				//   - [123b]
				0x01,
				0x00, 0x00, 0x00, 0x01,
				0x7B,
				//   - ["Test"]
				0x08,
				0x00, 0x00, 0x00, 0x01,
				0x00, 0x04,
				0x54, 0x65, 0x73, 0x74,
			},
		},
	},
	{
		name: `positive case: ListTag - Compound`,
		nbt: nbtTestCase[*ListPayload]{
			tagType: ListType,
			tagName: `List`,
			payload: NewListPayload(
				NewCompoundPayload(NewByteTag(NewTagName(`Byte`), NewBytePayload(123)), NewEndTag()),
				NewCompoundPayload(NewStringTag(NewTagName(`String`), NewStringPayload(`Hello`)), NewEndTag()),
			),
		},
		snbt: snbtTestCase{
			tagName: `List`,
			payload: stringifyType{
				typeDefault: `[{Byte: 123b}, {String: "Hello"}]`,
				typeCompact: `[{Byte:123b},{String:"Hello"}]`,
				typePretty: `[
  {
    Byte: 123b
  },
  {
    String: "Hello"
  }
]`,
			},
		},
		json: jsonTestCase{
			tagName: `"List"`,
			payload: stringifyType{
				typeDefault: `[{"Byte": 123}, {"String": "Hello"}]`,
				typeCompact: `[{"Byte":123},{"String":"Hello"}]`,
				typePretty: `[
  {
    "Byte": 123
  },
  {
    "String": "Hello"
  }
]`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: List(=9)
				0x09,
			},
			tagName: []byte{
				// Name Length: 4
				0x00, 0x04,
				// Name: List
				0x4C, 0x69, 0x73, 0x74,
			},
			payload: []byte{
				// Payload Type: TagCompound(=10)
				0x0A,
				// Payload Length: 2
				0x00, 0x00, 0x00, 0x02,
				// Payload:
				//   - - Type: Byte(=1)
				//       Name: Byte
				//       Payload: 123b
				0x01,
				0x00, 0x04,
				0x42, 0x79, 0x74, 0x65,
				0x7B,
				//     - Type: End(=0)
				0x00,
				//   - - Type: String(=8)
				//       Name: String
				//       Payload: "Hello"
				0x08,
				0x00, 0x06,
				0x53, 0x74, 0x72, 0x69, 0x6E, 0x67,
				0x00, 0x05,
				0x48, 0x65, 0x6C, 0x6C, 0x6F,
				//     - Type: End(=0)
				0x00,
			},
		},
	},
	{
		name: `positive case: ListTag - empty`,
		nbt: nbtTestCase[*ListPayload]{
			tagType: ListType,
			tagName: `List`,
			payload: NewListPayload(),
		},
		snbt: snbtTestCase{
			tagName: `List`,
			payload: stringifyType{
				typeDefault: `[]`,
				typeCompact: `[]`,
				typePretty:  `[]`,
			},
		},
		json: jsonTestCase{
			tagName: `"List"`,
			payload: stringifyType{
				typeDefault: `[]`,
				typeCompact: `[]`,
				typePretty:  `[]`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: List(=9)
				0x09,
			},
			tagName: []byte{
				// Name Length: 4
				0x00, 0x04,
				// Name: List
				0x4C, 0x69, 0x73, 0x74,
			},
			payload: []byte{
				// Payload Type: TagEnd(=0)
				0x00,
				// Payload Length: 0
				0x00, 0x00, 0x00, 0x00,
				// Payload: []
			},
		},
	},
}

var compoundTagCases = []tagTestCase[*CompoundPayload]{
	{
		name: `positive case: CompoundTag - has items`,
		nbt: nbtTestCase[*CompoundPayload]{
			tagType: CompoundType,
			tagName: `Compound`,
			payload: NewCompoundPayload(
				NewShortTag(NewTagName(`Short`), NewShortPayload(12345)),
				NewByteArrayTag(NewTagName(`ByteArray`), NewByteArrayPayload(0, 1)),
				NewStringTag(NewTagName(`String`), NewStringPayload(`Hello`)),
				NewListTag(NewTagName(`List`), NewListPayload(NewBytePayload(123))),
				NewCompoundTag(NewTagName(`Compound`), NewCompoundPayload(NewStringTag(NewTagName(`String`), NewStringPayload("World")), NewEndTag())),
				NewEndTag(),
			),
		},
		snbt: snbtTestCase{
			tagName: `Compound`,
			payload: stringifyType{
				typeDefault: `{ByteArray: [B; 0b, 1b], Compound: {String: "World"}, List: [123b], Short: 12345s, String: "Hello"}`,
				typeCompact: `{ByteArray:[B;0b,1b],Compound:{String:"World"},List:[123b],Short:12345s,String:"Hello"}`,
				typePretty: `{
  ByteArray: [B; 0b, 1b],
  Compound: {
    String: "World"
  },
  List: [
    123b
  ],
  Short: 12345s,
  String: "Hello"
}`,
			},
		},
		json: jsonTestCase{
			tagName: `"Compound"`,
			payload: stringifyType{
				typeDefault: `{"ByteArray": [0, 1], "Compound": {"String": "World"}, "List": [123], "Short": 12345, "String": "Hello"}`,
				typeCompact: `{"ByteArray":[0,1],"Compound":{"String":"World"},"List":[123],"Short":12345,"String":"Hello"}`,
				typePretty: `{
  "ByteArray": [
    0,
    1
  ],
  "Compound": {
    "String": "World"
  },
  "List": [
    123
  ],
  "Short": 12345,
  "String": "Hello"
}`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: Compound(=10)
				0x0A,
			},
			tagName: []byte{
				// Name Length: 8
				0x00, 0x08,
				// Name: Compound
				0x43, 0x6F, 0x6D, 0x70, 0x6F, 0x75, 0x6E, 0x64,
			},
			payload: []byte{
				// Payload:
				//   - Type: Short(=2)
				//     Name: Short
				//     Payload: 12345s
				0x02,
				0x00, 0x05,
				0x53, 0x68, 0x6f, 0x72, 0x74,
				0x30, 0x39,
				//   - Type: ByteArray(=7)
				//     Name: ByteArray
				//     Payload: [B; 0b, 1b]
				0x07,
				0x00, 0x09,
				0x42, 0x79, 0x74, 0x65, 0x41, 0x72, 0x72, 0x61, 0x79,
				0x00, 0x00, 0x00, 0x02,
				0x00, 0x01,
				//   - Type: String(=8)
				//     Name: String
				//     Payload: "Hello"
				0x08,
				0x00, 0x06,
				0x53, 0x74, 0x72, 0x69, 0x6E, 0x67,
				0x00, 0x05,
				0x48, 0x65, 0x6C, 0x6C, 0x6F,
				//   - Type: List(=9)
				//     Name: List
				//     Payload: [123b]
				0x09,
				0x00, 0x04,
				0x4C, 0x69, 0x73, 0x74,
				0x01,
				0x00, 0x00, 0x00, 0x01,
				0x7B,
				//   - Type: Compound(=10)
				//     Name: Compound
				0x0A,
				0x00, 0x08,
				0x43, 0x6F, 0x6D, 0x70, 0x6F, 0x75, 0x6E, 0x64,
				//     Payload:
				//       - Type: String(=8)
				//         Name: String
				//         Payload: "World"
				0x08,
				0x00, 0x06,
				0x53, 0x74, 0x72, 0x69, 0x6E, 0x67,
				0x00, 0x05,
				0x57, 0x6F, 0x72, 0x6C, 0x64,
				//       - Type: End(=0)
				0x00,
				//   - Type: End(=0)
				0x00,
			},
		},
	},
	{
		name: `positive case: CompoundTag - empty`,
		nbt: nbtTestCase[*CompoundPayload]{
			tagType: CompoundType,
			tagName: `Compound`,
			payload: NewCompoundPayload(NewEndTag()),
		},
		snbt: snbtTestCase{
			tagName: `Compound`,
			payload: stringifyType{
				typeDefault: `{}`,
				typeCompact: `{}`,
				typePretty:  `{}`,
			},
		},
		json: jsonTestCase{
			tagName: `"Compound"`,
			payload: stringifyType{
				typeDefault: `{}`,
				typeCompact: `{}`,
				typePretty:  `{}`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: Compound(=10)
				0x0A,
			},
			tagName: []byte{
				// Name Length: 8
				0x00, 0x08,
				// Name: Compound
				0x43, 0x6F, 0x6D, 0x70, 0x6F, 0x75, 0x6E, 0x64,
			},
			payload: []byte{
				// Payload:
				//   - Type: End(=0)
				0x00,
			},
		},
	},
}

var intArrayTagCases = []tagTestCase[*IntArrayPayload]{
	{
		name: `positive case: IntArrayTag - has items`,
		nbt: nbtTestCase[*IntArrayPayload]{
			tagType: IntArrayType,
			tagName: `IntArray`,
			payload: NewIntArrayPayload(0, 1, 2, 3),
		},
		snbt: snbtTestCase{
			tagName: `IntArray`,
			payload: stringifyType{
				typeDefault: `[I; 0, 1, 2, 3]`,
				typeCompact: `[I;0,1,2,3]`,
				typePretty:  `[I; 0, 1, 2, 3]`,
			},
		},
		json: jsonTestCase{
			tagName: `"IntArray"`,
			payload: stringifyType{
				typeDefault: `[0, 1, 2, 3]`,
				typeCompact: `[0,1,2,3]`,
				typePretty: `[
  0,
  1,
  2,
  3
]`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: IntArray(=11)
				0x0B,
			},
			tagName: []byte{
				// Name Length: 8
				0x00, 0x08,
				// Name: IntArray
				0x49, 0x6E, 0x74, 0x41, 0x72, 0x72, 0x61, 0x79,
			},
			payload: []byte{
				// Payload Length: 4
				0x00, 0x00, 0x00, 0x04,
				// Payload: [I; 0, 1, 2, 3]
				0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x01,
				0x00, 0x00, 0x00, 0x02,
				0x00, 0x00, 0x00, 0x03,
			},
		},
	},
	{
		name: `positive case: IntArrayTag - empty`,
		nbt: nbtTestCase[*IntArrayPayload]{
			tagType: IntArrayType,
			tagName: `IntArray`,
			payload: NewIntArrayPayload(),
		},
		snbt: snbtTestCase{
			tagName: `IntArray`,
			payload: stringifyType{
				typeDefault: `[I; ]`,
				typeCompact: `[I;]`,
				typePretty:  `[I; ]`,
			},
		},
		json: jsonTestCase{
			tagName: `"IntArray"`,
			payload: stringifyType{
				typeDefault: `[]`,
				typeCompact: `[]`,
				typePretty:  `[]`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: IntArray(=11)
				0x0B,
			},
			tagName: []byte{
				// Name Length: 8
				0x00, 0x08,
				// Name: IntArray
				0x49, 0x6E, 0x74, 0x41, 0x72, 0x72, 0x61, 0x79,
			},
			payload: []byte{
				// Payload Length: 0
				0x00, 0x00, 0x00, 0x00,
				// Payload: [I; ]
			},
		},
	},
}

var longArrayTagCases = []tagTestCase[*LongArrayPayload]{
	{
		name: `positive case: LongArrayTag - has items`,
		nbt: nbtTestCase[*LongArrayPayload]{
			tagType: LongArrayType,
			tagName: `LongArray`,
			payload: NewLongArrayPayload(0, 1, 2, 3),
		},
		snbt: snbtTestCase{
			tagName: `LongArray`,
			payload: stringifyType{
				typeDefault: `[L; 0L, 1L, 2L, 3L]`,
				typeCompact: `[L;0L,1L,2L,3L]`,
				typePretty:  `[L; 0L, 1L, 2L, 3L]`,
			},
		},
		json: jsonTestCase{
			tagName: `"LongArray"`,
			payload: stringifyType{
				typeDefault: `[0, 1, 2, 3]`,
				typeCompact: `[0,1,2,3]`,
				typePretty: `[
  0,
  1,
  2,
  3
]`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: LongArray(=12)
				0x0C,
			},
			tagName: []byte{
				// Name Length: 9
				0x00, 0x09,
				// Name: LongArray
				0x4C, 0x6F, 0x6E, 0x67, 0x41, 0x72, 0x72, 0x61, 0x79,
			},
			payload: []byte{
				// Payload Length: 4
				0x00, 0x00, 0x00, 0x04,
				// Payload: [L; 0L, 1L, 2L, 3L]
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03,
			},
		},
	},
	{
		name: `positive case: LongArrayTag - empty`,
		nbt: nbtTestCase[*LongArrayPayload]{
			tagType: LongArrayType,
			tagName: `LongArray`,
			payload: NewLongArrayPayload(),
		},
		snbt: snbtTestCase{
			tagName: `LongArray`,
			payload: stringifyType{
				typeDefault: `[L; ]`,
				typeCompact: `[L;]`,
				typePretty:  `[L; ]`,
			},
		},
		json: jsonTestCase{
			tagName: `"LongArray"`,
			payload: stringifyType{
				typeDefault: `[]`,
				typeCompact: `[]`,
				typePretty:  `[]`,
			},
		},
		raw: rawTestCase{
			tagType: []byte{
				// Type: LongArray(=12)
				0x0C,
			},
			tagName: []byte{
				// Name Length: 9
				0x00, 0x09,
				// Name: LongArray
				0x4C, 0x6F, 0x6E, 0x67, 0x41, 0x72, 0x72, 0x61, 0x79,
			},
			payload: []byte{
				// Payload Length: 0
				0x00, 0x00, 0x00, 0x00,
				// Payload: [L; ]
			},
		},
	},
}

var (
	excludeEndTagCases = slices.Concat(
		interfacedTagTestCases(byteTagCases),
		interfacedTagTestCases(shortTagCases),
		interfacedTagTestCases(intTagCases),
		interfacedTagTestCases(longTagCases),
		interfacedTagTestCases(floatTagCases),
		interfacedTagTestCases(doubleTagCases),
		interfacedTagTestCases(byteArrayTagCases),
		interfacedTagTestCases(stringTagCases),
		interfacedTagTestCases(listTagCases),
		interfacedTagTestCases(compoundTagCases),
		interfacedTagTestCases(intArrayTagCases),
		interfacedTagTestCases(longArrayTagCases),
	)
	tagCases = slices.Concat(endTagCases, excludeEndTagCases)
)
