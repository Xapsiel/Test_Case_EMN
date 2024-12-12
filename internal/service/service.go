package service

import "mobileTest_Case/internal/repository"

type User interface {
	Register(nickname, name, email string) error
	VerifyEmail(token string) error
	SendToken(email string) error
}

type Service struct {
	User
}

func NewService(repo *repository.Repository) *Service {
	return &Service{User: NewUserService(repo.User)}
}
