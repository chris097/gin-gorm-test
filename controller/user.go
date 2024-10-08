package controller

import (
	"net/http"
	"strings"

	"github.com/chris097/gin-gorm-test/config"
	"github.com/chris097/gin-gorm-test/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := []models.User{}
	query := config.DB

	// Retrieve filter parameters from the query string
	name := c.Query("name")
	email := c.Query("email")
	description := c.Query("description")

	// Apply filters based on provided parameters with case-insensitive comparison
	if name != "" {
		query = query.Where("LOWER(name) LIKE ?", strings.ToLower(name))
	}
	if email != "" {
		query = query.Where("LOWER(email) LIKE ?", strings.ToLower(email))
	}
	if description != "" {
		query = query.Where("LOWER(description) LIKE ?", strings.ToLower(description))
	}

	// Execute the query
	if err := query.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	// Return success message along with the data
	c.JSON(http.StatusOK, gin.H{
		"message": "Blogs retrieved successfully",
		"data":    users,
	})
}

// CreateUsers creates a new user
func CreateUsers(c *gin.Context) {
	var user models.User

	// Bind JSON input to the user struct
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Check if all required fields are provided
	if user.Name == "" || user.Email == "" || user.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	// Attempt to create the user in the database
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Return success message along with the created user
	c.JSON(http.StatusOK, gin.H{
		"message": "Blog created successfully",
		"user":    user,
	})
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
	c.JSON(http.StatusOK, gin.H{
		"message": "Blog deleted successfully",
		"user":    user,
	})
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
	c.JSON(http.StatusOK, gin.H{
		"message": "Blog updated successfully",
		"user":    user,
	})
}

// GetUserByID retrieves a user by ID
func GetUserByID(c *gin.Context) {
	var user models.User
	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Returned blog successfully",
		"user":    user,
	})
}
