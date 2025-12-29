package errs

import "errors"

var (
	ErrMovieNotFound = errors.New("movie not found")

	ErrMovieIDRequired          = errors.New("movie ID is required")
	ErrMovieTitleRequired       = errors.New("movie title is required")
	ErrMovieTitleTooShort       = errors.New("movie title must be at least 1 character long")
	ErrMovieTitleTooLong        = errors.New("movie title cannot exceed 500 characters")
	ErrMovieDescriptionRequired = errors.New("movie description is required")
	ErrMovieDescriptionTooLong  = errors.New("movie description cannot exceed 2000 characters")
	ErrMovieTypeRequired        = errors.New("movie content type is required")
	ErrMovieInvalidContentType  = errors.New("invalid content type for movie")

	ErrMovieAgeRecommendationInvalid = errors.New("age recommendation must be between 0 and 21")

	ErrMovieExternalRatingInvalid = errors.New("external rating must be between 0.0 and 10.0")

	ErrCreatedAtRequired        = errors.New("created at timestamp is required")
	ErrUpdatedAtRequired        = errors.New("updated at timestamp is required")
	ErrUpdatedAtBeforeCreatedAt = errors.New("updated at timestamp cannot be before created at timestamp")
)
