package render

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/KXLAA/chars/assets"
	"github.com/KXLAA/chars/pkg/funcs"
)

func Page(w http.ResponseWriter, status int, data any, pagePath string) error {
	patterns := []string{"base.html", "partials/*.html", pagePath}
	templateName := "base"

	for i := range patterns {
		patterns[i] = "templates/" + patterns[i]
	}

	ts, err := template.New("").Funcs(funcs.TemplateFuncs).ParseFS(assets.EmbeddedFiles, patterns...)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)

	err = ts.ExecuteTemplate(buf, templateName, data)
	if err != nil {
		return err
	}

	w.WriteHeader(status)
	buf.WriteTo(w)

	return nil
}
