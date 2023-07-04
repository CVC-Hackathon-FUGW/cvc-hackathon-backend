package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Loan struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	LoanId       int                `bson:"loan_id" json:"loan_id"`
	Lender       string             `bson:"lender" json:"lender"`
	Borrower     string             `bson:"borrower" json:"borrower"`
	Amount       int                `bson:"amount" json:"amount"`
	StartTime    int                `bson:"start_time" json:"start_time"`
	Duration     int                `bson:"duration" json:"duration"`
	TokenId      int                `bson:"token_id" json:"token_id"`
	PoolId       int                `bson:"pool_id" json:"pool_id"`
	TokenAddress string             `bson:"token_address" json:"token_address"`
	State        bool               `bson:"state" json:"state"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
}
