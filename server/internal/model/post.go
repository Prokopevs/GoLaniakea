package model

type Post struct {
	Id           int    `json:"id" db:"id"`
	ImageUrl     string `json:"imageUrl" db:"image_url"`
	Name         string `json:"name" db:"name"`
	Description  string `json:"description" db:"description"`
	Date         string `json:"date" db:"date"`
	Category     int    `json:"category" db:"category"`
	CategoryName string `json:"categoryName" db:"category_name"`
	LikeCount    int    `json:"likeCount" db:"like_count"`
	Text         string `json:"text" db:"text"`
}

type RankPost struct {
	Id           int    `json:"id" db:"id"`
	ImageUrl     string `json:"imageUrl" db:"image_url"`
	Name         string `json:"name" db:"name"`
	Description  string `json:"description" db:"description"`
	Date         string `json:"date" db:"date"`
	Category     int    `json:"category" db:"category"`
	CategoryName string `json:"categoryName" db:"category_name"`
	LikeCount    int    `json:"likeCount" db:"like_count"`
	Text         string `json:"text" db:"text"`
	R            int
}

type Total struct {
	Id         int `json:"id"`
	TotalCount int `json:"totalCount"`
}

type Interesting struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Date string `json:"date" db:"date"`
}
