package task

type task struct {
	Id            uint64
	Championat_id uint64
	Difficulty    uint64
	Title         string
	Description   string
}

func (t *task) IsEmpty() bool {
	return t.Id == 0
}

type taskModel []task

var taskEntities *taskModel

func (t *taskModel) Init() {

	p := taskModel{
		task{Id: 1, Championat_id: 1, Difficulty: 5, Title: "First product", Description: "first product desc"},
		task{Id: 2, Championat_id: 1, Difficulty: 5, Title: "Second product", Description: "Second product desc"},
		task{Id: 3, Championat_id: 1, Difficulty: 5, Title: "Third product", Description: "Third product desc"},
		task{Id: 4, Championat_id: 1, Difficulty: 5, Title: "Fourth product", Description: "Fourth product desc"},
		task{Id: 5, Championat_id: 1, Difficulty: 5, Title: "Fifth product", Description: "Fifth product desc"},
		task{Id: 6, Championat_id: 1, Difficulty: 5, Title: "Sixth product", Description: "Sixth product desc"},
		task{Id: 7, Championat_id: 1, Difficulty: 5, Title: "Seventh product", Description: "Seventh product desc"},
		task{Id: 8, Championat_id: 1, Difficulty: 5, Title: "Eighth product", Description: "Eighth product desc"},
		task{Id: 9, Championat_id: 1, Difficulty: 5, Title: "Ninth product", Description: "Ninth product desc"},
		task{Id: 10, Championat_id: 1, Difficulty: 5, Title: "Tenth product", Description: "Tenth product desc"},
	}

	taskEntities = &p
}
