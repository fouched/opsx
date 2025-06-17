package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/fouched/opsx/internal/config"
	"github.com/fouched/opsx/internal/driver"
	"github.com/fouched/opsx/internal/handlers"
	"github.com/fouched/opsx/internal/render"
	"github.com/fouched/opsx/internal/repo"
	"github.com/gorilla/schema"
	"github.com/pelletier/go-toml/v2"
	"log"
	"net/http"
	"os"
	"time"
)

// session - must be available in main package for middleware
var session *scs.SessionManager

func main() {
	app, err := initApp()
	if err != nil {
		log.Fatal(err)
	}

	// we have database connectivity, close it after app stops
	defer app.DB.Close()

	// Create handlers instance with dependency
	h := handlers.NewHandlers(app)

	mux := routes(h)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.Port),
		Handler: mux,
	}

	app.InfoLog.Println("Starting server on", app.Port)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func initApp() (*config.App, error) {
	app := &config.App{}

	setConfig(app)
	initLoggers(app)
	initFormDecoder(app)

	err := initDatabase(app)
	if err != nil {
		return nil, err
	}

	app.Repo = config.Repo{
		RemoteAppRepo: repo.NewRemoteAppRepo(app.DB),
	}

	session = initSession()
	app.Session = session

	// set up template rendering
	render.NewRenderer(app)

	return app, nil
}

func setConfig(app *config.App) {
	env := os.Getenv("OPSX_ENV")
	if env == "" {
		env = "development"
	}

	configFile := "config." + env + ".toml"
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		configFile = "config.toml" // Fallback to default
	}

	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	if err := toml.Unmarshal(data, app); err != nil {
		log.Fatalf("Error parsing TOML: %v", err)
	}

}

func initLoggers(app *config.App) {
	app.DebugLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.ErrorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
}

func initFormDecoder(app *config.App) {
	app.Decoder = schema.NewDecoder()
}

func initDatabase(app *config.App) error {
	dbPool, err := driver.ConnectSQL(app.DSN)
	if err != nil {
		return err
	}
	app.InfoLog.Println("Connected to database")
	app.DB = dbPool
	return nil
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
