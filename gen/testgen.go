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

package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	nbt "github.com/Aton-Kish/gonbt"
	"github.com/iancoleman/strcase"
	"golang.org/x/tools/imports"
)

var (
	//go:embed test.tmpl
	testTemplate string
)

type params struct {
	Type nbt.TagType
}

func testgen() error {
	funcMap := template.FuncMap{
		"public":  public,
		"private": private,
		"typeof":  typeof,
	}

	tmpl, err := template.New("Test").Funcs(funcMap).Parse(testTemplate)
	if err != nil {
		return err
	}

	for _, typ := range nbt.TagTypes {
		p := params{Type: typ}

		raw, err := render(tmpl, p)
		if err != nil {
			return err
		}

		filename := filepath.Join("..", fmt.Sprintf("tag_%02d_%s_test.go", typ, strcase.ToSnake(typ.String())))
		if err := os.WriteFile(filename, raw, os.ModePerm); err != nil {
			return err
		}
	}

	return nil
}

func render(tmpl *template.Template, params params) ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := tmpl.Execute(buf, params); err != nil {
		return nil, err
	}

	raw, err := imports.Process("", buf.Bytes(), nil)
	if err != nil {
		return nil, err
	}

	return raw, nil
}
