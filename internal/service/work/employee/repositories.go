package employee

type Repository struct {
	storage map[int]Employee
}

func NewRepository() Repository {
	storage := map[int]Employee{
		1: {Id: 1, Title: "Artem"},
		2: {Id: 2, Title: "Vasya"},
		3: {Id: 3, Title: "Petya"},
	}

	return Repository{storage}
}

func (repo Repository) all() map[int]Employee {
	return repo.storage
}

func (repo Repository) delete(id int) {
	delete(repo.storage, id)
}

func (repo Repository) existsById(id int) bool {
	_, state := repo.storage[id]

	return state
}

func (repo Repository) create(title string) Employee {
	id := len(repo.storage) + 1

	var employee = Employee{
		Id:    id,
		Title: title,
	}

	repo.storage[id] = employee

	return employee
}

func (repo Repository) find(id int) Employee {
	return repo.storage[id]
}

func (repo Repository) update(employee Employee) {
	repo.storage[employee.Id] = employee
}
