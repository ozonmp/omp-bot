package incident

import "encoding/json"

var allEntities = map[int]*Incident{
	1: {Id: 1, Title: "Timeout happened", Author: "Петров"},
	2: {Id: 2, Title: "Service is unavailable", Author: "Иванов"},
	3: {Id: 3, Title: "Programmer is retarded", Author: "Сидоров"},
	4: {Id: 4, Title: "Something going wrong", Author: "Сидоренко"},
	5: {Id: 5, Title: "Internal server error", Author: "Кириленко"},
}

type Incident struct {
	Id     int    `json:"Id"`
	Title  string `json:"Title"`
	Author string `json:"Author"`
}

func (sub *Incident) String() (string, error) {
	rawSub, err := json.Marshal(sub)
	return string(rawSub), err
}

func ConvertStringToIncident(rawEntity string) (Incident, error) {
	var entity Incident
	err := json.Unmarshal([]byte(rawEntity), &entity)
	if err != nil {
		return Incident{}, err
	}
	return entity, nil
}
