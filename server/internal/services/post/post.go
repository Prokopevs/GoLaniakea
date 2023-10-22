package services

import (
	"context"
	"github/Prokopevs/GoLaniakea/internal/model"
)

type PostRepo interface {
	CreatePost(ctx context.Context, post *model.Post) (*model.Post, error)
	GetPosts(category, page, limit string) ([]*model.Post, error)
	GetPostById(ctx context.Context, id string) (*model.Post, error)
	DeletePostById(ctx context.Context, id string) (*string, error)
}

func NewPost(imageUrl, name, description, date string, category int64, categoryName string, likeCount int64, liked bool, text string) *model.Post {
	return &model.Post{
		ImageUrl:     imageUrl,
		Name:         name,
		Description:  description,
		Date:         date,
		Category:     category,
		CategoryName: categoryName,
		LikeCount:    likeCount,
		Liked:        liked,
		Text:         text,
	}
}
