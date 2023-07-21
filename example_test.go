package rerror_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/cehnhy/rerror"
)

func ExampleNew() {
	code := "FILE_NOT_FOUND"
	err := errors.New("file not found")
	re := rerror.New(http.StatusNotFound, "file %s not found", "rerror.go").C(code).E(err)
	data, _ := json.Marshal(re)
	fmt.Println(string(data))
	fmt.Println(re.Error())
	fmt.Println(re.HTTPStatus())
	// Output:
	// {"code":"FILE_NOT_FOUND","message":"file rerror.go not found"}
	// file not found
	// 404
}
