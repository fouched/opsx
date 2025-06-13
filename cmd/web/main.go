package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/fouched/opsx/internal/config"
	"github.com/fouched/opsx/internal/handlers"
	"github.com/fouched/opsx/internal/render"
	"github.com/gorilla/schema"
	"log"
	"net/http"
	"os"
	"time"
)

// session - must be available in main package for middleware
var session *scs.SessionManager

func main() {
	app, err := newApp()
	if err != nil {
		log.Fatal(err)
	}

	// Create handlers instance with dependency
	h := handlers.NewHandlers(app)

	mux := routes(h)
	srv := &http.Server{
		Addr:    app.Port,
		Handler: mux,
	}

	app.InfoLog.Println("Starting server on", app.Port)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func newApp() (*config.App, error) {

	session = initSession()

	debugLog, infoLog, errorLog := initLogger()

	app := &config.App{
		Port:     ":9080",
		Decoder:  initDecoder(),
		DebugLog: debugLog,
		InfoLog:  infoLog,
		ErrorLog: errorLog,
		Session:  session,
	}

	// set up template rendering
	render.NewRenderer(app)

	return app, nil
}

func initDecoder() *schema.Decoder {
	return schema.NewDecoder()
}

func initLogger() (*log.Logger, *log.Logger, *log.Logger) {
	debugLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	return debugLog, infoLog, errorLog
}

func initSession() *scs.SessionManager {
	s := scs.New()
	s.Lifetime = 24 * time.Hour
	s.Cookie.Persist = true
	s.Cookie.SameSite = http.SameSiteLaxMode

	//we can use persistent storage iso cookies for session data, this allows us to
	//restart the server without users losing the login / session information
	//https://github.com/alexedwards/scs has various options available e.g.
	//session.Store = pgxstore.New(db)
	return s
}
