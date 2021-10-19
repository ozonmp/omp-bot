package employee

type Service struct{}

var repository = NewEmployeeRepository()

func NewService() *Service {
	return &Service{}
}

func (service *Service) List() map[int]Employee {
	return repository.all()
}

func (service *Service) Get(idx int) (Employee, error) {
	return repository.find(idx), nil
}

func (service *Service) Delete(id int) bool {
	if repository.existsById(id) {
		repository.delete(id)

		return true
	}

	return false
}

func (service *Service) Create(title string) Employee {
	return repository.create(title)
}
