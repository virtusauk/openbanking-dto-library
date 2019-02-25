package dto

import (
	"time"
)

type Accountlink struct {
	AccountLinkId int            `db:"AccountLinkId"`
	FromAccountId int            `db:"FromAccountId"`
	ToAccountId   int            `db:"ToAccountId"`
	Type          string         `db:"Type"`
	Description   string `db:"Description"`
	StartDate     time.Time      `db:"StartDate"`
	EndDate       time.Time    `db:"EndDate"`
	BankId        int            `db:"BankId"`
	MakerDate     time.Time      `db:"MakerDate"`
	CheckerDate   time.Time    `db:"CheckerDate"`
	MakerId       string         `db:"MakerId"`
	CheckerId     string `db:"CheckerId"`
	ModifiedBy    string `db:"ModifiedBy"`
	ModifiedDate  time.Time    `db:"ModifiedDate"`
}
