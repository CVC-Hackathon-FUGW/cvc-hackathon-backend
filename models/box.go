package models

type Box struct {
	BoxId      int     `bson:"box_id" json:"box_id"`
	TokenId    int     `bson:"token_id" json:"token_id"`
	BoxAddress *string `bson:"box_address" json:"box_address"`
	BoxPrice   *int64  `bson:"box_price" json:"box_price"`
	IsOpened   *bool   `bson:"is_opened" json:"is_opened"`
	Owner      *string `bson:"owner" json:"owner"`
	IsActive   bool    `bson:"is_active" json:"is_active"`
}
