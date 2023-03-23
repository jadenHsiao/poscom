package src

type PoscomError struct {
	Code      string
	Message   string
	ErrorList map[string]string
}

func NewError(code, message string) *PoscomError {
	return &PoscomError{
		Code:    code,
		Message: message,
	}
}
