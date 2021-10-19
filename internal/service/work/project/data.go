package project

import (
	"github.com/ozonmp/omp-bot/internal/model/work"
	"time"
)

var projectData = map[uint64]work.Project{
	1: {ID: 1, Name: "first", TeamID: 1, CreatedAt: time.Now()},
	2: {ID: 2, Name: "second", TeamID: 1, CreatedAt: time.Now()},
	3: {ID: 3, Name: "third", TeamID: 1, CreatedAt: time.Now()},
}
