package handlers

import (
	"github.com/fouched/opsx/internal/config"
	"github.com/gorilla/schema"
)

type Handlers struct {
	App     *config.App
	Decoder *schema.Decoder
}

func NewHandlers(app *config.App) *Handlers {
	return &Handlers{App: app}
}
