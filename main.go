package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

// TODOs:
// - support variadic parameters
// - support searching through nested directories
// - quiet flag not used yet.

func main() {
	config := parseConfigFromArgs(os.Args)

	structInfo, err := parseDir(config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("StructInfo: %v\n", structInfo)

	// create the output directory if it doesn't exist
	err = os.MkdirAll(config.outputDir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	// Create the file
	file, err := os.Create(config.outputFile())
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = executeTemplate(structInfo, file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Generated %s\n", config.outputFile())
}

type Config struct {
	structName string
	dir        string
	outputDir  string
	outputPkg  string
	quiet      bool
}

func (c *Config) outputFile() string {
	return fmt.Sprintf("%s/%s_builder.go", c.outputDir, strings.ToLower(c.structName))
}

func parseConfigFromArgs(args []string) *Config {
	config := &Config{}

	flagSet := flag.NewFlagSet(args[0], flag.ExitOnError)

	flagSet.StringVar(&config.structName, "structname", "", "Name of the struct to generate builder for")
	flagSet.StringVar(&config.outputDir, "outputdir", "", "Directory to write the generated builder to")
	flagSet.StringVar(&config.outputPkg, "outpkg", "builders", "Name of the generated package")
	flagSet.StringVar(&config.dir, "dir", ".", "Directory to search for the struct definition")
	flagSet.BoolVar(&config.quiet, "quiet", false, "suppress output to stdout")

	flagSet.Parse(args[1:])

	if config.structName == "" {
		log.Fatal("Usage: structBuilder -structname=<StructName>")
	}

	if config.outputDir == "" {
		config.outputDir = path.Join(config.dir, "builders")
	}

	return config
}
