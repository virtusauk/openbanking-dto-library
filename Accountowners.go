package dto

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type Accountowners struct {
	AccountOwnerId        int            `db:"AccountOwnerId"`
	AccountId             int            `db:"AccountId"`
	BankId                int            `db:"BankId"`
	PartyId               int            `db:"PartyId"`
	SchemeName            sql.NullString `db:"SchemeName"`
	AccountIdentification string         `db:"AccountIdentification"`
	IsOnlineAccessEnabled sql.NullString `db:"IsOnlineAccessEnabled"`
	AccountJoiningDate    time.Time      `db:"AccountJoiningDate"`
	AccountLeavingDate    pq.NullTime    `db:"AccountLeavingDate"`
	IsPrimaryOwner        string         `db:"IsPrimaryOwner"`
	PercentageOfShare     float64        `db:"PercentageOfShare"`
	Status                string         `db:"Status"`
	StartDate             time.Time      `db:"StartDate"`
	EndDate               pq.NullTime    `db:"EndDate"`
	MakerDate             time.Time      `db:"MakerDate"`
	CheckerDate           pq.NullTime    `db:"CheckerDate"`
	MakerId               string         `db:"MakerId"`
	CheckerId             sql.NullString `db:"CheckerId"`
	ModifiedBy            sql.NullString `db:"ModifiedBy"`
	ModifiedDate          pq.NullTime    `db:"ModifiedDate"`
}
