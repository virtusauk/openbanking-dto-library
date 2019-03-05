package dto

import (
	"database/sql"

	"github.com/lib/pq"
)

type Standingorders struct {
	StandingOrderId                        int             `db:"StandingOrderId"`
	DebtorAccountId                        int             `db:"DebtorAccountId"`
	BankId                                 int             `db:"BankId"`
	StandingOrderAccountRefId              string          `db:"StandingOrderAccountRefId"`
	StandingOrderIdRef                     string          `db:"StandingOrderIdRef"`
	Frequency                              string          `db:"Frequency"`
	Reference                              string          `db:"Reference"`
	FirstPaymentDateTime                   time.Time     `db:"FirstPaymentDateTime"`
	FirstPaymentAmount                     float64 `db:"FirstPaymentAmount"`
	FirstPaymentCurrency                   string  `db:"FirstPaymentCurrency"`
	NextPaymentDateTime                    time.Time     `db:"NextPaymentDateTime"`
	NextPaymentAmount                      float64 `db:"NextPaymentAmount"`
	NextPaymentCurrency                    string  `db:"NextPaymentCurrency"`
	FinalPaymentDateTime                   time.Time     `db:"FinalPaymentDateTime"`
	FinalPaymentAmount                     float64 `db:"FinalPaymentAmount"`
	FinalPaymentCurrency                   string  `db:"FinalPaymentCurrency"`
	CreditorBankCode                       string  `db:"CreditorBankCode"`
	CreditorBankName                       string  `db:"CreditorBankName"`
	ServicerSchemeName                     string  `db:"ServicerSchemeName"`
	ServicerIdentification                 string  `db:"ServicerIdentification"`
	CreditorAccountSchemeName              string  `db:"CreditorAccountSchemeName"`
	CreditorAccountIdentification          string          `db:"CreditorAccountIdentification"`
	CreditorAccountName                    string          `db:"CreditorAccountName"`
	CreditorAccountSecondaryIdentification string          `db:"CreditorAccountSecondaryIdentification"`
	CreditorAccountId                      int   `db:"CreditorAccountId"`
	MakerDate                              time.Time       `db:"MakerDate"`
	CheckerDate                            time.Time     `db:"CheckerDate"`
	MakerId                                int             `db:"MakerId"`
	CheckerId                              int   `db:"CheckerId"`
	ModifiedBy                             string  `db:"ModifiedBy"`
	ModifiedDate                           time.Time     `db:"ModifiedDate"`
	Banks								   Banks
}
func (Standingorders) TableName() string {
	return "StandingOrders"
}
