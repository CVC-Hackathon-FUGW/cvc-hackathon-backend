package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Lender struct {
	ID            primitive.ObjectID `bson:"_id" json:"_id"`
	WalletAddress string             `bson:"wallet_address" json:"wallet_address"`
	Offers        []Loan             `bson:"offers" json:"offers"`
}
