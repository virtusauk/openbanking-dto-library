package dto

import (
	"time"
)

type Account struct {
	AccountId               int    `db:"AccountId"`
	SchemeName              string `db:"SchemeName"`
	AccountIdentification   string `db:"AccountIdentification"`
	SecondaryIdentification string `db:"SecondaryIdentification"`
	BankId                  int    `db:"BankId"`
	//Bank					Bank
	BranchId int `db:"BranchId"`
	//Branches				Branches
	AccountName string `db:"AccountName"`
	Nickname    string `db:"Nickname"`
	ProductId   int    `db:"ProductId"`
	//Product					Product
	AccountCurrency       string    `db:"AccountCurrency"`
	Balance               float64   `db:"Balance"`
	Status                string    `db:"Status"`
	TypeOfBalance         string    `db:"TypeOfBalance"`
	AccountOpeningDate    time.Time `db:"AccountOpeningDate"`
	AccountClosingDate    time.Time `db:"AccountClosingDate"`
	AccountTypeId         int       `db:"AccountTypeId"`
	Accounttype           Accounttype
	Accountrefnumber      string    `db:"Accountrefnumber"`
	Isjointaccount        string    `db:"Isjointaccount"`
	Isonlineaccessenabled string    `db:"Isonlineaccessenabled"`
	Usage                 string    `db:"Usage"`
	NomineeName           string    `db:"NomineeName"`
	NomineeAddress        string    `db:"NomineeAddress"`
	NomineePhoneNo        string    `db:"NomineePhoneNo"`
	NomineeDob            time.Time `db:"NomineeDob"`
	NomineeRelatonship    string    `db:"NomineeRelatonship"`
	CardFacility          string    `db:"CardFacility"`
	PassbookFacility      string    `db:"PassbookFacility"`
	ChequebookFacility    string    `db:"ChequebookFacility"`
	Frozen                string    `db:"Frozen"`
	NoDebit               string    `db:"NoDebit"`
	NoCredit              string    `db:"NoCredit"`
	CreditDebitIndicator  string    `db:"CreditDebitIndicator"`
	CreditLineIncluded    string    `db:"CreditLineIncluded"`
	CreditLineAmount      float64   `db:"CreditLineAmount"`
	CreditLineType        string    `db:"CreditLineType"`
	MakerDate             time.Time `db:"MakerDate"`
	CheckerDate           time.Time `db:"CheckerDate"`
	MakerId               string    `db:"MakerId"`
	CheckerId             string    `db:"CheckerId"`
	ModifiedBy            string    `db:"ModifiedBy"`
	ModifiedDate          time.Time `db:"ModifiedDate"`
}

func (Account) TableName() string {
	return "Account"
}
