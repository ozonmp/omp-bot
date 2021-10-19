package employee

type EmployeeRepository struct {
	storage map[int]Employee
}

func NewEmployeeRepository() EmployeeRepository {
	storage := map[int]Employee{
		1: {Id: 1, Title: "Artem"},
		2: {Id: 2, Title: "Vasya"},
		3: {Id: 3, Title: "Petya"},
	}

	return EmployeeRepository{storage}
}

func (repo EmployeeRepository) all() map[int]Employee {
	return repo.storage
}

func (repo EmployeeRepository) delete(id int) {
	delete(repo.storage, id)
}

func (repo EmployeeRepository) existsById(id int) bool {
	_, state := repo.storage[id]

	return state
}

func (repo EmployeeRepository) create(title string) Employee {
	id := len(repo.storage) + 1

	var employee = Employee{
		Id:    id,
		Title: title,
	}

	repo.storage[id] = employee

	return employee
}

func (repo EmployeeRepository) find(id int) Employee {
	return repo.storage[id]
}
