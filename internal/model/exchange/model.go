package exchange

type Exchange struct {
	Id      uint64
	Package string
	From    string
	To      string
	Status  string
}

func (exchange Exchange) NewExchange() *Exchange {
	return &Exchange{}
}

func (exchange Exchange) String() string {
	return "domain: exchange, subdomain: exchange\n"
}
