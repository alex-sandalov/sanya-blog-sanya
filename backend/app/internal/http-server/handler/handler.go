package handler

import (
	"github.com/gin-gonic/gin"

	_ "blog-app/app/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	// TODO
	// ...
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			external := v1.Group("/external")

			users := external.Group("/users")
			h.initUsersHandler(users)

			posts := external.Group("/posts")
			h.initPostsHandler(posts)

			auth := external.Group("/auth")
			h.initAuthHandlers(auth)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler()))

	return router
}
