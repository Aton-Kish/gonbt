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

package nbt_test

import (
	"bytes"
	"fmt"
	"log"

	nbt "github.com/Aton-Kish/gonbt"
)

func ExampleEncode() {
	dat := nbt.NewCompoundTag(nbt.NewTagName("Data"), nbt.NewCompoundPayload(
		nbt.NewIntTag(nbt.NewTagName("GameType"), nbt.NewIntPayload(0)),
		nbt.NewStringTag(nbt.NewTagName("LevelName"), nbt.NewStringPayload("Go NBT")),
		nbt.NewCompoundTag(nbt.NewTagName("Version"), nbt.NewCompoundPayload(
			nbt.NewIntTag(nbt.NewTagName("Id"), nbt.NewIntPayload(3120)),
			nbt.NewStringTag(nbt.NewTagName("Name"), nbt.NewStringPayload("1.19.2")),
			nbt.NewStringTag(nbt.NewTagName("Series"), nbt.NewStringPayload("main")),
			nbt.NewByteTag(nbt.NewTagName("Snapshot"), nbt.NewBytePayload(0)),
			nbt.NewEndTag(),
		)),
		nbt.NewEndTag(),
	))

	buf := new(bytes.Buffer)
	if err := nbt.Encode(buf, dat); err != nil {
		log.Fatal(err)
	}

	// f, err := os.Create("level.dat")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()

	// zw := gzip.NewWriter(f)
	// defer zw.Close()

	// if _, err := zw.Write(buf.Bytes()); err != nil {
	// 	log.Fatal(err)
	// }
}

