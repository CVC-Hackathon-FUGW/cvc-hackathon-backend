package models

type Participant struct {
	ParticipantId      int     `bson:"participant_id" json:"participant_id"`
	ProjectName        *string `bson:"project_name" json:"project_name"`
	FundAttended       *int64  `bson:"fund_attended" json:"fund_attended"`
	ParticipantAddress *string `bson:"participant_address" json:"participant_address"`
	IsActive           bool    `bson:"is_active" json:"is_active"`
}
