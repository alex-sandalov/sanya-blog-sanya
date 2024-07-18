package main

import "blog-app/app/internal/http-server/handler"

// @title Blog App API
// @version 1.0
// @description API Server for Blog App

// @host localhost:8080
// @BasePath /api/v1/external

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	handler := handler.NewHandler().InitRoutes()
	handler.Run(":8080")
}
