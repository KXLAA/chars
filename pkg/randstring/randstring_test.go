package randstring

import "testing"

func TestRandomString(t *testing.T) {

	t.Run("Generates the correct number of characters based on count", func(t *testing.T) {
		config := Config{
			Count:             50,
			LowerCase:         true,
			UpperCase:         true,
			Numbers:           true,
			SpecialCharacters: true,
		}

		result, _ := RandomString(&config)
		if len(result) != 50 {
			t.Errorf("Expected string of length 10, got %d", len(result))
		}
	})

	t.Run("Generates string with lower case, upper case, numbers and special characters", func(t *testing.T) {
		config := Config{
			Count:             50,
			LowerCase:         true,
			UpperCase:         true,
			Numbers:           true,
			SpecialCharacters: true,
		}
		result, _ := RandomString(&config)
		if !containsLowerCase(result) {
			t.Errorf("Expected string to contain lower case characters, got %s", result)
		}

		if !containsUpperCase(result) {
			t.Errorf("Expected string to contain upper case characters, got %s", result)
		}

		if !containsNumbers(result) {
			t.Errorf("Expected string to contain numbers, got %s", result)
		}

		if !containsSpecialCharacters(result) {
			t.Errorf("Expected string to contain special characters, got %s", result)
		}
	})

	t.Run("Generates random string with only lower case characters ", func(t *testing.T) {
		config := Config{
			Count:             50,
			LowerCase:         true,
			UpperCase:         false,
			Numbers:           false,
			SpecialCharacters: false,
		}

		result, _ := RandomString(&config)
		if !containsLowerCase(result) {
			t.Errorf("Expected string to contain lower case characters, got %s", result)
		}
	})

	t.Run("Generates random string with only upper case characters ", func(t *testing.T) {
		config := Config{
			Count:             50,
			LowerCase:         false,
			UpperCase:         true,
			Numbers:           false,
			SpecialCharacters: false,
		}

		result, _ := RandomString(&config)
		if !containsUpperCase(result) {
			t.Errorf("Expected string to contain upper case characters, got %s", result)
		}
	})

	t.Run("Generates random string with only numbers ", func(t *testing.T) {
		config := Config{
			Count:             50,
			LowerCase:         false,
			UpperCase:         false,
			Numbers:           true,
			SpecialCharacters: false,
		}

		result, _ := RandomString(&config)
		if !containsNumbers(result) {
			t.Errorf("Expected string to contain numbers, got %s", result)
		}
	})

	t.Run("Generates random string with only special characters ", func(t *testing.T) {
		config := Config{
			Count:             50,
			LowerCase:         false,
			UpperCase:         false,
			Numbers:           false,
			SpecialCharacters: true,
		}

		result, _ := RandomString(&config)
		if !containsSpecialCharacters(result) {
			t.Errorf("Expected string to contain special characters, got %s", result)
		}
	})

}

// function that checks if the string contains lowercase characters
func containsUpperCase(s string) bool {
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			return true
		}
	}
	return false
}

// function that checks if the string contains uppercase characters
func containsLowerCase(s string) bool {
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			return true
		}
	}
	return false
}

// function that checks if the string contains numbers
func containsNumbers(s string) bool {
	for _, c := range s {
		if c >= '0' && c <= '9' {
			return true
		}
	}
	return false
}

// function that checks if the string contains special characters
func containsSpecialCharacters(s string) bool {
	for _, c := range s {
		if c >= '!' && c <= '+' {
			return true
		}
	}
	return false
}
