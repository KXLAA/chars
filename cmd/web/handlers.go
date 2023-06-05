package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/KXLAA/chars/pkg/randstring"
	"github.com/KXLAA/chars/pkg/response"
)

type configForm struct {
	Count             int  `form:"count"`
	Length            int  `form:"length"`
	LowerCase         bool `form:"lowerCase"`
	UpperCase         bool `form:"upperCase"`
	Numbers           bool `form:"numbers"`
	SpecialCharacters bool `form:"symbols"`
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	var form configForm

	form.Length = randomNumber(32, 54)
	form.Count = 1
	form.LowerCase = true
	form.UpperCase = true
	form.Numbers = true
	form.SpecialCharacters = true

	result, err := randstring.RandomString(&randstring.Config{
		Length:            form.Length,
		Count:             form.Count,
		LowerCase:         form.LowerCase,
		UpperCase:         form.UpperCase,
		Numbers:           form.Numbers,
		SpecialCharacters: form.SpecialCharacters,
	})

	fmt.Println(err)

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
	values, err := parseUrlQuery(r)

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
		SpecialCharacters: values.SpecialCharacters,
	})

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	template := app.newTemplateData(r)
	template["RandomString"] = result[0]
	template["Form"] = configForm{
		Count:             1,
		Length:            values.Length,
		LowerCase:         values.LowerCase,
		UpperCase:         values.UpperCase,
		Numbers:           values.Numbers,
		SpecialCharacters: values.SpecialCharacters,
	}

	err = response.Page(w, http.StatusOK, template, "pages/home.html")
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) generateBulk(w http.ResponseWriter, r *http.Request) {}

func (app *application) apiGenerate(w http.ResponseWriter, r *http.Request) {
	values, err := parseUrlQuery(r)

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
		SpecialCharacters: values.SpecialCharacters,
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
	values, err := parseUrlQuery(r)

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
		SpecialCharacters: values.SpecialCharacters,
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

func parseUrlQuery(r *http.Request) (configForm, error) {
	length := resolveIntQuery(r, "length", 32)
	count := resolveIntQuery(r, "count", 1)
	lowercase := convertToBool(r.URL.Query().Get("lowercase"))
	uppercase := convertToBool(r.URL.Query().Get("uppercase"))
	numbers := convertToBool(r.URL.Query().Get("numbers"))
	special := convertToBool(r.URL.Query().Get("special"))

	return configForm{
		Count:             count,
		Length:            length,
		LowerCase:         lowercase,
		UpperCase:         uppercase,
		Numbers:           numbers,
		SpecialCharacters: special,
	}, nil
}

func resolveIntQuery(r *http.Request, key string, defaultValue int) int {
	value, err := strconv.Atoi(r.URL.Query().Get(key))
	if err != nil {
		return defaultValue
	}
	return value
}

func convertToBool(value string) bool {
	if value == "on" {
		return true
	}
	return value == "true"
}

func randomNumber(min, max int) int {
	return min + rand.Intn(max-min)
}
