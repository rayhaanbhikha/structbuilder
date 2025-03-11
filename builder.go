package main

import (
	"os"
	"text/template"
)

type StructField struct {
	Name string
	Type string
}

type StructInfo struct {
	PackageName string
	TypeName    string
	Fields      []StructField
}

// Go template for generating a builder
const builderTemplate = `package {{.PackageName}}
type {{.TypeName}}Builder struct {
	{{- range .Fields}}
	{{.Name}} {{.Type}}
	{{- end}}
}

func New{{.TypeName}}Builder() *{{.TypeName}}Builder {
	return &{{.TypeName}}Builder{}
}

{{- range .Fields}}
func (b *{{$.TypeName}}Builder) With{{.Name}}(value {{.Type}}) *{{$.TypeName}}Builder {
	b.{{.Name}} = value
	return b
}
{{- end}}

func (b *{{.TypeName}}Builder) Build() *{{.TypeName}} {
	return &{{.TypeName}}{
		{{- range .Fields}}
		{{.Name}}: b.{{.Name}},
		{{- end}}
	}
}
`

func create(filePath string, structInfo *StructInfo) error {
	tmpl, err := template.New("builder").Parse(builderTemplate)
	if err != nil {
		return err
	}

	// Create the file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	// Execute the template
	err = tmpl.Execute(file, structInfo)
	if err != nil {
		return err
	}

	return nil
}
