package main

import (
	"github.com/uptrace/bun"
	"log"
	"time"
)

type Appliance struct {
	bun.BaseModel `bun:"table:bosque.appliances"`

	Appliance  string  `bun:"appliance,pk,notnull" json:"appliance"`
	SetTemp    float64 `bun:"settemp,notnull" json:"settemp"`
	ActualTemp float64 `bun:"actualtemp,notnull" json:"actualtemp"`
}

type ApplianceList struct {
	Appliances []Appliance `json:"appliances"`
}
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

type TempLog struct {
	bun.BaseModel `bun:"table:bosque.templogs"`

	ID         uint      `bun:"id,pk,autoincrement"`
	Appliance  string    `bun:"appliance,notnull"`
	ActualTemp float64   `bun:"actualtemp,notnull"`
	Time       time.Time `bun:"time,nullzero,default:now()"`
}
