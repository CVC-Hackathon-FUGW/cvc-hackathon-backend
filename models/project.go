package models

type Project struct {
	ProjectId          int     `bson:"project_id" json:"project_id"`
	ProjectOwner       *string `bson:"project_owner" json:"project_owner"`
	TotalRaiseAmount   *int64  `bson:"total_raise_amount" json:"total_raise_amount"`
	DueTime            *int    `bson:"due_time" json:"due_time"`
	ProjectName        *string `bson:"project_name" json:"project_name"`
	ProjectImage       *string `bson:"project_image" json:"project_image"`
	ProjectDescription *string `bson:"project_description" json:"project_description"`
	ProjectAddress     *string `bson:"project_address" json:"project_address"`
	TotalFundRaised    *int64  `bson:"total_fund_raised" json:"total_fund_raised"`
	IsActive           bool    `bson:"is_active" json:"is_active"`
}
