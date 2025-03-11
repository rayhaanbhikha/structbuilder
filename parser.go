package main

import (
	"fmt"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func parseDir(config *Config) (*StructInfo, error) {
	cfg := &packages.Config{
		Mode:  packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedName,
		Tests: false,
		Dir:   config.dir,
	}

	// Load all packages in the module or a specific subset
	// You can specify patterns like "./..." for the current module recursively or specific paths.
	pkgs, err := packages.Load(cfg, "./...")
	if err != nil {
		return nil, err
	}

	for _, pkg := range pkgs {

		// Iterate over each file in the package
		for _, syn := range pkg.Syntax {
			structInfo := parseFileDeclarations(syn, config.structName)
			if structInfo != nil {

				structInfo.OutputPackageName = config.outputPkg
				structInfo.ImportPath = pkg.PkgPath
				structInfo.OutputTypeName = fmt.Sprintf("%s.%s", pkg.Name, config.structName)

				return structInfo, nil
			}
		}
	}

	return nil, fmt.Errorf("struct %s not found in directory %s", config.structName, config.dir)
}

func parseFileDeclarations(file *ast.File, structName string) *StructInfo {
	for _, decl := range file.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range genDecl.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}

				structType, ok := typeSpec.Type.(*ast.StructType)
				if !ok {
					continue
				}

				if typeSpec.Name.Name != structName {
					continue
				}

				var fields []StructField

				for _, field := range structType.Fields.List {
					for _, name := range field.Names {
						// do we care if the field is Uppercased or not?
						fields = append(fields, StructField{Name: name.Name, Type: fmt.Sprint(field.Type)})
					}
				}

				return &StructInfo{
					BuilderTypeName: structName,
					Fields:          fields,
				}
			}
		}
	}

	return nil
}
