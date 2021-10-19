package exchange

var entities = []Exchange{}

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
