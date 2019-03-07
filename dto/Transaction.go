package dto

import (
	"time"
)

type Transaction struct {
	TransactionId                     int       `db:"TransactionId"`
	TransactionReference              string    `db:"TransactionReference"`
	BankId                            int       `db:"BankId"`
	//Banks							  Banks
	AccountId                         int       `db:"AccountId"`
	//Account							  Account
	CreditDebitIndicator              string    `db:"CreditDebitIndicator"`
	PaymentId                         int       `db:"PaymentId"`
	TransferId                        string    `db:"TransferId"`
	PaymentRefId                      string    `db:"PaymentRefId"`
	CounterpartyAccountId             int       `db:"CounterpartyAccountId"`
	TransactionAmount                 float64   `db:"TransactionAmount"`
	CurrencyCode                      string    `db:"CurrencyCode"`
	Currencymaster					  Currencymaster
	ChargeAmount                      float64   `db:"ChargeAmount"`
	ChargeCurrency                    string    `db:"ChargeCurrency"`
	TransactionType                   string    `db:"TransactionType"`
	CounterpartyBankId                int64     `db:"CounterpartyBankId"`
	BankLocation                      string    `db:"BankLocation"`
	CounterpartyBankLocation          string    `db:"CounterpartyBankLocation"`
	SwiftCode                         string    `db:"SwiftCode"`
	SwiftCodeTrace                    string    `db:"SwiftCodeTrace"`
	Purpose                           string    `db:"Purpose"`
	Status                            string    `db:"Status"`
	Balance                           float64   `db:"Balance"`
	BalanceType                       string    `db:"BalanceType"`
	BalanceCreditDebitIndicator       string    `db:"BalanceCreditDebitIndicator"`
	ObTransactionType                 string    `db:"ObTransactionType"`
	OverdraftFeesAmount               float64   `db:"OverdraftFeesAmount"`
	InterestAmount                    float64   `db:"InterestAmount"`
	OverdraftAmount                   float64   `db:"OverdraftAmount"`
	FundsAmount                       float64   `db:"FundsAmount"`
	OverdraftInterestAmount           float64   `db:"OverdraftInterestAmount"`
	InterestRate                      float64   `db:"InterestRate"`
	OverdraftInterestRate             float64   `db:"OverdraftInterestRate"`
	Comment                           string    `db:"Comment"`
	AddressLine                       string    `db:"AddressLine"`
	TransactionName                   string    `db:"TransactionName"`
	ExpectedDisbursementDate          time.Time `db:"ExpectedDisbursementDate"`
	ExpectedDisbursementTime          []uint8   `db:"ExpectedDisbursementTime"`
	BookingDateTime                   time.Time `db:"BookingDateTime"`
	ValueDateTime                     time.Time `db:"ValueDateTime"`
	CounterpartyBankCode              string    `db:"CounterpartyBankCode"`
	CounterpartyBankName              string    `db:"CounterpartyBankName"`
	BankTransactionCode               string    `db:"BankTransactionCode"`
	BankTransactionSubcode            string    `db:"BankTransactionSubcode"`
	ProprietaryBankTransactionCode    string    `db:"ProprietaryBankTransactionCode"`
	ProprietaryBankTransactionSubcode string    `db:"ProprietaryBankTransactionSubcode"`
	IssuerBank                        string    `db:"IssuerBank"`
	MerchantName                      string    `db:"MerchantName"`
	MerchantCategoryCode              string    `db:"MerchantCategoryCode"`
	MakerDate                         time.Time `db:"MakerDate"`
	CheckerDate                       time.Time `db:"CheckerDate"`
	MakerId                           string    `db:"MakerId"`
	CheckerId                         string    `db:"CheckerId"`
	ModifiedBy                        string    `db:"ModifiedBy"`
	ModifiedDate                      time.Time `db:"ModifiedDate"`
}
func (Transaction) TableName() string {
	return "Transaction"
}
