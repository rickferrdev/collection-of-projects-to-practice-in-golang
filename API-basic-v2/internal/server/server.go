package server

import (
	"github.com/gin-gonic/gin"
	"github.com/rickferrdev/api-basic-v2/internal/users"
)

func Server() *gin.Engine {
	server := gin.Default()
	server.SetTrustedProxies([]string{})

	users.Router(server)

	return server
}
