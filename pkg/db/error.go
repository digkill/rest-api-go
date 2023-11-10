package db

import (
	"database/sql"
	"errors"
	"fmt"
)

func HandleError(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return ErrObjectNotFound{}
	}
	return err
}

type ErrObjectNotFound struct{}

func (ErrObjectNotFound) Error() string {
	return "object not found"
}
func (ErrObjectNotFound) Unwrap() error {
	return fmt.Errorf("object not found")
}
