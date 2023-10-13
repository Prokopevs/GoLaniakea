package post

import (
	"context"
	"github/Prokopevs/GoLaniakea/internal/model"
)

type PostRepo interface {
	CreatePost(ctx context.Context, post *model.Post) (*model.Post, error)
}

type CreatePostReq struct {
	ImageUrl 		string	`json:"imageUrl" db:"imageUrl"`
	Name 			string	`json:"name" db:"name"`
	Description 	string	`json:"description" db:"description"`
	Date 			string	`json:"date" db:"date"`
	Category 		int64	`json:"category" db:"category"`
	CategoryName 	string	`json:"categoryName" db:"categoryName"`
	LikeCount 		int64	`json:"likeCount" db:"likeCount"`
	Liked 			bool	`json:"liked" db:"liked"`
	Text 			string	`json:"text" db:"text"`
}

type CreatePostRes struct {
	ID 				int64	`json:"id" db:"id"`
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
		Liked: 	      liked,
		Text:         text,
	}
}