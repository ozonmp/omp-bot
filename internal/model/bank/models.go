package bank

import (
	"fmt"
	"time"
)

type Operation struct {
	ID            uint64
	OperationType string
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
		op.CreatedAt.Format("01.02.2006 15:04:05"),
		op.ProceedAt.Format("01.02.2006 15:04:05"),
		op.Status)
}

func NewOperation(operationType string, transactionID uint64) Operation {
	return Operation{
		OperationType: operationType,
		TransactionID: transactionID,
		CreatedAt:     time.Now(),
		Status:        1,
	}
}
