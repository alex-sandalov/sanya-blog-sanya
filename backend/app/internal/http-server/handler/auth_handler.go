package handler

import (
	_ "blog-app/app/internal/models"

	"github.com/gin-gonic/gin"
)

// initAuthHandlers initializes the routes for handling authentication.
//
// It mounts the following routes to the given router group:
// - POST /auth/login: Authenticates a user.
// - POST /auth/register: Registers a new user.
// - POST /auth/refresh: Refreshes a user's access token.
func (h *Handler) initAuthHandlers(router *gin.RouterGroup) {
	router.POST("/login", h.handleLogin)
	router.POST("/register", h.handleRegister)
	router.POST("/refresh", h.handleRefresh)
}

// @Summary Login
// @Description Authenticates a user.
// @Tags auth
// @Accept  json
// @Param   request body models.LoginRequest true "login request body"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /auth/login [post]
func (h *Handler) handleLogin(c *gin.Context) {
	// TODO
	// ...
}

// @Summary Register
// @Description Registers a new user.
// @Tags auth
// @Accept  json
// @Param   request body models.RegistersRequest true "register request body"
// @Success 200 {object} models.RegistersResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /auth/register [post]
func (h *Handler) handleRegister(c *gin.Context) {
	// TODO
	// ...
}

// @Summary Refresh
// @Description Refreshes a user's access token.
// @Tags auth
// @Accept  json
// @Param   request body models.RefreshRequest true "refresh request body"
// @Success 200 {object} models.RefreshResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /auth/refresh [post]
// @Security ApiKeyAuth
func (h *Handler) handleRefresh(c *gin.Context) {
	// TODO
	// ...
}
