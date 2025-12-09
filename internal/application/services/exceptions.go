package services

import (
	"project/internal/domain/constants"
	exceptions "project/internal/domain/exception"
)

func GetStatusCodeFromError(err error) int {
	if err == nil {
		return 500
	}

	switch err := err.(type) {
	case *exceptions.BaseRepository:
		if err.Reason == constants.RepositoryNotFoundError {
			return 404
		}

		return 500
	case *exceptions.BaseUsecase:
		return err.StatusCode
	default:
		return 500
	}
}
