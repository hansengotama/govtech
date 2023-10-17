package domain

import "time"

type ProductReview struct {
	ID        int
	ProductID int
	Review    string
	Rating    int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

const (
	ProductMinRating = 0
	ProductMaxRating = 5
)
