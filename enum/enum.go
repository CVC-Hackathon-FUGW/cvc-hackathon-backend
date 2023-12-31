package enum

type PoolParams struct {
	Name string
}

type MarketCollectionsParams struct {
	Name string
}

type CheckinParams struct {
	Name string
	Sort string
}

type CountLoans struct {
	TotalLoanInPool int
	TotalLoanGot    int
}

type LoanParams struct {
	WithPool bool
}
