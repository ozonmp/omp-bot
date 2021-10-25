package access

var allEntities = []Access{
	{ID: 1, Role_ID: 5, Resource_ID: 3},
	{ID: 1, Role_ID: 6, Resource_ID: 3},
	{ID: 1, Role_ID: 3, Resource_ID: 5},
	{ID: 1, Role_ID: 2, Resource_ID: 1},
	{ID: 1, Role_ID: 2, Resource_ID: 3},
	{ID: 1, Role_ID: 7, Resource_ID: 2},
}

type Access struct {
	ID          uint64
	Role_ID     uint64
	Resource_ID uint64
}
