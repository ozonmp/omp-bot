package work

import "fmt"

type Course struct {
	Id          uint64
	Name        string
	Timing      string
	Description string
	Level       string
}

func (c *Course) String() string {
	return fmt.Sprintf("Id: %d\nName: %s\nTiming: %s\nDescription:\n%s\nLevel: %s\n", c.Id, c.Name, c.Timing, c.Description, c.Level)
}

var AllCourses = []Course{
	{Id: 1, Name: "Golang", Timing: "12.10.2021-31.12.2021", Description: "Курс_по_Golang_для_новичков", Level: "beginner"},
	{Id: 2, Name: "Java", Timing: "1.12.2021-1.1.2022", Description: "Курс_по_Java", Level: "intermediate"},
	{Id: 3, Name: "Kotlin", Timing: "1.1.2021-1.2.2022", Description: "Курс_по_Kotlin_для_профи", Level: "pro"},
	{Id: 4, Name: "Flutter", Timing: "1.2.2021-1.3.2022", Description: "Курс_по_Flutter", Level: "intermediate"},
	{Id: 5, Name: "Python", Timing: "1.3.2021-1.4.2022", Description: "Курс_по_Python", Level: "intermediate"},
	{Id: 6, Name: "C++", Timing: "1.4.2021-1.5.2022", Description: "Курс_по_C++_для_профи", Level: "pro"},
	{Id: 7, Name: "Rust", Timing: "1.5.2021-1.6.2022", Description: "Курс_по_Rust_для профи", Level: "pro"},
	{Id: 8, Name: "C#", Timing: "1.6.2021-1.7.2022", Description: "Курс_по_C#_для_новичков", Level: "beginner"},
	{Id: 9, Name: "ML", Timing: "1.7.2021-1.8.2022", Description: "Курс_по_ML_для_профи", Level: "pro"},
}

