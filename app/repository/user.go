package repository

import (
	//	"errors"
	"todo/app/models"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type userLayer struct {
	db *gorm.DB
}

func newUserRepoLayer(db *gorm.DB) *userLayer {
	return &userLayer{
		db: db,
	}
}

func (ul *userLayer) Create(user *models.User) error {
	if err := ul.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (ul *userLayer) Fetch(user *[]models.User) error {

	if err := ul.db.Find(&user).Error; err != nil {
		log.Error("error -->", err)
		return err
	}
	return nil
}

func (ul *userLayer) Get(user *models.User, id int) error {

	if err := ul.db.Preload("Tasks").Find(&user, id).First(&user).Error; err != nil {

		return err
	}
	return nil
}

func (ul *userLayer) Delete(user *models.User, id int) error {

	if err := ul.db.Delete(&user, id).Error; err != nil {
		return err
	}
	return nil
}

//work on user update here

func (ul *userLayer) Update(user *models.User) (*models.User, error) {
	if err := ul.db.Save(&user).Error; err != nil {
		log.Info(err)
		return nil, err
	}
	return user, nil
}
