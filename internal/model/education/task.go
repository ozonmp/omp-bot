package education

import (
	"errors"
	"fmt"
)

type Task struct {
	Id            uint64 `json:"id"`
	Championat_id uint64 `json:"championatID"`
	Difficulty    uint64 `json:"difficulty"`
	Title         string `json:"title"`
	Description   string `json:"description"`
}

func (t *Task) String() string {
	return fmt.Sprintf("ProductID: %d. Title: %s. Description: %s.", t.Id, t.Title, t.Description)
}

func (t *Task) IsEmpty() bool {
	return t.Id == 0
}

type TaskModel []Task

func (t *TaskModel) FindID(taskID uint64) (int, error) {

	for i, v := range *t {
		if v.Id == taskID {
			return i, nil
		}
	}

	return 0, errors.New("ProductID not found")

}

func (t *TaskModel) MaxID() (MaxId uint64) {

	for _, v := range *t {
		if MaxId < v.Id {
			MaxId = v.Id
		}
	}
	MaxId += 1

	return
}

func (t *TaskModel) Count() int {
	return len(*t)
}

var TaskEntities TaskModel

func TaskEntitiesInit() {

	TaskEntities = TaskModel{
		Task{Id: 1, Championat_id: 1, Difficulty: 5, Title: "First product", Description: "first product desc"},
		Task{Id: 2, Championat_id: 1, Difficulty: 5, Title: "Second product", Description: "Second product desc"},
		Task{Id: 3, Championat_id: 1, Difficulty: 5, Title: "Third product", Description: "Third product desc"},
		Task{Id: 4, Championat_id: 1, Difficulty: 5, Title: "Fourth product", Description: "Fourth product desc"},
		Task{Id: 5, Championat_id: 1, Difficulty: 5, Title: "Fifth product", Description: "Fifth product desc"},
		Task{Id: 6, Championat_id: 1, Difficulty: 5, Title: "Sixth product", Description: "Sixth product desc"},
		Task{Id: 7, Championat_id: 1, Difficulty: 5, Title: "Seventh product", Description: "Seventh product desc"},
		Task{Id: 8, Championat_id: 1, Difficulty: 5, Title: "Eighth product", Description: "Eighth product desc"},
		Task{Id: 9, Championat_id: 1, Difficulty: 5, Title: "Ninth product", Description: "Ninth product desc"},
		Task{Id: 10, Championat_id: 1, Difficulty: 5, Title: "Tenth product", Description: "Tenth product desc"},
	}

}
