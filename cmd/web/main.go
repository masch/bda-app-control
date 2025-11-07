package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/uptrace/bun/extra/bundebug"
)

func main() {

	var env = godotenv.Load(".env")
	if env != nil {
		log.Fatal("Error loading .env file")
	}

	var AppVersion = os.Getenv("VERSION")

	cfg := new(Config)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	data, err := os.ReadFile("./cmd/db/appliances.json")
	var extracted ApplianceList
	err = json.Unmarshal(data, &extracted)
	if err != nil {
		errorLog.Fatal(err)
	}
	appliances := extracted.Appliances

	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	//Webserver Flags
	flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP Network address")
	flag.StringVar(&cfg.BasePath, "base", "", "Base path for the application")
	flag.StringVar(&cfg.StaticDir, "static", "./static", "Path to static files")
	flag.Parse()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		config:   cfg,
	}

	srv := &http.Server{
		Addr:     cfg.Addr,
		ErrorLog: errorLog,
		Handler:  app.routes(cfg.BasePath),
	}

	app.createSchema()

	println("Checking for tables")
	//Check and create tables in DB if not present
	existTableAppliances := app.checkTable("appliances")
	existTableLogs := app.checkTable("templogs")

	if !existTableAppliances {
		app.createTable(&Appliance{})
	}
	if !existTableLogs {
		app.createTable(&TempLog{})
	}

	//Check and fill appliances table with defaults if not present
	applianceList := app.readAppliances()
	if applianceList == nil {
		app.fillTables(appliances)
	}

	//Initialize and read Appliances from DB
	applianceList = app.readAppliances()
	for _, appliance := range applianceList {
		Appliances[appliance] = true
	}

	//WebServer initialization
	infoLog.Printf("\n - App de Control Ver. %s -\n", AppVersion)
	infoLog.Printf("Starting server on localhost: http://%s", cfg.Addr+cfg.BasePath)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)

}