func ExampleDecode() {
	// f, err := os.Open("level.dat")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()

	// zr, err := gzip.NewReader(f)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer zr.Close()

	// fake NBT
	zr := bytes.NewBuffer([]byte{
		// CompoundTag(Data):
		0x0A,
		0x00, 0x04,
		0x44, 0x61, 0x74, 0x61,
		//   - IntTag(GameType): 0
		0x03,
		0x00, 0x08,
		0x47, 0x61, 0x6D, 0x65, 0x54, 0x79, 0x70, 0x65,
		0x00, 0x00, 0x00, 0x00,
		//   - StringTag(LevelName): "Go NBT"
		0x08,
		0x00, 0x09,
		0x4C, 0x65, 0x76, 0x65, 0x6C, 0x4E, 0x61, 0x6D, 0x65,
		0x00, 0x06,
		0x47, 0x6F, 0x20, 0x4E, 0x42, 0x54,
		//   - CompoundTag(Version):
		0x0A,
		0x00, 0x07,
		0x56, 0x65, 0x72, 0x73, 0x69, 0x6F, 0x6E,
		//     - IntTag(Id): 3120
		0x03,
		0x00, 0x02,
		0x49, 0x64,
		0x00, 0x00, 0x0C, 0x30,
		//     - StringTag(Name): "1.19.2"
		0x08,
		0x00, 0x04,
		0x4E, 0x61, 0x6D, 0x65,
		0x00, 0x06,
		0x31, 0x2E, 0x31, 0x39, 0x2E, 0x32,
		//     - StringTag(Series): "main"
		0x08,
		0x00, 0x06,
		0x53, 0x65, 0x72, 0x69, 0x65, 0x73,
		0x00, 0x04,
		0x6D, 0x61, 0x69, 0x6E,
		//     - ByteTag(Snapshot): 0b
		0x01,
		0x00, 0x08,
		0x53, 0x6E, 0x61, 0x70, 0x73, 0x68, 0x6F, 0x74,
		0x00,
		//     - EndTag
		0x00,
		//   - EndTag
		0x00,
	})

	dat, err := nbt.Decode(zr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(nbt.PrettyStringify(dat, "  "))
	// Output:
	// {
	//   Data: {
	//     GameType: 0,
	//     LevelName: "Go NBT",
	//     Version: {
	//       Id: 3120,
	//       Name: "1.19.2",
	//       Series: "main",
	//       Snapshot: 0b
	//     }
	//   }
	// }
}

func ExampleStringify() {
	dat := nbt.NewCompoundTag(nbt.NewTagName("Data"), nbt.NewCompoundPayload(
		nbt.NewIntTag(nbt.NewTagName("GameType"), nbt.NewIntPayload(0)),
		nbt.NewStringTag(nbt.NewTagName("LevelName"), nbt.NewStringPayload("Go NBT")),
		nbt.NewCompoundTag(nbt.NewTagName("Version"), nbt.NewCompoundPayload(
			nbt.NewIntTag(nbt.NewTagName("Id"), nbt.NewIntPayload(3120)),
			nbt.NewStringTag(nbt.NewTagName("Name"), nbt.NewStringPayload("1.19.2")),
			nbt.NewStringTag(nbt.NewTagName("Series"), nbt.NewStringPayload("main")),
			nbt.NewByteTag(nbt.NewTagName("Snapshot"), nbt.NewBytePayload(0)),
			nbt.NewEndTag(),
		)),
		nbt.NewEndTag(),
	))

	fmt.Println(nbt.Stringify(dat))
	// Output:
	// {Data: {GameType: 0, LevelName: "Go NBT", Version: {Id: 3120, Name: "1.19.2", Series: "main", Snapshot: 0b}}}
}

func ExampleCompactStringify() {
	dat := nbt.NewCompoundTag(nbt.NewTagName("Data"), nbt.NewCompoundPayload(
		nbt.NewIntTag(nbt.NewTagName("GameType"), nbt.NewIntPayload(0)),
		nbt.NewStringTag(nbt.NewTagName("LevelName"), nbt.NewStringPayload("Go NBT")),
		nbt.NewCompoundTag(nbt.NewTagName("Version"), nbt.NewCompoundPayload(
			nbt.NewIntTag(nbt.NewTagName("Id"), nbt.NewIntPayload(3120)),
			nbt.NewStringTag(nbt.NewTagName("Name"), nbt.NewStringPayload("1.19.2")),
			nbt.NewStringTag(nbt.NewTagName("Series"), nbt.NewStringPayload("main")),
			nbt.NewByteTag(nbt.NewTagName("Snapshot"), nbt.NewBytePayload(0)),
			nbt.NewEndTag(),
		)),
		nbt.NewEndTag(),
	))

	fmt.Println(nbt.CompactStringify(dat))
	// Output:
	// {Data:{GameType:0,LevelName:"Go NBT",Version:{Id:3120,Name:"1.19.2",Series:"main",Snapshot:0b}}}
}

func ExamplePrettyStringify_space2() {
	dat := nbt.NewCompoundTag(nbt.NewTagName("Data"), nbt.NewCompoundPayload(
		nbt.NewIntTag(nbt.NewTagName("GameType"), nbt.NewIntPayload(0)),
		nbt.NewStringTag(nbt.NewTagName("LevelName"), nbt.NewStringPayload("Go NBT")),
		nbt.NewCompoundTag(nbt.NewTagName("Version"), nbt.NewCompoundPayload(
			nbt.NewIntTag(nbt.NewTagName("Id"), nbt.NewIntPayload(3120)),
			nbt.NewStringTag(nbt.NewTagName("Name"), nbt.NewStringPayload("1.19.2")),
			nbt.NewStringTag(nbt.NewTagName("Series"), nbt.NewStringPayload("main")),
			nbt.NewByteTag(nbt.NewTagName("Snapshot"), nbt.NewBytePayload(0)),
			nbt.NewEndTag(),
		)),
		nbt.NewEndTag(),
	))

	fmt.Println(nbt.PrettyStringify(dat, "  "))
	// Output:
	// {
	//   Data: {
	//     GameType: 0,
	//     LevelName: "Go NBT",
	//     Version: {
	//       Id: 3120,
	//       Name: "1.19.2",
	//       Series: "main",
	//       Snapshot: 0b
	//     }
	//   }
	// }
}

func ExamplePrettyStringify_space4() {
	dat := nbt.NewCompoundTag(nbt.NewTagName("Data"), nbt.NewCompoundPayload(
		nbt.NewIntTag(nbt.NewTagName("GameType"), nbt.NewIntPayload(0)),
		nbt.NewStringTag(nbt.NewTagName("LevelName"), nbt.NewStringPayload("Go NBT")),
		nbt.NewCompoundTag(nbt.NewTagName("Version"), nbt.NewCompoundPayload(
			nbt.NewIntTag(nbt.NewTagName("Id"), nbt.NewIntPayload(3120)),
			nbt.NewStringTag(nbt.NewTagName("Name"), nbt.NewStringPayload("1.19.2")),
			nbt.NewStringTag(nbt.NewTagName("Series"), nbt.NewStringPayload("main")),
			nbt.NewByteTag(nbt.NewTagName("Snapshot"), nbt.NewBytePayload(0)),
			nbt.NewEndTag(),
		)),
		nbt.NewEndTag(),
	))

	fmt.Println(nbt.PrettyStringify(dat, "    "))
	// Output:
	// {
	//     Data: {
	//         GameType: 0,
	//         LevelName: "Go NBT",
	//         Version: {
	//             Id: 3120,
	//             Name: "1.19.2",
	//             Series: "main",
	//             Snapshot: 0b
	//         }
	//     }
	// }
}

func ExamplePrettyStringify_tab() {
	dat := nbt.NewCompoundTag(nbt.NewTagName("Data"), nbt.NewCompoundPayload(
		nbt.NewIntTag(nbt.NewTagName("GameType"), nbt.NewIntPayload(0)),
		nbt.NewStringTag(nbt.NewTagName("LevelName"), nbt.NewStringPayload("Go NBT")),
		nbt.NewCompoundTag(nbt.NewTagName("Version"), nbt.NewCompoundPayload(
			nbt.NewIntTag(nbt.NewTagName("Id"), nbt.NewIntPayload(3120)),
			nbt.NewStringTag(nbt.NewTagName("Name"), nbt.NewStringPayload("1.19.2")),
			nbt.NewStringTag(nbt.NewTagName("Series"), nbt.NewStringPayload("main")),
			nbt.NewByteTag(nbt.NewTagName("Snapshot"), nbt.NewBytePayload(0)),
			nbt.NewEndTag(),
		)),
		nbt.NewEndTag(),
	))

	fmt.Println(nbt.PrettyStringify(dat, "\t"))
	// Output:
	// {
	// 	Data: {
	// 		GameType: 0,
	// 		LevelName: "Go NBT",
	// 		Version: {
	// 			Id: 3120,
	// 			Name: "1.19.2",
	// 			Series: "main",
	// 			Snapshot: 0b
	// 		}
	// 	}
	// }
}

func ExampleParse() {
	str := `{Data: {GameType: 0, LevelName: "Go NBT", Version: {Id: 3120, Name: "1.19.2", Series: "main", Snapshot: 0b}}}`

	dat, err := nbt.Parse(str)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(nbt.Stringify(dat))
	// Output:
	// {Data: {GameType: 0, LevelName: "Go NBT", Version: {Id: 3120, Name: "1.19.2", Series: "main", Snapshot: 0b}}}
}

func ExampleJson() {
	dat := nbt.NewCompoundTag(nbt.NewTagName("Data"), nbt.NewCompoundPayload(
		nbt.NewIntTag(nbt.NewTagName("GameType"), nbt.NewIntPayload(0)),
		nbt.NewStringTag(nbt.NewTagName("LevelName"), nbt.NewStringPayload("Go NBT")),
		nbt.NewCompoundTag(nbt.NewTagName("Version"), nbt.NewCompoundPayload(
			nbt.NewIntTag(nbt.NewTagName("Id"), nbt.NewIntPayload(3120)),
			nbt.NewStringTag(nbt.NewTagName("Name"), nbt.NewStringPayload("1.19.2")),
			nbt.NewStringTag(nbt.NewTagName("Series"), nbt.NewStringPayload("main")),
			nbt.NewByteTag(nbt.NewTagName("Snapshot"), nbt.NewBytePayload(0)),
			nbt.NewEndTag(),
		)),
		nbt.NewEndTag(),
	))

	fmt.Println(nbt.Json(dat))
	// Output:
	// {"Data": {"GameType": 0, "LevelName": "Go NBT", "Version": {"Id": 3120, "Name": "1.19.2", "Series": "main", "Snapshot": 0}}}
}

func ExampleCompactJson() {
	dat := nbt.NewCompoundTag(nbt.NewTagName("Data"), nbt.NewCompoundPayload(
		nbt.NewIntTag(nbt.NewTagName("GameType"), nbt.NewIntPayload(0)),
		nbt.NewStringTag(nbt.NewTagName("LevelName"), nbt.NewStringPayload("Go NBT")),
		nbt.NewCompoundTag(nbt.NewTagName("Version"), nbt.NewCompoundPayload(
			nbt.NewIntTag(nbt.NewTagName("Id"), nbt.NewIntPayload(3120)),
			nbt.NewStringTag(nbt.NewTagName("Name"), nbt.NewStringPayload("1.19.2")),
			nbt.NewStringTag(nbt.NewTagName("Series"), nbt.NewStringPayload("main")),
			nbt.NewByteTag(nbt.NewTagName("Snapshot"), nbt.NewBytePayload(0)),
			nbt.NewEndTag(),
		)),
		nbt.NewEndTag(),
	))

	fmt.Println(nbt.CompactJson(dat))
	// Output:
	// {"Data":{"GameType":0,"LevelName":"Go NBT","Version":{"Id":3120,"Name":"1.19.2","Series":"main","Snapshot":0}}}
}

func ExamplePrettyJson_space2() {
	dat := nbt.NewCompoundTag(nbt.NewTagName("Data"), nbt.NewCompoundPayload(
		nbt.NewIntTag(nbt.NewTagName("GameType"), nbt.NewIntPayload(0)),
		nbt.NewStringTag(nbt.NewTagName("LevelName"), nbt.NewStringPayload("Go NBT")),
		nbt.NewCompoundTag(nbt.NewTagName("Version"), nbt.NewCompoundPayload(
			nbt.NewIntTag(nbt.NewTagName("Id"), nbt.NewIntPayload(3120)),
			nbt.NewStringTag(nbt.NewTagName("Name"), nbt.NewStringPayload("1.19.2")),
			nbt.NewStringTag(nbt.NewTagName("Series"), nbt.NewStringPayload("main")),
			nbt.NewByteTag(nbt.NewTagName("Snapshot"), nbt.NewBytePayload(0)),
			nbt.NewEndTag(),
		)),
		nbt.NewEndTag(),
	))

	fmt.Println(nbt.PrettyJson(dat, "  "))
	// Output:
	// {
	//   "Data": {
	//     "GameType": 0,
	//     "LevelName": "Go NBT",
	//     "Version": {
	//       "Id": 3120,
	//       "Name": "1.19.2",
	//       "Series": "main",
	//       "Snapshot": 0
	//     }
	//   }
	// }
}

func ExamplePrettyJson_space4() {
	dat := nbt.NewCompoundTag(nbt.NewTagName("Data"), nbt.NewCompoundPayload(
		nbt.NewIntTag(nbt.NewTagName("GameType"), nbt.NewIntPayload(0)),
		nbt.NewStringTag(nbt.NewTagName("LevelName"), nbt.NewStringPayload("Go NBT")),
		nbt.NewCompoundTag(nbt.NewTagName("Version"), nbt.NewCompoundPayload(
			nbt.NewIntTag(nbt.NewTagName("Id"), nbt.NewIntPayload(3120)),
			nbt.NewStringTag(nbt.NewTagName("Name"), nbt.NewStringPayload("1.19.2")),
			nbt.NewStringTag(nbt.NewTagName("Series"), nbt.NewStringPayload("main")),
			nbt.NewByteTag(nbt.NewTagName("Snapshot"), nbt.NewBytePayload(0)),
			nbt.NewEndTag(),
		)),
		nbt.NewEndTag(),
	))

	fmt.Println(nbt.PrettyJson(dat, "    "))
	// Output:
	// {
	//     "Data": {
	//         "GameType": 0,
	//         "LevelName": "Go NBT",
	//         "Version": {
	//             "Id": 3120,
	//             "Name": "1.19.2",
	//             "Series": "main",
	//             "Snapshot": 0
	//         }
	//     }
	// }
}

func ExamplePrettyJson_tab() {
	dat := nbt.NewCompoundTag(nbt.NewTagName("Data"), nbt.NewCompoundPayload(
		nbt.NewIntTag(nbt.NewTagName("GameType"), nbt.NewIntPayload(0)),
		nbt.NewStringTag(nbt.NewTagName("LevelName"), nbt.NewStringPayload("Go NBT")),
		nbt.NewCompoundTag(nbt.NewTagName("Version"), nbt.NewCompoundPayload(
			nbt.NewIntTag(nbt.NewTagName("Id"), nbt.NewIntPayload(3120)),
			nbt.NewStringTag(nbt.NewTagName("Name"), nbt.NewStringPayload("1.19.2")),
			nbt.NewStringTag(nbt.NewTagName("Series"), nbt.NewStringPayload("main")),
			nbt.NewByteTag(nbt.NewTagName("Snapshot"), nbt.NewBytePayload(0)),
			nbt.NewEndTag(),
		)),
		nbt.NewEndTag(),
	))

	fmt.Println(nbt.PrettyJson(dat, "\t"))
	// Output:
	// {
	// 	"Data": {
	// 		"GameType": 0,
	// 		"LevelName": "Go NBT",
	// 		"Version": {
	// 			"Id": 3120,
	// 			"Name": "1.19.2",
	// 			"Series": "main",
	// 			"Snapshot": 0
	// 		}
	// 	}
	// }
}
