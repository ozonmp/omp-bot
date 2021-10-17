package exchange

type Exchange struct {}

func (exchange Exchange) NewExchange() *Exchange {
	return &Exchange{}
}

func (exchange Exchange) String() string {
	return "domain: exchange, subdomain: exchange\n"
}
