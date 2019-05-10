// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

var TableNames = struct {
	Account                        string
	AccountCreditLine              string
	AccountLink                    string
	AccountOwners                  string
	AccountRequest                 string
	AccountRequestHistory          string
	AccountType                    string
	Address                        string
	Application                    string
	AssetCondition                 string
	Assets                         string
	AssetsType                     string
	Banks                          string
	Beneficiary                    string
	Biller                         string
	Borrower                       string
	Branches                       string
	Budget                         string
	Centre                         string
	CodeData                       string
	Collateral                     string
	Commodity                      string
	Country                        string
	CreditOrginationAssetHistory   string
	CreditOrigination              string
	Cryptography                   string
	CurrencyMaster                 string
	CurrentHolding                 string
	CustomerGoal                   string
	CustomerGoalSummary            string
	CustomerOffer                  string
	CustomerRating                 string
	DerivedTxn                     string
	Devices                        string
	DirectDebits                   string
	Document                       string
	DocumentType                   string
	Emails                         string
	ExchangeRate                   string
	FilePayment                    string
	FiscalYear                     string
	FundConfirmation               string
	FundDescription                string
	FundHistoricalReturn           string
	FundHistoricalReturnCovariance string
	FundPrice                      string
	FxDeals                        string
	FxPayment                      string
	GLAccount                      string
	Gateway                        string
	GoalMaster                     string
	Goals                          string
	InterestProduct                string
	InvestmentInLending            string
	InvestmentType                 string
	JointAccount                   string
	JournalEntry                   string
	JournalEntryLine               string
	KYC                            string
	LineOfCredit                   string
	Loan                           string
	LoanInterest                   string
	LoanSchedule                   string
	MccCode                        string
	Notifications                  string
	OBPFundConfirmationConsent     string
	OBPPaymentConsent              string
	Offer                          string
	Organization                   string
	OrganizationContacts           string
	OrganizationFinancialInfo      string
	OrganizationType               string
	Parties                        string
	Payment                        string
	PaymentAch                     string
	PaymentBill                    string
	PaymentCharges                 string
	PaymentCheque                  string
	PaymentFxTrade                 string
	PaymentGateway                 string
	PaymentInitiation              string
	PaymentLimit                   string
	PaymentLoan                    string
	PaymentMessage                 string
	PaymentMethod                  string
	PaymentRtp                     string
	PaymentWire                    string
	Person                         string
	PersonFinancialInfo            string
	PersonType                     string
	PhoneNumbers                   string
	Portfolio                      string
	PortfolioGoalSummary           string
	Product                        string
	Recommendations                string
	RegisteredBiller               string
	Resources                      string
	ResourcesType                  string
	Risk                           string
	RiskType                       string
	RoleResources                  string
	Roles                          string
	RTP                            string
	RtpRequest                     string
	RtpToken                       string
	SchedulePayment                string
	Shipment                       string
	SocialInfo                     string
	StandingOrders                 string
	Statement                      string
	StatementAmount                string
	StatementBenefit               string
	StatementFee                   string
	StatementInterest              string
	TFBankGuarantee                string
	TFChargeCodes                  string
	TFCharges                      string
	TFChargesTxn                   string
	TFCommodity                    string
	TFExportImportsBill            string
	TFInterest                     string
	TFInterestTxn                  string
	TFLcLetter                     string
	TFLcLetterTxn                  string
	TFLetterOfCredit               string
	TFLetterOfCreditTxn            string
	TFParty                        string
	TFPartyTxn                     string
	TFRisk                         string
	TFRiskType                     string
	TFSettlement                   string
	TFShipment                     string
	TFTracer                       string
	TFTrade                        string
	TFTradeItem                    string
	TierBand                       string
	TierBandSet                    string
	Trade                          string
	Transaction                    string
	TransactionFxInfo              string
	TransactionType                string
	Transfer                       string
	TransferAch                    string
	TransferMethod                 string
	TransferWire                   string
	UserLogin                      string
}{
	Account:                        "Account",
	AccountCreditLine:              "AccountCreditLine",
	AccountLink:                    "AccountLink",
	AccountOwners:                  "AccountOwners",
	AccountRequest:                 "AccountRequest",
	AccountRequestHistory:          "AccountRequestHistory",
	AccountType:                    "AccountType",
	Address:                        "Address",
	Application:                    "Application",
	AssetCondition:                 "AssetCondition",
	Assets:                         "Assets",
	AssetsType:                     "AssetsType",
	Banks:                          "Banks",
	Beneficiary:                    "Beneficiary",
	Biller:                         "Biller",
	Borrower:                       "Borrower",
	Branches:                       "Branches",
	Budget:                         "Budget",
	Centre:                         "Centre",
	CodeData:                       "CodeData",
	Collateral:                     "Collateral",
	Commodity:                      "Commodity",
	Country:                        "Country",
	CreditOrginationAssetHistory:   "CreditOrginationAssetHistory",
	CreditOrigination:              "CreditOrigination",
	Cryptography:                   "Cryptography",
	CurrencyMaster:                 "CurrencyMaster",
	CurrentHolding:                 "CurrentHolding",
	CustomerGoal:                   "CustomerGoal",
	CustomerGoalSummary:            "CustomerGoalSummary",
	CustomerOffer:                  "CustomerOffer",
	CustomerRating:                 "CustomerRating",
	DerivedTxn:                     "DerivedTxn",
	Devices:                        "Devices",
	DirectDebits:                   "DirectDebits",
	Document:                       "Document",
	DocumentType:                   "DocumentType",
	Emails:                         "Emails",
	ExchangeRate:                   "ExchangeRate",
	FilePayment:                    "FilePayment",
	FiscalYear:                     "FiscalYear",
	FundConfirmation:               "FundConfirmation",
	FundDescription:                "FundDescription",
	FundHistoricalReturn:           "FundHistoricalReturn",
	FundHistoricalReturnCovariance: "FundHistoricalReturnCovariance",
	FundPrice:                      "FundPrice",
	FxDeals:                        "FxDeals",
	FxPayment:                      "FxPayment",
	GLAccount:                      "GLAccount",
	Gateway:                        "Gateway",
	GoalMaster:                     "GoalMaster",
	Goals:                          "Goals",
	InterestProduct:                "InterestProduct",
	InvestmentInLending:            "InvestmentInLending",
	InvestmentType:                 "InvestmentType",
	JointAccount:                   "JointAccount",
	JournalEntry:                   "JournalEntry",
	JournalEntryLine:               "JournalEntryLine",
	KYC:                            "KYC",
	LineOfCredit:                   "LineOfCredit",
	Loan:                           "Loan",
	LoanInterest:                   "LoanInterest",
	LoanSchedule:                   "LoanSchedule",
	MccCode:                        "MccCode",
	Notifications:                  "Notifications",
	OBPFundConfirmationConsent:     "OBPFundConfirmationConsent",
	OBPPaymentConsent:              "OBPPaymentConsent",
	Offer:                          "Offer",
	Organization:                   "Organization",
	OrganizationContacts:           "OrganizationContacts",
	OrganizationFinancialInfo:      "OrganizationFinancialInfo",
	OrganizationType:               "OrganizationType",
	Parties:                        "Parties",
	Payment:                        "Payment",
	PaymentAch:                     "PaymentAch",
	PaymentBill:                    "PaymentBill",
	PaymentCharges:                 "PaymentCharges",
	PaymentCheque:                  "PaymentCheque",
	PaymentFxTrade:                 "PaymentFxTrade",
	PaymentGateway:                 "PaymentGateway",
	PaymentInitiation:              "PaymentInitiation",
	PaymentLimit:                   "PaymentLimit",
	PaymentLoan:                    "PaymentLoan",
	PaymentMessage:                 "PaymentMessage",
	PaymentMethod:                  "PaymentMethod",
	PaymentRtp:                     "PaymentRtp",
	PaymentWire:                    "PaymentWire",
	Person:                         "Person",
	PersonFinancialInfo:            "PersonFinancialInfo",
	PersonType:                     "PersonType",
	PhoneNumbers:                   "PhoneNumbers",
	Portfolio:                      "Portfolio",
	PortfolioGoalSummary:           "PortfolioGoalSummary",
	Product:                        "Product",
	Recommendations:                "Recommendations",
	RegisteredBiller:               "RegisteredBiller",
	Resources:                      "Resources",
	ResourcesType:                  "ResourcesType",
	Risk:                           "Risk",
	RiskType:                       "RiskType",
	RoleResources:                  "RoleResources",
	Roles:                          "Roles",
	RTP:                            "Rtp",
	RtpRequest:                     "RtpRequest",
	RtpToken:                       "RtpToken",
	SchedulePayment:                "SchedulePayment",
	Shipment:                       "Shipment",
	SocialInfo:                     "SocialInfo",
	StandingOrders:                 "StandingOrders",
	Statement:                      "Statement",
	StatementAmount:                "StatementAmount",
	StatementBenefit:               "StatementBenefit",
	StatementFee:                   "StatementFee",
	StatementInterest:              "StatementInterest",
	TFBankGuarantee:                "TFBankGuarantee",
	TFChargeCodes:                  "TFChargeCodes",
	TFCharges:                      "TFCharges",
	TFChargesTxn:                   "TFChargesTxn",
	TFCommodity:                    "TFCommodity",
	TFExportImportsBill:            "TFExportImportsBill",
	TFInterest:                     "TFInterest",
	TFInterestTxn:                  "TFInterestTxn",
	TFLcLetter:                     "TFLcLetter",
	TFLcLetterTxn:                  "TFLcLetterTxn",
	TFLetterOfCredit:               "TFLetterOfCredit",
	TFLetterOfCreditTxn:            "TFLetterOfCreditTxn",
	TFParty:                        "TFParty",
	TFPartyTxn:                     "TFPartyTxn",
	TFRisk:                         "TFRisk",
	TFRiskType:                     "TFRiskType",
	TFSettlement:                   "TFSettlement",
	TFShipment:                     "TFShipment",
	TFTracer:                       "TFTracer",
	TFTrade:                        "TFTrade",
	TFTradeItem:                    "TFTradeItem",
	TierBand:                       "TierBand",
	TierBandSet:                    "TierBandSet",
	Trade:                          "Trade",
	Transaction:                    "Transaction",
	TransactionFxInfo:              "TransactionFxInfo",
	TransactionType:                "TransactionType",
	Transfer:                       "Transfer",
	TransferAch:                    "TransferAch",
	TransferMethod:                 "TransferMethod",
	TransferWire:                   "TransferWire",
	UserLogin:                      "UserLogin",
}
