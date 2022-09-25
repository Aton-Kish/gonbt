# Go NBT

The library for Minecraft NBT

## Getting Started

Use `go get` to install the library

```shell
go get github.com/Aton-Kish/gonbt
```

Import in your application

```go
import (
	nbt "github.com/Aton-Kish/gonbt"
)
```

## Usage

```go
package main

import (
	"compress/gzip"
	"fmt"
	"log"
	"os"

	nbt "github.com/Aton-Kish/gonbt"
)

func main() {
  // Load NBT Data
	f, err := os.Open("level.dat")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	zr, err := gzip.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}
	defer zr.Close()

	// Decode
	dat, err := nbt.Decode(zr)
	if err != nil {
		log.Fatal(err)
	}

	// "Reversible" Stringify
	snbt := nbt.Stringify(dat)
	// snbt := nbt.CompactStringify(dat)
	// snbt := nbt.PrettyStringify(dat, "  ")
	fmt.Println(snbt)

	// Parse
	dat2, err := nbt.Parse(snbt)
	if err != nil {
		log.Fatal(err)
	}

	// "Irreversible" Stringify
	json := nbt.Json(dat2)
	// json := nbt.CompactJson(dat2)
	// json := nbt.PrettyJson(dat2, "  ")
	fmt.Println(json)
}
```

## License

This library is licensed under the MIT License, see [LICENSE](./LICENSE).
