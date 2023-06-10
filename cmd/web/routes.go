package main

import (
	"net/http"

	"github.com/KXLAA/chars/assets"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mw := alice.New(app.recoverPanic, app.logRequest, app.securityHeaders)
	mux := httprouter.New()
	mux.NotFound = http.HandlerFunc(app.notFound)

	fileServer := http.FileServer(http.FS(assets.EmbeddedFiles))
	mux.Handler("GET", "/static/*filepath", fileServer)
	mux.HandlerFunc(http.MethodGet, "/", app.home)
	mux.HandlerFunc(http.MethodGet, "/generate", app.generate)
	mux.HandlerFunc(http.MethodGet, "/api/v1/generate", app.apiGenerate)

	return mw.Then(mux)
}
