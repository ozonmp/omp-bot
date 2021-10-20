package course

import (
	"errors"
	"github.com/ozonmp/omp-bot/internal/model/work"
)

type CourseService interface {
	Describe(courseID uint64) (*work.Course, error)
	List(cursor uint64, limit uint64) ([]work.Course, error)
	Create(work.Course) (uint64, error)
	Update(courseID uint64, course work.Course) error
	Remove(courseID uint64) (bool, error)
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List(cursor uint64, limit uint64) ([]work.Course, error) {
	lenght := uint64(len(work.AllCourses))
	if cursor >= lenght {
		return nil, errors.New("cursor out of range")
	}

	if cursor+limit < lenght {
		return work.AllCourses[cursor : limit+cursor], nil
	} else {
		return work.AllCourses[cursor:], nil
	}

}

func (s *Service) Describe(courseID uint64) (*work.Course, error) {
	for i := 0; i < len(work.AllCourses); i++ {
		if courseID == work.AllCourses[i].Id {
			return &work.AllCourses[i], nil
		}
	}
	err := errors.New("course not found")
	return nil, err
}

func (s *Service) Create(Course work.Course) (uint64, error) {
	var max uint64 = 0
	for _, v := range work.AllCourses {
		if v.Id > max {
			max = v.Id
		}
	}
	max += 1
	Course.Id = max
	work.AllCourses = append(work.AllCourses, Course)

	return Course.Id, nil
}

func (s *Service) Remove(courseID uint64) (bool, error) {
	for i := range work.AllCourses {
		if courseID == work.AllCourses[i].Id {
			work.AllCourses = append(work.AllCourses[:i], work.AllCourses[i+1:]...)
			return true, nil
		}
	}
	return false, nil
}

func (s *Service) Update(Course work.Course) error {
	id := Course.Id
	for i := 0; i < len(work.AllCourses); i++ {
		if id == work.AllCourses[i].Id {
			work.AllCourses[i] = Course
			return nil
		}
	}
	return errors.New("not found id")
}
