package services

import (
	"github.com/fouched/opsx/internal/config"
	"github.com/fouched/opsx/internal/models"
	"github.com/fouched/opsx/internal/repo"
)

type RemoteApp struct {
	App           *config.App
	RemoteAppRepo repo.RemoteAppRepoInterface
}

func NewRemoteAppService(app *config.App) RemoteAppServiceInterface {
	return &RemoteApp{
		App:           app,
		RemoteAppRepo: app.Repo.RemoteAppRepo,
	}
}

func (r *RemoteApp) GetAll() ([]models.RemoteApp, error) {
	rs, err := r.RemoteAppRepo.SelectAll()
	if err != nil {
		return nil, err
	}

	return rs, nil
}
