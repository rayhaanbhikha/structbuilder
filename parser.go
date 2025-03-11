package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func parseGoFile(filename, structName string) (*StructInfo, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, nil, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	for _, decl := range file.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					if structType, ok := typeSpec.Type.(*ast.StructType); ok && typeSpec.Name.Name == structName {
						var fields []StructField
						for _, field := range structType.Fields.List {
							for _, name := range field.Names {
								// do we care if the field is Uppercased or not?
								fields = append(fields, StructField{Name: name.Name, Type: fmt.Sprint(field.Type)})
							}
						}
						return &StructInfo{
							PackageName: file.Name.Name,
							TypeName:    structName,
							Fields:      fields,
						}, nil
					}
				}
			}
		}
	}
	return nil, fmt.Errorf("struct %s not found", structName)
}
