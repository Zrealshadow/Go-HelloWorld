package errcode

import "fmt"

type Error struct {
	code int
	msg  string
}

var _codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := _codes[code]; ok {
		if _, ok := _codes[code]; ok {
			panic(fmt.Sprintf("Error code %d Exist, Please change", code))
		}

	}
	_codes[code] = msg
	return &Error{code: code, msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error code %d Error Msg %s", e.code, e.msg)
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}
