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
	Count             int  `form:"Count"`
	LowerCase         bool `form:"LowerCase"`
	UpperCase         bool `form:"UpperCase"`
	Numbers           bool `form:"Numbers"`
	SpecialCharacters bool `form:"SpecialCharacters"`
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	var form configForm

	form.Count = randomNumber(32, 54)
	form.LowerCase = true
	form.UpperCase = true
	form.Numbers = true
	form.SpecialCharacters = true

	result, err := randstring.RandomString(&randstring.Config{
		Count:             form.Count,
		LowerCase:         form.LowerCase,
		UpperCase:         form.UpperCase,
		Numbers:           form.Numbers,
		SpecialCharacters: form.SpecialCharacters,
	})

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data["RandomString"] = result
	data["Form"] = form

	err = response.Page(w, http.StatusOK, data, "pages/home.html")
	if err != nil {
		app.serverError(w, r, err)
	}

}

func (app *application) generate(w http.ResponseWriter, r *http.Request) {
	count, err := strconv.Atoi(r.URL.Query().Get("count"))
	lowercase := convertToBool(r.URL.Query().Get("lowercase"))
	uppercase := convertToBool(r.URL.Query().Get("uppercase"))
	numbers := convertToBool(r.URL.Query().Get("numbers"))
	symbols := convertToBool(r.URL.Query().Get("symbols"))

	fmt.Println(count, lowercase, uppercase, numbers, symbols)

	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	result, err := randstring.RandomString(&randstring.Config{
		Count:             count,
		LowerCase:         lowercase,
		UpperCase:         uppercase,
		Numbers:           numbers,
		SpecialCharacters: symbols,
	})

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	template := app.newTemplateData(r)
	template["RandomString"] = result
	template["Form"] = configForm{
		Count:             count,
		LowerCase:         lowercase,
		UpperCase:         uppercase,
		Numbers:           numbers,
		SpecialCharacters: symbols,
	}

	err = response.Page(w, http.StatusOK, template, "pages/home.html")
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) generateBulk(w http.ResponseWriter, r *http.Request) {}

func (app *application) apiGenerate(w http.ResponseWriter, r *http.Request) {
	count, err := strconv.Atoi(r.URL.Query().Get("count"))
	lowercase := convertToBool(r.URL.Query().Get("lowercase"))
	uppercase := convertToBool(r.URL.Query().Get("uppercase"))
	numbers := convertToBool(r.URL.Query().Get("numbers"))
	symbols := convertToBool(r.URL.Query().Get("symbols"))

	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	result, err := randstring.RandomString(&randstring.Config{
		Count:             count,
		LowerCase:         lowercase,
		UpperCase:         uppercase,
		Numbers:           numbers,
		SpecialCharacters: symbols,
	})

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusOK, map[string]string{
		"randomString": result,
	})
	if err != nil {
		app.serverError(w, r, err)
	}
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
