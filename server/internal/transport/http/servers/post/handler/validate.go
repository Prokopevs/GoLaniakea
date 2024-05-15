package handler

import (
	"fmt"
	"strconv"
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

func (p *PostJSON) validate(endpoint string) (Code, error) {
	if endpoint == "update" {
		if p.Id == 0 {
			return codeContentEmpty, fmt.Errorf("id cannot be zero")
		}
	}
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

	if p.Category == ""  {
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

	return codeOK, nil
}

func updateGetParams(category, page, limit string) (int, int, int, error) {
	var categoryInt int
	var pageInt int
	var limitInt int

	var err error
	
	if category == "" {
		categoryInt = -1
	} else {
		categoryInt, err = strconv.Atoi(category)
		if err != nil {
			return 0, 0, 0, fmt.Errorf("category: %w", err.Error())
		}
	}

	if page == "" {
		pageInt = -1
	} else {
		pageInt, err = strconv.Atoi(page)
		if err != nil {
			return 0, 0, 0, fmt.Errorf("page: %w", err.Error())
		}
	}

	if limit == "" {
		limitInt = -1
	} else {
		limitInt, err = strconv.Atoi(limit)
		if err != nil {
			return 0, 0, 0, fmt.Errorf("limit: %w", err.Error())
		}
	}
	
	return categoryInt, pageInt, limitInt, nil
}
