package access

var allEntities = []Access{
	{User: "Test1", Path: "TestPath1", ID: 1},
	{User: "Test2", Path: "TestPath2", ID: 2},
	{User: "Test3", Path: "TestPath3", ID: 3},
}

type Access struct {
	ID   uint64
	User string
	Path string
}
