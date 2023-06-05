package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/KXLAA/chars/pkg/randstring"
)

func main() {
	lowerCase := flag.Bool("lc", true, "whether the random text contains lowercase characters")
	upperCase := flag.Bool("uc", false, "whether the random text contains uppercase characters")
	numbers := flag.Bool("num", false, "whether the random text contains numbers")
	specialCharacters := flag.Bool("sc", true, "whether the random text contains special characters")
	count := flag.Int("c", 1, "the number of random strings to generate")
	length := flag.Int("l", 32, "the length of the random string")

	flag.Parse()

	config := randstring.Config{
		Length:            *length,
		Count:             *count,
		LowerCase:         *lowerCase,
		UpperCase:         *upperCase,
		Numbers:           *numbers,
		SpecialCharacters: *specialCharacters,
	}

	text, err := randstring.RandomString(&config)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
		return
	}

	for _, t := range text {
		fmt.Println(t)
	}

}
