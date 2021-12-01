package services

import (
	//	"errors"
	"todo/app/models"
	"todo/app/repository"
	"todo/core"
	//	"gorm.io/gorm"
)

type userServiceLayer struct {
	repository repository.Repo
	config     *core.Config
}

func newUserServiceLayer(r repository.Repo, c *core.Config) *userServiceLayer {
	return &userServiceLayer{
		repository: r,
		config:     c,
	}
}

func (u *userServiceLayer) CreateUser(req core.CreateUserRequest) core.Response {
	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	if err := u.repository.Users.Create(&user); err != nil {
		return core.Error(err, nil)
	}

	return core.Success(&map[string]interface{}{
		"user": user,
	}, core.String("user created successfully"))
}

func (u *userServiceLayer) FetchUsers() core.Response {
	user := models.User{}
	if err := u.repository.Users.Fetch(&user); err != nil {
		return core.Error(err, nil)
	}

	return core.Success(&map[string]interface{}{
		"user": user,
	}, core.String("users found successfully"))
}

func (u *userServiceLayer) GetUser(id int) core.Response {
	user := models.User{}
	if err := u.repository.Users.Get(&user, id); err != nil {
		return core.Error(err, nil)
	}

	return core.Success(&map[string]interface{}{
		"user": user,
	}, core.String("users found successfully"))
}
