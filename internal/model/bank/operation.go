package bank

import (
	"fmt"
	"time"
)

type OperationType string

const (
	Salary      OperationType = "Salary"
	Scholarship OperationType = "Scholarship"
	Restaurant  OperationType = "Restaurant"
	Transfer    OperationType = "Transfer"
	Electricity OperationType = "Electricity"
	Tax         OperationType = "Tax"
	Tickets     OperationType = "Tickets"
)

type Operation struct {
	ID            uint64
	OperationType OperationType
	TransactionID uint64
	CreatedAt     time.Time
	ProceedAt     time.Time
	Status        uint64
}

func (op *Operation) String() string {
	return fmt.Sprintf("OperationID: %d\nType: %s\nTransactionID: %d\nCreatedAt: %s\nProceedAt: %s\nStatus: %d",
		op.ID,
		op.OperationType,
		op.TransactionID,
		op.CreatedAt.Format("02.01.2006 15:04:05"),
		op.ProceedAt.Format("02.01.2006 15:04:05"),
		op.Status)
}

func NewOperation(operationType OperationType, transactionID uint64) Operation {
	return Operation{
		OperationType: operationType,
		TransactionID: transactionID,
		CreatedAt:     time.Now(),
		Status:        1,
	}
}
