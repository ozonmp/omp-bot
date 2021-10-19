package verification

type Verification struct {
	Title string
}

func (s Verification) String() string {
	return "Test domain, test subdomain, " + s.Title
}
