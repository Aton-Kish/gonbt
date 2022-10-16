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
	//go:embed tag_test.tmpl
	tagTestTemplate string
	//go:embed payload_test.tmpl
	payloadTestTemplate string
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

	tagTmpl, err := template.New("TagTest").Funcs(funcMap).Parse(tagTestTemplate)
	if err != nil {
		return err
	}

	payloadTmpl, err := template.New("PayloadTest").Funcs(funcMap).Parse(payloadTestTemplate)
	if err != nil {
		return err
	}

	for _, typ := range nbt.TagTypes {
		p := params{Type: typ}

		tagRaw, err := render(tagTmpl, p)
		if err != nil {
			return err
		}

		tagName := filepath.Join("..", fmt.Sprintf("tag_%s_test.go", strcase.ToSnake(typ.String())))
		if err := os.WriteFile(tagName, tagRaw, os.ModePerm); err != nil {
			return err
		}

		if typ == nbt.TagTypeEnd {
			continue
		}

		payloadRaw, err := render(payloadTmpl, p)
		if err != nil {
			return err
		}

		payloadName := filepath.Join("..", fmt.Sprintf("payload_%s_test.go", strcase.ToSnake(typ.String())))
		if err := os.WriteFile(payloadName, payloadRaw, os.ModePerm); err != nil {
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
