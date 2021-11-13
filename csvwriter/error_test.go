package csvwriter

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_wrapError(t *testing.T) {
	err := errors.New("test")
	errW := wrapError(err, "abcd%s", "efgh")

	assert.Equal(t, err, errW.cause)
	assert.Equal(t, "csvwriter.Test_wrapError", errW.stack.Name)
}
