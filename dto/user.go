package dto

import "go-boiler-clean/entity"

type (
	User struct {
		entity.User
		Password string `json:"-"`
	}
)
