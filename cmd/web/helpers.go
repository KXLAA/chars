package main

import (
	"net/http"

	"github.com/KXLAA/chars/pkg/version"
)

func (app *application) newTemplateData(r *http.Request) map[string]any {
	data := map[string]any{
		"Version": version.Get(),
	}

	return data
}
