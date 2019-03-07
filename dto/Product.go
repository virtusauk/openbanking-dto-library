package dto

import (
	"time"
)

type Product struct {
	ProductId                          int             `db:"ProductId"`
	ProductName                        string          `db:"ProductName"`
	CreationDate                       time.Time       `db:"CreationDate"`
	BankId                             int             `db:"BankId"`
	Description                        string  `db:"Description"`
	LastModifiedDate                   time.Time     `db:"LastModifiedDate"`
	SecondaryProductId                 string  `db:"SecondaryProductId"`
	MarketingStateId                   string  `db:"MarketingStateId"`
	ProductType                        string  `db:"ProductType"`
	ProductSubType                     string          `db:"ProductSubType"`
	ProductGroup                       string  `db:"ProductGroup"`
	InterestCalculationBalance         string  `db:"InterestCalculationBalance"`
	Activated                          string  `db:"Activated"`
	InterestPaymentPoint               string  `db:"InterestPaymentPoint"`
	CollectInterestWhenLocked          string  `db:"CollectInterestWhenLocked"`
	RecommendedDepositAmount           float64 `db:"RecommendedDepositAmount"`
	MaxWidthdrawlAmount                float64 `db:"MaxWidthdrawlAmount"`
	MinOpeningBalance                  float64 `db:"MinOpeningBalance"`
	MaxOpeningBalance                  float64 `db:"MaxOpeningBalance"`
	DefaultOpeningBalance              float64 `db:"DefaultOpeningBalance"`
	MinMaturityPeriod                  int   `db:"MinMaturityPeriod"`
	MaxMaturityPeriod                  int   `db:"MaxMaturityPeriod"`
	DefaultMaturityPeriod              int   `db:"DefaultMaturityPeriod"`
	MaxOverdraftLimit                  float64 `db:"MaxOverdraftLimit"`
	AllowOverdraft                     string  `db:"AllowOverdraft"`
	MaturityPeriodUnit                 string  `db:"MaturityPeriodUnit"`
	SavingsFees                        string  `db:"SavingsFees"`
	AllowArbitraryFees                 string  `db:"AllowArbitraryFees"`
	IdGeneratorType                    string  `db:"IdGeneratorType"`
	IdPattern                          string  `db:"IdPattern"`
	AccountingMethod                   string  `db:"AccountingMethod"`
	SavingsProductRules                string  `db:"SavingsProductRules"`
	DormancyPeriodDays                 int   `db:"DormancyPeriodDays"`
	OverdraftDaysInyear                string  `db:"OverdraftDaysInyear"`
	WithholdingTaxenabled              string  `db:"WithholdingTaxenabled"`
	LineOfCreditRequirement            string  `db:"LineOfCreditRequirement"`
	ForGroups                          string  `db:"ForGroups"`
	RefTemplates                       int   `db:"RefTemplates"`
	ForIndividuals                     string  `db:"ForIndividuals"`
	InterestPaidIntoAccount            string  `db:"InterestPaidIntoAccount"`
	RefInterestAccruedAccountingMethod int   `db:"RefInterestAccruedAccountingMethod"`
	InterestDaysInyear                 int   `db:"InterestDaysInyear"`
	ProductIdentifier                  string  `db:"ProductIdentifier"`
	Segment                            string  `db:"Segment"`
	FeeFreeLengthPeriod                int   `db:"FeeFreeLengthPeriod"`
	Note                               string  `db:"Note"`
	MakerDate                          time.Time       `db:"MakerDate"`
	CheckerDate                        time.Time     `db:"CheckerDate"`
	MakerId                            string          `db:"MakerId"`
	CheckerId                          string  `db:"CheckerId"`
	ModifiedBy                         string  `db:"ModifiedBy"`
}
