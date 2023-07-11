package models

import (
	"context"

	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/enum"
)

type DatastorePool interface {
	Create(ctx context.Context, params *Pool) (*Pool, error)
	FindByID(ctx context.Context, id *string) (*Pool, error)
	List(ctx context.Context, params enum.PoolParams) ([]*Pool, error)
	Update(ctx context.Context, params *Pool) (*Pool, error)
	Delete(ctx context.Context, params *string) error
}

type DatastoreLoan interface {
	Create(ctx context.Context, params *Loan) (*Loan, error)
	FindByID(ctx context.Context, id *string) (*Loan, error)
	List(ctx context.Context) ([]*Loan, error)
	Update(ctx context.Context, params *Loan) (*Loan, error)
	Delete(ctx context.Context, params *string) error
	MaxAmount(ctx context.Context, poolId *string) ([]*Loan, error)
	CountLoans(ctx context.Context, poolId *string) (*enum.CountLoans, error)
}

type DatastoreLender interface {
	Create(ctx context.Context, params *Lender) (*Lender, error)
	FindByID(ctx context.Context, id *string) (*Lender, error)
	List(ctx context.Context) ([]*Lender, error)
	Update(ctx context.Context, params *Lender) (*Lender, error)
	Delete(ctx context.Context, params *string) error
}

type DatastoreBorrower interface {
	Create(ctx context.Context, params *Borrower) (*Borrower, error)
	FindByID(ctx context.Context, id *string) (*Borrower, error)
	List(ctx context.Context) ([]*Borrower, error)
	Update(ctx context.Context, params *Borrower) (*Borrower, error)
	Delete(ctx context.Context, params *string) error
}

type DatastoreMarketItem interface {
	Create(ctx context.Context, params *MarketItem) (*MarketItem, error)
	FindByID(ctx context.Context, id *string) (*MarketItem, error)
	List(ctx context.Context) ([]*MarketItem, error)
	Update(ctx context.Context, params *MarketItem) (*MarketItem, error)
	Delete(ctx context.Context, params *string) error
}

type DatastoreMarketCollection interface {
	Create(ctx context.Context, params *MarketCollection) (*MarketCollection, error)
	FindByID(ctx context.Context, id *string) (*MarketCollection, error)
	List(ctx context.Context, params enum.MarketCollectionsParams) ([]*MarketCollection, error)
	Update(ctx context.Context, params *MarketCollection) (*MarketCollection, error)
	Delete(ctx context.Context, params *string) error
}
