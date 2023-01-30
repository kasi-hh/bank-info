package main

import (
	"database/sql"
	"fmt"
	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/go-chi/cors"
	_ "github.com/go-sql-driver/mysql"

	"log"
)

const (
	DBPORT = "7090"      // Port of the database
	DBHOST = "localhost" // Host of the database
	DBUSER = "root"      // User of the database
	DBPASS = "banken"    // Password of the database
)

func main() {

	mux := chi.NewRouter()

	// specify who is allowed to connect
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Use(middleware.Heartbeat("/ping"))

	dbConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/banken", DBUSER, DBPASS, DBHOST, DBPORT)
	fmt.Println("dbConn: ", dbConn)
	db, err := sql.Open("mysql", dbConn)

	// if there is an error opening the dbConn, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT * FROM bankinfo where bankleitzahl = ?", "10000000")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var bankInfo BankInfo
		// for each row, scan the result into our tag composite object

		err = results.Scan(&bankInfo.Bankleitzahl, &bankInfo.Merkmal, &bankInfo.Bezeichnung, &bankInfo.Plz, &bankInfo.Ort, &bankInfo.Kurzbezeichnung, &bankInfo.Pan, &bankInfo.Bic, &bankInfo.PruefzifferMethode, &bankInfo.Datensatznummer, &bankInfo.Aenderungskennzeichen, &bankInfo.Bankleitzahlloeschung, &bankInfo.Nachfolgebankleitzahl)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Println(bankInfo.Bankleitzahl, bankInfo.Bezeichnung)
	}
}
