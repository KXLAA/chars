package main

import (
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
)

var (
	ErrIncompleteQueryParams = errors.New("looks like your query params are incomplete. Please make sure lowercase, uppercase, numbers, or special are set to true in your query params")
	ErrInvalidQueryParams    = errors.New("looks like your query params are invalid. Please make sure the length and count query params are numbers")
)

func (app *application) reportError(err error) {
	trace := debug.Stack()
	app.logger.Error(err, trace)

}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	app.reportError(err)
	fmt.Println("server error", err.Error())
	message := "The server encountered a problem and could not process your request"
	http.Error(w, message, http.StatusInternalServerError)
}

func (app *application) notFound(w http.ResponseWriter, r *http.Request) {
	message := "The requested resource could not be found"
	http.Error(w, message, http.StatusNotFound)
}

func (app *application) badRequest(w http.ResponseWriter, r *http.Request, err error) {
	app.reportError(err)
	http.Error(w, err.Error(), http.StatusBadRequest)
}
