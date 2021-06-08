package db

import (
	"github.com/go-sandbox/bookstore_oauth-api/clients/cassandra"
	"github.com/go-sandbox/bookstore_oauth-api/domain/access_token"
	"github.com/go-sandbox/bookstore_oauth-api/utils/errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

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

	// TODO: implement get access token from CassandraDB
	return nil, errors.NewInternalServerError("database connection not implemented by yet")
}
