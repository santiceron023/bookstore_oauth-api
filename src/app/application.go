package app

import (
	"github.com/gin-gonic/gin"
	"github.com/santiceron023/bookstore_oauth-api/src/clients/cassandra"
	"github.com/santiceron023/bookstore_oauth-api/src/domain/access_token"
	"github.com/santiceron023/bookstore_oauth-api/src/http"
	"github.com/santiceron023/bookstore_oauth-api/src/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session, dbErr := cassandra.GetSesion()
	if dbErr != nil {
		panic(dbErr)
	}
	session.Close()
	atService := access_token.NewService(db.NewRepository())
	atHandler := http.NewHandler(atService)

	router.GET("oauth/access_token/:accessTokenId", atHandler.GetById)
	router.Run(":8080")
}
