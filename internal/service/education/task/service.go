package task

import (
	"errors"

	"github.com/ozonmp/omp-bot/internal/model/education"
)

type TaskService interface {
	Describe(taskID uint64) (*education.Task, error)
	List(cursor uint64, limit uint64) ([]education.Task, error)
	Create(education.Task) (uint64, error)
	Update(taskID uint64, task education.Task) error
	Remove(taskID uint64) (bool, error)
}

type DummyTaskService struct {
	data *education.TaskModel
}

func NewDummyTaskService() *DummyTaskService {

	education.TaskEntities.Init()

	return &DummyTaskService{education.TaskEntities}
}

func (s *DummyTaskService) Describe(taskID uint64) (*education.Task, error) {

	data := *s.data

	id, err := data.FindID(taskID)
	if err != nil {
		return &education.Task{}, err
	}

	return &data[id], err

}

func (s *DummyTaskService) List(cursor, limit uint64) ([]education.Task, error) {

	lenSlice := uint64(len(*s.data))

	if cursor > lenSlice {
		cursor = 0
	}

	data := *s.data

	if cursor+limit < lenSlice {
		return data[cursor : cursor+limit], nil
	}

	return data[cursor:], nil

}

func (s *DummyTaskService) Create(Task education.Task) (uint64, error) {

	data := *s.data

	if _, err := data.FindID(Task.Id); err == nil {
		return 0, errors.New("product id found on another product")
	}

	Task.Id = data.MaxID()

	data = append(data, Task)

	s.data = &data

	return Task.Id, nil
}

func (s *DummyTaskService) Update(taskID uint64, Task education.Task) error {

	data := *s.data

	id, err := data.FindID(taskID)
	if err != nil {
		return err
	}

	data[id].Title = Task.Title
	data[id].Description = Task.Description

	return nil
}

func (s *DummyTaskService) Remove(taskID uint64) (bool, error) {

	data := *s.data

	id, err := data.FindID(taskID)
	if err != nil {
		return false, err
	}

	data = append(data[:id], data[id+1:]...)

	s.data = &data

	return true, nil
}

// TODO - Как такое лучше организовывать ?
func (s *DummyTaskService) CountData() int {
	return s.data.Count()
}
