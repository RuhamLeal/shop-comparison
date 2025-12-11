package sqlite

import (
	"database/sql"
	"project/internal/domain/constants"
	"project/internal/domain/types"
	"strings"
)

func Reason(err error) types.RepositoryErrorReason {
	if err == nil {
		return constants.RepositoryUnknownError
	}

	var sqlErrorMapping = map[error]types.RepositoryErrorReason{
		sql.ErrNoRows: constants.RepositoryNotFoundError,
	}

	if mappedErr, exists := sqlErrorMapping[err]; exists {
		return mappedErr
	}

	errorMappings := map[string]types.RepositoryErrorReason{
		"foreign key":            constants.RepositoryForeignKeyViolationError,
		"Duplicate entry":        constants.RepositoryUniqueConstraintError,
		"unique constraint":      constants.RepositoryUniqueConstraintError,
		"index":                  constants.RepositoryIndexError,
		"syntax error":           constants.RepositoryQuerySyntaxError,
		"timeout":                constants.RepositoryTimeoutError,
		" no rows in result set": constants.RepositoryNotFoundError,
	}

	errMsg := err.Error()

	for key, dbErr := range errorMappings {
		if strings.Contains(errMsg, key) {
			return dbErr
		}
	}

	return constants.RepositoryUnknownError
}
