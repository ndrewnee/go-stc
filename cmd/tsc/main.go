package main

import (
	"log"

	"github.com/ndrewnee/go-stc/compiler"
)

func main() {
	program := "(add 10 (subtract 10 6))"
	output, err := compiler.Compile(program)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(output)
}
