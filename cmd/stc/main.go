package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ndrewnee/go-stc/compiler"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println(`Usage: stc "program code". For example: stc "(add 10 (subtract 10 6))".`)
		os.Exit(0)
	}

	program := os.Args[1]
	output, err := compiler.Compile(program)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)
}
