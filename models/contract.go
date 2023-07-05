package models

import "context"

type DatastorePool interface {
	Create(ctx context.Context, params *Pool) (*Pool, error)
	FindByID(ctx context.Context, id *string) (*Pool, error)
	List(ctx context.Context) ([]*Pool, error)
	Update(ctx context.Context, params *Pool) (*Pool, error)
	Delete(ctx context.Context, params *string) error
}