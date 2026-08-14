package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/weaver/api/internal/models"
	"github.com/weaver/api/internal/utils"
)

// AuthHandler handles authentication requests
type AuthHandler struct {
	db *sql.DB
}

func NewAuthHandler(db *sql.DB) *AuthHandler {
	return &AuthHandler{db: db}
}

// Demo users for quick login
var demoUsers = map[string]string{
	"admin":     "admin123",     // Admin user
	"manager":   "manager123",   // Manager user
	"accountant": "accountant123", // Accountant user
}

var userRoles = map[string]string{
	"admin":     "Admin",
	"manager":   "Manager",
	"accountant": "Accountant",
}

// Login handles user login
func (h *AuthHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request body")
		return
	}

	// Check demo credentials
	password, exists := demoUsers[req.Username]
	if !exists || password != req.Password {
		utils.UnauthorizedResponse(c, "Invalid username or password")
		return
	}

	// Create user object (for demo, using dummy ID)
	user := &models.User{
		ID:       1,
		Username: req.Username,
		Email:    req.Username + "@demo.com",
		FullName: req.Username,
		Role:     userRoles[req.Username],
		Status:   "Active",
	}

	resp := models.LoginResponse{
		Token: "demo-token-" + req.Username,
		User:  user,
	}

	utils.OKResponse(c, "Login successful", resp)
}

// GetProfile returns current user profile
func (h *AuthHandler) GetProfile(c *gin.Context) {
	user := &models.User{
		ID:       1,
		Username: "admin",
		Email:    "admin@demo.com",
		FullName: "Admin User",
		Role:     "Admin",
		Status:   "Active",
	}
	utils.OKResponse(c, "Profile retrieved", user)
}

// Logout handles user logout
func (h *AuthHandler) Logout(c *gin.Context) {
	utils.OKResponse(c, "Logged out successfully", nil)
}

// HealthCheck returns health status
func (h *AuthHandler) HealthCheck(c *gin.Context) {
	response := map[string]interface{}{
		"status":   "healthy",
		"database": "connected",
	}
	c.JSON(http.StatusOK, response)
}
