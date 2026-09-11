package services

import (
	"user-service/repositories"
	services "user-service/services/user"
)

type Registry struct {
	Repository repositories.IRepositoryRegistry
}

type IServiceRegistry interface {
	GetUser() services.IUserService
}

func NewServiceRegistry(repository repositories.IRepositoryRegistry) IServiceRegistry {
	return &Registry{
		Repository: repository,
	}
}

func (r *Registry) GetUser() services.IUserService {
	return services.NewUserService(r.Repository)
}
