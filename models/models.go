package models

import "github.com/shopspring/decimal"

type Product struct {
	Id    string          `json:id`
	Desc  string          `json:desc`
	Price decimal.Decimal `json:price`
	Cat   string          `json:cat`
	Src   string          `json:src`
}
