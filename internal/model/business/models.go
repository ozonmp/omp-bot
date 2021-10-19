package business

import "fmt"

type Company struct {
	Name    string
	ZipCode int64
}

var AllEntities = []Company{
	{Name: "Company_01", ZipCode: 100001},
	{Name: "Company_02", ZipCode: 100002},
	{Name: "Company_03", ZipCode: 100003},
	{Name: "Company_04", ZipCode: 100004},
	{Name: "Company_05", ZipCode: 100005},
	{Name: "Company_06", ZipCode: 100006},
	{Name: "Company_07", ZipCode: 100007},
	{Name: "Company_08", ZipCode: 100008},
	{Name: "Company_09", ZipCode: 100009},
	{Name: "Company_10", ZipCode: 100010},
	{Name: "Company_11", ZipCode: 100011},
	{Name: "Company_12", ZipCode: 100012},
}

func (c *Company) String() string {
	return fmt.Sprintf("Company %s [ZipCode: %d]\n", c.Name, c.ZipCode)
}
