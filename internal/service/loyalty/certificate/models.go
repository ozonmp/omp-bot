package certificate

import "time"

var allEntities = []Сertificate{
	{
		Id: 0,
		SellerTitle: "Seller 0",
		Amount: 3000,
		ExpireDate: time.Date(2021, 11, 18, 23, 59, 59, 0, time.Local),
	},
	{
		Id: 1,
		SellerTitle: "Seller 1",
		Amount: 3000,
		ExpireDate: time.Date(2021, 11, 19, 23, 59, 59, 0, time.Local),
	},
	{
		Id: 2,
		SellerTitle: "Seller 2",
		Amount: 3000,
		ExpireDate: time.Date(2021, 11, 20, 23, 59, 59, 0, time.Local),
	},
	{
		Id: 3,
		SellerTitle: "Seller 3",
		Amount: 3000,
		ExpireDate: time.Date(2021, 11, 21, 23, 59, 59, 0, time.Local),
	},
	{
		Id: 4,
		SellerTitle: "Seller 4",
		Amount: 3000,
		ExpireDate: time.Date(2021, 11, 22, 23, 59, 59, 0, time.Local),
	},
	{
		Id: 5,
		SellerTitle: "Seller 5",
		Amount: 3000,
		ExpireDate: time.Date(2021, 11, 23, 23, 59, 59, 0, time.Local),
	},
	{
		Id: 6,
		SellerTitle: "Seller 6",
		Amount: 3000,
		ExpireDate: time.Date(2021, 11, 24, 23, 59, 59, 0, time.Local),
	},
	{
		Id: 7,
		SellerTitle: "Seller 7",
		Amount: 3000,
		ExpireDate: time.Date(2021, 12, 25, 23, 59, 59, 0, time.Local),
	},
	{
		Id: 8,
		SellerTitle: "Seller 8",
		Amount: 3000,
		ExpireDate: time.Date(2021, 12, 26, 23, 59, 59, 0, time.Local),
	},
	{
		Id: 9,
		SellerTitle: "Seller 9",
		Amount: 3000,
		ExpireDate: time.Date(2021, 12, 27, 23, 59, 59, 0, time.Local),
	},
}

type Сertificate struct {
	Id int64
	SellerTitle string
	Amount int64
	ExpireDate time.Time
}

