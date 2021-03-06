package csvwriter

import (
	"encoding/json"
	"fmt"

	"github.com/tigorlazuardi/tokopedia-web-scrap/pkg"
)

type Error struct {
	message  string
	cause    error
	origin   string
	location pkg.CallerTrace
}

func wrapError(cause error, msgFormat string, args ...interface{}) Error {
	trace := pkg.GetCallerInfo(2)
	return Error{fmt.Sprintf(msgFormat, args...), cause, "csvwriter", trace}
}

// implementes error interface
func (e Error) Error() string {
	if e.cause != nil {
		return e.message + ": " + e.cause.Error()
	}
	return e.message + ": " + "<nil>"
}

// implements the silent errors.Unwrap interface
func (e Error) Unwrap() error {
	return e.cause
}

// implements json.Marshaler interface.
// because we want to know what the actual error is even after marshaled.
func (e Error) MarhsalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"message":  e.message,
		"origin":   e.origin,
		"location": e.location,
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
