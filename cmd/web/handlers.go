package main

import (
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
	//handle GET request to /generate with no query params
	if len(r.URL.Query()) == 0 {
		data, err := app.parseUrlQueryWithDefaults(r)
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

func (app *application) apiGenerate(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query()) == 0 {
		values, err := app.parseUrlQueryWithDefaults(r)

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

		err = response.JSON(w, http.StatusOK, map[string]string{
			"randomString": result[0],
		})

		if err != nil {
			app.serverError(w, r, err)
			return
		}

		return
	}

	form := configForm{}
	length, err := strconv.Atoi(r.URL.Query().Get("length"))

	if err != nil {
		err = response.JSON(w, http.StatusBadRequest, map[string]string{
			"error": ErrInvalidQueryParams.Error(),
		})

		if err != nil {
			app.badRequest(w, r, err)
			return
		}

		return
	}

	var count int
	if r.URL.Query().Get("count") == "" {
		count = 1
	} else {
		count, err = strconv.Atoi(r.URL.Query().Get("count"))
		if err != nil {
			err = response.JSON(w, http.StatusBadRequest, map[string]string{
				"error": ErrInvalidQueryParams.Error(),
			})

			if err != nil {
				app.badRequest(w, r, err)
				return
			}

			return
		}
	}

	form.Length = length
	form.Count = count
	form.LowerCase = resolveBoolQuery(r, "lowercase")
	form.UpperCase = resolveBoolQuery(r, "uppercase")
	form.Numbers = resolveBoolQuery(r, "numbers")
	form.Special = resolveBoolQuery(r, "special")

	if !form.LowerCase && !form.UpperCase && !form.Numbers && !form.Special {
		err = response.JSON(w, http.StatusBadRequest, map[string]string{
			"error": ErrIncompleteQueryParams.Error(),
		})

		if err != nil {
			app.badRequest(w, r, err)
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

	if form.Count > 1 {
		err = response.JSON(w, http.StatusOK, map[string][]string{
			"randomString": result,
		})
	} else {
		err = response.JSON(w, http.StatusOK, map[string]string{
			"randomString": result[0],
		})
	}

	if err != nil {
		app.serverError(w, r, err)
		return
	}
}
