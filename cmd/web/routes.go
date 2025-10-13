package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	// Web pages for each appliance
	mux.HandleFunc("/bosquesdeagua", app.home)
	mux.HandleFunc("/bosquesdeagua/horno", app.horno)
	mux.HandleFunc("/bosquesdeagua/heladera", app.heladera)
	mux.HandleFunc("/bosquesdeagua/heladera-doble-1", app.heladeraDoble1)
	mux.HandleFunc("/bosquesdeagua/heladera-doble-2", app.heladeraDoble2)
	
	// API endpoints for each appliance
	mux.HandleFunc("/bosquesdeagua/v1/{id}/settemp", app.setTemp)
	mux.HandleFunc("/bosquesdeagua/v1/{id}/updatetemp", app.updateTemp)
	mux.HandleFunc("/bosquesdeagua/v1/{id}/gettemp", app.getTemp)
	mux.HandleFunc("/bosquesdeagua/status", app.status)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
