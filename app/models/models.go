package models

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"time"
)

type Status string

var (
	Pending   Status = "pending"
	Completed Status = "completed"
)

type Model struct {
	Id        int        `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
}

type User struct {
	Model
	Name     string `json:"name"`
	Email    string `gorm:"index:,unique" json:"email"`
	Password string `json:"password"`
	Tasks    []Task
}

type Task struct {
	Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      Status `json:"status"`
	UserId      int    `json:"user_id"`
}

func RunSeeds(db *gorm.DB) {
	user := User{
		Name:     "kwamekert",
		Email:    "kwamekert@gmail.com",
		Password: "password",
	}

	if err := db.Model(&User{}).Where("name=?", user.Name).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			db.Create(&user)
		} else {
			log.Println("err: ", err)
		}
	}

	task := Task{
		UserId:      user.Id,
		Name:        "Check database",
		Description: "Review database description and scripts to query",
	}

	if err := db.Model(&Task{}).Where("name=?", task.Name).First(&task).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			db.Create(&task)
		} else {
			log.Println("err: ", err)
		}
	}

}
