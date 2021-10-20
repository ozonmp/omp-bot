package transaction

import (
	"testing"
	"time"
)

var mockTransactions = []Transaction{
	{ID: 1, BankAccountFrom: "A", BankAccountTo: "B", Amount: 1.0, CreatedAt: time.Now(), Currency: 1},
	{ID: 2, BankAccountFrom: "B", BankAccountTo: "A", Amount: 2.0, CreatedAt: time.Now().AddDate(0, 0, -1), Currency: 1},
	{ID: 3, BankAccountFrom: "A", BankAccountTo: "C", Amount: 3.0, CreatedAt: time.Now().AddDate(0, 0, -2), Currency: 1},
}

func upFixtures() {
	allTransactions = mockTransactions
}

func TestDummyService_List(t *testing.T) {
	upFixtures()

	s := NewDummyTransactionService()
	tx, err := s.List(0, 3)
	if err != nil {
		t.Fatalf("excepted no err")
	}
	if len(tx) != 3 {
		t.Fatalf("expected 3 txs, got %v", len(tx))
	}

	_, err = s.List(10000, 1)
	if err == nil {
		t.Fatalf("expected out of bound")
	}

	tx, _ = s.List(1, 1)

	if tx[0] != mockTransactions[1] {
		t.Fatalf("got wrong tx")
	}
}

func TestDummyService_Remove(t *testing.T) {
	upFixtures()

	s := NewDummyTransactionService()
	result, _ := s.Remove(1)
	if !result {
		t.Fatalf("excepted success")
	}

	tx, _ := s.List(0, 3)
	if len(tx) != 2 {
		t.Fatalf("expected 2 txs, got %v", len(tx))
	}

	result, _ = s.Remove(111)
	if result {
		t.Fatalf("excepted error")
	}
}

func TestDummyService_Describe(t *testing.T) {
	upFixtures()

	s := NewDummyTransactionService()
	tx, _ := s.Describe(3)
	if tx.ID != 3 {
		t.Fatalf("expected tx with ID 3, got %v", tx.ID)
	}

	_, err := s.Describe(333)
	if err == nil {
		t.Fatalf("excepted not exists error")
	}
}
