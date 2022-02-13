package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/muesli/termenv"
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

const titleTmpl = `{{define "title"}}Title: {{.Title}}{{end}}`

const itemTmpl = `{{define "item" -}}
[{{ if .Done }}x{{ else }} {{ end }}] {{ if .Active }}{{ .Name }}{{ else }}{{ Faint .Name }}{{ end }}
{{- end}}`

const listTmpl = `{{define "list"}}{{template "title" .}}
{{- range .Items }}
  {{template "item" .}}
{{- else}}
  (No items in list.)
{{- end}}
{{end}}`

const tmpl = `{{ template "list" . }}`

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
	t := template.New("my-list")
	t = t.Funcs(template.FuncMap{
		"toUpper": strings.ToUpper,
	})

	// Add the termenv template functions
	cp := termenv.ColorProfile()
	fmt.Printf("Terminal color profile: %+v\n", cp)
	f := termenv.TemplateFuncs(cp)
	t = t.Funcs(f)

	t, err := t.Parse(titleTmpl)
	if err != nil {
		panic(err)
	}
	t, err = t.Parse(itemTmpl)
	if err != nil {
		panic(err)
	}
	t, err = t.Parse(listTmpl)
	if err != nil {
		panic(err)
	}
	t, err = t.Parse(tmpl)
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, list)
	if err != nil {
		panic(err)
	}
}
