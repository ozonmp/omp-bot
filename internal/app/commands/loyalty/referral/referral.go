package referral

import "fmt"

type Referral struct {
	id        uint64 // поля структуры взял от балды - могу поправить если нужно :)
	firstName string
	lastName  string
	invitedBy uint64
}

func (referral *Referral) String() string {
	return fmt.Sprintf("Id: %d\nfirstName: %s\nlastName: %s\nInvitedBy:%d\n", referral.id, referral.firstName, referral.lastName, referral.invitedBy)
}

var tempReferrals = []Referral{
	{id: 1, firstName: "Ivan", lastName: "Ivanov", invitedBy: 0},
	{id: 2, firstName: "Ivan", lastName: "Smirnov", invitedBy: 1},
	{id: 3, firstName: "Ivan", lastName: "Sidorov", invitedBy: 2},
	{id: 4, firstName: "Ivan", lastName: "Petrov", invitedBy: 3},
	{id: 5, firstName: "Ivan", lastName: "Borzunov", invitedBy: 4},
}
