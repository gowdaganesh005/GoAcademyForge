package main

import "net/http"

func (app *application) serverError(w http.ResponseWriter, err error) {
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)

}
func (app *application) notfound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
