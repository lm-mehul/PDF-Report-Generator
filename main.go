package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"

	u "github.com/c-seeger/Golang-HTML-TO-PDF-Converter"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var tpl *template.Template
var err error

type Data struct {
	Player   string
	Channel  string
	Playout  string
	Duration string
}

type Page struct {
	Network_name     string
	Generated_on     string
	Date_range       string
	Report_type      string
	Total_players    string
	Media_items      string
	Total_duration   string
	Total_impression string
	Table_data       []Data
}

type Report struct {
	Pages []Page
}

func displayerror(e error) {
	if e != nil {
		panic(e.Error())
	}
}

func main() {

	db, err = sql.Open("mysql", "lemma:admin@tcp(localhost:3306)/PDF")
	displayerror(err)
	defer db.Close()
	err = db.Ping()
	displayerror(err)
	fmt.Println("Successful Connection to Database!")

	stmt := "SELECT * FROM PDF.data ;"

	rows, err := db.Query(stmt)
	displayerror(err)
	defer rows.Close()

	var data []Data

	for rows.Next() {
		var p Data

		err = rows.Scan(&p.Player, &p.Channel, &p.Playout, &p.Duration)
		displayerror(err)
		data = append(data, p)
	}

	no_of_rows := 30
	array_size := len(data)

	pages := int(array_size / no_of_rows)

	var pdf []Page
	var p Page

	for i := 0; i < pages; i++ {
		var temp []Data
		for j := 0; j < no_of_rows; j++ {
			temp = append(temp, data[i*no_of_rows+j])
		}
		p = Page{
			Network_name:     "Test",
			Generated_on:     "33538957933",
			Date_range:       "453535435-34535355345",
			Report_type:      "jnjsnfsnfjksnknjknjdsjkvndskvndfkjvjndjvbkvsd",
			Total_players:    "4323242",
			Media_items:      "fefgmsgmksdmnvjknsvjkvnjkvnrjkvnrjkvnkvjnvnnkdjn",
			Total_duration:   "23423424324",
			Total_impression: "3242422344234242422243",
			Table_data:       temp}
		pdf = append(pdf, p)
	}
	if array_size%no_of_rows > 0 {
		var temp []Data
		for j := 0; j < array_size%no_of_rows; j++ {
			temp = append(temp, data[pages*no_of_rows+j])
		}
		p = Page{
			Network_name:     "Test",
			Generated_on:     "33538957933",
			Date_range:       "453535435-34535355345",
			Report_type:      "jnjsnfsnfjksnknjknjdsjkvndskvndfkjvjndjvbkvsd",
			Total_players:    "4323242",
			Media_items:      "fefgmsgmksdmnvjfefgmsgmksdmnvjknsvjkvnjkvnrjkvnrjkvnkvjnvnnkdjn",
			Total_duration:   "23423424324",
			Total_impression: "3242422344234242422243",
			Table_data:       temp}

		pdf = append(pdf, p)
	}

	report := Report{
		Pages: pdf,
	}

	new_pdf := u.NewRequestPdf("")

	templatePath := "report.html"

	outputPath := "example.pdf"

	if err := new_pdf.ParseTemplate(templatePath, report); err != nil {
		log.Fatal(err)
	}
	if err := new_pdf.GeneratePDF(outputPath); err != nil {
		log.Fatal(err)
	}
	fmt.Println("pdf generated successfully")

}
