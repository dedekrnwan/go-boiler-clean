package adapter

import (
	"go-boiler-clean/internal/adapter/inbound/rest"
	"go-boiler-clean/internal/usecase"
)

type (
	InBound struct {
		Rest rest.Rest
	}
)

func NewInBound(
	usecaseUser usecase.User,
	usecasAuth usecase.Auth,
) *InBound {
	restInstance := rest.New(usecaseUser, usecasAuth)

	return &InBound{
		Rest: restInstance,
	}
}
