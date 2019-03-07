package dto

import (
	"time"
)

type Banks struct {
	BankId                 int            `db:"BankId"`
	BankRegNumber          string         `db:"BankRegNumber"`
	BankIdentificationCode string         `db:"BankIdentificationCode"`
	BankName               string         `db:"BankName"`
	Country                string         `db:"Country"`
	RegisteredAddress      string         `db:"RegisteredAddress"`
	MakerDate              time.Time      `db:"MakerDate"`
	CheckerDate            time.Time    `db:"CheckerDate"`
	MakerId                string         `db:"MakerId"`
	CheckerId              string `db:"CheckerId"`
	ModifiedBy             string `db:"ModifiedBy"`
	ModifiedDate           time.Time    `db:"ModifiedDate"`
}
