//go:generate stringer -type=Type
package errors

import "fmt"

type Type int

const (
	badRequest Type = iota
	dbConnErr
	unauthorized
	notFound
)

type ErrorType struct {
	Type    Type
	Code    int
	Message string
	Err     error
}

func (e ErrorType) String() string {

	return [...]string{"badRequest", "unauthorized", "dbConnErr", "notFound"}[e.Type]
}
func (e ErrorType) Int() int {
	return int(e.Type)
}
func (e ErrorType) Error() string {
	return fmt.Sprintf("error type was %v, code: %d, message: %v", e.String(), e.Type, e.Message)
}

func TypeFromName(typeStr string) (Type, error) {
	switch typeStr {
	case "badRequest":
		return badRequest, nil
	case "unauthorized":
		return unauthorized, nil
	case "dbConnErr":
		return dbConnErr, nil
	case "notFound":
		return notFound, nil
	default:
		return -1, fmt.Errorf("%v is not a valid error type", typeStr)
	}
}
