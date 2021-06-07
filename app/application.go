package app

import (
	"github.com/gin-gonic/gin"
	"github.com/go-sandbox/bookstore_oauth-api/domain/access_token"
	"github.com/go-sandbox/bookstore_oauth-api/http"
	"github.com/go-sandbox/bookstore_oauth-api/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atService := access_token.NewService(db.NewRepository())
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.Run("127.0.0.1:8000")
}
