package models

type Lender struct {
	LenderID      string `bson:"lender_id" json:"lender_id"`
	WalletAddress string `bson:"wallet_address" json:"wallet_address"`
	Offers        []Loan `bson:"offers" json:"offers"`
}
