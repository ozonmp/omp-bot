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
		"Glen Quagmire",
		"Shibari shop",
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
		"Ultra-Violence",
	},
	{
		8,
		"PS5: Cyberpunk 2077 edition",
		"Sonyboy",
		"pleer.ru",
		"Registered",
	},
	{
		9,
		"Christian Jokes",
		"George Carlin",
		"The God",
		"Unknown",
	},
	{
		10,
		"Hydralisk den",
		"Sarah Kerrigan",
		"Overmind",
		"Declined",
	},
	{
		11,
		"Christian Jokes",
		"George Carlin",
		"The God",
		"Unknown",
	},
	{
		12,
		"CD: The Prodigy - The Fat Of The Land (1997)",
		"Konstantin Kondratenko",
		"Pirate CD Store",
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
