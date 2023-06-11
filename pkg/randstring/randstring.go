package randstring

import (
	"math/rand"
	"time"
)

type Config struct {
	Length            int
	Count             int
	LowerCase         bool
	UpperCase         bool
	Numbers           bool
	SpecialCharacters bool
}

func RandomString(config *Config) ([]string, error) {
	var result []string
	for i := 0; i < config.Count; i++ {
		s, err := generateRandomString(config)
		if err != nil {
			return nil, err
		}
		result = append(result, s)
	}
	return result, nil
}

func generateRandomString(config *Config) (string, error) {
	// Create a new random number generator with a time-based seed
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// The pool of characters to choose from
	pool := resolveCharacters(config)

	// Generate a random string of the specified length
	result := make([]byte, config.Length)

	for i := range result {
		// A random index into the character pool
		result[i] = pool[r.Intn(len(pool))]
	}

	return string(result), nil
}

func resolveCharacters(config *Config) string {
	var finalPool string

	loweCasePool := "abcdefghijklmnopqrstuvwxyz"
	upperCasePool := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbersPool := "0123456789"
	specialCharactersPool := "!@#$%^&*()_-+=~`[]{}\\|:;\"'<>,.?/"

	if config.LowerCase {
		finalPool += loweCasePool
	}

	if config.UpperCase {
		finalPool += upperCasePool
	}

	if config.Numbers {
		finalPool += numbersPool
	}

	if config.SpecialCharacters {
		finalPool += specialCharactersPool
	}

	return finalPool

}
