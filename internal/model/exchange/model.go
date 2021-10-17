package exchange

type Exchange struct {}

func (xchg Exchange) NewExchange() *Exchange {
	return &Exchange{}
}

func (xchg Exchange) String() string {
	return "domain: exchange, subdomain: exchange\n"
}
