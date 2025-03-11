# Struct Builder
structbuilder is a CLI tool for Go that automates the generation of "builder" structs and methods for your existing structs

## Installation
To install the structBuilder tool, you'll need to ensure that your Go environment is set up with the necessary paths. You can install the tool using the following command:
```sh
go install github.com/yourusername/structBuilder@latest
```

## Usage
structBuilder is meant to be integrated directly into your Go project's build process using the go:generate directive. This allows you to automatically generate builder structs whenever you run go generate.

### How to Use
1. Annotate your structs with the go:generate directive specifying the -structname flag. This will tell structBuilder which structs need builders.
2. Execute go generate in your package directory to automatically generate the builder structs and methods.

### Example
Consider you have the following structs in your models package:

```go
package models

//go:generate structBuilder -structname=Person
type Person struct {
    FirstName string
    LastName  string
    Age       int
    Height    float64
    Weight    float64
}

//go:generate structBuilder -structname=Animal
type Animal struct {
    Name string
    Age  int
}
```

The directives `//go:generate structBuilder -structname=Person` and `//go:generate structBuilder -structname=Animal` specify that you want to create builder structs for Person and Animal.

After setting up the directives:
1. Open a terminal and navigate to the directory containing your models package. 
2. Run the following command: `go generate`

This will execute the structBuilder tool for each of the specified structs, generating files with builder methods.

### Output

For the Person struct, the PersonBuilder struct would then look something like this:

```go
// Code generated by structbuilder. DO NOT EDIT.

package builders

import "structbuilder/models"

type PersonBuilder struct {
	FirstName string
	LastName string
	Age int
	Height float64
	Weight float64
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{}
}

func (b *PersonBuilder) WithFirstName(value string) *PersonBuilder {
	b.FirstName = value
	return b
}

func (b *PersonBuilder) WithLastName(value string) *PersonBuilder {
	b.LastName = value
	return b
}

func (b *PersonBuilder) WithAge(value int) *PersonBuilder {
	b.Age = value
	return b
}

func (b *PersonBuilder) WithHeight(value float64) *PersonBuilder {
	b.Height = value
	return b
}

func (b *PersonBuilder) WithWeight(value float64) *PersonBuilder {
	b.Weight = value
	return b
}


func (b *PersonBuilder) Build() *models.Person {
	return &models.Person{
		FirstName: b.FirstName,
		LastName: b.LastName,
		Age: b.Age,
		Height: b.Height,
		Weight: b.Weight,
	}
}
```