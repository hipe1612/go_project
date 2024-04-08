package common

type Error struct {
	Code    string
	Message string
}

var Success = &Error{Code: "200", Message: "Success!"}

func (e *Error) Error() string {
	return e.Message
}
