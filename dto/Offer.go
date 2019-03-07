package dto

import (
      "time"	
)

type Offer struct {
	OfferId       int             `db:"OfferId"`
	AccountId     int             `db:"AccountId"`
	OfferType     string          `db:"OfferType"`
	Description   string          `db:"Description"`
	StartDateTime time.Time      `db:"StartDateTime"`
	EndDateTime   time.Time      `db:"EndDateTime"`
	Rate          string         `db:"Rate"`
	Value         float64        `db:"Value"`
	Term          string         `db:"Term"`
	Url           string         `db:"Url"`
	Amount        float64        `db:"Amount"`
	Currency      string          `db:"Currency"`
	FeeAmount     float64        `db:"FeeAmount"`
	FeeCurency    string         `db:"FeeCurency"`
	MakerDate     time.Time      `db:"MakerDate"`
	CheckerDate   time.Time      `db:"CheckerDate"`
	MakerId       string          `db:"MakerId"`
	CheckerId     string          `db:"CheckerId"`
	ModifiedBy    string          `db:"ModifiedBy"`
	ModifiedDate  time.Time      `db:"ModifiedDate"`
}
func (Offer) TableName() string {
	return "Offer"
}
