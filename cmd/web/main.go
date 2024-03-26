package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"nebil/golang/internal/models"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *models.SnippetModel
}

func main() {

	port := os.Getenv("PORT")

	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQl data source name")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)

	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	app := application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &models.SnippetModel{DB: db},
	}

	infoLog.Printf("Listen and Serve in port :%v", port)
	// err := http.ListenAndServe(":"+port, mux)
	srv := &http.Server{
		Addr:     port,
		Handler:  app.routes(),
		ErrorLog: errorLog}
	err = srv.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	fmt.Println(dsn)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
