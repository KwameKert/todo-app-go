package services

import (
	"todo/app/repository"
	"todo/core"
)

type Services struct {
	UserService    *userServiceLayer
	SettingService *settingServiceLayer
	TaskService    *taskServiceLayer
}

func NewService(r repository.Repo, c *core.Config) Services {
	return Services{
		UserService:    newUserServiceLayer(r, c),
		TaskService:    newTaskServiceLayer(r, c),
		SettingService: newSettingServiceLayer(c),
	}
}
