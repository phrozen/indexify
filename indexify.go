package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
)

// Directory type for recursive/template purposes
type Directory struct {
	Path string
	List []os.FileInfo
}

// NewDirectory returns a new isntance of a directory type
func NewDirectory(path string) *Directory {
	return &Directory{Path: path}
}

func main() {

	funcMap := template.FuncMap{
		"formatSize": FormatBytes,
		"type":       DetectType,
		"title":      NameFromPath,
	}

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dir := NewDirectory(path)
	dir.List, err = ioutil.ReadDir(dir.Path)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(fmt.Sprintf("%s/index.html", dir.Path))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	t := template.Must(template.New("index").Funcs(funcMap).Parse(tmpl))
	err = t.Execute(f, dir)
	if err != nil {
		panic(err)
	}

}
