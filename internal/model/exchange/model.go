package exchange

var entities = []Exchange {
	{
		1,
		"Book: Scott Meyers: Effective and Modern C++",
		"Evgeniy Elizarov",
		"Labirinth",
		"In progress",
	},
}

type Exchange struct {
	Id      uint64
	Package string
	From    string
	To      string
	Status  string
}

func (exchange Exchange) String() string {
	return "domain: exchange, subdomain: exchange\n"
}
