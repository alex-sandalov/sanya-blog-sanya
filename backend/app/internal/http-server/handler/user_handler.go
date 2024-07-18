package handler

import "github.com/gin-gonic/gin"

// initUsersHandler initializes the routes for handling users.
//
// It mounts the following routes to the given router group:
// - GET /users: Retrieves all users.
// - GET /users/:user_id: Retrieves a specific user by its ID.
// - POST /users: Creates a new user.
// - DELETE /users/:user_id: Deletes a specific user by its ID.
func (h *Handler) initUsersHandler(router *gin.RouterGroup) {
	router.GET("/", h.handleAllUsersGet)
	router.GET("/:user_id", h.handleUserGet)
	router.POST("/", h.handleUserCreate)
	router.DELETE("/:user_id", h.handleUserDelete)
}

// @Summary Get all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} models.AllUsersGetResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /users [get]
// @Security ApiKeyAuth
func (h *Handler) handleAllUsersGet(c *gin.Context) {
	// TODO
	// ...
}

// @Summary Get user
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} models.UserGetResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /users/{user_id} [get]
// @Security ApiKeyAuth
func (h *Handler) handleUserGet(c *gin.Context) {
	// TODO
	// ...
}

// @Summary Create user
// @Tags users
// @Accept  json
// @Produce  json
// @Param   request body models.UserCreateRequest true "create user request body"
// @Success 200 {object} models.UserCreateResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /users [post]
// @Security ApiKeyAuth
func (h *Handler) handleUserCreate(c *gin.Context) {
	// TODO
	// ...
}

// @Summary Delete user
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} models.UserDeleteResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /users/{user_id} [delete]
// @Security ApiKeyAuth
func (h *Handler) handleUserDelete(c *gin.Context) {
	// TODO
	// ...
}
