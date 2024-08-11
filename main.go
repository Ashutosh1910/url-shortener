package main

import (
	"fmt"
	//"html/template"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/Ashutosh1910/url-shortner/models"
	shortner "github.com/Ashutosh1910/url-shortner/service"
	///"github.com/Ashutosh1910/url-shortner/templates"
)

func main() {

	router := shortner.NewRouter()
	db, err := sql.Open("sqlite3", "./urls.db")
	//db.SetMaxOpenConns(1)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router.Service = &shortner.ShortnerService{Db: db}
	models.MigrateUrlPair(db)
	fmt.Println("Server is Running")
	log.Fatalln(http.ListenAndServe(":8000", router))

}
