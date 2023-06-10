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

func (app *application) parseForm(w http.ResponseWriter, r *http.Request) configForm {
	form := configForm{}
	form.Length = resolveIntQuery(r, "length", 0)
	form.Count = resolveIntQuery(r, "count", 1)
	form.LowerCase = resolveBoolQuery(r, "lowercase")
	form.UpperCase = resolveBoolQuery(r, "uppercase")
	form.Numbers = resolveBoolQuery(r, "numbers")
	form.Special = resolveBoolQuery(r, "special")
	form.Validator.CheckField(form.Length > 0, "Length", "length must be greater than 0")
	if !form.LowerCase && !form.UpperCase && !form.Numbers && !form.Special {
		form.Validator.AddFieldError("Empty", "At least one option must be selected")
	}

	return form
}

// func (app *application) parseUrlQueries(w http.ResponseWriter, r *http.Request) {
// 	form := configForm{}

// }

func (app *application) parseUrlQueriesWithDefaults(r *http.Request) (configForm, error) {
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
