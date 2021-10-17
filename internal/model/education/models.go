package education

import "fmt"

type Solution struct {
	Id uint64
	TaskID uint64
	Title string
	Autor string
}

func (c Solution) String() string{
	return fmt.Sprintf("ID: %d TaskID: %d Autor: %s Title: %s", c.Id, c.TaskID, c.Autor, c.Title)
}

var Data map[uint64]Solution

func init() {
	Data = make(map[uint64]Solution, 100)
	Data[1] = Solution{1, 100, "Cool Title 1", "Unknown1"}
	Data[2] = Solution{2, 200, "Cool Title 2", "Unknown2"}
	Data[3] = Solution{3, 300, "Cool Title 3", "Unknown3"}
	Data[4] = Solution{4, 400, "Cool Title 4", "Unknown4"}
	Data[5] = Solution{5, 500, "Cool Title 5", "Unknown5"}
	Data[6] = Solution{6, 600, "Cool Title 6", "Unknown6"}
	Data[7] = Solution{7, 700, "Cool Title 7", "Unknown7"}
}

