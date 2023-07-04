package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Borrower struct {
	ID            primitive.ObjectID `bson:"_id" json:"_id"`
	WalletAddress string             `bson:"wallet_address" json:"wallet_address"`
	Loans         []Loan             `bson:"loans" json:"loans"`
}
