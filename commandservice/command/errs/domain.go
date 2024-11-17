package errs

type DomainError struct {
	Message string
}

func (e DomainError) Error() string {
	return e.Message
}

func NewDomainError(message string) *DomainError {
	return &DomainError{Message: message}
}
