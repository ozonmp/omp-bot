package employee

import "strconv"

type EmployeeService interface {
	List() string

	Get(idx int) string

	Delete(id int) string

	Create(title string) string

	Update(id int, title string) string
}

type Service struct{}

var repository = NewRepository()

func NewService() EmployeeService {
	return &Service{}
}

func (service *Service) List() string {
	outputMsgText := "Success -> Here all the employees: \n\n"

	for _, p := range repository.all() {
		outputMsgText += p.idAsString() + " | " + p.Title
		outputMsgText += "\n"
	}

	return outputMsgText
}

func (service *Service) Get(idx int) string {
	if !repository.existsById(idx) {
		return "Error -> employee[" + strconv.Itoa(idx) + "] not found!"
	}

	employee := repository.find(idx)

	return "Success -> Employee[" + employee.idAsString() + "] - " + employee.Title
}

func (service *Service) Delete(id int) string {
	if repository.existsById(id) {
		repository.delete(id)

		return "Success -> Deleted Employee[" + strconv.Itoa(id) + "]"
	}

	return "Error -> Employee[" + strconv.Itoa(id) + "] not found"
}

func (service *Service) Create(title string) string {
	employee := repository.create(title)

	return "Error -> Was created employee[" + employee.idAsString() + "] " + employee.Title
}

func (service *Service) Update(id int, title string) string {
	if !repository.existsById(id) {
		return "Error -> Employee[" + strconv.Itoa(id) + "] not found"
	}

	employee := repository.find(id)
	employee.Title = title

	repository.update(employee)

	return "Success -> Employee[" + employee.idAsString() + "] was updated"
}
