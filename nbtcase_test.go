package gonbt

import (
	"go.uber.org/thriftrw/ptr"
)

var NBTValidTestCases = []struct {
	name        string
	raw         []byte
	nbt         Tag
	snbt        *string
	snbtCompact *string
	snbtPretty  *string
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
		snbtCompact: ptr.String(`{"Hello World":{Name:"Steve"}}`),
		snbtPretty: ptr.String(`{
    "Hello World": {
        Name: "Steve"
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
				&ListTag{TagName(`List`), ListPayload{PayloadType: TagByte, Payloads: []Payload{BytePayloadPtr(123)}}},
				&CompoundTag{TagName(`Compound`), CompoundPayload{&StringTag{TagName(`String`), StringPayload(`World`)}, &EndTag{}}},
				&EndTag{},
			},
		},
		snbt:        ptr.String(`{Compound: {ByteArray: [B; 0b, 1b], Compound: {String: "World"}, List: [123b], Short: 12345s, String: "Hello"}}`),
		snbtCompact: ptr.String(`{Compound:{ByteArray:[B;0b,1b],Compound:{String:"World"},List:[123b],Short:12345s,String:"Hello"}}`),
		snbtPretty: ptr.String(`{
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
	},
}

var TagNameValidTestCases = []struct {
	name string
	raw  []byte
	nbt  *TagName
	snbt *string
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
	},
}

