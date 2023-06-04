package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/KXLAA/chars/pkg/randstring"
)

func main() {
	lowerCase := flag.Bool("lc", true, "whether the random text contains lowercase characters")
	upperCase := flag.Bool("uc", false, "whether the random text contains uppercase characters")
	numbers := flag.Bool("num", false, "whether the random text contains numbers")
	specialCharacters := flag.Bool("sc", true, "whether the random text contains special characters")
	count := flag.Int("c", 1, "the number of random strings to generate")

	flag.Parse()

	//Get the length of characters to generate from the command line
	countArgs := os.Args[1:][0]

	//convert string to int
	length, err := strconv.Atoi(countArgs)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	config := randstring.Config{
		Length:            length,
		Count:             *count,
		LowerCase:         *lowerCase,
		UpperCase:         *upperCase,
		Numbers:           *numbers,
		SpecialCharacters: *specialCharacters,
	}

	text, error := randstring.RandomString(&config)

	if error != nil {
		fmt.Println("Error: ", error)
		os.Exit(1)
		return
	}

	fmt.Println(text)
}
