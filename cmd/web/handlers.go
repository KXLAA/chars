package main

import (
	"net/http"

	"github.com/KXLAA/chars/pkg/forms"
	"github.com/KXLAA/chars/pkg/randstring"
	"github.com/KXLAA/chars/pkg/render"
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

	form.Count = 32
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

	err = render.Page(w, http.StatusOK, data, "pages/home.html")
	if err != nil {
		app.serverError(w, r, err)
	}

}

func (app *application) generate(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	var form configForm

	err := forms.DecodePostForm(r, &form)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

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

	data["Generated"] = result
	data["Form"] = form

	err = render.Page(w, http.StatusOK, data, "pages/generate.html")
	if err != nil {
		app.serverError(w, r, err)
	}
}
