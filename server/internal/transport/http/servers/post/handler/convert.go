package handler

import (
	"strconv"

	"github.com/Prokopevs/GoLaniakea/server/internal/model"
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
		Id:           p.Id,
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

type interestingResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Date string `json:"date"`
}

func convertPostsToInteresting(p []*model.Post) []*interestingResponse {
	var interesting []*interestingResponse
	for _, i := range p {
		interesting = append(interesting, &interestingResponse{
			Id:   i.Id,
			Name: i.Name,
			Date: i.Date,
		})
	}
	return interesting
}

func convertRankPostsToPosts(p []*model.RankPost) []*model.Post {
	var posts []*model.Post
	for _, i := range p {
		posts = append(posts, &model.Post{
			Id:           i.Id,
			ImageUrl:     i.ImageUrl,
			Name:         i.Name,
			Description:  i.Description,
			Date:         i.Date,
			Category:     i.Category,
			CategoryName: i.CategoryName,
			LikeCount:    i.LikeCount,
			Text:         i.Text,
		})
	}
	return posts
}
