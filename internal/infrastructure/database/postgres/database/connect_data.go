package database

import (
	"fmt"
)

type connectData struct {
	host     string
	port     string
	user     string
	password string
	dbName   string
}

func getConnectData() *connectData {
	connData := &connectData{
		//host:     os.Getenv("DB_HOST"),
		//port:     os.Getenv("DB_PORT"),
		//user:     os.Getenv("DB_USER"),
		//password: os.Getenv("DB_PASS"),
		//dbName:   os.Getenv("DB_NAME"),
		host:     "localhost",
		port:     "5432",
		user:     "postgres",
		password: "1111",
		dbName:   "WB",
	}
	fmt.Printf("Connect Data========================: %+v\n", connData) // Отладочная печать
	return connData
}

//DB_HOST=localhost
//DB_PORT=5432
//DB_USER=postgres
//DB_PASS=1111
//DB_NAME=WB
