package gonbt

import (
	"go.uber.org/thriftrw/ptr"
)

var NBTValidTestCases = []struct {
	name        string
	raw         []byte
	nbt         Tag
	snbt        *string
	compactSnbt *string
	prettySnbt  *string
	json        *string
	compactJson *string
	prettyJson  *string
}{
	{
		name: `Valid Case: Simple`,
		raw: []byte{
			// CompoundTag(""):
			0x0A,
			0x00, 0x00,
			//   - CompoundTag("Hello World"):
			0x0A,
			0x00, 0x0B,
			0x48, 0x65, 0x6C, 0x6C, 0x6F, 0x20, 0x57, 0x6F, 0x72, 0x6C, 0x64,
			//       - StringTag("Name"): "Steve"
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
		nbt: &CompoundTag{
			TagName(``),
			CompoundPayload{
				&CompoundTag{
					TagName(`Hello World`),
					CompoundPayload{
						&StringTag{TagName(`Name`), StringPayload(`Steve`)},
						&EndTag{},
					},
				},
				&EndTag{},
			},
		},
		snbt:        ptr.String(`{"Hello World": {Name: "Steve"}}`),
		compactSnbt: ptr.String(`{"Hello World":{Name:"Steve"}}`),
		prettySnbt: ptr.String(`{
    "Hello World": {
        Name: "Steve"
    }
}`),
		json:        ptr.String(`{"Hello World": {"Name": "Steve"}}`),
		compactJson: ptr.String(`{"Hello World":{"Name":"Steve"}}`),
		prettyJson: ptr.String(`{
    "Hello World": {
        "Name": "Steve"
    }
}`),
	},
	{
		name: `Valid Case: Tag Check`,
		raw: []byte{
			// CompoundTag("Compound"):
			0x0A,
			0x00, 0x08,
			0x43, 0x6F, 0x6D, 0x70, 0x6F, 0x75, 0x6E, 0x64,
			//   - ShortTag("Short"): 12345
			0x02,
			0x00, 0x05,
			0x53, 0x68, 0x6f, 0x72, 0x74,
			0x30, 0x39,
			//   - ByteArrayTag("ByteArray"): [B; 0b, 1b]
			0x07,
			0x00, 0x09,
			0x42, 0x79, 0x74, 0x65, 0x41, 0x72, 0x72, 0x61, 0x79,
			0x00, 0x00, 0x00, 0x02,
			0x00, 0x01,
			//   - StringTag("String"): "Hello"
			0x08,
			0x00, 0x06,
			0x53, 0x74, 0x72, 0x69, 0x6E, 0x67,
			0x00, 0x05,
			0x48, 0x65, 0x6C, 0x6C, 0x6F,
			//   - ListTag("List"): <Byte>[123]
			0x09,
			0x00, 0x04,
			0x4C, 0x69, 0x73, 0x74,
			0x01,
			0x00, 0x00, 0x00, 0x01,
			0x7B,
			//   - CompoundTag("Compound"):
			0x0A,
			0x00, 0x08,
			0x43, 0x6F, 0x6D, 0x70, 0x6F, 0x75, 0x6E, 0x64,
			//       - StringTag("String"): "World"
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
		nbt: &CompoundTag{
			TagName(`Compound`),
			CompoundPayload{
				&ShortTag{TagName(`Short`), ShortPayload(12345)},
				&ByteArrayTag{TagName(`ByteArray`), ByteArrayPayload{0, 1}},
				&StringTag{TagName(`String`), StringPayload(`Hello`)},
				&ListTag{TagName(`List`), ListPayload{BytePayloadPtr(123)}},
				&CompoundTag{TagName(`Compound`), CompoundPayload{&StringTag{TagName(`String`), StringPayload(`World`)}, &EndTag{}}},
				&EndTag{},
			},
		},
		snbt:        ptr.String(`{Compound: {ByteArray: [B; 0b, 1b], Compound: {String: "World"}, List: [123b], Short: 12345s, String: "Hello"}}`),
		compactSnbt: ptr.String(`{Compound:{ByteArray:[B;0b,1b],Compound:{String:"World"},List:[123b],Short:12345s,String:"Hello"}}`),
		prettySnbt: ptr.String(`{
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
		json:        ptr.String(`{"Compound": {"ByteArray": [0, 1], "Compound": {"String": "World"}, "List": [123], "Short": 12345, "String": "Hello"}}`),
		compactJson: ptr.String(`{"Compound":{"ByteArray":[0,1],"Compound":{"String":"World"},"List":[123],"Short":12345,"String":"Hello"}}`),
		prettyJson: ptr.String(`{
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
}`),
	},
}

var TagNameValidTestCases = []struct {
	name string
	raw  []byte
	nbt  *TagName
	snbt *string
	json *string
}{
	{
		name: `Valid Case: Test`,
		raw: []byte{
			// Name Length: 4
			0x00, 0x04,
			// Name: Test
			0x54, 0x65, 0x73, 0x74,
		},
		nbt:  TagNamePtr(`Test`),
		snbt: ptr.String(`Test`),
		json: ptr.String(`"Test"`),
	},
	{
		name: `Valid Case: "Test`,
		raw: []byte{
			// Name Length: 5
			0x00, 0x05,
			// Name: "Test
			0x22, 0x54, 0x65, 0x73, 0x74,
		},
		nbt:  TagNamePtr(`"Test`),
		snbt: ptr.String(`'"Test'`),
		json: ptr.String(`"\"Test"`),
	},
	{
		name: `Valid Case: 'Test`,
		raw: []byte{
			// Name Length: 5
			0x00, 0x05,
			// Name: 'Test
			0x27, 0x54, 0x65, 0x73, 0x74,
		},
		nbt:  TagNamePtr(`'Test`),
		snbt: ptr.String(`"'Test"`),
		json: ptr.String(`"'Test"`),
	},
	{
		name: `Valid Case: "'Test`,
		raw: []byte{
			// Name Length: 6
			0x00, 0x06,
			// Name: "'Test
			0x22, 0x27, 0x54, 0x65, 0x73, 0x74,
		},
		nbt:  TagNamePtr(`"'Test`),
		snbt: ptr.String(`"\"'Test"`),
		json: ptr.String(`"\"'Test"`),
	},
	{
		name: `Valid Case: minecraft:the_end`,
		raw: []byte{
			// Name Length: 17
			0x00, 0x11,
			// Name: "minecraft:the_end"
			0x6D, 0x69, 0x6E, 0x65, 0x63, 0x72, 0x61, 0x66, 0x74, 0x3A,
			0x74, 0x68, 0x65, 0x5F, 0x65, 0x6E, 0x64,
		},
		nbt:  TagNamePtr(`minecraft:the_end`),
		snbt: ptr.String(`"minecraft:the_end"`),
		json: ptr.String(`"minecraft:the_end"`),
	},
	{
		name: `Valid Case: (empty)`,
		raw: []byte{
			// Name Length: 0
			0x00, 0x00,
			// Name: (empty)
		},
		nbt:  TagNamePtr(``),
		snbt: ptr.String(``),
		json: ptr.String(`""`),
	},
	{
		name: `Valid Case: マインクラフト`,
		raw: []byte{
			// Name Length: 21
			0x00, 0x15,
			// Name: マインクラフト
			0xE3, 0x83, 0x9E, 0xE3, 0x82, 0xA4, 0xE3, 0x83, 0xB3, 0xE3,
			0x82, 0xAF, 0xE3, 0x83, 0xA9, 0xE3, 0x83, 0x95, 0xE3, 0x83,
			0x88,
		},
		nbt:  TagNamePtr(`マインクラフト`),
		snbt: ptr.String(`マインクラフト`),
		json: ptr.String(`"マインクラフト"`),
	},
}

var PayloadValidTestCases = []struct {
	name        string
	raw         []byte
	nbt         Payload
	snbt        *string
	compactSnbt *string
	prettySnbt  *string
	json        *string
	compactJson *string
	prettyJson  *string
}{
	{
		name: `Valid Case: BytePayload`,
		raw: []byte{
			// Payload: 123
			0x7B,
		},
		nbt:         BytePayloadPtr(123),
		snbt:        ptr.String(`123b`),
		compactSnbt: ptr.String(`123b`),
		prettySnbt:  ptr.String(`123b`),
		json:        ptr.String(`123`),
		compactJson: ptr.String(`123`),
		prettyJson:  ptr.String(`123`),
	},
	{
		name: `Valid Case: ShortPayload`,
		raw: []byte{
			// Payload: 12345
			0x30, 0x39,
		},
		nbt:         ShortPayloadPtr(12345),
		snbt:        ptr.String(`12345s`),
		compactSnbt: ptr.String(`12345s`),
		prettySnbt:  ptr.String(`12345s`),
		json:        ptr.String(`12345`),
		compactJson: ptr.String(`12345`),
		prettyJson:  ptr.String(`12345`),
	},
	{
		name: `Valid Case: IntPayload`,
		raw: []byte{
			// Payload: 123456789
			0x07, 0x5B, 0xCD, 0x15,
		},
		nbt:         IntPayloadPtr(123456789),
		snbt:        ptr.String(`123456789`),
		compactSnbt: ptr.String(`123456789`),
		prettySnbt:  ptr.String(`123456789`),
		json:        ptr.String(`123456789`),
		compactJson: ptr.String(`123456789`),
		prettyJson:  ptr.String(`123456789`),
	},
	{
		name: `Valid Case: LongPayload`,
		raw: []byte{
			// Payload: 123456789123456789
			0x01, 0xB6, 0x9B, 0x4B, 0xAC, 0xD0, 0x5F, 0x15,
		},
		nbt:         LongPayloadPtr(123456789123456789),
		snbt:        ptr.String(`123456789123456789L`),
		compactSnbt: ptr.String(`123456789123456789L`),
		prettySnbt:  ptr.String(`123456789123456789L`),
		json:        ptr.String(`123456789123456789`),
		compactJson: ptr.String(`123456789123456789`),
		prettyJson:  ptr.String(`123456789123456789`),
	},
	{
		name: `Valid Case: FloatPayload`,
		raw: []byte{
			// Payload: 0.12345678
			0x3D, 0xFC, 0xD6, 0xE9,
		},
		nbt:         FloatPayloadPtr(0.12345678),
		snbt:        ptr.String(`0.12345678f`),
		compactSnbt: ptr.String(`0.12345678f`),
		prettySnbt:  ptr.String(`0.12345678f`),
		json:        ptr.String(`0.12345678`),
		compactJson: ptr.String(`0.12345678`),
		prettyJson:  ptr.String(`0.12345678`),
	},
	{
		name: `Valid Case: DoublePayload`,
		raw: []byte{
			// Payload: 0.123456789
			0x3F, 0xBF, 0x9A, 0xDD, 0x37, 0x39, 0x63, 0x5F,
		},
		nbt:         DoublePayloadPtr(0.123456789),
		snbt:        ptr.String(`0.123456789d`),
		compactSnbt: ptr.String(`0.123456789d`),
		prettySnbt:  ptr.String(`0.123456789d`),
		json:        ptr.String(`0.123456789`),
		compactJson: ptr.String(`0.123456789`),
		prettyJson:  ptr.String(`0.123456789`),
	},
	{
		name: `Valid Case: ByteArrayPayload`,
		raw: []byte{
			// Payload Length: 10
			0x00, 0x00, 0x00, 0x0A,
			// Payload: [B; 0b, 1b, 2b, 3b, 4b, 5b, 6b, 7b, 8b, 9b]
			0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09,
		},
		nbt:         &ByteArrayPayload{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		snbt:        ptr.String(`[B; 0b, 1b, 2b, 3b, 4b, 5b, 6b, 7b, 8b, 9b]`),
		compactSnbt: ptr.String(`[B;0b,1b,2b,3b,4b,5b,6b,7b,8b,9b]`),
		prettySnbt:  ptr.String(`[B; 0b, 1b, 2b, 3b, 4b, 5b, 6b, 7b, 8b, 9b]`),
		json:        ptr.String(`[0, 1, 2, 3, 4, 5, 6, 7, 8, 9]`),
		compactJson: ptr.String(`[0,1,2,3,4,5,6,7,8,9]`),
		prettyJson: ptr.String(`[
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
]`),
	},
	{
		name: `Valid Case: ByteArrayPayload - Empty`,
		raw: []byte{
			// Payload Length: 0
			0x00, 0x00, 0x00, 0x00,
			// Payload: [B; ]
		},
		nbt:         &ByteArrayPayload{},
		snbt:        ptr.String(`[B; ]`),
		compactSnbt: ptr.String(`[B;]`),
		prettySnbt:  ptr.String(`[B; ]`),
		json:        ptr.String(`[]`),
		compactJson: ptr.String(`[]`),
		prettyJson:  ptr.String(`[]`),
	},
	{
		name: `Valid Case: StringPayload "Test"`,
		raw: []byte{
			// Payload Length: 4
			0x00, 0x04,
			// Payload: "Test"
			0x54, 0x65, 0x73, 0x74,
		},
		nbt:         StringPayloadPtr(`Test`),
		snbt:        ptr.String(`"Test"`),
		compactSnbt: ptr.String(`"Test"`),
		prettySnbt:  ptr.String(`"Test"`),
		json:        ptr.String(`"Test"`),
		compactJson: ptr.String(`"Test"`),
		prettyJson:  ptr.String(`"Test"`),
	},
	{
		name: `Valid Case: StringPayload '"Test'`,
		raw: []byte{
			// Payload Length: 5
			0x00, 0x05,
			// Payload: '"Test'
			0x22, 0x54, 0x65, 0x73, 0x74,
		},
		nbt:         StringPayloadPtr(`"Test`),
		snbt:        ptr.String(`'"Test'`),
		compactSnbt: ptr.String(`'"Test'`),
		prettySnbt:  ptr.String(`'"Test'`),
		json:        ptr.String(`"\"Test"`),
		compactJson: ptr.String(`"\"Test"`),
		prettyJson:  ptr.String(`"\"Test"`),
	},
	{
		name: `Valid Case: StringPayload "'Test"`,
		raw: []byte{
			// Payload Length: 5
			0x00, 0x05,
			// Payload: "'Test"
			0x27, 0x54, 0x65, 0x73, 0x74,
		},
		nbt:         StringPayloadPtr(`'Test`),
		snbt:        ptr.String(`"'Test"`),
		compactSnbt: ptr.String(`"'Test"`),
		prettySnbt:  ptr.String(`"'Test"`),
		json:        ptr.String(`"'Test"`),
		compactJson: ptr.String(`"'Test"`),
		prettyJson:  ptr.String(`"'Test"`),
	},
	{
		name: `Valid Case: StringPayload '"Test'`,
		raw: []byte{
			// Payload Length: 6
			0x00, 0x06,
			// Payload: "\"'Test"
			0x22, 0x27, 0x54, 0x65, 0x73, 0x74,
		},
		nbt:         StringPayloadPtr(`"'Test`),
		snbt:        ptr.String(`"\"'Test"`),
		compactSnbt: ptr.String(`"\"'Test"`),
		prettySnbt:  ptr.String(`"\"'Test"`),
		json:        ptr.String(`"\"'Test"`),
		compactJson: ptr.String(`"\"'Test"`),
		prettyJson:  ptr.String(`"\"'Test"`),
	},
	{
		name: `Valid Case: StringPayload "minecraft:the_end"`,
		raw: []byte{
			// Payload Length: 17
			0x00, 0x11,
			// Payload: "minecraft:the_end"
			0x6D, 0x69, 0x6E, 0x65, 0x63, 0x72, 0x61, 0x66, 0x74, 0x3A,
			0x74, 0x68, 0x65, 0x5F, 0x65, 0x6E, 0x64,
		},
		nbt:         StringPayloadPtr(`minecraft:the_end`),
		snbt:        ptr.String(`"minecraft:the_end"`),
		compactSnbt: ptr.String(`"minecraft:the_end"`),
		prettySnbt:  ptr.String(`"minecraft:the_end"`),
		json:        ptr.String(`"minecraft:the_end"`),
		compactJson: ptr.String(`"minecraft:the_end"`),
		prettyJson:  ptr.String(`"minecraft:the_end"`),
	},
	{
		name: `Valid Case: StringPayload ""`,
		raw: []byte{
			// Payload Length: 0
			0x00, 0x00,
			// Payload: ""
		},
		nbt:         StringPayloadPtr(``),
		snbt:        ptr.String(`""`),
		compactSnbt: ptr.String(`""`),
		prettySnbt:  ptr.String(`""`),
		json:        ptr.String(`""`),
		compactJson: ptr.String(`""`),
		prettyJson:  ptr.String(`""`),
	},
	{
		name: `Valid Case: StringPayload "マインクラフト"`,
		raw: []byte{
			// Payload Length: 21
			0x00, 0x15,
			// Payload: "マインクラフト"
			0xE3, 0x83, 0x9E, 0xE3, 0x82, 0xA4, 0xE3, 0x83, 0xB3,
			0xE3, 0x82, 0xAF, 0xE3, 0x83, 0xA9, 0xE3, 0x83, 0x95,
			0xE3, 0x83, 0x88,
		},
		nbt:         StringPayloadPtr(`マインクラフト`),
		snbt:        ptr.String(`"マインクラフト"`),
		compactSnbt: ptr.String(`"マインクラフト"`),
		prettySnbt:  ptr.String(`"マインクラフト"`),
		json:        ptr.String(`"マインクラフト"`),
		compactJson: ptr.String(`"マインクラフト"`),
		prettyJson:  ptr.String(`"マインクラフト"`),
	},
	{
		name: `Valid Case: ListPayload - Short`,
		raw: []byte{
			// Payload Type: TagShort(=2)
			0x02,
			// Payload Length: 2
			0x00, 0x00, 0x00, 0x02,
			// Payload:
			//   - ShortPayload: 12345
			0x30, 0x39,
			//   - ShortPayload: 6789
			0x1A, 0x85,
		},
		nbt:         &ListPayload{ShortPayloadPtr(12345), ShortPayloadPtr(6789)},
		snbt:        ptr.String(`[12345s, 6789s]`),
		compactSnbt: ptr.String(`[12345s,6789s]`),
		prettySnbt: ptr.String(`[
    12345s,
    6789s
]`),
		json:        ptr.String(`[12345, 6789]`),
		compactJson: ptr.String(`[12345,6789]`),
		prettyJson: ptr.String(`[
    12345,
    6789
]`),
	},
	{
		name: `Valid Case: ListPayload - ByteArray`,
		raw: []byte{
			// Payload Type: TagByteArray(=7)
			0x07,
			// Payload Length: 2
			0x00, 0x00, 0x00, 0x02,
			// Payload:
			//   - ByteArrayPayload: [B; 0b, 1b]
			0x00, 0x00, 0x00, 0x02,
			0x00, 0x01,
			//   - ByteArrayPayload: [B; 2b, 3b]
			0x00, 0x00, 0x00, 0x02,
			0x02, 0x03,
		},
		nbt:         &ListPayload{&ByteArrayPayload{0, 1}, &ByteArrayPayload{2, 3}},
		snbt:        ptr.String(`[[B; 0b, 1b], [B; 2b, 3b]]`),
		compactSnbt: ptr.String(`[[B;0b,1b],[B;2b,3b]]`),
		prettySnbt: ptr.String(`[
    [B; 0b, 1b],
    [B; 2b, 3b]
]`),
		json:        ptr.String(`[[0, 1], [2, 3]]`),
		compactJson: ptr.String(`[[0,1],[2,3]]`),
		prettyJson: ptr.String(`[
    [
        0,
        1
    ],
    [
        2,
        3
    ]
]`),
	},
	{
		name: `Valid Case: ListPayload - String`,
		raw: []byte{
			// Payload Type: TagString(=8)
			0x08,
			// Payload Length: 2
			0x00, 0x00, 0x00, 0x02,
			// Payload:
			//   - StringPayload: "Hello"
			0x00, 0x05,
			0x48, 0x65, 0x6C, 0x6C, 0x6F,
			//   - StringPayload: "World"
			0x00, 0x05,
			0x57, 0x6F, 0x72, 0x6C, 0x64,
		},
		nbt:         &ListPayload{StringPayloadPtr(`Hello`), StringPayloadPtr(`World`)},
		snbt:        ptr.String(`["Hello", "World"]`),
		compactSnbt: ptr.String(`["Hello","World"]`),
		prettySnbt: ptr.String(`[
    "Hello",
    "World"
]`),
		json:        ptr.String(`["Hello", "World"]`),
		compactJson: ptr.String(`["Hello","World"]`),
		prettyJson: ptr.String(`[
    "Hello",
    "World"
]`),
	},
	{
		name: `Valid Case: ListPayload - List`,
		raw: []byte{
			// Payload Type: TagList(=9)
			0x09,
			// Payload Length: 2
			0x00, 0x00, 0x00, 0x02,
			// Payload:
			//   - ListPayload: <Byte>[123]
			0x01,
			0x00, 0x00, 0x00, 0x01,
			0x7B,
			//   - ListPayload: <String>["Test"]
			0x08,
			0x00, 0x00, 0x00, 0x01,
			0x00, 0x04,
			0x54, 0x65, 0x73, 0x74,
		},
		nbt: &ListPayload{
			&ListPayload{BytePayloadPtr(123)},
			&ListPayload{StringPayloadPtr(`Test`)},
		},
		snbt:        ptr.String(`[[123b], ["Test"]]`),
		compactSnbt: ptr.String(`[[123b],["Test"]]`),
		prettySnbt: ptr.String(`[
    [
        123b
    ],
    [
        "Test"
    ]
]`),
		json:        ptr.String(`[[123], ["Test"]]`),
		compactJson: ptr.String(`[[123],["Test"]]`),
		prettyJson: ptr.String(`[
    [
        123
    ],
    [
        "Test"
    ]
]`),
	},
	{
		name: `Valid Case: ListPayload - Compound`,
		raw: []byte{
			// Payload Type: TagCompound(=10)
			0x0A,
			// Payload Length: 2
			0x00, 0x00, 0x00, 0x02,
			// Payload:
			//   - CompoundPayload:
			//       - ByteTag("Byte"): 123
			0x01,
			0x00, 0x04,
			0x42, 0x79, 0x74, 0x65,
			0x7B,
			//       - EndTag
			0x00,
			//   - CompoundPayload:
			//       - StringTag("String"): "Hello"
			0x08,
			0x00, 0x06,
			0x53, 0x74, 0x72, 0x69, 0x6E, 0x67,
			0x00, 0x05,
			0x48, 0x65, 0x6C, 0x6C, 0x6F,
			//       - EndTag
			0x00,
		},
		nbt: &ListPayload{
			&CompoundPayload{&ByteTag{TagName(`Byte`), BytePayload(123)}, &EndTag{}},
			&CompoundPayload{&StringTag{TagName(`String`), StringPayload(`Hello`)}, &EndTag{}},
		},
		snbt:        ptr.String(`[{Byte: 123b}, {String: "Hello"}]`),
		compactSnbt: ptr.String(`[{Byte:123b},{String:"Hello"}]`),
		prettySnbt: ptr.String(`[
    {
        Byte: 123b
    },
    {
        String: "Hello"
    }
]`),
		json:        ptr.String(`[{"Byte": 123}, {"String": "Hello"}]`),
		compactJson: ptr.String(`[{"Byte":123},{"String":"Hello"}]`),
		prettyJson: ptr.String(`[
    {
        "Byte": 123
    },
    {
        "String": "Hello"
    }
]`),
	},
	{
		name: `Valid Case: ListPayload - Empty`,
		raw: []byte{
			// Payload Type: TagEnd(=0)
			0x00,
			// Payload Length: 0
			0x00, 0x00, 0x00, 0x00,
			// Payload: []
		},
		nbt:         &ListPayload{},
		snbt:        ptr.String(`[]`),
		compactSnbt: ptr.String(`[]`),
		prettySnbt:  ptr.String(`[]`),
		json:        ptr.String(`[]`),
		compactJson: ptr.String(`[]`),
		prettyJson:  ptr.String(`[]`),
	},
	{
		name: `Valid Case: CompoundPayload`,
		raw: []byte{
			// Payload:
			//   - ShortTag("Short"): 12345
			0x02,
			0x00, 0x05,
			0x53, 0x68, 0x6f, 0x72, 0x74,
			0x30, 0x39,
			//   - ByteArrayTag("ByteArray"): [B; 0b, 1b]
			0x07,
			0x00, 0x09,
			0x42, 0x79, 0x74, 0x65, 0x41, 0x72, 0x72, 0x61, 0x79,
			0x00, 0x00, 0x00, 0x02,
			0x00, 0x01,
			//   - StringTag("String"): "Hello"
			0x08,
			0x00, 0x06,
			0x53, 0x74, 0x72, 0x69, 0x6E, 0x67,
			0x00, 0x05,
			0x48, 0x65, 0x6C, 0x6C, 0x6F,
			//   - ListTag("List"): <Byte>[123]
			0x09,
			0x00, 0x04,
			0x4C, 0x69, 0x73, 0x74,
			0x01,
			0x00, 0x00, 0x00, 0x01,
			0x7B,
			//   - CompoundTag("Compound"):
			0x0A,
			0x00, 0x08,
			0x43, 0x6F, 0x6D, 0x70, 0x6F, 0x75, 0x6E, 0x64,
			//       - StringTag("String"): "World"
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
		nbt: &CompoundPayload{
			&ShortTag{TagName(`Short`), ShortPayload(12345)},
			&ByteArrayTag{TagName(`ByteArray`), ByteArrayPayload{0, 1}},
			&StringTag{TagName(`String`), StringPayload(`Hello`)},
			&ListTag{TagName(`List`), ListPayload{BytePayloadPtr(123)}},
			&CompoundTag{TagName(`Compound`), CompoundPayload{&StringTag{TagName(`String`), StringPayload(`World`)}, &EndTag{}}},
			&EndTag{},
		},
		snbt:        ptr.String(`{ByteArray: [B; 0b, 1b], Compound: {String: "World"}, List: [123b], Short: 12345s, String: "Hello"}`),
		compactSnbt: ptr.String(`{ByteArray:[B;0b,1b],Compound:{String:"World"},List:[123b],Short:12345s,String:"Hello"}`),
		prettySnbt: ptr.String(`{
    ByteArray: [B; 0b, 1b],
    Compound: {
        String: "World"
    },
    List: [
        123b
    ],
    Short: 12345s,
    String: "Hello"
}`),
		json:        ptr.String(`{"ByteArray": [0, 1], "Compound": {"String": "World"}, "List": [123], "Short": 12345, "String": "Hello"}`),
		compactJson: ptr.String(`{"ByteArray":[0,1],"Compound":{"String":"World"},"List":[123],"Short":12345,"String":"Hello"}`),
		prettyJson: ptr.String(`{
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
}`),
	},
	{
		name: `Valid Case: CompoundPayload - Empty`,
		raw: []byte{
			// Payload:
			//   - EndTag
			0x00,
		},
		nbt:         &CompoundPayload{&EndTag{}},
		snbt:        ptr.String(`{}`),
		compactSnbt: ptr.String(`{}`),
		prettySnbt:  ptr.String(`{}`),
		json:        ptr.String(`{}`),
		compactJson: ptr.String(`{}`),
		prettyJson:  ptr.String(`{}`),
	},
	{
		name: `Valid Case: IntArrayPayload`,
		raw: []byte{
			// Payload Length: 4
			0x00, 0x00, 0x00, 0x04,
			// Payload: [I; 0, 1, 2, 3]
			0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x01,
			0x00, 0x00, 0x00, 0x02,
			0x00, 0x00, 0x00, 0x03,
		},
		nbt:         &IntArrayPayload{0, 1, 2, 3},
		snbt:        ptr.String(`[I; 0, 1, 2, 3]`),
		compactSnbt: ptr.String(`[I;0,1,2,3]`),
		prettySnbt:  ptr.String(`[I; 0, 1, 2, 3]`),
		json:        ptr.String(`[0, 1, 2, 3]`),
		compactJson: ptr.String(`[0,1,2,3]`),
		prettyJson: ptr.String(`[
    0,
    1,
    2,
    3
]`),
	},
	{
		name: `Valid Case: IntArrayPayload - Empty`,
		raw: []byte{
			// Payload Length: 0
			0x00, 0x00, 0x00, 0x00,
			// Payload: [I; ]
		},
		nbt:         &IntArrayPayload{},
		snbt:        ptr.String(`[I; ]`),
		compactSnbt: ptr.String(`[I;]`),
		prettySnbt:  ptr.String(`[I; ]`),
		json:        ptr.String(`[]`),
		compactJson: ptr.String(`[]`),
		prettyJson:  ptr.String(`[]`),
	},
	{
		name: `Valid Case: LongArrayPayload`,
		raw: []byte{
			// Payload Length: 4
			0x00, 0x00, 0x00, 0x04,
			// Payload: [L; 0L, 1L, 2L, 3L]
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03,
		},
		nbt:         &LongArrayPayload{0, 1, 2, 3},
		snbt:        ptr.String(`[L; 0L, 1L, 2L, 3L]`),
		compactSnbt: ptr.String(`[L;0L,1L,2L,3L]`),
		prettySnbt:  ptr.String(`[L; 0L, 1L, 2L, 3L]`),
		json:        ptr.String(`[0, 1, 2, 3]`),
		compactJson: ptr.String(`[0,1,2,3]`),
		prettyJson: ptr.String(`[
    0,
    1,
    2,
    3
]`),
	},
	{
		name: `Valid Case: LongArrayPayload - Empty`,
		raw: []byte{
			// Payload Length: 0
			0x00, 0x00, 0x00, 0x00,
			// Payload: [L; ]
		},
		nbt:         &LongArrayPayload{},
		snbt:        ptr.String(`[L; ]`),
		compactSnbt: ptr.String(`[L;]`),
		prettySnbt:  ptr.String(`[L; ]`),
		json:        ptr.String(`[]`),
		compactJson: ptr.String(`[]`),
		prettyJson:  ptr.String(`[]`),
	},
}

var TagValidTestCases = []struct {
	name        string
	raw         []byte
	nbt         Tag
	snbt        *string
	compactSnbt *string
	prettySnbt  *string
	json        *string
	compactJson *string
	prettyJson  *string
}{
	{
		name:        `Valid Case: EndTag`,
		raw:         []byte{},
		nbt:         new(EndTag),
		snbt:        ptr.String(``),
		compactSnbt: ptr.String(``),
		prettySnbt:  ptr.String(``),
		json:        ptr.String(``),
		compactJson: ptr.String(``),
		prettyJson:  ptr.String(``),
	},
	{
		name: `Valid Case: ByteTag`,
		raw: []byte{
			// Name Length: 4
			0x00, 0x04,
			// Name: Byte
			0x42, 0x79, 0x74, 0x65,
			// Payload: 123
			0x7B,
		},
		nbt:         &ByteTag{TagName(`Byte`), BytePayload(123)},
		snbt:        ptr.String(`Byte: 123b`),
		compactSnbt: ptr.String(`Byte:123b`),
		prettySnbt:  ptr.String(`Byte: 123b`),
		json:        ptr.String(`"Byte": 123`),
		compactJson: ptr.String(`"Byte":123`),
		prettyJson:  ptr.String(`"Byte": 123`),
	},
	{
		name: `Valid Case: ShortTag`,
		raw: []byte{
			// Name Length: 5
			0x00, 0x05,
			// Name: Short
			0x53, 0x68, 0x6f, 0x72, 0x74,
			// Payload: 12345
			0x30, 0x39,
		},
		nbt:         &ShortTag{TagName(`Short`), ShortPayload(12345)},
		snbt:        ptr.String(`Short: 12345s`),
		compactSnbt: ptr.String(`Short:12345s`),
		prettySnbt:  ptr.String(`Short: 12345s`),
		json:        ptr.String(`"Short": 12345`),
		compactJson: ptr.String(`"Short":12345`),
		prettyJson:  ptr.String(`"Short": 12345`),
	},
	{
		name: `Valid Case: IntTag`,
		raw: []byte{
			// Name Length: 3
			0x00, 0x03,
			// Name: Int
			0x49, 0x6E, 0x74,
			// Payload: 123456789
			0x07, 0x5B, 0xCD, 0x15,
		},
		nbt:         &IntTag{TagName(`Int`), IntPayload(123456789)},
		snbt:        ptr.String(`Int: 123456789`),
		compactSnbt: ptr.String(`Int:123456789`),
		prettySnbt:  ptr.String(`Int: 123456789`),
		json:        ptr.String(`"Int": 123456789`),
		compactJson: ptr.String(`"Int":123456789`),
		prettyJson:  ptr.String(`"Int": 123456789`),
	},
	{
		name: `Valid Case: LongTag`,
		raw: []byte{
			// Name Length: 4
			0x00, 0x04,
			// Name: Long
			0x4C, 0x6F, 0x6E, 0x67,
			// Payload: 123456789123456789
			0x01, 0xB6, 0x9B, 0x4B, 0xAC, 0xD0, 0x5F, 0x15,
		},
		nbt:         &LongTag{TagName(`Long`), LongPayload(123456789123456789)},
		snbt:        ptr.String(`Long: 123456789123456789L`),
		compactSnbt: ptr.String(`Long:123456789123456789L`),
		prettySnbt:  ptr.String(`Long: 123456789123456789L`),
		json:        ptr.String(`"Long": 123456789123456789`),
		compactJson: ptr.String(`"Long":123456789123456789`),
		prettyJson:  ptr.String(`"Long": 123456789123456789`),
	},
	{
		name: `Valid Case: FloatTag`,
		raw: []byte{
			// Name Length: 5
			0x00, 0x05,
			// Name: Float
			0x46, 0x6C, 0x6F, 0x61, 0x74,
			// Payload: 0.12345678
			0x3D, 0xFC, 0xD6, 0xE9,
		},
		nbt:         &FloatTag{TagName(`Float`), FloatPayload(0.12345678)},
		snbt:        ptr.String(`Float: 0.12345678f`),
		compactSnbt: ptr.String(`Float:0.12345678f`),
		prettySnbt:  ptr.String(`Float: 0.12345678f`),
		json:        ptr.String(`"Float": 0.12345678`),
		compactJson: ptr.String(`"Float":0.12345678`),
		prettyJson:  ptr.String(`"Float": 0.12345678`),
	},
	{
		name: `Valid Case: DoubleTag`,
		raw: []byte{
			// Name Length: 6
			0x00, 0x06,
			// Name: Double
			0x44, 0x6F, 0x75, 0x62, 0x6C, 0x65,
			// Payload: 0.123456789
			0x3F, 0xBF, 0x9A, 0xDD, 0x37, 0x39, 0x63, 0x5F,
		},
		nbt:         &DoubleTag{TagName(`Double`), DoublePayload(0.123456789)},
		snbt:        ptr.String(`Double: 0.123456789d`),
		compactSnbt: ptr.String(`Double:0.123456789d`),
		prettySnbt:  ptr.String(`Double: 0.123456789d`),
		json:        ptr.String(`"Double": 0.123456789`),
		compactJson: ptr.String(`"Double":0.123456789`),
		prettyJson:  ptr.String(`"Double": 0.123456789`),
	},
	{
		name: `Valid Case: ByteArrayTag`,
		raw: []byte{
			// Name Length: 9
			0x00, 0x09,
			// Name: ByteArray
			0x42, 0x79, 0x74, 0x65, 0x41, 0x72, 0x72, 0x61, 0x79,
			// Payload Length: 10
			0x00, 0x00, 0x00, 0x0A,
			// Payload: [B; 0b, 1b, 2b, 3b, 4b, 5b, 6b, 7b, 8b, 9b]
			0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09,
		},
		nbt:         &ByteArrayTag{TagName(`ByteArray`), ByteArrayPayload{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		snbt:        ptr.String(`ByteArray: [B; 0b, 1b, 2b, 3b, 4b, 5b, 6b, 7b, 8b, 9b]`),
		compactSnbt: ptr.String(`ByteArray:[B;0b,1b,2b,3b,4b,5b,6b,7b,8b,9b]`),
		prettySnbt:  ptr.String(`ByteArray: [B; 0b, 1b, 2b, 3b, 4b, 5b, 6b, 7b, 8b, 9b]`),
		json:        ptr.String(`"ByteArray": [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]`),
		compactJson: ptr.String(`"ByteArray":[0,1,2,3,4,5,6,7,8,9]`),
		prettyJson: ptr.String(`"ByteArray": [
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
]`),
	},
	{
		name: `Valid Case: StringTag`,
		raw: []byte{
			// Name Length: 6
			0x00, 0x06,
			// Name: String
			0x53, 0x74, 0x72, 0x69, 0x6E, 0x67,
			// Payload Length: 11
			0x00, 0x0B,
			// Payload: "Hello World"
			0x48, 0x65, 0x6C, 0x6C, 0x6F, 0x20, 0x57, 0x6F, 0x72, 0x6C, 0x64,
		},
		nbt:         &StringTag{TagName(`String`), StringPayload(`Hello World`)},
		snbt:        ptr.String(`String: "Hello World"`),
		compactSnbt: ptr.String(`String:"Hello World"`),
		prettySnbt:  ptr.String(`String: "Hello World"`),
		json:        ptr.String(`"String": "Hello World"`),
		compactJson: ptr.String(`"String":"Hello World"`),
		prettyJson:  ptr.String(`"String": "Hello World"`),
	},
	{
		name: `Valid Case: ListTag`,
		raw: []byte{
			// Name Length: 4
			0x00, 0x04,
			// Name: List
			0x4C, 0x69, 0x73, 0x74,
			// Payload Type: TagShort(=2)
			0x02,
			// Payload Length: 2
			0x00, 0x00, 0x00, 0x02,
			// Payload:
			//   - ShortPayload: 12345
			0x30, 0x39,
			//   - ShortPayload: 6789
			0x1A, 0x85,
		},
		nbt:         &ListTag{TagName(`List`), ListPayload{ShortPayloadPtr(12345), ShortPayloadPtr(6789)}},
		snbt:        ptr.String(`List: [12345s, 6789s]`),
		compactSnbt: ptr.String(`List:[12345s,6789s]`),
		prettySnbt: ptr.String(`List: [
    12345s,
    6789s
]`),
		json:        ptr.String(`"List": [12345, 6789]`),
		compactJson: ptr.String(`"List":[12345,6789]`),
		prettyJson: ptr.String(`"List": [
    12345,
    6789
]`),
	},
	{
		name: `Valid Case: CompoundTag`,
		raw: []byte{
			// Name Length: 8
			0x00, 0x08,
			// Name: Compound
			0x43, 0x6F, 0x6D, 0x70, 0x6F, 0x75, 0x6E, 0x64,
			// Payload:
			//   - ShortTag("Short"): 12345
			0x02,
			0x00, 0x05,
			0x53, 0x68, 0x6f, 0x72, 0x74,
			0x30, 0x39,
			//   - ByteArrayTag("ByteArray"): [B; 0b, 1b]
			0x07,
			0x00, 0x09,
			0x42, 0x79, 0x74, 0x65, 0x41, 0x72, 0x72, 0x61, 0x79,
			0x00, 0x00, 0x00, 0x02,
			0x00, 0x01,
			//   - StringTag("String"): "Hello"
			0x08,
			0x00, 0x06,
			0x53, 0x74, 0x72, 0x69, 0x6E, 0x67,
			0x00, 0x05,
			0x48, 0x65, 0x6C, 0x6C, 0x6F,
			//   - ListTag("List"): <Byte>[123]
			0x09,
			0x00, 0x04,
			0x4C, 0x69, 0x73, 0x74,
			0x01,
			0x00, 0x00, 0x00, 0x01,
			0x7B,
			//   - CompoundTag("Compound"):
			0x0A,
			0x00, 0x08,
			0x43, 0x6F, 0x6D, 0x70, 0x6F, 0x75, 0x6E, 0x64,
			//       - StringTag("String"): "World"
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
		nbt: &CompoundTag{
			TagName(`Compound`),
			CompoundPayload{
				&ShortTag{TagName(`Short`), ShortPayload(12345)},
				&ByteArrayTag{TagName(`ByteArray`), ByteArrayPayload{0, 1}},
				&StringTag{TagName(`String`), StringPayload(`Hello`)},
				&ListTag{TagName(`List`), ListPayload{BytePayloadPtr(123)}},
				&CompoundTag{TagName(`Compound`), CompoundPayload{&StringTag{TagName(`String`), StringPayload(`World`)}, &EndTag{}}},
				&EndTag{},
			},
		},
		snbt:        ptr.String(`Compound: {ByteArray: [B; 0b, 1b], Compound: {String: "World"}, List: [123b], Short: 12345s, String: "Hello"}`),
		compactSnbt: ptr.String(`Compound:{ByteArray:[B;0b,1b],Compound:{String:"World"},List:[123b],Short:12345s,String:"Hello"}`),
		prettySnbt: ptr.String(`Compound: {
    ByteArray: [B; 0b, 1b],
    Compound: {
        String: "World"
    },
    List: [
        123b
    ],
    Short: 12345s,
    String: "Hello"
}`),
		json:        ptr.String(`"Compound": {"ByteArray": [0, 1], "Compound": {"String": "World"}, "List": [123], "Short": 12345, "String": "Hello"}`),
		compactJson: ptr.String(`"Compound":{"ByteArray":[0,1],"Compound":{"String":"World"},"List":[123],"Short":12345,"String":"Hello"}`),
		prettyJson: ptr.String(`"Compound": {
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
}`),
	},
	{
		name: `Valid Case: IntArrayTag`,
		raw: []byte{
			// Name Length: 8
			0x00, 0x08,
			// Name: IntArray
			0x49, 0x6E, 0x74, 0x41, 0x72, 0x72, 0x61, 0x79,
			// Payload Length: 4
			0x00, 0x00, 0x00, 0x04,
			// Payload: [I; 0, 1, 2, 3]
			0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x01,
			0x00, 0x00, 0x00, 0x02,
			0x00, 0x00, 0x00, 0x03,
		},
		nbt:         &IntArrayTag{TagName(`IntArray`), IntArrayPayload{0, 1, 2, 3}},
		snbt:        ptr.String(`IntArray: [I; 0, 1, 2, 3]`),
		compactSnbt: ptr.String(`IntArray:[I;0,1,2,3]`),
		prettySnbt:  ptr.String(`IntArray: [I; 0, 1, 2, 3]`),
		json:        ptr.String(`"IntArray": [0, 1, 2, 3]`),
		compactJson: ptr.String(`"IntArray":[0,1,2,3]`),
		prettyJson: ptr.String(`"IntArray": [
    0,
    1,
    2,
    3
]`),
	},
	{
		name: `Valid Case: LongArrayTag`,
		raw: []byte{
			// Name Length: 9
			0x00, 0x09,
			// Name: LongArray
			0x4C, 0x6F, 0x6E, 0x67, 0x41, 0x72, 0x72, 0x61, 0x79,
			// Payload Length: 4
			0x00, 0x00, 0x00, 0x04,
			// Payload: [L; 0L, 1L, 2L, 3L]
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03,
		},
		nbt:         &LongArrayTag{TagName(`LongArray`), LongArrayPayload{0, 1, 2, 3}},
		snbt:        ptr.String(`LongArray: [L; 0L, 1L, 2L, 3L]`),
		compactSnbt: ptr.String(`LongArray:[L;0L,1L,2L,3L]`),
		prettySnbt:  ptr.String(`LongArray: [L; 0L, 1L, 2L, 3L]`),
		json:        ptr.String(`"LongArray": [0, 1, 2, 3]`),
		compactJson: ptr.String(`"LongArray":[0,1,2,3]`),
		prettyJson: ptr.String(`"LongArray": [
    0,
    1,
    2,
    3
]`),
	},
}
