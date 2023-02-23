package db

import (
	"github.com/Snehasish7080/bookstore_oauth-api/src/domain/access_token"
	"github.com/Snehasish7080/bookstore_oauth-api/src/utils/errors"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func (r *dbRepository) GetById(string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, nil
}

func NewRepository() DbRepository {
	return &dbRepository{}

}
