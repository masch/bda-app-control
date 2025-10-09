package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var Appliances = make(map[string]bool)
var DbHost = os.Getenv("DB_HOST")
var DbName = os.Getenv("DB_NAME")
var DbUsername = os.Getenv("DB_USERNAME")
var DbPassword = os.Getenv("DB_PASSWORD")

var ctx = context.Background()
var dsn = fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", DbUsername, DbPassword, DbHost, DbName)
var pgdb = sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
var db = bun.NewDB(pgdb, pgdialect.New())

// Web Functions
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) methodNotAllowed(w http.ResponseWriter) {
	app.clientError(w, http.StatusMethodNotAllowed)
}

func (app *application) badRequest(w http.ResponseWriter) {
	app.clientError(w, http.StatusBadRequest)
}

// Create schema in database
func (app *application) createSchema() {
	_, err := db.NewRaw("CREATE SCHEMA IF NOT EXISTS bosque").Exec(ctx)
	if err != nil {
		app.errorLog.Fatal(err)
	}
	app.infoLog.Println("Schema 'bosque' created or already exists.")
}

// Check for table name in database
func (app *application) checkTable(table string) bool {
	exist, err := db.NewSelect().
		Table("information_schema.tables").
		Where("table_schema = 'bosque' AND table_name = ?", table).
		Exists(ctx)
	if err != nil {
		println(err.Error())
	}

	if !exist {
		println("Table " + table + " does not exist in DB, creating...")
		return false
	}
	return true
}

// Create table in database
func (app *application) createTable(basemodel interface{}) {
	_, create := db.NewCreateTable().Model(basemodel).IfNotExists().Exec(ctx)

	if create != nil {
		panic(create)
	}
	log.Println("Created table")
}

// Fill table with model data
func (app *application) fillTables(model []Appliance) {
	res, err := db.NewInsert().Model(&model).Exec(ctx)

	if err != nil {
		panic(err)
	}
	log.Println("Inserted data, transaction: ", res)
}

// Log actual temperature into database
func (app *application) logTemp(appliance string, temp float64) {
	templog := TempLog{Appliance: appliance, ActualTemp: temp, Time: time.Now()}
	res, err := db.NewUpdate().Model(&templog).WherePK(appliance).Exec(ctx)

	if err != nil {
		panic(err)
	}
	log.Println(res)
}

// Read temperature from Database
func (app *application) readTempDB(appliance string, temptype string) float64 {
	var temp float64

	err := db.NewSelect().
		ColumnExpr("?", bun.Safe(temptype)).
		Table("bosque.appliances").
		Where("appliance = ? ", appliance).
		Scan(ctx, &temp)

	if err != nil {
		panic(err)
	}

	return temp
}

// Write temperature to Database
func (app *application) updateTempDB(w http.ResponseWriter, appliance string, temptype string, temp float64) {
	var newdata Appliance

	switch temptype {
	case "actualtemp":
		newdata = Appliance{Appliance: appliance, ActualTemp: temp}
		break
	case "settemp":
		newdata = Appliance{Appliance: appliance, SetTemp: temp}
		break
	default:
		app.clientError(w, http.StatusBadRequest)
		return
	}

	update, err := db.NewUpdate().Model(&newdata).
		Column(temptype).
		WherePK("appliance").
		Exec(ctx)

	if err != nil {
		panic(err)
	}
	app.infoLog.Println(update)

	log.Printf("%s Temp. updated on %s \n", temptype, appliance)
}

// Read appliance list from Database
func (app *application) readAppliances() []string {
	var m []string
	err := db.NewSelect().
		ColumnExpr("appliances.appliance").
		Table("bosque.appliances").
		Scan(ctx, &m)

	if err != nil {
		panic(err)
	}

	return m
}
