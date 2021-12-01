package services

import (
	"todo/app/repository"
	"todo/core"
)

type Services struct {
	UserService    *userServiceLayer
	SettingService *settingServiceLayer
}

func NewService(r repository.Repo, c *core.Config) Services {
	return Services{
		UserService:    newUserServiceLayer(r, c),
		SettingService: newSettingServiceLayer(c),
	}
}
