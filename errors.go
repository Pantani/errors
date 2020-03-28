package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"runtime"
	"strings"
)

type Params map[string]interface{}

// Error represents a error's specification.
type Error struct {
	Err   error
	meta  map[string]interface{}
	stack []string
}

var (
	_ error = (*Error)(nil)
)

// IsEmpty verify if the error object is empty, without meta and the original error.
// It returns true if is empty.
func (e *Error) IsEmpty() bool {
	return e.meta == nil && e.Err == nil
}

// Error satisfy the built-in interface type is the conventional error interface.
// It returns a JSON object from an error object.
func (e *Error) Error() string {
	r, err := e.MarshalJSON()
	if err != nil {
		return e.Err.Error()
	}
	return string(r)
}

// String satisfy the built-in interface type is the conventional error interface.
// It returns a string from an error object.
func (e *Error) String() string {
	msg := e.Err.Error()
	if len(e.Meta()) > 0 {
		msg = fmt.Sprintf("%s | Meta: %s", msg, e.Meta())
	}
	if len(e.stack) > 0 {
		msg = fmt.Sprintf("%s | Stack: %s", msg, e.stack)
	}
	return msg
}

// SetMeta sets the error's meta data.
// It return the error itself
func (e *Error) SetMeta(data Params) *Error {
	e.meta = data
	return e
}

// Meta gets the error's meta data into a string.
// It return the string representation of error.
func (e *Error) Meta() string {
	r, err := json.Marshal(e.meta)
	if err != nil {
		return ""
	}
	return string(r)
}

// JSON creates a properly formatted JSON
// It return a JSON interface representation of error.
func (e *Error) JSON() interface{} {
	p := Params{}
	if e.meta != nil {
		p["meta"] = e.meta
	}
	if e.Err != nil {
		p["error"] = e.Err.Error()
	}
	if len(e.stack) > 0 {
		p["stack"] = e.stack
	}
	return p
}

// MarshalJSON implements the json.Marshaller interface.
func (e *Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.JSON())
}

// T create a new error with runtime stack trace.
func T(args ...interface{}) *Error {
	e := E(args...)
	for i := 1; i <= 5; i++ {
		_, fn, line, ok := runtime.Caller(i)
		if ok {
			e.stack = append(e.stack, fmt.Sprintf("%s:%d", fn, line))
		}
	}
	return e
}

// E create a new error.
// It returns the new error object
func E(args ...interface{}) *Error {
	e := &Error{meta: make(Params)}
	var message []string
	for _, arg := range args {
		switch arg := arg.(type) {
		case nil:
			continue
		case string:
			message = append(message, arg)
		case *Error:
			message = append([]string{arg.Err.Error()}, message...)
			appendMap(e.meta, arg.meta)
		case error:
			message = append([]string{arg.Error()}, message...)
		case Params:
			appendMap(e.meta, arg)
		case map[string]interface{}:
			appendMap(e.meta, arg)
		default:
			continue
		}
	}
	if len(message) > 0 {
		msg := strings.Join(message[:], ": ")
		e.Err = errors.New(msg)
	}
	return e
}

// appendMap append the new map into a old map
func appendMap(root map[string]interface{}, tmp map[string]interface{}) {
	for k, v := range tmp {
		root[k] = v
	}
}
