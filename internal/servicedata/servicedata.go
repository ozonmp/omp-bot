package servicedata

const (
	NewoperationData  = iota
	EditoperationData
)
type operationData struct{
	ProductID uint64
	OperationType int
}
var EditedChat = make(map[int64]operationData, 100)

func GetOperationData(productID uint64, operationType int) *operationData{
	return &operationData{
		ProductID: productID,
		OperationType: operationType,
	}
}
