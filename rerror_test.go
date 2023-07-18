package rerror

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponseError(t *testing.T) {
	errFileNotFound := errors.New("file not found")

	err := New(http.StatusNotFound, "file %s not found", "rerror.go")
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, errNil)
	assert.Equal(t, "nil error", err.Error())
	assert.Equal(t, http.StatusNotFound, err.HTTPStatus())
	assert.Equal(t, "", err.Code)
	assert.Equal(t, "file rerror.go not found", err.Message)

	err = err.E(errFileNotFound).C("FILE_NOT_FOUND")
	assert.ErrorIs(t, err, errFileNotFound)
	assert.Equal(t, "file not found", err.Error())
	assert.NotEmpty(t, err.Stack())
	assert.Equal(t, "FILE_NOT_FOUND", err.Code)
}
