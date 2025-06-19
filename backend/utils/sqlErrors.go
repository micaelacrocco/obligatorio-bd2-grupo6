package utils

import (
	"errors"
	"github.com/VividCortex/mysqlerr"
	"github.com/go-sql-driver/mysql"
)

var ErrForeignKeyNotFound = errors.New("duplicated FK")

func ForeignKeyNotFoundError(err error) error {
	var driverErr *mysql.MySQLError
	if errors.As(err, &driverErr) {
		if driverErr.Number == mysqlerr.ER_NO_REFERENCED_ROW {
			return ErrForeignKeyNotFound
		}
	}
	return err
}
