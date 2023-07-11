package models

type MarketCollection struct {
	CollectionId int    `bson:"collection_id" json:"collection_id"`
	TokenAddress string `bson:"token_address" json:"token_address"`
	Image        string `bson:"image" json:"image"`
}
