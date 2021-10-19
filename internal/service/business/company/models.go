package company

type Company struct {
	Name string
	ZipCode int64
}

var allEntities = []Company {
	{ Name: "Company_one", ZipCode: 100001 },
	{ Name: "Company_two", ZipCode: 100002 },
	{ Name: "Company_three", ZipCode: 100003 },
}