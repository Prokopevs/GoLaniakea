package handler

import (
	"context"
	"github/Prokopevs/GoLaniakea/internal/model"
)

type PostService interface {
	CreatePost(c context.Context, req *CreatePostReq) (*CreatePostRes, error)
	GetPosts(category, page, limit string) ([]*model.Post, string, error)
	GetPostById(c context.Context, id string) (*model.Post, error)
	DeletePostById(ctx context.Context, id string) (*string, error)
	UpdatePost(c context.Context, req *model.Post) (*string, error)
}

type CreatePostReq struct {
	ImageUrl     string `json:"imageUrl" db:"imageUrl"`
	Name         string `json:"name" db:"name"`
	Description  string `json:"description" db:"description"`
	Date         string `json:"date" db:"date"`
	Category     int64  `json:"category" db:"category"`
	CategoryName string `json:"categoryName" db:"categoryName"`
	LikeCount    int64  `json:"likeCount" db:"likeCount"`
	Liked        bool   `json:"liked" db:"liked"`
	Text         string `json:"text" db:"text"`
}

type CreatePostRes struct {
	ID int64 `json:"id" db:"id"`
}

type PostsResponse struct {
	Posts []*model.Post `json:"posts"`
	Total string `json:"total"`
}
