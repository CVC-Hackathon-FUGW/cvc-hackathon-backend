package models

type Checkin struct {
	CheckinId          int     `json:"checkin_id" bson:"checkin_id"`
	Wallet             *string `json:"wallet" bson:"wallet"`
	NumberChecked      *int    `json:"number_checked" bson:"number_checked"`
	CurrentTokenAmount *int64  `json:"current_token_amount" bson:"current_token_amount"`
	IsActive           bool    `json:"is_active" bson:"is_active"`
}
