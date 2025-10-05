package service

import (
	"github.com/brij812/go-Foundry/internal/lib/job"
	"github.com/brij812/go-Foundry/internal/repository"
	"github.com/brij812/go-Foundry/internal/server"
)

type Services struct {
	Auth *AuthService
	Job  *job.JobService
}

func NewServices(s *server.Server, repos *repository.Repositories) (*Services, error) {
	authService := NewAuthService(s)

	return &Services{
		Job:  s.Job,
		Auth: authService,
	}, nil
}
