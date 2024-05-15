package db

import (
	"context"

	"github.com/Prokopevs/GoLaniakea/server/internal/model"
)

func (r *database) CreatePost(ctx context.Context, post *model.Post) (int, error) {
	const query = `INSERT INTO posts(image_url, name, description, date, category, category_name, like_count, text) VALUES 
	($1, $2, $3, $4, $5, $6, $7, $8) returning id`

	var id int

	err := r.db.QueryRowContext(ctx, query, post.ImageUrl, post.Name, post.Description, post.Date, post.Category, post.CategoryName, post.LikeCount, post.Text).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *database) GetPosts(ctx context.Context, category, page, limit int) ([]*model.RankPost, error) {
    const (
        qWithoutCategory = "SELECT * FROM (SELECT *, RANK() OVER (ORDER BY id) AS r FROM posts) AS ranked_posts WHERE r > $1 AND r <= $2"
        qWithCategory    = "SELECT * FROM (SELECT *, RANK() OVER (ORDER BY id) AS r FROM posts WHERE category = $1) AS ranked_posts WHERE r > $2 AND r <= $3"
    )
    posts := []*model.RankPost{}

    if category == -1 {
        startRank := (page - 1) * limit
        endRank := startRank + limit

        err := r.db.SelectContext(ctx, &posts, qWithoutCategory, startRank, endRank)
        if err != nil {
            return nil, err
        }
    } else {
        startRank := (page - 1) * limit
        endRank := startRank + limit

        err := r.db.SelectContext(ctx, &posts, qWithCategory, category, startRank, endRank)
        if err != nil {
            return nil, err
        }
    }

    return posts, nil
}

func (r *database) GetPostById(ctx context.Context, id int) (*model.Post, error) {
	const q = "SELECT * FROM posts WHERE id = $1"

	p := &model.Post{}

	err := r.db.GetContext(ctx, p, q, id)

	return p, err
}

func (r *database) IsPostWithIdExist(ctx context.Context, id int) (bool, error) {
	const q = "select exists(select from posts where id=$1)"

	exists := false
	err := r.db.GetContext(ctx, &exists, q, id)

	return exists, err
}

func (r *database) DeletePostById(ctx context.Context, id int) error {
	const query = "DELETE FROM posts WHERE id = $1"

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *database) UpdatePost(ctx context.Context, post *model.Post) error {
	const UpdatePost = `UPDATE posts SET image_url=$1, name=$2, description=$3, date=$4, category=$5, category_name=$6, like_count=$7, text=$8 WHERE id = $9`

	_, err := r.db.ExecContext(ctx, UpdatePost, post.ImageUrl, post.Name, post.Description, post.Date, post.Category, post.CategoryName, post.LikeCount, post.Text, post.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *database) GetTotalCount(ctx context.Context) ([]*model.Total, error) {
	const totalCountQ = "SELECT category AS id, COUNT(*) AS totalCount FROM posts GROUP BY category ORDER BY category;"
		
	total := []*model.Total{}

	err := r.db.SelectContext(ctx, &total, totalCountQ)
	if err != nil {
		return nil, err
	}
	
	return total, nil
}

func (r *database) GetInteresting(ctx context.Context) ([]*model.Post, error) {
	const interestingQ = "SELECT id, name, date FROM posts;"
		
	posts := []*model.Post{}

	err := r.db.SelectContext(ctx, &posts, interestingQ)
	if err != nil {
		return nil, err
	}
	
	return posts, nil
}



