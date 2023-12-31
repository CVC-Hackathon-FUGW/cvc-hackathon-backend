package models

type Pool struct {
	PoolId          int     `bson:"pool_id" json:"pool_id"`
	TokenAddress    *string `bson:"token_address" json:"token_address"`
	CollectionName  *string `bson:"collection_name" json:"collection_name"`
	TotalPoolAmount *int64  `bson:"total_pool_amount" json:"total_pool_amount"`
	APY             *int    `bson:"apy" json:"apy"`
	Duration        *int    `bson:"duration" json:"duration"`
	State           *bool   `bson:"state" json:"state"`
	Image           *string `bson:"image" json:"image"`
	IsActive        bool    `bson:"is_active" json:"is_active"`
}

type PoolWithLoanDetails struct {
	Pool          *Pool  `json:"pool"`
	LoanMaxAmount *int64 `json:"loan_max_amount"`
	LoanCount     string `json:"loan_count"`
}
