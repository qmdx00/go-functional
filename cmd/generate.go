package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"text/template"
)

// types generate func param type slice ...
var types = []reflect.Kind{
	reflect.Uint,
	reflect.Uint8,
	reflect.Uint16,
	reflect.Uint32,
	reflect.Uint64,
	reflect.Int,
	reflect.Int8,
	reflect.Int16,
	reflect.Int32,
	reflect.Int64,
	reflect.Float32,
	reflect.Float64,
	reflect.Bool,
	reflect.String,
}

// findExtFiles ...
func findExtFiles(fsys fs.FS, dir, ext string) (map[string][]byte, error) {
	var err error
	var bts []byte
	var entries []fs.DirEntry
	files := make(map[string][]byte)

	entries, err = fs.ReadDir(fsys, dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if filepath.Ext(entry.Name()) == ext && !entry.IsDir() {
			bts, err = fs.ReadFile(fsys, fmt.Sprintf("%s/%s", dir, entry.Name()))
			if err != nil {
				return nil, err
			}
			files[entry.Name()] = bts
		}
	}
	return files, err
}

// renderTemplate ...
func renderTemplate(name, content string, fm template.FuncMap) ([]byte, error) {
	buff := bytes.NewBufferString("")
	temp, err := template.New(name).Funcs(fm).Parse(content)
	if err != nil {
		return nil, err
	}
	if err = temp.Execute(buff, map[string][]string{
		"types": func() []string {
			tps := make([]string, 0)
			for _, typ := range types {
				tps = append(tps, typ.String())
			}
			return tps
		}(),
	}); err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

// sourceFormat ...
func sourceFormat(source []byte) ([]byte, error) {
	src, err := format.Source(source)
	if err != nil {
		return nil, err
	}
	return src, nil
}

// generateFile ...
func generateFile(name string, content []byte) error {
	return ioutil.WriteFile(fmt.Sprintf("%s.go", strings.Split(name, ".")[0]), content, 0644)
}

// customFunc
func customFunc() template.FuncMap {
	return template.FuncMap{
		"firstToUpper": func(s string) string {
			if len(s) > 1 {
				return strings.ToUpper(string(s[0])) + s[1:]
			} else {
				return s
			}
		},
	}
}

func main() {
	files, _ := findExtFiles(os.DirFS("."), "template", ".tmpl")
	for name, content := range files {
		source, _ := renderTemplate(name, string(content), customFunc())
		formatted, _ := sourceFormat(source)
		_ = generateFile(name, formatted)
	}
}
