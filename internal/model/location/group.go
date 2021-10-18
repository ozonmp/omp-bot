package location

import "fmt"

type Group struct {
	Id,
	CountOfLocation int
	Type string
}

func (g Group) String() (res string) {
	res = fmt.Sprintf("%#+v\n", g)
	return
}
