package product

import (
	"context"
	"github.com/google/uuid"
	"github.com/hansengotama/govtech/internal/domain"
	"github.com/hansengotama/govtech/internal/lib/array"
	"github.com/hansengotama/govtech/internal/lib/errcustom"
	"strings"
)

var (
	defaultSortBy      = "created_at"
	validSortBy        = []string{"created_at", "rating"}
	defaultOrderAction = "desc"
	validOrderAction   = []string{"asc", "desc"}
)

type GetProductsQueryParam struct {
	SKU          string
	Title        string
	CategoryGUID uuid.UUID
	Etalase      string
	Rating       float32

	// sort
	SortBy string
	Order  string
}

func (param GetProductsQueryParam) Validate() error {
	if param.Rating < domain.ProductMinRating {
		return errcustom.ErrorFieldLessThanMinValue{
			Field:    "rating",
			MinValue: domain.ProductMinRating,
		}
	}

	if param.Rating > domain.ProductMaxRating {
		return errcustom.ErrorFieldMoreThanMaxValue{
			Field:    "rating",
			MaxValue: domain.ProductMaxRating,
		}
	}

	if param.Etalase != "" {
		_, isValidEtalase := domain.ConvertToProductEtalase(param.Etalase)
		if !isValidEtalase {
			return errcustom.ErrorInvalidField{
				Field:         "etalase",
				AcceptedValue: strings.Join(domain.ProductEtalaseStrs, ", "),
			}
		}
	}

	if param.SortBy != "" {
		isValidSortBy := array.IsContain(validSortBy, param.SortBy)
		if !isValidSortBy {
			return errcustom.ErrorInvalidField{
				Field:         "sort by",
				AcceptedValue: strings.Join(validSortBy, ", "),
			}
		}
	}

	if param.Order != "" {
		isValidOrder := array.IsContain(validOrderAction, param.Order)
		if !isValidOrder {
			return errcustom.ErrorInvalidField{
				Field:         "order",
				AcceptedValue: strings.Join(validOrderAction, ", "),
			}
		}
	}

	return nil
}

type GetProductsResponse struct {
	Title       string
	Description string
	Category    string
	Etalase     []domain.ProductEtalase
	Images      []domain.ProductImage
	Weight      float32
	Price       float32
	Rating      float32
}

func GetProducts(ctx context.Context, param GetProductsQueryParam) ([]GetProductsResponse, error) {
	err := param.Validate()
	if err != nil {
		return nil, err
	}

	if param.SortBy == "" {
		param.SortBy = defaultSortBy
	}

	if param.Order == "" {
		param.Order = defaultOrderAction
	}

	// If category id is more than 0
	// Get category by id
	// Check is err not found

	// Get products by filter

	return nil, nil
}
