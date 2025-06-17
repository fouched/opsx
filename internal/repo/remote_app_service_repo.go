package repo

import (
	"database/sql"
	"fmt"
	"github.com/fouched/opsx/internal/models"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

type RemoteAppRepo struct {
	Session db.Session
}

func NewRemoteAppRepo(db *sql.DB) RemoteAppRepoInterface {
	session, _ := postgresql.New(db)
	return &RemoteAppRepo{
		Session: session,
	}
}

func (r RemoteAppRepo) Table() string {
	return "remote_services"
}

func (r RemoteAppRepo) Insert(remoteApp models.RemoteApp) (string, error) {
	collection := r.Session.Collection(r.Table())

	rs, err := collection.Insert(remoteApp)
	if err != nil {
		return "", fmt.Errorf("failed to insert remote service: %w", err)
	}

	return fmt.Sprintf("%v", rs), nil
}

func (r RemoteAppRepo) SelectById(id string) (models.RemoteApp, error) {
	var app models.RemoteApp
	collection := r.Session.Collection(r.Table())

	if err := collection.Find(db.Cond{"id": id}).One(&app); err != nil {
		return app, fmt.Errorf("failed to fetch remote app: %w", err)
	}
	return app, nil
}

func (r RemoteAppRepo) SelectAll() ([]models.RemoteApp, error) {
	var apps []models.RemoteApp
	collection := r.Session.Collection(r.Table())

	rs := collection.Find().OrderBy("name")
	if err := rs.All(&apps); err != nil {
		return nil, err
	}

	return apps, nil
}
