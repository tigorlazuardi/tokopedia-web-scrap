package pkg

import (
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

func TestGetCallerInfo(t *testing.T) {
	trace := GetCallerInfo(1)
	assert.Equal(t, "pkg.TestGetCallerInfo", trace.Name)
	assert.Equal(t, true, trace.Line > 10, "Expected line to be higher than this test definition location")
	s := strings.Split(trace.File, "/")
	lastFileName := s[len(s)-1]
	assert.Equal(t, "call_stack_test.go", lastFileName)
}
