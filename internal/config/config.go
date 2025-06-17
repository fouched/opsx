package config

import (
	"database/sql"
	"github.com/alexedwards/scs/v2"
	"github.com/fouched/opsx/internal/repo"
	"github.com/gorilla/schema"
	"log"
)

type App struct {
	Port     int                 `toml:"port"`
	DSN      string              `toml:"dsn"`
	DB       *sql.DB             `toml:"-"`
	Decoder  *schema.Decoder     `toml:"-"`
	DebugLog *log.Logger         `toml:"-"`
	InfoLog  *log.Logger         `toml:"-"`
	ErrorLog *log.Logger         `toml:"-"`
	Repo     Repo                `toml:"-"`
	Session  *scs.SessionManager `toml:"-"`
}

type Repo struct {
	RemoteAppRepo repo.RemoteAppRepoInterface
}
