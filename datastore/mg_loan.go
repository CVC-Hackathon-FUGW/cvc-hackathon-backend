package datastore

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DatastoreLoanMG struct {
	loanCollection *mongo.Collection
}

func NewDatastoreLoanMG(loanCollection *mongo.Collection) *DatastoreLoanMG {
	return &DatastoreLoanMG{loanCollection}
}

var _ models.DatastoreLoan = (*DatastoreLoanMG)(nil)

func (ds DatastoreLoanMG) Create(ctx context.Context, params *models.Loan) (*models.Loan, error) {
	count, err := ds.loanCollection.CountDocuments(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	params.LoanId = int(count) + 1
	_, err = ds.loanCollection.InsertOne(ctx, params)
	if err != nil {
		return nil, err
	}

	return params, nil
}

func (ds DatastoreLoanMG) FindByID(ctx context.Context, id *string) (*models.Loan, error) {
	var loan *models.Loan
	query := bson.D{bson.E{Key: "loan_id", Value: id}}
	err := ds.loanCollection.FindOne(ctx, query).Decode(&loan)
	return loan, err
}

func (ds DatastoreLoanMG) List(ctx context.Context) ([]*models.Loan, error) {
	var loans []*models.Loan
	cursor, err := ds.loanCollection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var loan models.Loan
		err := cursor.Decode(&loan)
		if err != nil {
			return nil, err
		}
		loans = append(loans, &loan)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	if len(loans) == 0 {
		return nil, errors.New("documents not found")
	}
	return loans, nil
}

func (ds DatastoreLoanMG) Update(ctx context.Context, params *models.Loan) (*models.Loan, error) {
	filter := bson.D{primitive.E{Key: "loan_id", Value: params.LoanId}}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "loan_id", Value: params.LoanId},
			primitive.E{Key: "lender", Value: params.Lender},
			primitive.E{Key: "borrower", Value: params.Borrower},
			primitive.E{Key: "amount", Value: params.Amount},
			primitive.E{Key: "start_time", Value: params.StartTime},
			primitive.E{Key: "duration", Value: params.Duration},
			primitive.E{Key: "token_id", Value: params.TokenId},
			primitive.E{Key: "pool_id", Value: params.PoolId},
			primitive.E{Key: "token_address", Value: params.TokenAddress},
			primitive.E{Key: "updated_at", Value: time.Now()},
		}}}

	var loanUpdated *models.Loan
	err := ds.loanCollection.FindOneAndUpdate(ctx, filter, update).Decode(&loanUpdated)
	if err != nil {
		return nil, errors.New("no matched document found for update")
	}

	return loanUpdated, nil
}

func (ds DatastoreLoanMG) Delete(ctx context.Context, id *string) error {
	filter := bson.D{primitive.E{Key: "loan_id", Value: id}}
	result, _ := ds.loanCollection.DeleteOne(ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}

func (ds DatastoreLoanMG) MaxAmount(ctx context.Context, poolId *string) ([]*models.Loan, error) {
	var loans []*models.Loan
	cursor, err := ds.loanCollection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	max := -1
	for cursor.Next(ctx) {
		var loan models.Loan
		err := cursor.Decode(&loan)
		if err != nil {
			return nil, err
		}
		if strconv.Itoa(loan.PoolId) != *poolId {
			continue
		}

		if loan.Amount > max {
			loans = []*models.Loan{&loan}
			max = loan.Amount
		}
		if loan.Amount == max {
			loans = append(loans, &loan)
		}
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	if len(loans) == 0 {
		return nil, errors.New("documents not found")
	}
	return loans, nil
}
