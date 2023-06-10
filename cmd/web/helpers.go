package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/KXLAA/chars/pkg/version"
)

func (app *application) newTemplateData(r *http.Request) map[string]any {
	data := map[string]any{
		"Version": version.Get(),
	}

	return data
}

func (app *application) parseUrlQueryWithDefaults(r *http.Request) (configForm, error) {
	length := resolveIntQuery(r, "length", 32)
	count := resolveIntQuery(r, "count", 1)
	lowercase := resolveBoolQueryWithDefaults(r, "lowercase", true)
	uppercase := resolveBoolQueryWithDefaults(r, "uppercase", false)
	numbers := resolveBoolQueryWithDefaults(r, "numbers", false)
	special := resolveBoolQueryWithDefaults(r, "special", true)

	return configForm{
		Count:     count,
		Length:    length,
		LowerCase: lowercase,
		UpperCase: uppercase,
		Numbers:   numbers,
		Special:   special,
	}, nil
}

func resolveIntQuery(r *http.Request, key string, defaultValue int) int {
	value, err := strconv.Atoi(r.URL.Query().Get(key))
	if err != nil {
		return defaultValue
	}
	return value
}

func resolveBoolQuery(r *http.Request, key string) bool {
	value := r.URL.Query().Get(key)
	if value == "" {
		return false
	}

	//convert to bool for consistency with checkbox values ion HTML
	return convertToBool(value)
}

func resolveBoolQueryWithDefaults(r *http.Request, key string, defaultValue bool) bool {
	value := r.URL.Query().Get(key)
	if value == "" {
		return defaultValue
	}

	//convert to bool for consistency with checkbox values ion HTML
	return convertToBool(value)
}

func randomNumber(min, max int) int {
	return min + rand.Intn(max-min)
}

func convertToBool(value string) bool {
	//convert to bool for consistency with checkbox values ion HTML
	if value == "on" {
		return true
	}

	return value == "true"
}
