package goerr

import (
	"fmt"
	"testing"
)

func TestServerError(t *testing.T) {
	t.Run("", func(t *testing.T) {
		err := ServerError([]string{})
		fmt.Println(err.Data)
	})
	t.Run("", func(t *testing.T) {
		err := ServerError(nil)
		fmt.Println(err.Data)
	})
}

func TestNoRecordFound(t *testing.T) {
	t.Run("", func(t *testing.T) {
		err := NoRecordFound([]string{})
		fmt.Println(err.Data)
	})
	t.Run("", func(t *testing.T) {
		err := NoRecordFound(nil)
		fmt.Println(err.Data)
	})
}

func TestBadRequest(t *testing.T) {
	t.Run("", func(t *testing.T) {
		err := BadRequest([]string{})
		fmt.Println(err.Data)
	})
	t.Run("", func(t *testing.T) {
		err := BadRequest(nil)
		fmt.Println(err.Data)
	})
}
