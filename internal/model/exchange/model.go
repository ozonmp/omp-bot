package exchange

var entities = []Exchange {
	{
		1,
		"Book: Scott Meyers: Effective and Modern C++",
		"Evgeniy Elizarov",
		"Labirinth",
		"In progress",
	},
	{
		2,
		"Ticket: Moscow - Vladivostok - Tokyo",
		"Tatiana Kuznetsova",
		"Aeroflot",
		"In progress",
	},
	{
		3,
		"Rope: 10 meters",
		"Glen",
		"Dirty shop",
		"Registred",
	},
	{
		4,
		"Nintendo Switch",
		"Yuri Petrov",
		"Nintendo",
		"In progress",
	},
	{
		5,
		"Item",
		"Person",
		"Organization",
		"Delivered",
	},
	{
		6,
		"Zygote",
		"X",
		"Y",
		"Aborted",
	},
	{
		7,
		"BFG9000",
		"Doomguy",
		"Hell on Earth",
		"Nightmare",
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
