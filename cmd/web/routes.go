package main

import (
	"net/http"
)

func (app *application) routes(prefixPath string) *http.ServeMux {
	mux := http.NewServeMux()
	// Web pages for each appliance
	mux.HandleFunc(prefixPath, app.home)
	mux.HandleFunc(prefixPath+"horno", app.horno)
	mux.HandleFunc(prefixPath+"heladera", app.heladera)
	mux.HandleFunc(prefixPath+"heladera-doble-1", app.heladeraDoble1)
	mux.HandleFunc(prefixPath+"heladera-doble-2", app.heladeraDoble2)

	// API endpoints for each appliance
	mux.HandleFunc(prefixPath+"v1/{id}/settemp", app.setTemp)
	mux.HandleFunc(prefixPath+"v1/{id}/updatetemp", app.updateTemp)
	mux.HandleFunc(prefixPath+"v1/{id}/gettemp", app.getTemp)
	mux.HandleFunc(prefixPath+"status", app.status)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle(prefixPath+"static/", http.StripPrefix(prefixPath+"static/", fileServer))

	return mux
}
