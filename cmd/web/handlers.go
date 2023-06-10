package main

import (
	"net/http"

	rnd "github.com/KXLAA/chars/pkg/randstring"
	"github.com/KXLAA/chars/pkg/response"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	form := &configForm{
		Length:    randomNumber(32, 54),
		Count:     1,
		LowerCase: true,
		UpperCase: true,
		Numbers:   true,
		Special:   true,
	}

	result, err := rnd.RandomString(&rnd.Config{
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
	//handle get requests to /generate ie no query params
	if len(r.URL.Query()) == 0 {
		data, err := app.parseUrlQueriesWithDefaults(r)
		if err != nil {
			app.badRequest(w, r, err)
			return
		}

		result, err := rnd.RandomString(&rnd.Config{
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

	form := app.parseForm(w, r)
	data := app.newTemplateData(r)

	if form.Validator.HasErrors() {
		data["Form"] = form
		data["RandomString"] = "AN ERROR OCCURRED"

		err := response.Page(w, http.StatusUnprocessableEntity, data, "pages/home.html")
		if err != nil {
			app.serverError(w, r, err)
			return
		}

		return
	}

	result, err := rnd.RandomString(&rnd.Config{
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
	// handle GET requests to /api/v1/generate
	if len(r.URL.Query()) == 0 {
		values, err := app.parseUrlQueriesWithDefaults(r)
		if err != nil {
			app.badRequest(w, r, err)
			return
		}

		result, err := rnd.RandomString(&rnd.Config{
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

	length, err := parseLengthValue(w, r)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}
	count, err := parseCount(w, r)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	form := configForm{
		Length:    length,
		Count:     count,
		LowerCase: resolveBoolQuery(r, "lowercase"),
		UpperCase: resolveBoolQuery(r, "uppercase"),
		Numbers:   resolveBoolQuery(r, "numbers"),
		Special:   resolveBoolQuery(r, "special"),
	}

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

	result, err := rnd.RandomString(&rnd.Config{
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
