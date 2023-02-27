package db

import (
	"github.com/Snehasish7080/bookstore_oauth-api/src/cllient/cassandra"
	"github.com/Snehasish7080/bookstore_oauth-api/src/domain/access_token"
	"github.com/Snehasish7080/bookstore_oauth-api/src/utils/errors"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func (r *dbRepository) GetById(string) (*access_token.AccessToken, *errors.RestErr) {

	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//Todo: implement get access token form cassandra
	return nil, errors.NewInternalServerError("database connection not implemented")
}

func NewRepository() DbRepository {
	return &dbRepository{}
}
