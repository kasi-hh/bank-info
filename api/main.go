package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"golang.org/x/text/encoding/charmap"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Config = struct {
	DbHost string
	DbPort string
	DbUser string
	DbPass string
	DbName string
	Port   string
}

var Cfg = Config{}
var bankInfos = map[string]BankInfo{}

func init() {
	// Cfg.DbHost = getEnv("DBHOST", "localhost")
	// Cfg.DbPort = getEnv("DBPORT", "7090")
	// Cfg.DbUser = getEnv("DBUSER", "root")
	// Cfg.DbPass = getEnv("DBPASS", "banken")
	// Cfg.DbName = getEnv("DBNAME", "banken")
	Cfg.Port = getEnv("PORT", "9060")
	loadData()
}
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
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
	mux.Get("/{bankleitzahl}", getBankInfo)
	log.Fatal(http.ListenAndServe(":"+Cfg.Port, mux))
}

// Load data from File
func loadData() {
	file, err := os.Open("./bankleitzahlen.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 150 {
			bankInfo := parseLine(line)
			bankInfos[bankInfo.Bankleitzahl] = bankInfo
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
func parseLine(line string) BankInfo {

	var bankInfo BankInfo
	merkmal, _ := strconv.ParseFloat(line[8:9], 64)
	bankInfo.Bankleitzahl = line[0:8]
	bankInfo.Merkmal = merkmal
	bankInfo.Bezeichnung = decodeISO8859_1(line[9:67])
	bankInfo.Plz = line[67:72]
	bankInfo.Ort = decodeISO8859_1(line[72:107])
	bankInfo.Kurzbezeichnung = decodeISO8859_1(line[107:134])
	bankInfo.Bic = line[139:150]
	bankInfo.Pan = line[134:139]
	fmt.Println(bankInfo)
	return bankInfo
}

func decodeISO8859_1(s string) string {
	dec := charmap.ISO8859_1.NewDecoder()
	out, _ := dec.String(s)
	return strings.Trim(out, " ")
}
func getBankInfo(w http.ResponseWriter, r *http.Request) {
	bankleitzahl := chi.URLParam(r, "bankleitzahl")
	if len(bankleitzahl) > 12 {
		bankleitzahl = bankleitzahl[4:12]
	}
	//dbConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", Cfg.DbUser, Cfg.DbPass, Cfg.DbHost, Cfg.DbPort, Cfg.DbName)
	//fmt.Println("dbConn: ", dbConn)
	//db, err := sql.Open("mysql", dbConn)

	// if there is an error opening the dbConn, handle it
	// if err != nil {
	//	panic(err.Error())
	//}

	// defer the close till after the main function has finished
	// executing
	// defer db.Close()

	// Execute the query
	// results, err := db.Query("SELECT * FROM bankinfo where bankleitzahl = ?", bankleitzahl)
	// if err != nil {
	// 	panic(err.Error()) // proper error handling instead of panic in your app
	// }
	// defer results.Close()
	// for results.Next() {
	//	var bankInfo BankInfo
	// for each row, scan the result into our tag composite object

	// 	err = results.Scan(&bankInfo.Bankleitzahl, &bankInfo.Merkmal, &bankInfo.Bezeichnung, &bankInfo.Plz, &bankInfo.Ort, &bankInfo.Kurzbezeichnung, &bankInfo.Pan, &bankInfo.Bic, &bankInfo.PruefzifferMethode, &bankInfo.Datensatznummer, &bankInfo.Aenderungskennzeichen, &bankInfo.Bankleitzahlloeschung, &bankInfo.Nachfolgebankleitzahl)
	// 	if err != nil {
	// 		panic(err.Error()) // proper error handling instead of panic in your app
	//	}
	// and then print out the tag's Name attribute
	// 	log.Println(bankInfo.Bankleitzahl, bankInfo.Bezeichnung, bankInfo)
	bankInfo, ok := bankInfos[bankleitzahl]
	if !ok {
		log.Println("Bankleitzahl not found: ", bankleitzahl)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	result, err := json.Marshal(bankInfo)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println(err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(result))
	log.Println("Bankleitzahl found: ", bankleitzahl, bankInfo)
	// }
}
