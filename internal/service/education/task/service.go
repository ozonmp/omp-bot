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
}

func NewDummyTaskService() *DummyTaskService {

	education.TaskEntitiesInit()

	return &DummyTaskService{}
}

func (s *DummyTaskService) Describe(taskID uint64) (*education.Task, error) {

	data := *education.TaskEntities

	id, err := data.FindID(taskID)
	if err != nil {
		return &education.Task{}, err
	}

	return &data[id], err

}

func (s *DummyTaskService) List(cursor, limit uint64) ([]education.Task, error) {

	data := *education.TaskEntities
	dataCount := data.Count()

	if cursor > uint64(dataCount) {
		cursor = 0
	}

	if cursor+limit < uint64(dataCount) {
		return data[cursor : cursor+limit], nil
	}

	return data[cursor:], nil

}

func (s *DummyTaskService) Create(Task education.Task) (uint64, error) {

	data := *education.TaskEntities

	if _, err := data.FindID(Task.Id); err == nil {
		return 0, errors.New("product id found on another product")
	}

	Task.Id = data.MaxID()

	data = append(data, Task)

	education.TaskEntities = &data

	return Task.Id, nil
}

func (s *DummyTaskService) Update(taskID uint64, Task education.Task) error {

	data := *education.TaskEntities

	id, err := data.FindID(taskID)
	if err != nil {
		return err
	}

	data[id].Title = Task.Title
	data[id].Description = Task.Description

	return nil
}

func (s *DummyTaskService) Remove(taskID uint64) (bool, error) {

	data := *education.TaskEntities

	id, err := data.FindID(taskID)
	if err != nil {
		return false, err
	}

	data = append(data[:id], data[id+1:]...)

	education.TaskEntities = &data

	return true, nil
}

func (s *DummyTaskService) CountData() int {
	return education.TaskEntities.Count()
}
