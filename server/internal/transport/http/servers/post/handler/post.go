package handler

import (
	"context"
	"github/Prokopevs/GoLaniakea/internal/model"
)

type PostService interface {
	CreatePost(c context.Context, req *CreatePostReq) (*CreatePostRes, error)
	GetPosts(page string, limit string) ([]*model.Post, error)
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
