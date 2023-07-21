package rerror_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cehnhy/rerror"
)

func TestResponseError(t *testing.T) {
	errFileNotFound := errors.New("file not found")
	errTestFileNotFound := fmt.Errorf("test: %w", errFileNotFound)
	codeFileNotFound := "FILE_NOT_FOUND"

	re := rerror.New(http.StatusNotFound, "file %s not found", "rerror.go")
	assert.NotNil(t, re)
	assert.ErrorIs(t, re, rerror.ErrNil)
	assert.Equal(t, "rerror: err nil", re.Error())
	assert.Equal(t, http.StatusNotFound, re.HTTPStatus())
	assert.Equal(t, "", re.Code)
	assert.Equal(t, "file rerror.go not found", re.Message)

	re = re.E(errFileNotFound)
	assert.ErrorIs(t, re, errFileNotFound)
	assert.Equal(t, "file not found", re.Error())
	assert.NotEmpty(t, re.Stack())

	re = re.E(errTestFileNotFound)
	assert.ErrorIs(t, re, errFileNotFound)
	assert.Equal(t, "test: file not found", re.Error())
	assert.NotEmpty(t, re.Stack())

	re = re.C(codeFileNotFound)
	assert.Equal(t, "FILE_NOT_FOUND", re.Code)

	data, err := json.Marshal(re)
	assert.Nil(t, err)
	assert.Equal(t, `{"code":"FILE_NOT_FOUND","message":"file rerror.go not found"}`, string(data))
}
