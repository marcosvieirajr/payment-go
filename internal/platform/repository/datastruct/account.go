package datastruct

const AccountTableName = "accounts"

type Account struct {
	ID             int64  `db:"id"`
	DocumentNumber string `db:"document_number"`
	Audit          Audit
}
