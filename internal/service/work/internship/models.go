package internship

var allEntities = []Internship{
	{Id: 1, Team_id: 1, Description: "Task-1 Lead", Period: "1-10/10/2021", Compensation: false},
	{Id: 2, Team_id: 1, Description: "Task-1 Pr.1", Period: "1-10/10/2021", Compensation: true},
	{Id: 3, Team_id: 1, Description: "Task-1 Pr.2", Period: "1-10/10/2021", Compensation: true},
	{Id: 4, Team_id: 2, Description: "Task-2&3 Lead", Period: "11-20/10/2021", Compensation: false},
	{Id: 5, Team_id: 2, Description: "Task-2&4 Pr.1", Period: "11-20/10/2021", Compensation: true},
	{Id: 6, Team_id: 2, Description: "Task-3&4 Pr.2", Period: "11-20/10/2021", Compensation: true},
	{Id: 7, Team_id: 3, Description: "Task-5 Pr.1", Period: "21-30/10/2021", Compensation: true},
	{Id: 8, Team_id: 3, Description: "Task-5 Pr.2", Period: "21-30/10/2021", Compensation: true},
}

type Internship struct {
	Id           uint64
	Team_id      uint64
	Description  string
	Period       string
	Compensation bool
}
