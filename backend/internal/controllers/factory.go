package controllers

import (
	"backend/internal/auth"
	"backend/internal/repositories"
	"backend/internal/services"
)

func NewAuthHandler(db repositories.Service, cache services.RedisService, jwtSecret string) *auth.Handler {
	return auth.NewHandler(db, cache, jwtSecret)
}
