package model

type Post struct {
	ID 				int64	`json:"id" db:"id"`
	ImageUrl 		string	`json:"imageUrl" db:"imageUrl"`
	Name 			string	`json:"name" db:"name"`
	Description 	string	`json:"description" db:"description"`
	Date 			string	`json:"date" db:"date"`
	Category 		int64	`json:"category" db:"category"`
	CategoryName 	string	`json:"categoryName" db:"categoryName"`
	LikeCount 		int64	`json:"likeCount" db:"likeCount"`
	Liked 			bool	`json:"liked" db:"liked"`
	Text 			string	`json:"text" db:"text"`
}