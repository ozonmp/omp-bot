package card

type EnumCardType uint8

const (
	DEBIT  EnumCardType = iota
	CREDIT
	UNDEF
)

func FromStrToEnum(in string) EnumCardType {
	switch in {
	case "DEBIT": return DEBIT
	case "CREDIT": return CREDIT
	default:
		return UNDEF
	}
}

func FromEnumToStr(in EnumCardType) string {
	switch in {
	case DEBIT: return "DEBIT"
	case CREDIT: return "CREDIT"
	default:
		return  "UNDEF"
	}
}

type Card struct {
	ownerId        uint64
	number         string
	cvv            string
	expirationDate string
	cardType       EnumCardType
}

func (p *Card) String() string {
	return "Card " + p.number + " expires " + p.expirationDate + " type " + FromEnumToStr(p.cardType)
}

var allCards = []Card{
	{ownerId: 1, number: "1234567843218765", cvv: "123", expirationDate: "1.01.2023", cardType: DEBIT},
	{ownerId: 1, number: "1234567843219876", cvv: "124", expirationDate: "1.02.2023", cardType: CREDIT},
	{ownerId: 2, number: "2345543267899876", cvv: "23A", expirationDate: "1.03.2022", cardType: DEBIT},
	{ownerId: 3, number: "1122334455667788", cvv: "ABC", expirationDate: "1.04.2023", cardType: DEBIT},
	{ownerId: 4, number: "9988776655443322", cvv: "789", expirationDate: "5.04.2023", cardType: CREDIT},
}
