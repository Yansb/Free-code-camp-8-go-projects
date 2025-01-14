package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Yansb/Go-study-projects/eshop/internal/driver"
	"github.com/Yansb/Go-study-projects/eshop/internal/models"
	"github.com/alexedwards/scs/v2"
)

const version = "1.0.0"
const cssVersion = "1"

var session *scs.SessionManager

type config struct {
	port int
	env  string
	api  string

	db struct {
		dsn string
	}

	stripe struct {
		secret string
		key    string
	}
}

type application struct {
	config        config
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache map[string]*template.Template
	version       string
	DB            models.DBModel
	Session       *scs.SessionManager
}

func (app *application) server() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	app.infoLog.Printf("Starting server in %s on port %d", app.config.env, app.config.port)

	return srv.ListenAndServe()
}

func main() {
	gob.Register(TransactionData{})
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment (development|production)")
	flag.StringVar(&cfg.api, "api", "http://localhost:4001", "URL to API")
	flag.StringVar(&cfg.db.dsn, "dsn", "user:password@tcp(localhost:3306)/widgets?parseTime=true&tls=false", "Connection to DB string")

	flag.Parse()

	cfg.stripe.key = "pk_test_sYfSqQMRBXLAzfFcoszflMuj0028QEtbKZ"
	cfg.stripe.secret = os.Getenv("STRIPE_SECRET")

	infoLog := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)

	conn, err := driver.OpenDB(cfg.db.dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer conn.Close()

	session = scs.New()
	session.Lifetime = 24 * time.Hour

	tc := make(map[string]*template.Template)

	app := &application{
		config:        cfg,
		infoLog:       infoLog,
		errorLog:      errorLog,
		templateCache: tc,
		version:       version,
		DB:            models.DBModel{DB: conn},
		Session:       session,
	}

	err = app.server()

	if err != nil {
		app.errorLog.Println(err)
		log.Fatal(err)
	}
}
