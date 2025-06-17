package repo

import (
	"github.com/fouched/opsx/internal/models"
	"time"
)

// DbTimeout should be in config and lowered in production - 30 secs max
const DbTimeout = time.Minute * 5

type RemoteAppRepoInterface interface {
	SelectById(id string) (models.RemoteApp, error)
	SelectAll() ([]models.RemoteApp, error)
	Insert(remoteService models.RemoteApp) (string, error)
}
