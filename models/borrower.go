package models

type Borrower struct {
	BorrowerID    int    `bson:"borrower_id" json:"borrower_id"`
	WalletAddress string `bson:"wallet_address" json:"wallet_address"`
	Loans         []Loan `bson:"loans" json:"loans"`
	IsActive      bool   `bson:"is_active" json:"is_active"`
}
