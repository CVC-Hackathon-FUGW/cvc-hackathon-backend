package models

type Participant struct {
	ParticipantId      int     `bson:"participant_id" json:"participant_id"`
	ProjectId          *int    `bson:"project_id" json:"project_id"`
	FundAttended       *int64  `bson:"fund_attended" json:"fund_attended"`
	ParticipantAddress *string `bson:"participant_address" json:"participant_address"`
	IsActive           bool    `bson:"is_active" json:"is_active"`
}
