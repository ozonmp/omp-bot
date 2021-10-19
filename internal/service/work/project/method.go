package project

import (
	"errors"
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/work"
)

func(p *Service) Describe(projectID uint64) (*work.Project, error) {

	if _, ok := projectData[projectID]; ok == false {
		return nil, fmt.Errorf("project ID %v is not found", projectID)
	}
	project, _ := projectData[projectID]

	return &project, nil
}

func(p *Service) List(cursor uint64, limit uint64) ([]work.Project, error) {

	length := len(projectData)
	if length == 0 {
		return nil, errors.New("projects are not found")
	}

	projectArray := make([]work.Project, 0, length)
	for _, v := range projectData {
		projectArray = append(projectArray, v)
	}

	return projectArray, nil
}

func(p *Service) Create(project work.Project) (uint64, error) {

	if _, ok := projectData[project.ID]; ok == true {
		return 0, fmt.Errorf("project ID %v is present", project.ID)
	}
	projectData[project.ID] = project
	return project.ID, nil
}

func(p *Service) Update(projectID uint64, project work.Project) error {

	if _, ok := projectData[projectID]; ok == false {
		return fmt.Errorf("project ID %v is not found", projectID)
	}
	projectData[projectID] = project
	return nil
}

func(p *Service) Remove(projectID uint64) (bool, error) {

	if _, ok := projectData[projectID]; ok == false {
		return false, fmt.Errorf("project ID %v is not found", projectID)
	}
	delete(projectData, projectID)

	return true, nil
}
