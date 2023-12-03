package utils

import (
	"errors"
	"fmt"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"net/http"
)

func ParseDBError(err error, ph string) (int, string) {
	var pgError *pgconn.PgError
	if errors.As(err, &pgError) {
		switch pgError.Code {
		case pgerrcode.ForeignKeyViolation:
			return http.StatusBadRequest, fmt.Sprintf("Failed to create %s", ph)
		case pgerrcode.UniqueViolation:
			return http.StatusBadRequest, fmt.Sprintf("%s with this Info already exists", ph)
		case pgerrcode.CaseNotFound:
			return http.StatusNotFound, fmt.Sprintf("%s doesn't exist", ph)
		case pgerrcode.UndefinedTable:
			return http.StatusInternalServerError, fmt.Sprintf("%s table doesn't exist", ph)
		case pgerrcode.InvalidTextRepresentation:
			return http.StatusBadRequest, fmt.Sprintf("%s invalid type", ph)
		case pgerrcode.UndefinedColumn:
			return http.StatusBadRequest, fmt.Sprintf("invalid column for table %s", ph)
		default:
			return http.StatusBadRequest, fmt.Sprintf("Something went wrong with %s", ph)
		}
	}

	return http.StatusBadRequest, "Something went wrong"
}
