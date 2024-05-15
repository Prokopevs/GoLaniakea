package handler

import (
	"context"
	"errors"

	"net/http"
	"github.com/Prokopevs/GoLaniakea/server/internal/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Service interface {
	CreatePost(ctx context.Context, post *model.Post) (int, error)
	GetPosts(ctx context.Context, category, page, limit int) ([]*model.RankPost, error)
	GetPostById(ctx context.Context, id int) (*model.Post, error)
	IsPostWithIdExist(ctx context.Context, id int) (bool, error)
	DeletePostById(ctx context.Context, id int) error
	UpdatePost(ctx context.Context, post *model.Post) error
	GetTotalCount(ctx context.Context) ([]*model.Total, error)
	GetInteresting(ctx context.Context) ([]*model.Post, error)
}

type HTTP struct {
	innerServer *http.Server

	log     *zap.SugaredLogger
	service Service
	password string
}

func (h *HTTP) Run(ctx context.Context) {
	h.log.Infow("HTTP server starting.", "addr", h.innerServer.Addr)

	go func() {
		err := h.innerServer.ListenAndServe()
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return
			}

			h.log.Errorw("Listen and serve HTTP", "addr", h.innerServer.Addr, "err", err)
		}
	}()

	<-ctx.Done()

	h.log.Info("Graceful server shutdown.")
	h.innerServer.Shutdown(context.Background())
}

func NewHTTP(addr string, password string, logger *zap.SugaredLogger, postSvc Service) *HTTP {
	h := &HTTP{
		log:     logger,
		service: postSvc,
		password: password,
	}

	r := gin.Default()
	h.setRoutes(r)

	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	h.innerServer = srv

	return h
}
