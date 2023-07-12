package models

type MarketItem struct {
	ItemId            int    `bson:"item_id" json:"item_id"`
	TokenAddress      string `bson:"address" json:"address"`
	TokenId           int    `bson:"token_id" json:"token_id"`
	Seller            string `bson:"seller" json:"seller"`
	Owner             string `bson:"owner" json:"owner"`
	Price             int64  `bson:"price" json:"price"`
	IsOfferable       bool   `bson:"is_offerable" json:"is_offerable"`
	AcceptVisaPayment bool   `bson:"accept_visa_payment" json:"accept_visa_payment"` //visa or paypal
	CurrentOfferValue int    `bson:"current_offer_value" json:"current_offer_value"`
	CurrentOfferer    string `bson:"current_offerer" json:"current_offerer"`
	Sold              bool   `bson:"sold" json:"sold"`
}
