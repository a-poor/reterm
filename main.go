package main

import (
	"html/template"
	"os"
	"strings"
)

type ListItem struct {
	Name   string
	Done   bool
	Active bool
}

type List struct {
	Title string
	Items []ListItem
}

const tmpl = `Title: {{ .Title }}
{{- range .Items }}
[{{ if .Done }}x{{ else }} {{ end }}] {{ if .Active }}{{ toUpper .Name }}{{ else }}{{ .Name }}{{ end }}
{{- end }}
`

func main() {
	// Create a list...
	list := List{
		Title: "My List",
		Items: []ListItem{
			{Name: "milk", Active: true, Done: false},
			{Name: "eggs"},
			{Name: "bread", Done: true},
			{Name: "liquid death"},
		},
	}

	// Parse the template & display...
	t := template.New("list")
	t.Funcs(template.FuncMap{
		"toUpper": strings.ToUpper,
	})
	t, err := t.Parse(tmpl)
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, list)
	if err != nil {
		panic(err)
	}
}
