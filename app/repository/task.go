package repository

import (
	//	"errors"

	"todo/app/models"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type taskLayer struct {
	db *gorm.DB
}

func newTaskRepoLayer(db *gorm.DB) *taskLayer {
	return &taskLayer{
		db: db,
	}
}

func (ul *taskLayer) Create(task *models.Task) (*models.Task, error) {
	if err := ul.db.Create(task).Error; err != nil {
		log.Info(err)
		return nil, err
	}
	return task, nil
}

func (ul *taskLayer) Update(task *models.Task) (*models.Task, error) {
	if err := ul.db.Save(&task).Error; err != nil {
		log.Info(err)
		return nil, err
	}
	return task, nil
}

func (ul *taskLayer) Get(task *models.Task, id int) error {

	if err := ul.db.Find(&task, id).First(&task).Error; err != nil {

		return err
	}
	return nil
}

func (ul *taskLayer) Fetch(tasks *[]models.Task) error {

	if err := ul.db.Find(&tasks).Error; err != nil {
		log.Error("error -->", err)
		return err
	}
	return nil
}

func (ul *taskLayer) FetchUserTasks(tasks *[]models.Task, userId int) error {

	if err := ul.db.Where("user_id = ?", userId).Find(&tasks).Error; err != nil {
		log.Error("error -->", err)
		return err
	}
	return nil
}
