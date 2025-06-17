package handlers

import (
	"github.com/fouched/opsx/internal/render"
	"github.com/fouched/opsx/internal/services"
	"github.com/fouched/opsx/internal/templates"
	"net/http"
)

func (h *Handlers) AllApps(w http.ResponseWriter, r *http.Request) {

	remoteAppService := services.NewRemoteAppService(h.App)
	apps, err := remoteAppService.GetAll()
	if err != nil {
		h.App.ErrorLog.Print(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	t := templates.RemoteAppGrid(apps)
	_ = render.Template(w, r, t)
}
