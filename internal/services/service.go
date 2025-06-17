package services

import "github.com/fouched/opsx/internal/models"

type RemoteAppServiceInterface interface {
	GetAll() ([]models.RemoteApp, error)
}
