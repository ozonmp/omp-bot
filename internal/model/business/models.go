package business

import "fmt"

type Company struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	ZipCode int64  `json:"zipcode"`
}

var AllEntities = []Company{
	{Name: "Company_01", ZipCode: 100001, Address: "City 1, Street 1"},
	{Name: "Company_02", ZipCode: 100002, Address: "City 2, Street 2"},
	{Name: "Company_03", ZipCode: 100003, Address: "City 3, Street 3"},
	{Name: "Company_04", ZipCode: 100004, Address: "City 4, Street 4"},
}

func (c *Company) String() string {
	return fmt.Sprintf("Company %s [Address: %s, ZipCode: %d]\n", c.Name, c.Address, c.ZipCode)
}
