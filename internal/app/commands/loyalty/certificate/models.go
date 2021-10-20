package certificate

const pageSize = 5

type CertificateData struct {
	Id			uint64		`json:"id"`
	SellerTitle string		`json:"seller_title"`
	Amount 		uint64		`json:"amount"`
	ExpireDate	int64		`json:"expire_date"`
}
