package main

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func mainL() {
	fmt.Println(cases.Title(language.German).String("goSAMples.dev is the best Go bLog in the world!"))
	fmt.Println(cases.Title(language.Und, cases.NoLower).String("goSAMples.dev is the best Go bLog in the world!"))
}
