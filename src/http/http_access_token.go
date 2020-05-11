package http

import (
	"github.com/gin-gonic/gin"
	"github.com/santiceron023/bookstore_oauth-api/src/domain/access_token"
	"net/http"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
}
type accessTokenHandler struct {
	service access_token.Service
}

func (handler *accessTokenHandler) GetById(context *gin.Context) {
	accessTokenId, err := handler.service.GetById(
			context.Param("accessTokenId"))
	if err != nil {
		context.JSON(err.Code, err)
		return
	}
	context.JSON(http.StatusOK, accessTokenId)
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}
