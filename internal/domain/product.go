package domain

import (
	"strings"
	"time"
)

type Product struct {
	ID          int
	CategoryID  int
	SKU         string
	Title       string
	Description string
	Etalase     []ProductEtalase
	Images      []ProductImage
	Weight      float32
	Price       float32
	Rating      float32

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type ProductImage struct {
	Url         string
	Description string
}

type ProductEtalase int

const (
	ProductEtalaseRecommendation ProductEtalase = iota
	ProductEtalaseTrending
	ProductEtalaseNew
	ProductEtalaseDiscount
)

var (
	ProductEtalaseStrs = []string{
		"recommendation",
		"trending",
		"new",
		"discount",
	}

	productEtalaseMap = map[string]ProductEtalase{
		"recommendation": ProductEtalaseRecommendation,
		"trending":       ProductEtalaseTrending,
		"new":            ProductEtalaseNew,
		"discount":       ProductEtalaseDiscount,
	}
)

func (pe ProductEtalase) String() string {
	return ProductEtalaseStrs[pe]
}

func (pe ProductEtalase) EnumIndex() int {
	return int(pe)
}

func ConvertToProductEtalase(str string) (ProductEtalase, bool) {
	pe, ok := productEtalaseMap[strings.ToLower(str)]
	return pe, ok
}
