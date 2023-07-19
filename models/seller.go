package models

type Seller struct {
	SellerID   int    `bson:"seller_id" json:"seller_id"`
	Address    string `bson:"address" json:"address"`
	MerchantId string `bson:"merchant_id" json:"merchant_id"`
}
