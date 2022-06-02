package app

type errorInfo string

func (e errorInfo) Error() string {
	return string(e)
}

const ErrInvalidOperationType = errorInfo("operation type is invalid")
const ErrDocumentNumberIsInvalid = errorInfo("document number is invalid")
const ErrInvalidAmount = errorInfo("operation's amount is invalid")
const ErrAccountNotFound = errorInfo("account not found")
const ErrAccountAlreadyExists = errorInfo("there is already an account with the given document number")
const ErrInvalidTransaction = errorInfo("transaction is invalid")
