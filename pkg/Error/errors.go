package errors

import (
	"fmt"
	"net/http"
)

const validatorOp string = "producer.handlers.validator"

var ValidatorError 			= New(validatorOp, "", 4000, http.StatusBadRequest)
var RabbitConnectionError   = New(validatorOp, "Rabbit Connection Error!", 1000, http.StatusInternalServerError)
var UnknownError   			= New(validatorOp, "Unknown Error!", 1001, http.StatusInternalServerError)


type (
	Error struct {
	Public     PublicError
	StatusCode int
	Internal   error
	Args       interface{}
	}

	PublicError struct {
		Op        string
		Desc      string
		ErrorCode int
	}
)

func (e *Error) Error() string {
	return fmt.Sprintf("Operation: %s, Description: %s, ErrorCode: %d, Internal: %v , Args: %v", e.Public.Op, e.Public.Desc, e.Public.ErrorCode, e.Internal, e.Args)
}


func New(op string, desc string, errorCode int, statusCode int) *Error {
	return &Error{Public: PublicError{
		Op:        op,
		Desc:      desc,
		ErrorCode: errorCode,
	}, StatusCode: statusCode}
}
func (e *Error) WrapDesc(desc string) *Error {
	return &Error{Public: PublicError{
		Op:        e.Public.Op,
		Desc:      desc,
		ErrorCode: e.Public.ErrorCode,
	},
		StatusCode: e.StatusCode,
	}
}


func (e *Error) Wrap(err error, args ...interface{}) *Error {
	if err == nil {
		return nil
	}

	return &Error{Public: PublicError{
		Op:        e.Public.Op,
		Desc:      e.Public.Desc,
		ErrorCode: e.Public.ErrorCode,
	},
		StatusCode: e.StatusCode,
		Internal:   err,
		Args:       args,
	}
}