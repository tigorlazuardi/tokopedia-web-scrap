package csvwriter

import (
	"encoding/json"
	"fmt"
)

type Error struct {
	message string
	cause   error
	origin  string
}

func wrapError(cause error, msgFormat string, args ...interface{}) Error {
	return Error{fmt.Sprintf(msgFormat, args...), cause, "csvwriter"}
}

func (e Error) Error() string {
	if e.cause != nil {
		return e.message + ": " + e.cause.Error()
	}
	return e.message + ": " + "<nil>"
}

func (e Error) Unwrap() error {
	return e.cause
}

func (e Error) MarhsalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"message": e.message,
		"origin":  e.origin,
	}
	if e.cause != nil {
		val, _ := json.Marshal(e.cause)
		if string(val) == "{}" {
			m["error"] = e.cause.Error()
		} else {
			m["error"] = e.cause
		}
	}
	return json.Marshal(m)
}
