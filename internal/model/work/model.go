package work

import (
	"encoding/json"
	"time"
)

type Project struct {
	ID uint64 `json:"ID"`
	Name string `json:"Name"`
	TeamID uint64 `json:"TeamID"`
	CreatedAt time.Time `json:"CreatedAt"`
}

func(p *Project) String() (string) {
	text, err := json.Marshal(p)
	if err != nil {
		return ""
	}

	return string(text)
}