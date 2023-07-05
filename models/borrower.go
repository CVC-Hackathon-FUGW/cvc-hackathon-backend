package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Borrower struct {
	BorrowerID    primitive.ObjectID `bson:"borrower_id" json:"borrower_id"`
	WalletAddress string             `bson:"wallet_address" json:"wallet_address"`
	Loans         []Loan             `bson:"loans" json:"loans"`
}
