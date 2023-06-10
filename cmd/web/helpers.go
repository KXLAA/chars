package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/KXLAA/chars/pkg/response"
	"github.com/KXLAA/chars/pkg/validator"
	"github.com/KXLAA/chars/pkg/version"
)

type configForm struct {
	Count     int                 `form:"count"`
	Length    int                 `form:"length"`
	LowerCase bool                `form:"lowercase"`
	UpperCase bool                `form:"uppercase"`
	Numbers   bool                `form:"numbers"`
	Special   bool                `form:"special"`
	Validator validator.Validator `form:"-"`
}

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
	form.Validator.CheckField(form.Length > 0, "Length", "Length must be greater than 0")
	form.Validator.CheckField(form.Length <= 100, "Length", "Length must be less than 100")
	if !form.LowerCase && !form.UpperCase && !form.Numbers && !form.Special {
		form.Validator.AddFieldError("Empty", "At least one option must be selected")
	}

	return form
}

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

func parseLengthValue(w http.ResponseWriter, r *http.Request) (int, error) {
	length, err := strconv.Atoi(r.URL.Query().Get("length"))

	if err != nil {
		err = response.JSON(w, http.StatusBadRequest, map[string]string{
			"error": ErrInvalidQueryParams.Error(),
		})

		return 0, err
	}

	return length, nil
}

func parseCount(w http.ResponseWriter, r *http.Request) (int, error) {
	var count int
	var err error
	if r.URL.Query().Get("count") == "" {
		count = 1
	} else {
		count, err = strconv.Atoi(r.URL.Query().Get("count"))
		if err != nil {
			err = response.JSON(w, http.StatusBadRequest, map[string]string{
				"error": ErrInvalidQueryParams.Error(),
			})
			return 0, err
		}
	}

	return count, nil
}
