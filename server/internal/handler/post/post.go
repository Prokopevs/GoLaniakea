package handler

import (
	"context"
	"github/Prokopevs/GoLaniakea/internal/services/post"
)

type PostService interface {
	CreatePost(c context.Context, req *services.CreatePostReq) (*services.CreatePostRes, error)
}