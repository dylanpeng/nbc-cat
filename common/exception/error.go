package exception

import (
	"cat/common"
	"fmt"
)

type Exception struct {
	code    int
	message string
}

func (e *Exception) GetCode() int {
	return e.code
}

func (e *Exception) GetMessage() string {
	return e.message
}

func (e *Exception) String() string {
	return fmt.Sprintf("(%d) %s", e.code, e.message)
}

func Desc(code int) string {
	if e, ok := Desces[code]; ok {
		return e
	}

	return "server internal error"
}

func New(code int, args ...interface{}) *Exception {
	message := Desc(code)

	if len(args) > 0 {
		if err, ok := args[0].(error); ok {
			common.Logger.Infof("Error: %s | Args: %+v", err, args[1:])
		}
	}

	return &Exception{code: code, message: message}
}
