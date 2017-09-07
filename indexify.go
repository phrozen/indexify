package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
)

type Directory struct {
	Name string
	Path string
	List []os.FileInfo
}

func main() {

	funcMap := template.FuncMap{
		"formatSize": FormatBytes,
		"type":       DetectType,
	}

	dir := Directory{Name: "", Path: "."}
	list, err := ioutil.ReadDir(dir.Path)
	if err != nil {
		panic(err)
	}
	dir.List = list

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
