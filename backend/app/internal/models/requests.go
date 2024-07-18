package models

type LoginRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegistersRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type PostCreateRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type PostUpdateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type AllPostsGetRequest struct {
	Offset int `form:"offset" binding:"min=1"`
	Limit  int `form:"limit" binding:"min=1"`
}

type AllUsersGetRequest struct {
	Offset int `form:"offset" binding:"min=1"`
	Limit  int `form:"limit" binding:"min=1"`
}

type UserCreateRequest struct {
	Login string `json:"login" binding:"required"`
}
