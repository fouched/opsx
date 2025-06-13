package config

import (
	"database/sql"
	"github.com/alexedwards/scs/v2"
	"github.com/gorilla/schema"
	"log"
)

type App struct {
	Port     string
	DSN      string
	DB       *sql.DB
	Decoder  *schema.Decoder
	DebugLog *log.Logger
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	Session  *scs.SessionManager
}
