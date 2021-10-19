package testsubdomain

type TestSubdomain struct {
	Title string
}

func (s TestSubdomain) String() string {
	return "Test domain, test subdomain, " + s.Title
}
