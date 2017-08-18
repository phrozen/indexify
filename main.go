package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
)

const tmpl = `
<!DOCTYPE html>
<html>
	<head>
	</head>
	<body>
		<table>
			<tr>
				<th>Size</th>
				<th>Name</th>
				<th>Last Modified</th>
			</tr>
			{{range .}}
				<tr>
					{{if .IsDir}}
						<td>[dir]</td>
					{{else}}
						<td>{{.Size}}</td>
					{{end}}
					<td>
						<a href="./{{.Name}}">{{.Name}}</a>
					</td>
					<td>{{.ModTime}}</td>
				</tr>
			{{end}}
		</table>
	</body>
</html>
`
const dir = "."

func main() {
	list, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	f, err := os.Create(fmt.Sprintf("%s/index.html", dir))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	t := template.Must(template.New("index").Parse(tmpl))
	err = t.Execute(f, list)
	if err != nil {
		panic(err)
	}
}
