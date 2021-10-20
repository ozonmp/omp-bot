package transaction

import (
	"fmt"
	"time"
)

var allTransactions = []Transaction{
	{ID: 1, BankAccountFrom: "A", BankAccountTo: "B", Amount: 11.1, CreatedAt: time.Now(), Currency: 1},
	{ID: 2, BankAccountFrom: "B", BankAccountTo: "A", Amount: 100.0, CreatedAt: time.Now().AddDate(0, 0, -1), Currency: 1},
	{ID: 3, BankAccountFrom: "A", BankAccountTo: "C", Amount: 555.55, CreatedAt: time.Now().AddDate(0, 0, -2), Currency: 1},
	{ID: 4, BankAccountFrom: "B", BankAccountTo: "C", Amount: 1234.5, CreatedAt: time.Now().AddDate(0, 0, -3), Currency: 1},
	{ID: 5, BankAccountFrom: "C", BankAccountTo: "A", Amount: 9999.9, CreatedAt: time.Now().AddDate(0, 0, -4), Currency: 1},
	{ID: 6, BankAccountFrom: "C", BankAccountTo: "B", Amount: 10.0, CreatedAt: time.Now().AddDate(0, 0, -5), Currency: 1},
}

type Transaction struct {
	ID              uint64
	BankAccountFrom string
	BankAccountTo   string
	CreatedAt       time.Time
	Amount          float64
	Currency        int
}

func (m Transaction) String() string {
	return fmt.Sprintf("Transaction: \nID: %v \nFrom: %v \nTo: %v \nAmount: %v \nCreatedAt: %v\nCurrency: %v", m.ID, m.BankAccountFrom, m.BankAccountTo, m.Amount, m.CreatedAt.Format("2006-01-02 15:04:05"), m.Currency)
}
