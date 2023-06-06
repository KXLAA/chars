package main

import (
	"fmt"
	"net/http"
	"regexp"
	"testing"
)

var (
	charRxAll = "^[a-zA-Z0-9!@#$%^&*()_+\\-=\\[\\]{};':\"\\\\|,.<>\\/?]*$"
	//charRx    = regexp.MustCompile(`<span class="generator__output__text" id="generator__output__text">(.+?)</span>`)
)

func TestHome(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/")

	_, err := rxMatch(body, charRxAll)

	if err != nil {
		t.Errorf("could not find matching random string in body")
	}

	if code != http.StatusOK {
		t.Errorf("want 200; got %d", code)
	}

}

func TestGenerate(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	tests :=
		[]struct {
			name     string
			url      string
			wantCode int
			matchRx  *regexp.Regexp
		}{
			{
				name:     "generates a random string when no query params are provided",
				url:      "/generate",
				wantCode: http.StatusOK,
				matchRx:  regexp.MustCompile(charRxAll),
			}, {
				name:     "generates a random string with a length of 10",
				url:      "/generate?length=10",
				wantCode: http.StatusOK,
				matchRx:  regexp.MustCompile(`^.{10}$`),
			}, {
				name:     "generates a random string with a length of 10 and only lowercase characters",
				url:      "/generate?length=10&lowercase=true",
				wantCode: http.StatusOK,
				matchRx:  regexp.MustCompile(`^[a-z]{10}$`),
			}, {
				name:     "generates a random string with a length of 10 and only uppercase characters",
				url:      "/generate?length=10&uppercase=true",
				wantCode: http.StatusOK,
				matchRx:  regexp.MustCompile(`^[A-Z]{10}$`),
			}, {
				name:     "generates a random string with a length of 10 and only numbers",
				url:      "/generate?length=10&numbers=true",
				wantCode: http.StatusOK,
				matchRx:  regexp.MustCompile(`^[0-9]{10}$`),
			}, {
				name:     "generates a random string with a length of 10 and only special characters",
				url:      "/generate?length=10&special=true",
				wantCode: http.StatusOK,
				matchRx:  regexp.MustCompile(`^[!@#$%^&*()_+\-=\[\]{};':\"\\|,.<>\/?]{10}$`),
			},
			{
				name:     "generates a random string with a combination of all config options",
				url:      "/generate?length=10&lowercase=true&uppercase=true&numbers=true&special=true",
				wantCode: http.StatusOK,
				matchRx:  regexp.MustCompile(charRxAll),
			},
		}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, body := ts.get(t, tt.url)
			char := extractChar(t, body)
			fmt.Println(char)

			if !tt.matchRx.MatchString(char) {
				t.Errorf("want %s; got %s", tt.matchRx, char)
			}

			if code != tt.wantCode {
				t.Errorf("want %d; got %d", tt.wantCode, code)
			}
		})
	}

}

func rxMatch(body []byte, regexPattern string) (bool, error) {
	match, err := regexp.Match(regexPattern, body)
	if err != nil {
		return false, err
	}

	fmt.Println(match)

	return match, nil
}

func extractChar(t *testing.T, body []byte) string {
	matches, err := findMatchedString(`<span class="generator__output__text" id="generator__output__text">(.+?)</span>`, body)
	if err != nil {
		t.Fatalf("could not find random character in body: %s", body)
	}

	return matches
}

func findMatchedString(pattern string, input []byte) (string, error) {
	// Compile the regex pattern
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}

	// Find the matching content within the input byte slice
	match := regex.FindSubmatch(input)
	fmt.Println(match, "match")

	// Check if a match is found
	if len(match) > 1 {
		return string(match[1]), nil
	}

	return "", fmt.Errorf("No match found")
}
