package handler

import (
	"github/Prokopevs/GoLaniakea/internal/model"
	"strconv"
)

func convertPostJSONToPost(p *PostJSON) (*model.Post, error) {
	categoryInt, err := strconv.Atoi(p.Category)
	if err != nil {
		return nil, err
	}

	LikeCountInt, err := strconv.Atoi(p.LikeCount)
	if err != nil {
		return nil, err
	}

	return &model.Post{
		ImageUrl:     p.ImageUrl,
		Name:         p.Name,
		Description:  p.Description,
		Date:         p.Date,
		Category:     categoryInt,
		CategoryName: p.CategoryName,
		LikeCount:    LikeCountInt,
		Text:         p.Text,
	}, err
}