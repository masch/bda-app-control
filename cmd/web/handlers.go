package main

import (
	"fmt"
	"html/template"
	"net/http"
	"reflect"
	"strconv"
)

// #GET# Server alive
func (app *application) status(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/bosquesdeagua/status" {
		w.WriteHeader(http.StatusOK)
	}
	return
}

// #PAGINA# Homepage
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/bosquesdeagua" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partials.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

// #PAGINA# Horno
func (app *application) horno(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/bosquesdeagua/horno" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/horno.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partials.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	var appliance = new(Appliance)
	appliance.ActualTemp = app.readTempDB("horno", "actualtemp")
	appliance.SetTemp = app.readTempDB("horno", "settemp")

	err = ts.Execute(w, appliance)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

// #PAGINA# Heladera simple
func (app *application) heladera(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/bosquesdeagua/heladera" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/heladera.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partials.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	var appliance = new(Appliance)
	appliance.ActualTemp = app.readTempDB("heladera", "actualtemp")
	appliance.SetTemp = app.readTempDB("heladera", "settemp")

	err = ts.Execute(w, appliance)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

// #PAGINA# Heladera Doble Superior
func (app *application) heladeraDoble1(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/bosquesdeagua/heladera-doble-1" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/heladera-doble-1.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partials.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	var appliance = new(Appliance)
	appliance.ActualTemp = app.readTempDB("heladera-doble-1", "actualtemp")
	appliance.SetTemp = app.readTempDB("heladera-doble-2", "settemp")

	err = ts.Execute(w, appliance)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

// #PAGINA# Heladera Doble Inferior
func (app *application) heladeraDoble2(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/bosquesdeagua/heladera-doble-2" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/heladera-doble-2.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partials.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	var appliance = new(Appliance)
	appliance.ActualTemp = app.readTempDB("heladera-doble-2", "actualtemp")
	appliance.SetTemp = app.readTempDB("heladera-doble-2", "settemp")

	err = ts.Execute(w, appliance)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

// #POST# Sets new setTemp on server (HMI)
func (app *application) setTemp(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		app.clientError(w, 405)
		return
	}

	appliance := r.PathValue("id")
	requestTemp := r.Header.Get("temp")
	if requestTemp != "" {
		fmt.Printf("New temp sent via API is: %s ", requestTemp)
	} else {
		app.clientError(w, 400)
		return
	}
	if !Appliances[appliance] {
		app.notFound(w)
		return
	}

	newTemp, err := strconv.ParseFloat(requestTemp, 32)
	if err != nil {
		app.infoLog.Println(err)
	}

	app.updateTempDB(w, appliance, "settemp", newTemp)

	fmt.Printf("Appliance - %s - has a new set temp of: %f", appliance, newTemp)

}

// #POST# Sets actual appliance temp to server (Appliance)
func (app *application) updateTemp(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		app.clientError(w, 405)
		return
	}

	appliance := r.PathValue("id")
	temp := r.Header.Get("temp")
	if temp == "" {
		app.clientError(w, 400)
		return
	}
	if !Appliances[appliance] {
		app.notFound(w)
		return
	}

	newTemp, err := strconv.ParseFloat(temp, 32)
	if err != nil {
		app.infoLog.Println(err)
	}

	app.updateTempDB(w, appliance, "actualtemp", newTemp)

	fmt.Printf("Appliance - %s - has reported a temp of: %f \n", appliance, newTemp)

}

// #GET# Gets set temp from server (Appliance)
func (app *application) getTemp(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		app.clientError(w, 405)
		return
	}

	appliance := r.PathValue("id")
	if !Appliances[appliance] {
		app.notFound(w)
		return
	}

	var temp = app.readTempDB(appliance, "settemp")
	if reflect.TypeOf(temp) != reflect.TypeOf(float64(0)) {
		app.serverError(w, fmt.Errorf("Expected to read float64 but got %v", reflect.TypeOf(temp)))
	}
	fmt.Fprintf(w, "%.2f", temp)

}
