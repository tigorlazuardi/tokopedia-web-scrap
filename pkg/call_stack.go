package pkg

import (
	"runtime"
	"strings"
)

type CallerTrace struct {
	Name string `json:"name"`
	File string `json:"file"`
	Line int    `json:"line"`
}

func GetCallerInfo(depth int) CallerTrace {
	stackIdx, file, line, ok := runtime.Caller(depth)

	trace := CallerTrace{
		Name: "Unknown Function",
		File: file,
		Line: line,
	}
	details := runtime.FuncForPC(stackIdx)
	if ok && details != nil {
		name := details.Name()
		s := strings.Split(name, "/")
		trace.Name = s[len(s)-1]
	}
	return trace
}
