package post

import (
	"context"
	"time"
)

type service struct {
	repo PostRepo
	timeout time.Duration
}

func NewServices(repo PostRepo) *service {
	return &service{
		repo,
		time.Duration(2) * time.Second,
	}
}

func (s *service) CreatePost(c context.Context, req *CreatePostReq) (*CreatePostRes, error) {
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