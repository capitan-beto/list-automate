package models

import (
	"github.com/shopspring/decimal"
)

type Product struct {
	ID    string          `json:"id"`
	Desc  string          `json:"desc"`
	Price decimal.Decimal `json:"price"`
	Date  string          `json:"date"`
}
