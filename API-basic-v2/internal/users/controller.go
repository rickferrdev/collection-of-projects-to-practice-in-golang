package users

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ControllerCreateUser(c *gin.Context) {
	var user CreateUserDTO

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	user = ServiceCreateUser(user)

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "data": user})
}

func ControllerObtainUserByID(c *gin.Context) {
	var user ObtainUserDTO
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	response := ServiceObtainUserByID(user.ID)

	c.JSON(http.StatusOK, gin.H{"message": "User retrieved successfully", "data": response})
}

func ControllerUpdateUser(c *gin.Context) {
	var user UpdateUserDTO
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	user = ServiceUpdateUser(user)

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "data": user})
}

func ControllerDeleteUser(c *gin.Context) {
	var user DeleteUserDTO
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	success := ServiceDeleteUser(user.ID)
	if !success {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully", "data": user})
}
