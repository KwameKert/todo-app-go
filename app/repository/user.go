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

func (ul *userLayer) Fetch(user *models.User) error {

	// 	result := ul.db.Find(&user)
	// 	if result.Error != nil {
	// 	   log.Error("error -->", result.Error)
	// 	   return result.Error
	//    }
	//    if result.RowsAffected >0 {
	// 	   return nil
	//    }
	//    return nil

	//check if user list is empty and return 204
	if err := ul.db.Find(&user).Error; err != nil {
		log.Error("error -->", err)
		return err
	}
	return nil
}

func (ul *userLayer) Get(user *models.User, id int) error {

	if err := ul.db.Find(&user, id).First(&user).Error; err != nil {

		return err
	}
	return nil
}
