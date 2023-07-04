package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Pool struct {
	ID              primitive.ObjectID `bson:"_id" json:"_id"`
	PoolId          int                `bson:"pool_id" json:"pool_id"`
	TokenAddress    string             `bson:"token_address" json:"token_address"`
	TotalPoolAmount int                `bson:"total_pool_amount" json:"total_pool_amount"`
	APY             int                `bson:"apy" json:"apy"`
	Duration        int                `bson:"duration" json:"duration"`
	State           bool               `bson:"state" json:"state"`
}
