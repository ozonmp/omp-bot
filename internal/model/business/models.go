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
	{Name: "Company_05", ZipCode: 100005, Address: "City 5, Street 5"},
	{Name: "Company_06", ZipCode: 100006, Address: "City 6, Street 6"},
	{Name: "Company_07", ZipCode: 100007, Address: "City 7, Street 7"},
}

var Limit = 2

func (c *Company) String() string {
	return fmt.Sprintf("Company %s [Address: %s, ZipCode: %d]", c.Name, c.Address, c.ZipCode)
}