var PayloadValidTestCases = []struct {
	name        string
	raw         []byte
	nbt         Payload
	snbt        *string
	snbtCompact *string
	snbtPretty  *string
}{
	{
		name: `Valid Case: BytePayload`,
		raw: []byte{
			// Payload: 123
			0x7B,
		},
		nbt:         BytePayloadPtr(123),
		snbt:        ptr.String(`123b`),
		snbtCompact: ptr.String(`123b`),
		snbtPretty:  ptr.String(`123b`),
	},
	{
		name: `Valid Case: ShortPayload`,
		raw: []byte{
			// Payload: 12345
			0x30, 0x39,
		},
		nbt:         ShortPayloadPtr(12345),
		snbt:        ptr.String(`12345s`),
		snbtCompact: ptr.String(`12345s`),
		snbtPretty:  ptr.String(`12345s`),
	},
	{
		name: `Valid Case: IntPayload`,
		raw: []byte{
			// Payload: 123456789
			0x07, 0x5B, 0xCD, 0x15,
		},
		nbt:         IntPayloadPtr(123456789),
		snbt:        ptr.String(`123456789`),
		snbtCompact: ptr.String(`123456789`),
		snbtPretty:  ptr.String(`123456789`),
	},
	{
		name: `Valid Case: LongPayload`,
		raw: []byte{
			// Payload: 123456789123456789
			0x01, 0xB6, 0x9B, 0x4B, 0xAC, 0xD0, 0x5F, 0x15,
		},
		nbt:         LongPayloadPtr(123456789123456789),
		snbt:        ptr.String(`123456789123456789L`),
		snbtCompact: ptr.String(`123456789123456789L`),
		snbtPretty:  ptr.String(`123456789123456789L`),
	},
	{
		name: `Valid Case: FloatPayload`,
		raw: []byte{
			// Payload: 0.12345678
			0x3D, 0xFC, 0xD6, 0xE9,
		},
		nbt:         FloatPayloadPtr(0.12345678),
		snbt:        ptr.String(`0.12345678f`),
		snbtCompact: ptr.String(`0.12345678f`),
		snbtPretty:  ptr.String(`0.12345678f`),
	},
	{
		name: `Valid Case: DoublePayload`,
		raw: []byte{
			// Payload: 0.123456789
			0x3F, 0xBF, 0x9A, 0xDD, 0x37, 0x39, 0x63, 0x5F,
		},
		nbt:         DoublePayloadPtr(0.123456789),
		snbt:        ptr.String(`0.123456789d`),
		snbtCompact: ptr.String(`0.123456789d`),
		snbtPretty:  ptr.String(`0.123456789d`),
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
		snbtCompact: ptr.String(`[B;0b,1b,2b,3b,4b,5b,6b,7b,8b,9b]`),
		snbtPretty:  ptr.String(`[B; 0b, 1b, 2b, 3b, 4b, 5b, 6b, 7b, 8b, 9b]`),
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
		snbtCompact: ptr.String(`"Test"`),
		snbtPretty:  ptr.String(`"Test"`),
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
		snbtCompact: ptr.String(`'"Test'`),
		snbtPretty:  ptr.String(`'"Test'`),
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
		snbtCompact: ptr.String(`"'Test"`),
		snbtPretty:  ptr.String(`"'Test"`),
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
		snbtCompact: ptr.String(`"\"'Test"`),
		snbtPretty:  ptr.String(`"\"'Test"`),
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
		snbtCompact: ptr.String(`"minecraft:the_end"`),
		snbtPretty:  ptr.String(`"minecraft:the_end"`),
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
		snbtCompact: ptr.String(`""`),
		snbtPretty:  ptr.String(`""`),
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
		snbtCompact: ptr.String(`"マインクラフト"`),
		snbtPretty:  ptr.String(`"マインクラフト"`),
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
		nbt:         &ListPayload{PayloadType: TagShort, Payloads: []Payload{ShortPayloadPtr(12345), ShortPayloadPtr(6789)}},
		snbt:        ptr.String(`[12345s, 6789s]`),
		snbtCompact: ptr.String(`[12345s,6789s]`),
		snbtPretty: ptr.String(`[
    12345s,
    6789s
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
		nbt:         &ListPayload{PayloadType: TagByteArray, Payloads: []Payload{&ByteArrayPayload{0, 1}, &ByteArrayPayload{2, 3}}},
		snbt:        ptr.String(`[[B; 0b, 1b], [B; 2b, 3b]]`),
		snbtCompact: ptr.String(`[[B;0b,1b],[B;2b,3b]]`),
		snbtPretty: ptr.String(`[
    [B; 0b, 1b],
    [B; 2b, 3b]
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
		nbt:         &ListPayload{PayloadType: TagString, Payloads: []Payload{StringPayloadPtr(`Hello`), StringPayloadPtr(`World`)}},
		snbt:        ptr.String(`["Hello", "World"]`),
		snbtCompact: ptr.String(`["Hello","World"]`),
		snbtPretty: ptr.String(`[
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
			PayloadType: TagList,
			Payloads: []Payload{
				&ListPayload{PayloadType: TagByte, Payloads: []Payload{BytePayloadPtr(123)}},
				&ListPayload{PayloadType: TagString, Payloads: []Payload{StringPayloadPtr(`Test`)}},
			},
		},
		snbt:        ptr.String(`[[123b], ["Test"]]`),
		snbtCompact: ptr.String(`[[123b],["Test"]]`),
		snbtPretty: ptr.String(`[
    [
        123b
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
			PayloadType: TagCompound,
			Payloads: []Payload{
				&CompoundPayload{&ByteTag{TagName(`Byte`), BytePayload(123)}, &EndTag{}},
				&CompoundPayload{&StringTag{TagName(`String`), StringPayload(`Hello`)}, &EndTag{}},
			},
		},
		snbt:        ptr.String(`[{Byte: 123b}, {String: "Hello"}]`),
		snbtCompact: ptr.String(`[{Byte:123b},{String:"Hello"}]`),
		snbtPretty: ptr.String(`[
    {
        Byte: 123b
    },
    {
        String: "Hello"
    }
]`),
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
			&ListTag{TagName(`List`), ListPayload{PayloadType: TagByte, Payloads: []Payload{BytePayloadPtr(123)}}},
			&CompoundTag{TagName(`Compound`), CompoundPayload{&StringTag{TagName(`String`), StringPayload(`World`)}, &EndTag{}}},
			&EndTag{},
		},
		snbt:        ptr.String(`{ByteArray: [B; 0b, 1b], Compound: {String: "World"}, List: [123b], Short: 12345s, String: "Hello"}`),
		snbtCompact: ptr.String(`{ByteArray:[B;0b,1b],Compound:{String:"World"},List:[123b],Short:12345s,String:"Hello"}`),
		snbtPretty: ptr.String(`{
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
		snbtCompact: ptr.String(`{}`),
		snbtPretty:  ptr.String(`{}`),
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
		snbtCompact: ptr.String(`[I;0,1,2,3]`),
		snbtPretty:  ptr.String(`[I; 0, 1, 2, 3]`),
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
		snbtCompact: ptr.String(`[L;0L,1L,2L,3L]`),
		snbtPretty:  ptr.String(`[L; 0L, 1L, 2L, 3L]`),
	},
}

var TagValidTestCases = []struct {
	name        string
	raw         []byte
	nbt         Tag
	snbt        *string
	snbtCompact *string
	snbtPretty  *string
}{
	{
		name:        `Valid Case: EndTag`,
		raw:         []byte{},
		nbt:         new(EndTag),
		snbt:        ptr.String(``),
		snbtCompact: ptr.String(``),
		snbtPretty:  ptr.String(``),
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
		snbtCompact: ptr.String(`Byte:123b`),
		snbtPretty:  ptr.String(`Byte: 123b`),
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
		snbtCompact: ptr.String(`Short:12345s`),
		snbtPretty:  ptr.String(`Short: 12345s`),
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
		snbtCompact: ptr.String(`Int:123456789`),
		snbtPretty:  ptr.String(`Int: 123456789`),
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
		snbtCompact: ptr.String(`Long:123456789123456789L`),
		snbtPretty:  ptr.String(`Long: 123456789123456789L`),
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
		snbtCompact: ptr.String(`Float:0.12345678f`),
		snbtPretty:  ptr.String(`Float: 0.12345678f`),
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
		snbtCompact: ptr.String(`Double:0.123456789d`),
		snbtPretty:  ptr.String(`Double: 0.123456789d`),
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
		snbtCompact: ptr.String(`ByteArray:[B;0b,1b,2b,3b,4b,5b,6b,7b,8b,9b]`),
		snbtPretty:  ptr.String(`ByteArray: [B; 0b, 1b, 2b, 3b, 4b, 5b, 6b, 7b, 8b, 9b]`),
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
		snbtCompact: ptr.String(`String:"Hello World"`),
		snbtPretty:  ptr.String(`String: "Hello World"`),
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
		nbt:         &ListTag{TagName(`List`), ListPayload{PayloadType: TagShort, Payloads: []Payload{ShortPayloadPtr(12345), ShortPayloadPtr(6789)}}},
		snbt:        ptr.String(`List: [12345s, 6789s]`),
		snbtCompact: ptr.String(`List:[12345s,6789s]`),
		snbtPretty: ptr.String(`List: [
    12345s,
    6789s
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
				&ListTag{TagName(`List`), ListPayload{PayloadType: TagByte, Payloads: []Payload{BytePayloadPtr(123)}}},
				&CompoundTag{TagName(`Compound`), CompoundPayload{&StringTag{TagName(`String`), StringPayload(`World`)}, &EndTag{}}},
				&EndTag{},
			},
		},
		snbt:        ptr.String(`Compound: {ByteArray: [B; 0b, 1b], Compound: {String: "World"}, List: [123b], Short: 12345s, String: "Hello"}`),
		snbtCompact: ptr.String(`Compound:{ByteArray:[B;0b,1b],Compound:{String:"World"},List:[123b],Short:12345s,String:"Hello"}`),
		snbtPretty: ptr.String(`Compound: {
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
		snbtCompact: ptr.String(`IntArray:[I;0,1,2,3]`),
		snbtPretty:  ptr.String(`IntArray: [I; 0, 1, 2, 3]`),
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
		snbtCompact: ptr.String(`LongArray:[L;0L,1L,2L,3L]`),
		snbtPretty:  ptr.String(`LongArray: [L; 0L, 1L, 2L, 3L]`),
	},
}
