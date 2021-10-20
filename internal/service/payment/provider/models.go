package provider

var allEntities = []Provider{
	{Id: 1, Name: "Payment provider 1", Description: "More specific information about payment provider 1", IsImplemented: true},
	{Id: 2, Name: "Payment provider 2", Description: "More specific information about payment provider 2", IsImplemented: false},
	{Id: 3, Name: "Payment provider 3", Description: "More specific information about payment provider 3", IsImplemented: true},
	{Id: 4, Name: "Payment provider 4", Description: "More specific information about payment provider 4", IsImplemented: false},
	{Id: 5, Name: "Payment provider 5", Description: "More specific information about payment provider 5", IsImplemented: true},
	{Id: 6, Name: "Payment provider 6", Description: "More specific information about payment provider 6", IsImplemented: false},
	{Id: 7, Name: "Payment provider 7", Description: "More specific information about payment provider 7", IsImplemented: true},
	{Id: 8, Name: "Payment provider 8", Description: "More specific information about payment provider 8", IsImplemented: false},
	{Id: 9, Name: "Payment provider 9", Description: "More specific information about payment provider 9", IsImplemented: true},
	{Id: 10, Name: "Payment provider 10", Description: "More specific information about payment provider 10", IsImplemented: false},
	{Id: 11, Name: "Payment provider 11", Description: "More specific information about payment provider 11", IsImplemented: true},
	{Id: 12, Name: "Payment provider 12", Description: "More specific information about payment provider 12", IsImplemented: false},
}

type Provider struct {
	Id            uint64
	Name          string
	Description   string
	IsImplemented bool
}
