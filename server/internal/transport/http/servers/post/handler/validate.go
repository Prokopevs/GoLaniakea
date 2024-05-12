package handler

import (
	"fmt"
)

const (
	codeImageUrlCannotBeEmpty    = "IMAGE_URL_CANNOT_BE_EMPTY"
	codeNameCannotBeEmpty        = "NAME_CANNOT_BE_EMPTY"
	codeDescriptionCannotBeEmpty = "DESCRIPTION_CANNOT_BE_EMPTY"
	codeDateCannotBeEmpty        = "DATE_CANNOT_BE_EMPTY"
	codeCategoryCannotBeEmpty    = "CATEGORY_CANNOT_BE_EMPTY"
	codeCategoryNameCannotBeEmpty = "CATEGORY_NAME_CANNOT_BE_EMPTY"
	codeLikeCountCannotBeEmpty = "LIKE_COUNT_CANNOT_BE_EMPTY"
	codeTextCannotBeEmpty = "TEXT_CANNOT_BE_EMPTY"
)

func (p *PostJSON) validate() (Code, error) {
	if p.ImageUrl == "" {
		return codeContentEmpty, fmt.Errorf("imageUrl cannot be empty")
	}
	if p.Name == "" {
		return codeNameCannotBeEmpty, fmt.Errorf("name cannot be empty")
	}

	if p.Description == "" {
		return codeNameCannotBeEmpty, fmt.Errorf("description cannot be empty")
	}

	if p.Date == "" {
		return codeDateCannotBeEmpty, fmt.Errorf("date cannot be empty")
	}

	if p.Category == "" {
		return codeCategoryCannotBeEmpty, fmt.Errorf("category cannot be empty")
	}

	if p.CategoryName == "" {
		return codeCategoryNameCannotBeEmpty, fmt.Errorf("categoryName cannot be empty")
	}

	if p.LikeCount == "" {
		return codeLikeCountCannotBeEmpty, fmt.Errorf("likeCount cannot be empty")
	}

	if p.Text == "" {
		return codeTextCannotBeEmpty, fmt.Errorf("text cannot be empty")
	}

	return CodeOK, nil
}
