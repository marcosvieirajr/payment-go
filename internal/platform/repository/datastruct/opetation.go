package datastruct

const OperationTableName = "operation_types"

type Operation struct {
	ID          int    `db:"id"`
	Description string `db:"description"`
}
