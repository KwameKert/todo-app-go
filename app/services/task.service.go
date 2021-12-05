package services

import (
	//	"errors"
	"todo/app/models"
	"todo/app/repository"
	"todo/core"
	//	"gorm.io/gorm"
)

type taskServiceLayer struct {
	repository repository.Repo
	config     *core.Config
}

func newTaskServiceLayer(r repository.Repo, c *core.Config) *taskServiceLayer {
	return &taskServiceLayer{
		repository: r,
		config:     c,
	}
}

func (t *taskServiceLayer) CreateTask(req core.CreateTaskRequest) core.Response {
	task := models.Task{
		Name:        req.Name,
		Description: req.Description,
		UserId:      req.UserId,
		Status:      models.Pending,
	}
	user := models.User{}

	if err := t.repository.Users.Get(&user, req.UserId); err != nil {
		return core.BadRequest(err, core.String("user not found"))
	}

	taskSaved, err := t.repository.Tasks.Create(&task)
	if err != nil {
		return core.Error(err, nil)
	}

	return core.Success(&map[string]interface{}{
		"task": taskSaved,
	}, core.String("task created successfully"))
}

func (t *taskServiceLayer) UpdateTask(task models.Task) core.Response {

	taskDTO := models.Task{}
	if err := t.repository.Tasks.Get(&taskDTO, task.Id); err != nil {
		return core.BadRequest(err, core.String("task not found"))
	}

	taskSaved, err := t.repository.Tasks.Update(&task)
	if err != nil {
		return core.Error(err, nil)
	}

	return core.Success(&map[string]interface{}{
		"task": taskSaved,
	}, core.String("task updated successfully"))
}

func (t *taskServiceLayer) GetTaskById(id int) core.Response {
	task := models.Task{}

	if err := t.repository.Tasks.Get(&task, id); err != nil {
		return core.Error(err, nil)
	}

	return core.Success(&map[string]interface{}{
		"task": task,
	}, core.String("task created successfully"))
}

func (t *taskServiceLayer) FetchAllTasks() core.Response {
	var tasks []models.Task

	if err := t.repository.Tasks.Fetch(&tasks); err != nil {
		return core.Error(err, nil)
	}

	if len(tasks) < 1 {
		return core.NoContentFound(nil, core.String("No tasks found"))
	}

	return core.Success(&map[string]interface{}{
		"tasks": tasks,
	}, core.String("tasks found successfully"))
}

func (t *taskServiceLayer) FetchAllUserTasks(userId int) core.Response {
	var tasks []models.Task

	if err := t.repository.Tasks.FetchUserTasks(&tasks, userId); err != nil {
		return core.Error(err, nil)
	}

	if len(tasks) < 1 {
		return core.NoContentFound(nil, core.String("No tasks found"))
	}

	return core.Success(&map[string]interface{}{
		"tasks": tasks,
	}, core.String("tasks found successfully"))
}

//modify the repo to have shared crud responsibilities
