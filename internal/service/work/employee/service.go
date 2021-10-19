package employee

import "strconv"

type Service struct{}

var repository = NewEmployeeRepository()

func NewService() *Service {
	return &Service{}
}

func (service *Service) List() string {
	outputMsgText := "Success -> Here all the employees: \n\n"

	employees := repository.all()
	for _, p := range employees {
		outputMsgText += p.Title
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
