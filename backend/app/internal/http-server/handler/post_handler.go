package handler

import "github.com/gin-gonic/gin"

// initPostsHandler initializes the routes for handling posts.
//
// It mounts the following routes to the given router group:
// - GET /posts: Retrieves all posts.
// - GET /posts/:post_id: Retrieves a specific post by its ID.
// - POST /posts: Creates a new post.
// - PATCH /posts/:post_id: Updates a specific post by its ID.
// - DELETE /posts/:post_id: Deletes a specific post by its ID.
func (h *Handler) initPostsHandler(router *gin.RouterGroup) {
	router.GET("/", h.handleAllPostsGet)
	router.GET("/:post_id", h.handlePostGet)
	router.POST("/", h.handlePostCreate)
	router.PATCH("/:post_id", h.handlePostUpdate)
	router.DELETE("/:post_id", h.handlePostDelete)
}

// @Summary Get all posts
// @Tags posts
// @Accept  json
// @Produce  json
// @Param   request query models.AllPostsGetRequest true "all posts request body"
// @Success 200 {object} models.AllPostsGetResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /posts [get]
// @Security ApiKeyAuth
func (h *Handler) handleAllPostsGet(c *gin.Context) {
	// TODO
	// ...
}

// @Summary Get post
// @Tags posts
// @Accept  json
// @Produce  json
// @Success 200 {object} models.PostGetResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /posts/{post_id} [get]
// @Security ApiKeyAuth
func (h *Handler) handlePostGet(c *gin.Context) {
	// TODO
	// ...
}

// @Summary Create post
// @Tags posts
// @Accept  json
// @Produce  json
// @Param   request body models.PostCreateRequest true "create post request body"
// @Success 200 {object} models.PostCreateResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /posts [post]
// @Security ApiKeyAuth
func (h *Handler) handlePostCreate(c *gin.Context) {
	// TODO
	// ...
}

// @Summary Update post
// @Tags posts
// @Accept  json
// @Produce  json
// @Param   request body models.PostUpdateRequest true "update post request body"
// @Success 200 {object} models.PostUpdateResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /posts [patch]
// @Security ApiKeyAuth
func (h *Handler) handlePostUpdate(c *gin.Context) {
	// TODO
	// ...
}

// @Summary Delete post
// @Tags posts
// @Accept  json
// @Produce  json
// @Success 200 {object} models.PostDeleteResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /posts/{post_id} [delete]
// @Security ApiKeyAuth
func (h *Handler) handlePostDelete(c *gin.Context) {
	// TODO
	// ...
}
