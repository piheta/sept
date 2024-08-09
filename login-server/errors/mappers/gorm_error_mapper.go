package errors

import (
	"errors"

	weberrors "github.com/piheta/sept/login-server/errors"

	"gorm.io/gorm"
)

func MapGormError(object string, resultError error) *weberrors.WebError {

	if errors.Is(resultError, gorm.ErrRecordNotFound) {
		return weberrors.NewError(404, object+" not found")
	}

	if errors.Is(resultError, gorm.ErrDuplicatedKey) {
		return weberrors.NewError(409, object+" with this id already exists")
	}

	if resultError != nil {
		return weberrors.NewError(500, resultError.Error())
	}

	return nil
}
