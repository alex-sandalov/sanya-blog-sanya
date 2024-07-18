package models

import (
	"github.com/gin-gonic/gin"
)

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RegistersResponse struct {
	AccessToken    string `json:"access_token"`
	RefreshesToken string `json:"refresh_token"`
}

type RefreshResponse struct {
	AccessToken string `json:"access_token"`
}

type AllPostsGetResponse struct {
	PostsTotal       int    `json:"count_total"`
	CountPostsOnPage int    `json:"count_posts_on_page"`
	Posts            []Post `json:"posts"`
}

type PostGetResponse struct {
	Post Post `json:"post"`
}

type PostCreateResponse struct {
	Post Post `json:"post"`
}

type PostUpdateResponse struct {
	Post Post `json:"post"`
}

type PostDeleteResponse struct {
	Post Post `json:"post"`
}

type AllUsersGetResponse struct {
	UsersTotal       int    `json:"count_total"`
	CountUsersOnPage int    `json:"count_users_on_page"`
	Users            []User `json:"users"`
}

type UserGetResponse struct {
	User User `json:"user"`
}

type UserCreateResponse struct {
	User User `json:"user"`
}

type UserDeleteResponse struct {
	User User `json:"user"`
}

type User struct {
	ID    string `json:"id"`
	Login string `json:"login"`
}

type Post struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	// logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, ErrorResponse{message})
}
