package datastore

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/enum"
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

	params.IsActive = true
	params.LoanId = int(count) + 1

	_, err = ds.loanCollection.InsertOne(ctx, params)
	if err != nil {
		return nil, err
	}

	return params, nil
}

func (ds DatastoreLoanMG) FindByID(ctx context.Context, id *string) (*models.Loan, error) {
	var loan *models.Loan
	idInt, err := strconv.Atoi(*id)
	if err != nil {
		return nil, err
	}

	query := bson.D{bson.E{Key: "loan_id", Value: idInt}}
	err = ds.loanCollection.FindOne(ctx, query).Decode(&loan)
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
		if loan.IsActive == true {
			loans = append(loans, &loan)
		}
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	return loans, nil
}

func (ds DatastoreLoanMG) Update(ctx context.Context, params *models.Loan) (*models.Loan, error) {
	var loanDB *models.Loan

	query := bson.D{bson.E{Key: "loan_id", Value: params.LoanId}}
	err := ds.loanCollection.FindOne(ctx, query).Decode(&loanDB)
	if err != nil {
		return nil, err
	}

	if params.Lender != nil {
		loanDB.Lender = params.Lender
	}

	if params.Borrower != nil {
		loanDB.Borrower = params.Borrower
	}

	if params.Amount != nil {
		loanDB.Amount = params.Amount
	}

	if params.StartTime != nil {
		loanDB.StartTime = params.StartTime
	}

	if params.Duration != nil {
		loanDB.Duration = params.Duration
	}

	if params.TokenId != nil {
		loanDB.TokenId = params.TokenId
	}

	if params.PoolId != nil {
		loanDB.PoolId = params.PoolId
	}

	if params.TokenAddress != nil {
		loanDB.TokenAddress = params.TokenAddress
	}

	if params.State != nil {
		loanDB.State = params.State
	}

	filter := bson.D{primitive.E{Key: "loan_id", Value: params.LoanId}}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "lender", Value: loanDB.Lender},
			primitive.E{Key: "borrower", Value: loanDB.Borrower},
			primitive.E{Key: "amount", Value: loanDB.Amount},
			primitive.E{Key: "start_time", Value: loanDB.StartTime},
			primitive.E{Key: "duration", Value: loanDB.Duration},
			primitive.E{Key: "token_id", Value: loanDB.TokenId},
			primitive.E{Key: "pool_id", Value: loanDB.PoolId},
			primitive.E{Key: "token_address", Value: loanDB.TokenAddress},
			primitive.E{Key: "state", Value: loanDB.State},
			primitive.E{Key: "updated_at", Value: time.Now()},
		}}}

	var loanUpdated *models.Loan
	err = ds.loanCollection.FindOneAndUpdate(ctx, filter, update).Decode(&loanUpdated)
	if err != nil {
		return nil, errors.New("no matched document found for update")
	}

	return loanUpdated, nil
}

func (ds DatastoreLoanMG) Delete(ctx context.Context, id *string) error {
	idInt, err := strconv.Atoi(*id)
	if err != nil {
		return err
	}

	filter := bson.D{primitive.E{Key: "loan_id", Value: idInt}}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "loan_id", Value: idInt},
			primitive.E{Key: "is_active", Value: false},
		}}}

	result, _ := ds.loanCollection.UpdateOne(ctx, filter, update)

	if result.MatchedCount != 1 {
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

	var max int64 = -1
	for cursor.Next(ctx) {
		var loan models.Loan
		err := cursor.Decode(&loan)
		if err != nil {
			return nil, err
		}
		if strconv.Itoa(*loan.PoolId) != *poolId {
			continue
		}

		if *loan.Amount > max {
			loans = []*models.Loan{&loan}
			max = *loan.Amount
		}
		if *loan.Amount == max {
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

func (ds DatastoreLoanMG) CountLoans(ctx context.Context, poolId *string) (*enum.CountLoans, error) {
	cursor, err := ds.loanCollection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	count := enum.CountLoans{
		TotalLoanInPool: 0,
		TotalLoanGot:    0,
	}

	for cursor.Next(ctx) {
		var loan models.Loan
		err := cursor.Decode(&loan)
		if err != nil {
			return nil, err
		}
		if strconv.Itoa(*loan.PoolId) != *poolId {
			continue
		}
		if *loan.State {
			count.TotalLoanGot++
		}
		count.TotalLoanInPool++
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	return &count, nil
}
