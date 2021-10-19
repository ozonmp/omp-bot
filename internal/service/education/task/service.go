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

	id, err := education.TaskEntities.FindID(taskID)
	if err != nil {
		return &education.Task{}, err
	}

	return &education.TaskEntities[id], err

}

func (s *DummyTaskService) List(cursor, limit uint64) (result []education.Task, err error) {

	dataCount := education.TaskEntities.Count()

	if cursor > uint64(dataCount) {
		return []education.Task{}, errors.New("cursor out of data")
	}

	if cursor+limit < uint64(dataCount) {
		result = education.TaskEntities[cursor : cursor+limit]
	} else {
		result = education.TaskEntities[cursor:]
	}

	return
}

func (s *DummyTaskService) Create(Task education.Task) (uint64, error) {

	if _, err := education.TaskEntities.FindID(Task.Id); err == nil {
		return 0, errors.New("product id found on another product")
	}

	Task.Id = education.TaskEntities.MaxID()

	education.TaskEntities = append(education.TaskEntities, Task)

	return Task.Id, nil
}

func (s *DummyTaskService) Update(taskID uint64, Task education.Task) error {

	id, err := education.TaskEntities.FindID(taskID)
	if err != nil {
		return err
	}

	education.TaskEntities[id].Title = Task.Title
	education.TaskEntities[id].Description = Task.Description

	return nil
}

func (s *DummyTaskService) Remove(taskID uint64) (bool, error) {

	id, err := education.TaskEntities.FindID(taskID)
	if err != nil {
		return false, err
	}

	education.TaskEntities = append(education.TaskEntities[:id], education.TaskEntities[id+1:]...)

	return true, nil
}

func (s *DummyTaskService) CountData() int {
	return education.TaskEntities.Count()
}
