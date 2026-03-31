package repositories

import (
	"database/sql"

	"backend/internal/database"
)

type Service interface {
	Health() map[string]string
	Close() error
	DB() *sql.DB
}

type service struct {
	db *sql.DB
}

var dbInstance *service

func (s *service) DB() *sql.DB {
	return s.db
}

func New() Service {
	if dbInstance != nil {
		return dbInstance
	}
	dbInstance = &service{
		db: database.GetDB(),
	}
	return dbInstance
}

func (s *service) Health() map[string]string {
	return database.Health()
}

func (s *service) Close() error {
	return nil
}

var _ Service = (*service)(nil)
