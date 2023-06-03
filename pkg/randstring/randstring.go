package randstring

import (
	"math/rand"
	"time"
)

type Config struct {
	Count             int
	LowerCase         bool
	UpperCase         bool
	Numbers           bool
	SpecialCharacters bool
}

func RandomString(config *Config) (string, error) {
	// Create a new random number generator with a time-based seed
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// The pool of characters to choose from
	pool := resolveCharacters(config.LowerCase, config.UpperCase, config.Numbers, config.SpecialCharacters)

	// Generate a random string of the specified length
	result := make([]byte, config.Count)

	for i := range result {
		// A random index into the character pool
		result[i] = pool[r.Intn(len(pool))]
	}

	return string(result), nil
}

func resolveCharacters(lowerCase bool, upperCase bool, numbers bool, specialCharacters bool) string {
	var finalPool string

	loweCasePool := "abcdefghijklmnopqrstuvwxyz"
	upperCasePool := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbersPool := "0123456789"
	specialCharactersPool := "!@#$%^&*()_+"

	if lowerCase {
		finalPool += loweCasePool
	}

	if upperCase {
		finalPool += upperCasePool
	}

	if numbers {
		finalPool += numbersPool
	}

	if specialCharacters {
		finalPool += specialCharactersPool
	}

	return finalPool

}
