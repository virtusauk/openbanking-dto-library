package dto

import (

	"time"


)

type Branches struct {
	BranchId               int            `db:"BranchId"`
	BankId                 int            `db:"BankId"`
	ServicerSchemename     string `db:"ServicerSchemename"`
	ServicerIdentification string `db:"ServicerIdentification"`
	UnitName               string         `db:"UnitName"`
	UnitType               string `db:"UnitType"`
	UnitLocation           string `db:"UnitLocation"`
	ParentUnit             int  `db:"ParentUnit"`
	AddressType            string `db:"AddressType"`
	Department             string `db:"Department"`
	Subdepartment          string `db:"Subdepartment"`
	StreetName             string `db:"StreetName"`
	BuildingNumber         string `db:"BuildingNumber"`
	Townname               string `db:"Townname"`
	Addressline1           string `db:"Addressline1"`
	Addressline2           string `db:"Addressline2"`
	Addressline3           string `db:"Addressline3"`
	Addressline4           string `db:"Addressline4"`
	Addressline5           string `db:"Addressline5"`
	Addressline6           string `db:"Addressline6"`
	Addressline7           string `db:"Addressline7"`
	Countrysubdivision     string `db:"Countrysubdivision"`
	Country                string `db:"Country"`
	Address                string `db:"Address"`
	Postcode               string         `db:"Postcode"`
	EmailAddress           string `db:"EmailAddress"`
	PhoneNumber            string `db:"PhoneNumber"`
	LangCode               string `db:"LangCode"`
	IsoCode                string `db:"IsoCode"`
	AppLink                string `db:"AppLink"`
	MakerDate              time.Time      `db:"MakerDate"`
	CheckerDate            time.Time    `db:"CheckerDate"`
	MakerId                string         `db:"MakerId"`
	CheckerId              string `db:"CheckerId"`
	ModifiedBy             string `db:"ModifiedBy"`
	ModifiedDate           time.Time    `db:"ModifiedDate"`
}
