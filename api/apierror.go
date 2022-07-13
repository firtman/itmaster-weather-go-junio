package api

import "fmt"

type ApiError struct {
	code    int
	message string
	extra   interface{}
}

func (error ApiError) Error() string {
	return fmt.Sprintf("%v: %v", error.code, error.message)
}
