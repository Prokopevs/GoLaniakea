package services

import (
	"context"
	"github/Prokopevs/GoLaniakea/internal/model"
	"github/Prokopevs/GoLaniakea/internal/transport/http/servers/post/handler"
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

func (s *serv) CreatePost(c context.Context, req *handler.CreatePostReq) (*handler.CreatePostRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	p := NewPost(req.ImageUrl, req.Name, req.Description, req.Date, req.Category, req.CategoryName, req.LikeCount, req.Liked, req.Text)
	r, err := s.repo.CreatePost(ctx, p)
	if err != nil {
		return nil, err
	}

	res := &handler.CreatePostRes{
		ID: r.ID,
	}

	return res, nil
}

func (s *serv) GetPosts(page string, limit string) ([]*model.Post, error) {
	r, err := s.repo.GetPosts(page, limit)
	if err != nil {
		return nil, err
	}

	return r, nil
}