package main

// Backinfo testing
type BankInfo struct {
	Bankleitzahl          string  `json:"bankleitzahl"`
	Merkmal               float64 `json:"merkmal"`
	Bezeichnung           string  `json:"bezeichnung"`
	Plz                   string  `json:"plz"`
	Ort                   string  `json:"ort"`
	Kurzbezeichnung       string  `json:"kurzbezeichnung"`
	Pan                   string  `json:"pan"`
	Bic                   string  `json:"bic"`
	PruefzifferMethode    string  `json:"pruefzifferMethode"`
	Datensatznummer       string  `json:"datensatznummer"`
	Aenderungskennzeichen string  `json:"aenderungskennzeichen"`
	Bankleitzahlloeschung string  `json:"bankleitzahlloeschung"`
	Nachfolgebankleitzahl string  `json:"nachfolgebankleitzahl"`
}
