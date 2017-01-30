package middlewares

import (
	"errors"
	"net/http/httptest"
	"testing"
)

func TestDisplayError(t *testing.T) {
	errObj := errors.New("TestDisplayError")
	errString := "TestDisplayError"

	rr := httptest.NewRecorder()
	DisplayError(rr, errObj, errString, 200)

	expected := `{"data":{"error":"TestDisplayError","message":"TestDisplayError","status":200}}`
	actual := rr.Body.String()

	if expected != actual {
		t.Errorf("DisplayError returned wrong response body: got %v want %v",
			actual, expected)
	}
}
