package usecase_test

import (
	"testing"
	"todoapp/model"
	"todoapp/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TaskRepositoryMock struct {
	mock.Mock
}

// repositoryのTaskRepositoryインターフェースを満たすmockを作成
func (m *TaskRepositoryMock) Create(task *model.Task) (int, error) {
	args := m.Called(task)
	return args.Int(0), args.Error(1)
}

func (m *TaskRepositoryMock) Read(id int) (*model.Task, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Task), args.Error(1)
}

func (m *TaskRepositoryMock) Update(task *model.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *TaskRepositoryMock) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestTaskUseCase(t *testing.T) {
	mockRepo := new(TaskRepositoryMock)
	taskUseCase := usecase.NewTaskUseCase(mockRepo)

	task := &model.Task{Title: "test"}

	// mockしたいメソッドをonで指定してそこにtaskを渡し返り値を指定する
	mockRepo.On("Create", task).Return(1, nil)

	id, err := taskUseCase.CreateTask(task.Title)

	assert.NoError(t, err)
	assert.Equal(t, 1, id)
}
