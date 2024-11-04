package repo

import "fmt"

type DuplicateNameError struct {
	name string
}

func NewDuplicateNameError(name string) DuplicateNameError {
	return DuplicateNameError{name: name}
}

func (e DuplicateNameError) Error() string {
	return fmt.Sprintf("survivor with name %s already exists", e.name)
}
