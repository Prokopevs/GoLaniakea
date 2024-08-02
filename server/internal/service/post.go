package service

import (
	"context"
	"github.com/Prokopevs/GoLaniakea/server/internal/model"
)

func (s *ServiceImpl) CreatePost(ctx context.Context, post *model.Post) (int, error) {
	id, err := s.db.CreatePost(ctx, post)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *ServiceImpl) GetPosts(ctx context.Context, category, page, limit int) ([]*model.RankPost, error) {
	if page == -1 {
		page = 0
	}
	if limit == -1 {
		limit = 3
	}

	r, err := s.db.GetPosts(ctx, category, page, limit)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *ServiceImpl) GetPostById(ctx context.Context, id int) (*model.Post, error) {
	exists, err := s.db.IsPostWithIdExist(ctx, id)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, ErrNoSuchPost
	}

	post, err := s.db.GetPostById(ctx, id)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (s *ServiceImpl) IsPostWithIdExist(ctx context.Context, id int) (bool, error) {
	return s.db.IsPostWithIdExist(ctx, id)
}

func (s *ServiceImpl) DeletePostById(ctx context.Context, id int) error {
	exists, err := s.db.IsPostWithIdExist(ctx, id)
	if err != nil {
		return err
	}

	if !exists {
		return ErrNoSuchPost
	}

	err = s.db.DeletePostById(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceImpl) UpdatePost(ctx context.Context, post *model.Post) error {
	exists, err := s.db.IsPostWithIdExist(ctx, post.Id)
	if err != nil {
		return err
	}

	if !exists {
		return ErrNoSuchPost
	}

	err = s.db.UpdatePost(ctx, post)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceImpl) GetTotalCount(ctx context.Context) ([]*model.Total, error) {
	r, err := s.db.GetTotalCount(ctx)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *ServiceImpl) GetInteresting(ctx context.Context) ([]*model.Post, error) {
	r, err := s.db.GetInteresting(ctx)
	if err != nil {
		return nil, err
	}

	return r, nil
}
