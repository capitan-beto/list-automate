package models

import (
	"github.com/shopspring/decimal"
)

type Product struct {
	ID       string          `json:"id"`
	Desc     string          `json:"desc"`
	Price    decimal.Decimal `json:"price"`
	Subcat   string          `json:"subcat"`
	Cat      string          `json:"cat"`
	Src      string          `json:"src"`
	Date     string          `json:"date"`
	AlternID string          `json:"altern-id"`
}
