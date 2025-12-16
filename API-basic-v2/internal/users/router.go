package users

import "github.com/gin-gonic/gin"

func Router(server *gin.Engine) {
	userGroup := server.Group("/users")
	{
		userGroup.POST("/", ControllerCreateUser)
		userGroup.GET("/", ControllerObtainUserByID)
		userGroup.PUT("/", ControllerUpdateUser)
		userGroup.DELETE("/", ControllerDeleteUser)
	}
}
