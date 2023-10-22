package repository

import (
	"context"
	"database/sql"
	"github/Prokopevs/GoLaniakea/internal/model"
)

type PostRepo struct {
	db *sql.DB
}

// NewPostRepository create new repository
func NewPostRepository(db *sql.DB) *PostRepo {
	return &PostRepo{
		db: db,
	}
}

func (r *PostRepo) CreatePost(ctx context.Context, post *model.Post) (*model.Post, error) {
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

func (r *PostRepo) GetPosts(category, page, limit string) ([]*model.Post, error) {
	var posts []*model.Post

	var query string
	var rows *sql.Rows
	var err error
	if page == "" || limit == "" {
		page = "0"
		limit = "3"
	}

	if category == "" {
		query = "SELECT * FROM posts OFFSET $1 LIMIT $2"
		rows, err = r.db.Query(query, page, limit)
	} else {
		query = "SELECT * FROM posts WHERE category = $1 OFFSET $2 LIMIT $3"
		rows, err = r.db.Query(query, category, page, limit)
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var p model.Post

		err = rows.Scan(&p.ID, &p.ImageUrl, &p.Name, &p.Description, &p.Date, &p.Category, &p.CategoryName, &p.LikeCount, &p.Liked, &p.Text)
		if err != nil {
			return nil, err
		}

		posts = append(posts, &p)
	}

	return posts, nil
}

func (r *PostRepo) GetPostById(ctx context.Context, id string) (*model.Post, error) {
	var p model.Post
	const query = "SELECT * FROM posts WHERE id = $1"

	err := r.db.QueryRowContext(ctx, query, id).Scan(&p.ID, &p.ImageUrl, &p.Name, &p.Description, &p.Date, &p.Category, &p.CategoryName, &p.LikeCount, &p.Liked, &p.Text)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *PostRepo) DeletePostById(ctx context.Context, id string) (*string, error) {
	var idFromDB string
	const query = "DELETE FROM posts WHERE id = $1 returning id"

	err := r.db.QueryRowContext(ctx, query, id).Scan(&idFromDB)
	if err != nil {
		return nil, err
	}

	return &idFromDB, nil
}