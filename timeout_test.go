package middlewares

import (
	"errors"
	"testing"
)

// implements the net error interface and represents a network error.
type timeoutError struct{}

func (e *timeoutError) Error() string   { return "i/o timeout" }
func (e *timeoutError) Timeout() bool   { return true }
func (e *timeoutError) Temporary() bool { return true }

func TestHasTimedOut(t *testing.T) {

	suite := []struct {
		err         error
		description string
		expected    bool
	}{
		// #1
		{
			err:         errors.New("use of closed network connection"),
			description: "Error text contains use of closed network connection",
			expected:    true,
		},
		// #2
		{
			err:         &timeoutError{},
			description: "net.Error returns true",
			expected:    true,
		},
	}

	for _, test := range suite {

		if actual := HasTimedOut(test.err); actual != test.expected {
			t.Errorf("Failed: %s", test.description)
		}
	}
}
