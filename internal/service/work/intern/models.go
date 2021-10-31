package intern

import (
	"fmt"
	"github.com/google/uuid"
)

var idIncrementer uint64

var allInterns = []Intern{
	{Name: "Иван Охлобыстин", UniqueKey: uuid.New(), InternshipID: getNextInternshipId()},
	{Name: "Кристина Асмус", UniqueKey: uuid.New(), InternshipID: getNextInternshipId()},
	{Name: "Аглая Тарасова", UniqueKey: uuid.New(), InternshipID: getNextInternshipId()},
	{Name: "Светлана Пермякова", UniqueKey: uuid.New(), InternshipID: getNextInternshipId()},
	{Name: "Илья Глинников", UniqueKey: uuid.New(), InternshipID: getNextInternshipId()},
	{Name: "Дмитрий Шаракоис", UniqueKey: uuid.New(), InternshipID: getNextInternshipId()},
	{Name: "Александр Ильин", UniqueKey: uuid.New(), InternshipID: getNextInternshipId()},
	{Name: "Один Ланд Байрон", UniqueKey: uuid.New(), InternshipID: getNextInternshipId()},
	{Name: "Вадим Демчог", UniqueKey: uuid.New(), InternshipID: getNextInternshipId()},
	{Name: "Азамат Мусагалиев", UniqueKey: uuid.New(), InternshipID: getNextInternshipId()},
}

type Intern struct {
	Name         string
	UniqueKey    uuid.UUID
	InternshipID uint64
}

func NewIntern(name string) *Intern {
	return &Intern{
		Name: name,
	}
}

func (i Intern) String() string {
	return fmt.Sprintf("Name: %s\nUnique key: %s\nInternship ID: %d", i.Name, i.UniqueKey, i.InternshipID)
}

func getNextInternshipId() uint64 {
	idIncrementer += 1
	return idIncrementer
}
