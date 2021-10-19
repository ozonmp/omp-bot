package lead

import "fmt"

var allEntities = map[uint64]Lead{
	1: {ID: 1, FirstName: "Tom", LastName: "Texeira", Project: "Destiny Realty Solutions"},
	2: {ID: 2, FirstName: "Leigh", LastName: "Hudson", Project: "Buehler Foods"},
	3: {ID: 3, FirstName: "Mitchell", LastName: "Hildebrant", Project: "Happy Bear Investment"},
	4: {ID: 4, FirstName: "Paul", LastName: "Roldan", Project: "Reliable Guidance"},
	5: {ID: 5, FirstName: "Esther", LastName: "Sturdivant", Project: "Sistemos"},
}

var idSeq = uint64(len(allEntities))

type Lead struct {
	ID        uint64
	FirstName string
	LastName  string
	Project   string
}

func (l *Lead) String() string {
	return fmt.Sprintf(`Lead #%d %s %s, %s`, l.ID, l.FirstName, l.LastName, l.Project)
}
