package usecase

import (
	// usecaseはrepositoryに依存している
	"todoapp/model"
	"todoapp/repository"
)

type TaskUseCase interface {
	CreateTask(title string) (int, error)
	GetTask(id int) (*model.Task, error)
	UpdateTask(id int, title string) error
	DeleteTask(id int) error
}

type taskUseCase struct {
	r repository.TaskRepository
}

func NewTaskUseCase(r repository.TaskRepository) *taskUseCase {
	return &taskUseCase{r: r}
}

func (u *taskUseCase) CreateTask(title string) (int, error) {
	task := model.Task{Title: title}
	err := task.Validate()
	if err != nil {
		return 0, err
	}
	id, err := u.r.Create(&task)
	println(id)
	return id, err
}

func (u *taskUseCase) GetTask(id int) (*model.Task, error) {
	t, err := u.r.Read(id)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (u *taskUseCase) UpdateTask(id int, title string) error {
	task := model.Task{ID: id, Title: title}
	err := u.r.Update(&task)
	return err
}

func (u *taskUseCase) DeleteTask(id int) error {
	err := u.r.Delete(id)
	return err
}
