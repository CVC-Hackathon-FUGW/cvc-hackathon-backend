package models

type MarketCollection struct {
	CollectionId   int     `bson:"collection_id" json:"collection_id"`
	CollectionName *string `bson:"collection_name" json:"collection_name"`
	TokenAddress   *string `bson:"token_address" json:"token_address"`
	Image          *string `bson:"image" json:"image"`
	Volume         int64   `bson:"volume" json:"volume"`
	IsActive       bool    `bson:"is_active" json:"is_active"`
}
