package dto

import (
	"time"
)

type Currencymaster struct {
	CurrencyId             string         `db:"CurrencyId"`
	CurrencyDesc           string         `db:"CurrencyDesc"`
	CurrencyCode           string         `db:"CurrencyCode"`
	CurrencyName           string         `db:"CurrencyName"`
	Active                 string         `db:"Active"`
	CurrencySymbol         string         `db:"CurrencySymbol"`
	DigitsAfterDecimal     int  `db:"DigitsAfterDecimal"`
	CurrencySymbolPosition string `db:"CurrencySymbolPosition"`
	IsBaseCurrency         string         `db:"IsBaseCurrency"`
	IsoCode                string `db:"IsoCode"`
	MajorName              string `db:"MajorName"`
	MinorName              string `db:"MinorName"`
	MakerDate              time.Time      `db:"MakerDate"`
	CheckerDate            time.Time    `db:"CheckerDate"`
	MakerId                string         `db:"MakerId"`
	CheckerId              string `db:"CheckerId"`
	ModifiedBy             string `db:"ModifiedBy"`
	ModifiedDate           time.Time    `db:"ModifiedDate"`
}
func (Currencymaster) TableName() string {
	return "CurrencyMaster"
}