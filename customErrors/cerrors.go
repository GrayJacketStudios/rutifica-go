package customerrors

import "fmt"

type EmptyInputError struct {
}

type InvalidInputError struct {
	Input string
}

// Implement the Error method to satisfy the error interface.
func (e *EmptyInputError) Error() string {
	return "El string no puede estar vacio."
}

func (e *InvalidInputError) Error() string {
	return fmt.Sprintf("El input ('%s') solo debe contener numeros.", e.Input)
}
