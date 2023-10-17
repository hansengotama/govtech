package domain

import (
	"github.com/google/uuid"
	"time"
)

type ProductCategory struct {
	ID   int
	GUID uuid.UUID
	Name string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
