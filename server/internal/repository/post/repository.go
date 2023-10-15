package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github/Prokopevs/GoLaniakea/internal/model"
)


type PostRepo struct {
	db *sql.DB
}

//NewPostRepository create new repository
func NewPostRepository(db *sql.DB) *PostRepo {
	return &PostRepo{
		db: db,
	}
}

func (r *PostRepo) CreatePost(ctx context.Context, post *model.Post) (*model.Post, error) {
	fmt.Println(post.Name)
	var lastInsertId int
	const query = `INSERT INTO posts(imageUrl, name, description, date, category, categoryName, likeCount, liked, text) VALUES 
	($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id`
	err := r.db.QueryRowContext(ctx, query, post.ImageUrl, post.Name, post.Description, post.Date, post.Category, post.CategoryName, 
		post.LikeCount, post.Liked, post.Text).Scan(&lastInsertId)
	if err != nil {
		return nil, err
	}

	post.ID = int64(lastInsertId)
	return post, nil
}