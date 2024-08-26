package controller

import (
	"net/http"

	"github.com/chris097/gin-gorm-test/config"
	"github.com/chris097/gin-gorm-test/models"
	"github.com/gin-gonic/gin"
)

// GetUsers retrieves all users
func GetUsers(c *gin.Context) {
	users := []models.User{}
	if err := config.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}
	c.JSON(http.StatusOK, &users)
}

// CreateUsers creates a new user
func CreateUsers(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusOK, &user)
}

// DeleteUser deletes a user by ID
func DeleteUser(c *gin.Context) {
	var user models.User
	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}
	if err := config.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete blog"})
		return
	}
	c.JSON(http.StatusOK, &user)
}

// UpdateUser updates a user by ID
func UpdateUser(c *gin.Context) {
	var user models.User
	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update blog"})
		return
	}
	c.JSON(http.StatusOK, &user)
}

// GetUserByID retrieves a user by ID
func GetUserByID(c *gin.Context) {
	var user models.User
	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}
	c.JSON(http.StatusOK, &user)
}

// package controller

// import (
// 	"net/http"

// 	"github.com/chris097/gin-gorm-test/config"
// 	"github.com/chris097/gin-gorm-test/models"
// 	"github.com/gin-gonic/gin"
// )

// func GetUsers(c *gin.Context) {
// 	users := []models.User{}
// 	config.DB.Find(&users)
// 	c.JSON(http.StatusOK, &users)
// }
// func CreateUsers(c *gin.Context) {
// 	var user models.User
// 	c.BindJSON(&user)
// 	config.DB.Create(&user)
// 	c.JSON(http.StatusOK, &user)
// }
// func DeleteUser(c *gin.Context) {
// 	var user models.User
// 	config.DB.Where("id = ?", c.Param("id")).Delete(&user)
// 	c.JSON(http.StatusOK, &user)
// }
// func UpdateUser(c *gin.Context) {
// 	var user models.User
// 	config.DB.Where("id = ?", c.Param("id")).First(&user)
// 	c.BindJSON(&user)
// 	c.JSON(http.StatusOK, &user)
// }

// // GetUserByID retrieves a user by ID
// func GetUserByID(c *gin.Context) {
// 	var user models.User
// 	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, &user)
// }
