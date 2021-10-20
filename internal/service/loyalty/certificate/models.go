package certificate

import (
	"strconv"
	"time"
)

var allEntities = []Certificate{
	{
		Id: 1,
		SellerTitle: "Seller 1",
		Amount: 1000,
		ExpireDate: time.Date(2021, 11, 18, 23, 59, 59, 0, time.Local).Unix(),
	},
	{
		Id: 2,
		SellerTitle: "Seller 2",
		Amount: 5000,
		ExpireDate: time.Date(2021, 11, 19, 23, 59, 59, 0, time.Local).Unix(),
	},
	{
		Id: 3,
		SellerTitle: "Seller 3",
		Amount: 3000,
		ExpireDate: time.Date(2021, 11, 20, 23, 59, 59, 0, time.Local).Unix(),
	},
	{
		Id: 4,
		SellerTitle: "Seller 4",
		Amount: 5000,
		ExpireDate: time.Date(2021, 11, 21, 23, 59, 59, 0, time.Local).Unix(),
	},
	{
		Id: 5,
		SellerTitle: "Seller 5",
		Amount: 3000,
		ExpireDate: time.Date(2021, 11, 22, 23, 59, 59, 0, time.Local).Unix(),
	},
	{
		Id: 6,
		SellerTitle: "Seller 6",
		Amount: 7000,
		ExpireDate: time.Date(2021, 11, 23, 23, 59, 59, 0, time.Local).Unix(),
	},
	{
		Id: 7,
		SellerTitle: "Seller 7",
		Amount: 1000,
		ExpireDate: time.Date(2021, 11, 24, 23, 59, 59, 0, time.Local).Unix(),
	},
	{
		Id: 8,
		SellerTitle: "Seller 8",
		Amount: 3000,
		ExpireDate: time.Date(2021, 12, 25, 23, 59, 59, 0, time.Local).Unix(),
	},
	{
		Id: 9,
		SellerTitle: "Seller 9",
		Amount: 5000,
		ExpireDate: time.Date(2021, 12, 26, 23, 59, 59, 0, time.Local).Unix(),
	},
	{
		Id: 10,
		SellerTitle: "Seller 10",
		Amount: 10000,
		ExpireDate: time.Date(2021, 12, 27, 23, 59, 59, 0, time.Local).Unix(),
	},
	{
		Id: 11,
		SellerTitle: "Seller 11",
		Amount: 10000,
		ExpireDate: time.Date(2021, 12, 10, 23, 59, 59, 0, time.Local).Unix(),
	},
	{
		Id: 12,
		SellerTitle: "Seller 12",
		Amount: 15000,
		ExpireDate: time.Date(2022, 12, 15, 23, 59, 59, 0, time.Local).Unix(),
	},
}

type Certificate struct {
	Id uint64
	SellerTitle string
	Amount uint64
	ExpireDate int64
}

func (c *Certificate) String() string {
	return "Id: " + strconv.Itoa(int(c.Id)) + "\n" +
		"SellerTitle: " + c.SellerTitle + "\n" +
		"Amount: " + strconv.Itoa(int(c.Amount)) + "\n" +
		"ExpireDate: " + strconv.Itoa(int(c.ExpireDate)) + "\n"
}

