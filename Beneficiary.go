package dto

import (
	"time"
)

type Beneficiary struct {
	BeneficiaryId                                     int            `db:"BeneficiaryId"`
	DebtorAccountId                                   int            `db:"DebtorAccountId"`
	DebtorAccountOwnerId                              int            `db:"DebtorAccountOwnerId"`
	BankId                                            int            `db:"BankId"`
	BeneficiaryAccountRefId                           string         `db:"BeneficiaryAccountRefId"`
	BeneficiaryRefId                                  string         `db:"BeneficiaryRefId"`
	BeneficiaryTranReference                          string         `db:"BeneficiaryTranReference"`
	BeneficiaryBankCode                               string         `db:"BeneficiaryBankCode"`
	BeneficiaryBankName                               string         `db:"BeneficiaryBankName"`
	BeneficiaryCreditorAgentSchemename                string         `db:"BeneficiaryCreditorAgentSchemename"`
	BeneficiaryCreditorAgentIdentification            string         `db:"BeneficiaryCreditorAgentIdentification"`
	BeneficiaryCreditorAccountSchemename              string `db:"BeneficiaryCreditorAccountSchemename"`
	BeneficiaryCreditorAccountIdentification          string         `db:"BeneficiaryCreditorAccountIdentification"`
	BeneficiaryCreditorAccountSecondaryIdentification string         `db:"BeneficiaryCreditorAccountSecondaryIdentification"`
	BeneficiaryCreditorAccountName                    string         `db:"BeneficiaryCreditorAccountName"`
	BeneficiaryServicerSchemename                    string         `db:"BeneficiaryServicerSchemename"`
	BeneficiaryServicerIdentification                string         `db:"BeneficiaryServicerIdentification"`
	BeneficiaryCreditorAgentAddressType               string        `db:"BeneficiaryCreditorAgentAddressType"`
	BeneficiaryCreditorAgentDepartment               string         `db:"BeneficiaryCreditorAgentDepartment"`
	BeneficiaryCreditorAgentSubdepartment            string         `db:"BeneficiaryCreditorAgentSubdepartment"`
	BeneficiaryCreditorAgentStreetName               string         `db:"BeneficiaryCreditorAgentStreetName"`
	BeneficiaryCreditorAgentBuildingNumber           string         `db:"BeneficiaryCreditorAgentBuildingNumber"`
	BeneficiaryCreditorAgentTownname                 string          `db:"BeneficiaryCreditorAgentTownname"`
	BeneficiaryCreditorAgentAddressline1              string         `db:"BeneficiaryCreditorAgentAddressline1"`
	BeneficiaryCreditorAgentAddressline2              string         `db:"BeneficiaryCreditorAgentAddressline2"`
	BeneficiaryCreditorAgentAddressline3              string        `db:"BeneficiaryCreditorAgentAddressline3"`
	BeneficiaryCreditorAgentAddressline4             string         `db:"BeneficiaryCreditorAgentAddressline4"`
	BeneficiaryCreditorAgentAddressline5             string         `db:"BeneficiaryCreditorAgentAddressline5"`
	BeneficiaryCreditorAgentAddressline6              string        `db:"BeneficiaryCreditorAgentAddressline6"`
	BeneficiaryCreditorAgentAddressline7             string       `db:"BeneficiaryCreditorAgentAddressline7"`
	BeneficiaryCreditorAgentPostcode                  string      `db:"BeneficiaryCreditorAgentPostcode"`
	BeneficiaryCreditorAgentCountrysubdivision        string      `db:"BeneficiaryCreditorAgentCountrysubdivision"`
	BeneficiaryCreditorAgentCountry                   string       `db:"BeneficiaryCreditorAgentCountry"`
	MakerDate                                         time.Time    `db:"MakerDate"`
	CheckerDate                                       time.Time    `db:"CheckerDate"`
	MakerId                                           string       `db:"MakerId"`
	CheckerId                                         string       `db:"CheckerId"`
	ModifiedBy                                        string       `db:"ModifiedBy"`
	ModifiedDate                                      time.Time    `db:"ModifiedDate"`
	
	Account  	          				  Account                  `gorm:"auto_preload"`
}

func (Beneficiary) TableName() string {
	return "Beneficiary"
}