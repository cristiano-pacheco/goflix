package enum

import (
	"errors"
)

const (
	EnumContentTypeMovie  string = "Movie"
	EnumContentTypeTvShow string = "TvShow"
)

type ContentTypeEnum struct {
	value string
}

func NewContentTypeEnum(value string) (ContentTypeEnum, error) {
	if err := validateContentTypeEnum(value); err != nil {
		return ContentTypeEnum{}, err
	}

	return ContentTypeEnum{value: value}, nil
}

func (s *ContentTypeEnum) String() string {
	return s.value
}

func validateContentTypeEnum(value string) error {
	allowedValues := map[string]struct{}{
		EnumContentTypeMovie:  {},
		EnumContentTypeTvShow: {},
	}

	if _, ok := allowedValues[value]; !ok {
		return errors.New("invalid content type: " + value)
	}

	return nil
}
