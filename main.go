package main

import (
	"flag"
	"fmt"
	"log"

	"strings"
)

func main() {
	structName := flag.String("type", "", "Name of the struct to generate builder for")
	// TODO: is this needed?
	inputFile := flag.String("file", "", "Go source file containing the struct")
	flag.Parse()

	if *structName == "" || *inputFile == "" {
		log.Fatal("Usage: structBuilder -type=StructName -file=source.go")
	}

	structInfo, err := parseGoFile(*inputFile, *structName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("StructInfo: ", structInfo)

	outputFile := fmt.Sprintf("%s_builder.go", strings.ToLower(*structName))
	err = create(outputFile, structInfo)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Generated %s\n", outputFile)
}
