package dto

type Accounttype struct {
	AccountTypeId  int    `db:"AccountTypeId"`
	BankId         int    `db:"BankId"`
	AccountType    string `db:"AccountType"`
	AccountSubType string `db:"AccountSubType"`
	Description    string `db:"Description"`
	Active         string `db:"Active"`
}

func (Accounttype) TableName() string {
	return "AccountType"
}
