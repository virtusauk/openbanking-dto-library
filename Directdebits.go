package dto

import (
	"time"
)

type Directdebits struct {
	DirectDebitId                          int             `db:"DirectDebitId"`
	BankId                                 int             `db:"BankId"`
	DebtorAccountId                        int             `db:"DebtorAccountId"`
	ServicerSchemeName                     string          `db:"ServicerSchemeName"`
	ServicerIdentification                 string          `db:"ServicerIdentification"`
	CreditorAccountSchemeName              string          `db:"CreditorAccountSchemeName"`
	CreditorAccountIdentification          string          `db:"CreditorAccountIdentification"`
	CreditorAccountName                    string          `db:"CreditorAccountName"`
	CreditorAccountSecondaryIdentification string          `db:"CreditorAccountSecondaryIdentification"`
	DirectDebitReferenceNumber             string          `db:"DirectDebitReferenceNumber"`
	MandateIdentification                  string          `db:"MandateIdentification"`
	DirectdebitStatusCode                  string          `db:"DirectdebitStatusCode"`
	NameServiceUser                        string          `db:"NameServiceUser"`
	PreviousPaymentDatetime                time.Time       `db:"PreviousPaymentDatetime"`
	PreviousPaymentAmount                  float           `db:"PreviousPaymentAmount"`
	Currency                               string          `db:"Currency"`
	MakerDate                              time.Time       `db:"MakerDate"`
	CheckerDate                            time.Time      `db:"CheckerDate"`
	MakerId                                int            `db:"MakerId"`
	CheckerId                              int            `db:"CheckerId"`
	ModifiedBy                             string         `db:"ModifiedBy"`
	ModifiedDate                           time.Time     `db:"ModifiedDate"`
	
	Account  	          				  Account                  `db:"Account"`
}
