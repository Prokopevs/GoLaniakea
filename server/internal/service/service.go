package service

import (
	"context"

	"github.com/Prokopevs/GoLaniakea/server/internal/model"
)

type DB interface {
	CreatePost(ctx context.Context, post *model.Post) (int, error)
	GetPosts(ctx context.Context, category, page, limit int) ([]*model.RankPost, error)
	GetPostById(ctx context.Context, id int) (*model.Post, error)
	IsPostWithIdExist(ctx context.Context, id int) (bool, error)
	DeletePostById(ctx context.Context, id int) error
	UpdatePost(ctx context.Context, post *model.Post) error
	GetTotalCount(ctx context.Context) ([]*model.Total, error)
	GetInteresting(ctx context.Context) ([]*model.Post, error)
}

type ServiceImpl struct {
	db DB
}

func NewServiceImpl(db DB) *ServiceImpl {
	return &ServiceImpl{
		db: db,
	}
}
