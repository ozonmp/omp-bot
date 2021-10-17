package task

import (
	"errors"
)

type DummyTaskService struct {
	data *taskModel
}

func NewDummyTaskService() *DummyTaskService {

	taskEntities.Init()

	return &DummyTaskService{taskEntities}
}

func (s *DummyTaskService) List(cursor, limit uint64) (result taskModel) {

	lenSlice := uint64(len(*s.data))

	if cursor > lenSlice {
		cursor = 0
	}

	data := *s.data

	if cursor+limit > lenSlice {
		result = data[cursor:]
	} else {
		result = data[cursor : cursor+limit]
	}

	return result
}

func (s *DummyTaskService) Create(TaskId uint64, name, desc string) error {

	if _, err := s.Find(TaskId); err == nil {
		return errors.New("product id found on another product")
	}

	data := *s.data

	data = append(data, task{Id: TaskId, Championat_id: 1, Difficulty: 5, Title: name, Description: desc})

	s.data = &data

	return nil
}

func (s *DummyTaskService) Update(TaskId uint64, Title, Desc string) error {

	id, err := s.Find(TaskId)
	if err != nil {
		return err
	}

	data := *s.data
	data[id-1].Title = Title
	data[id-1].Description = Desc

	return nil
}

func (s *DummyTaskService) Remove(TaskId uint64) error {

	id, err := s.Find(TaskId)
	if err != nil {
		return err
	}

	data := *s.data
	data = append(data[:id], data[id+1:]...)

	s.data = &data

	return nil
}

func (s *DummyTaskService) Count() int {
	return len(*s.data)
}

func (s *DummyTaskService) Get(TaskId uint64) task {

	id, err := s.Find(TaskId)
	if err != nil {
		return task{}
	}

	result := *s.data

	return result[id]

}

func (s *DummyTaskService) Find(TaskId uint64) (int, error) {

	data := *s.data

	for i, v := range data {
		if v.Id == TaskId {
			return i, nil
		}
	}

	return 0, errors.New("id not found")
}
