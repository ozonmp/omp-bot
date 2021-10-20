package card

type cardType uint8

const (
	DEBIT  cardType = iota
	CREDIT
)

type Card struct {
	OwnerId        uint64
	Number         string
	Cvv            string
	ExpirationDate string
	CardType       cardType
}

func (p *Card) String() string {
	return "Card " + p.Number + " expires " + p.ExpirationDate
}

var allCards = []Card{
	{OwnerId: 1, Number: "1234567843218765", Cvv: "123", ExpirationDate: "1.01.2023", CardType: DEBIT},
	{OwnerId: 1, Number: "1234567843219876", Cvv: "124", ExpirationDate: "1.02.2023", CardType: CREDIT},
	{OwnerId: 2, Number: "2345543267899876", Cvv: "23A", ExpirationDate: "1.03.2022", CardType: DEBIT},
	{OwnerId: 3, Number: "1122334455667788", Cvv: "ABC", ExpirationDate: "1.04.2023", CardType: DEBIT},
	{OwnerId: 4, Number: "9988776655443322", Cvv: "789", ExpirationDate: "5.04.2023", CardType: CREDIT},
}
