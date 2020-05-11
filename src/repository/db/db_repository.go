package db

import (
	"github.com/gocql/gocql"
	"github.com/santiceron023/bookstore_oauth-api/src/clients/cassandra"
	"github.com/santiceron023/bookstore_oauth-api/src/domain/access_token"
	"github.com/santiceron023/bookstore_oauth-api/src/utils/rest_errors"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	queryUpdateExpires     = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

type dbRepository struct {
}

//porque ssi ya está en service!
type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *rest_errors.RestError)
	Create(access_token.AccessToken) rest_errors.RestError
	UpdateExpirationTime(access_token.AccessToken) rest_errors.RestError
}

func (d *dbRepository) GetById(id string) (*access_token.AccessToken, *rest_errors.RestError) {
	session, err := cassandra.GetSesion()
	if err != nil {
		return nil, rest_errors.NewInternalServerError(err.Error())
	}
	defer session.Close()

	var result access_token.AccessToken

	if err := session.Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires,
	); err != nil{
		if err == gocql.ErrNotFound {
			return nil, rest_errors.NewNotFoundError("no access token found with given id")
		}
		return nil, rest_errors.NewInternalServerError("error when trying to get current id")
	}

	return &result,nil
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

//docker run --name some-cassandra -p 9042:9042 -v /my/own/datadir:/var/lib/cassandra -d cassandra:latest
//docker exec -it some-cassandra bash
//cqlsh
//describe keyspaces
//CREATE KEYSPACE oauth WITH REPLICATION = {'class':'SimpleStrategy','replication_factor':1};
//USE oauth;
//describe tables;
//CREATE TABLE access_tokens( access_token varchar PRIMARY KEY, user_id bigint,client_id bigint, expires bigint);
//
//NEVER!!
//select * from access_tokens where user_id=3;
//select * from access_tokens where access_token='dfr';
