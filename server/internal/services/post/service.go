package services

import (
	"context"
	"time"
)

type serv struct {
	repo PostRepo
	timeout time.Duration
}

func NewService(repo PostRepo) *serv {
	return &serv{
		repo: repo,
		timeout: 2 * time.Second,
	}
}

func (s *serv) CreatePost(c context.Context, req *CreatePostReq) (*CreatePostRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	p := NewPost(req.ImageUrl, req.Name, req.Description, req.Date, req.Category, req.CategoryName, req.LikeCount, req.Liked, req.Text)
	r, err := s.repo.CreatePost(ctx, p)
	if err != nil {
		return nil, err
	}

	res := &CreatePostRes{
		ID: r.ID,
	}

	return res, nil
}