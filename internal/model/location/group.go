package location

import "fmt"

type Group struct {
	Id,
	CountOfLocation int
	Type string
}

func (g Group) String() (res string) {
	res = fmt.Sprint(g.Id, g.CountOfLocation, g.Type)
	return
}
