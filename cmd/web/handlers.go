package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/KXLAA/chars/pkg/randstring"
	"github.com/KXLAA/chars/pkg/response"
	"github.com/KXLAA/chars/pkg/validator"
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

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	var form configForm

	form.Length = randomNumber(32, 54)
	form.Count = 1
	form.LowerCase = true
	form.UpperCase = true
	form.Numbers = true
	form.Special = true

	result, err := randstring.RandomString(&randstring.Config{
		Length:            form.Length,
		Count:             form.Count,
		LowerCase:         form.LowerCase,
		UpperCase:         form.UpperCase,
		Numbers:           form.Numbers,
		SpecialCharacters: form.Special,
	})

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data["RandomString"] = result[0]
	data["Form"] = form

	err = response.Page(w, http.StatusOK, data, "pages/home.html")
	if err != nil {
		app.serverError(w, r, err)
	}

}

func (app *application) generate(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query()) == 0 {
		data, err := parseUrlQueryWithDefaults(r)
		if err != nil {
			app.badRequest(w, r, err)
			return
		}

		result, err := randstring.RandomString(&randstring.Config{
			Count:             data.Count,
			Length:            data.Length,
			LowerCase:         data.LowerCase,
			UpperCase:         data.UpperCase,
			Numbers:           data.Numbers,
			SpecialCharacters: data.Special,
		})

		if err != nil {
			app.serverError(w, r, err)
			return
		}

		template := app.newTemplateData(r)
		template["RandomString"] = result[0]
		template["Form"] = data
		err = response.Page(w, http.StatusOK, template, "pages/home.html")
		if err != nil {
			app.serverError(w, r, err)
			return
		}

		return
	}

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

	if form.Validator.HasErrors() {
		data := app.newTemplateData(r)
		data["Form"] = form
		data["RandomString"] = "AN ERROR OCCURRED"

		err := response.Page(w, http.StatusUnprocessableEntity, data, "pages/home.html")
		if err != nil {
			app.serverError(w, r, err)
			return
		}

		return
	}

	result, err := randstring.RandomString(&randstring.Config{
		Count:             form.Count,
		Length:            form.Length,
		LowerCase:         form.LowerCase,
		UpperCase:         form.UpperCase,
		Numbers:           form.Numbers,
		SpecialCharacters: form.Special,
	})

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	template := app.newTemplateData(r)
	template["RandomString"] = result[0]
	template["Form"] = form

	err = response.Page(w, http.StatusOK, template, "pages/home.html")
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) generateBulk(w http.ResponseWriter, r *http.Request) {}

func (app *application) apiGenerate(w http.ResponseWriter, r *http.Request) {
	values, err := parseUrlQueryWithDefaults(r)

	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	result, err := randstring.RandomString(&randstring.Config{
		Count:             1,
		Length:            values.Length,
		LowerCase:         values.LowerCase,
		UpperCase:         values.UpperCase,
		Numbers:           values.Numbers,
		SpecialCharacters: values.Special,
	})

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusOK, map[string]string{
		"randomString": result[0],
	})
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) apiGenerateBulk(w http.ResponseWriter, r *http.Request) {
	values, err := parseUrlQueryWithDefaults(r)

	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	result, err := randstring.RandomString(&randstring.Config{
		Count:             values.Count,
		Length:            values.Length,
		LowerCase:         values.LowerCase,
		UpperCase:         values.UpperCase,
		Numbers:           values.Numbers,
		SpecialCharacters: values.Special,
	})

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusOK, map[string][]string{
		"randomString": result,
	})
	if err != nil {
		app.serverError(w, r, err)
	}
}

func parseUrlQueryWithDefaults(r *http.Request) (configForm, error) {
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
