package handlers

import (
	"github.com/fouched/opsx/internal/render"
	"github.com/fouched/opsx/internal/templates"
	"net/http"
)

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {

	t := templates.Home()
	_ = render.Template(w, r, t)
}
