package main

import (
	"encoding/json"
	"html"
	"net/http"
	"regexp"
	"testing"
)

var (
	randomCharRx      = regexp.MustCompile("^[a-zA-Z0-9!@#$%^&*()_+\\-=\\[\\]{};':\"\\\\|,.<>\\/?]*$")
	charRx            = regexp.MustCompile(`<span class="generator__output__text" id="generator__output__text">(.+?)</span>`)
	lowercase10CharRx = regexp.MustCompile(`^[a-z]{10}$`)
	uppercase10CharRx = regexp.MustCompile(`^[A-Z]{10}$`)
	numbers10CharRx   = regexp.MustCompile(`^[0-9]{10}$`)
	special10CharRx   = regexp.MustCompile(`^[!@#$%^&*()_+\-=\[\]{};':\"\\|,.<>\/?]{10}$`)
)

func TestHome(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/")
	_, err := rxMatch(t, body, randomCharRx)
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
				matchRx:  randomCharRx,
			}, {
				name:     "generates a random string with a length of 10 and only lowercase characters",
				url:      "/generate?length=10&lowercase=true",
				wantCode: http.StatusOK,
				matchRx:  lowercase10CharRx,
			}, {
				name:     "generates a random string with a length of 10 and only uppercase characters",
				url:      "/generate?length=10&uppercase=true",
				wantCode: http.StatusOK,
				matchRx:  uppercase10CharRx,
			}, {
				name:     "generates a random string with a length of 10 and only numbers",
				url:      "/generate?length=10&numbers=true",
				wantCode: http.StatusOK,
				matchRx:  numbers10CharRx,
			}, {
				name:     "generates a random string with a length of 10 and only special characters",
				url:      "/generate?length=10&special=true",
				wantCode: http.StatusOK,
				matchRx:  special10CharRx,
			},
			{
				name:     "generates a random string with a combination of all config options",
				url:      "/generate?length=10&lowercase=true&uppercase=true&numbers=true&special=true",
				wantCode: http.StatusOK,
				matchRx:  randomCharRx,
			},
		}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, body := ts.get(t, tt.url)

			char := extractChar(t, body)
			if !tt.matchRx.MatchString(char) {
				t.Errorf("want %s; got %s", tt.matchRx, char)
			}

			if code != tt.wantCode {
				t.Errorf("want %d; got %d", tt.wantCode, code)
			}
		})
	}
}

func TestApiGenerate(t *testing.T) {
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
				url:      "/api/v1/generate",
				wantCode: http.StatusOK,
				matchRx:  randomCharRx,
			}, {
				name:     "generates a random string with a length of 10 and only lowercase characters",
				url:      "/api/v1/generate?length=10&lowercase=true",
				wantCode: http.StatusOK,
				matchRx:  lowercase10CharRx,
			}, {
				name:     "generates a random string with a length of 10 and only uppercase characters",
				url:      "/api/v1/generate?length=10&uppercase=true",
				wantCode: http.StatusOK,
				matchRx:  uppercase10CharRx,
			}, {
				name:     "generates a random string with a length of 10 and only numbers",
				url:      "/api/v1/generate?length=10&numbers=true",
				wantCode: http.StatusOK,
				matchRx:  numbers10CharRx,
			}, {
				name:     "generates a random string with a length of 10 and only special characters",
				url:      "/api/v1/generate?length=10&special=true",
				wantCode: http.StatusOK,
				matchRx:  special10CharRx,
			},
			{
				name:     "generates a random string with a combination of all config options",
				url:      "/api/v1/generate?length=10&lowercase=true&uppercase=true&numbers=true&special=true",
				wantCode: http.StatusOK,
				matchRx:  randomCharRx,
			},
		}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input struct {
				RandomString string `json:"randomString"`
			}
			code, _, body := ts.get(t, tt.url)
			err := json.Unmarshal(body, &input)
			if err != nil {
				t.Errorf("could not unmarshal json: %v", err)
			}

			char := input.RandomString

			if !tt.matchRx.MatchString(char) {
				t.Errorf("want %s; got %s", tt.matchRx, char)
			}

			if code != tt.wantCode {
				t.Errorf("want %d; got %d", tt.wantCode, code)
			}
		})
	}

	t.Run("returns a 400 when invalid values for length or count in query params", func(t *testing.T) {
		code, _, _ := ts.get(t, "/api/v1/generate?length=invalid&count=invalid")

		if code != http.StatusBadRequest {
			t.Errorf("want %d; got %d", http.StatusBadRequest, code)
		}
	})

	t.Run("returns a 400 when incomplete config is passed in quey param", func(t *testing.T) {
		code, _, _ := ts.get(t, "/api/v1/generate?length=10&count=100")

		if code != http.StatusBadRequest {
			t.Errorf("want %d; got %d", http.StatusBadRequest, code)
		}
	})

	t.Run("returns an array of random strings when count is greater than 1", func(t *testing.T) {
		var input struct {
			RandomString []string `json:"randomString"`
		}
		code, _, body := ts.get(t, "/api/v1/generate?count=3&length=10&lowercase=true&uppercase=true&numbers=true&special=true")
		err := json.Unmarshal(body, &input)
		if err != nil {
			t.Errorf("could not unmarshal json: %v", err)
		}

		if len(input.RandomString) != 3 {
			t.Errorf("want 3; got %d", len(input.RandomString))
		}

		for _, char := range input.RandomString {
			if !randomCharRx.MatchString(char) {
				t.Errorf("want %s; got %s", randomCharRx, char)
			}
		}

		if code != http.StatusOK {
			t.Errorf("want %d; got %d", http.StatusOK, code)
		}
	})

}

func rxMatch(t *testing.T, body []byte, rx *regexp.Regexp) (bool, error) {
	match, err := regexp.Match(rx.String(), body)
	if err != nil {
		return false, err
	}

	return match, nil
}

func extractChar(t *testing.T, body []byte) string {
	matches := charRx.FindSubmatch(body)

	if len(matches) < 2 {
		t.Fatalf("no match for %q in body %s", charRx, body)
	}

	return html.UnescapeString(string(matches[1]))
}
