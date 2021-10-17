package service

type BadRequestError struct {
	s string
}

func NewBadRequestError(text string) error {
	return &BadRequestError{text}
}

func (e *BadRequestError) Error() string {
	return e.s
}
