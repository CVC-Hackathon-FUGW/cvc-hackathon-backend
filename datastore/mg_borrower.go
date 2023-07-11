package datastore

import (
	"context"
	"errors"

	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DatastoreBorrowerMG struct {
	borrowerCollection *mongo.Collection
}

func NewDatastoreBorrowerMG(borrowerCollection *mongo.Collection) *DatastoreBorrowerMG {
	return &DatastoreBorrowerMG{borrowerCollection}
}

var _ models.DatastoreBorrower = (*DatastoreBorrowerMG)(nil)

func (ds DatastoreBorrowerMG) Create(ctx context.Context, params *models.Borrower) (*models.Borrower, error) {
	count, err := ds.borrowerCollection.CountDocuments(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	params.BorrowerID = int(count) + 1
	_, err = ds.borrowerCollection.InsertOne(ctx, params)
	if err != nil {
		return nil, err
	}

	return params, nil
}

func (ds DatastoreBorrowerMG) FindByID(ctx context.Context, id *string) (*models.Borrower, error) {
	var Borrower *models.Borrower
	query := bson.D{bson.E{Key: "borrower_id", Value: id}}
	err := ds.borrowerCollection.FindOne(ctx, query).Decode(&Borrower)
	return Borrower, err
}

func (ds DatastoreBorrowerMG) List(ctx context.Context) ([]*models.Borrower, error) {
	var borrowers []*models.Borrower
	cursor, err := ds.borrowerCollection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var borrower models.Borrower
		err := cursor.Decode(&borrower)
		if err != nil {
			return nil, err
		}
		borrowers = append(borrowers, &borrower)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	if len(borrowers) == 0 {
		return nil, errors.New("documents not found")
	}
	return borrowers, nil
}

func (ds DatastoreBorrowerMG) Update(ctx context.Context, params *models.Borrower) (*models.Borrower, error) {
	filter := bson.D{primitive.E{Key: "borrower_id", Value: params.BorrowerID}}
	update := bson.D{
		primitive.E{Key: "loans", Value: params.Loans},
	}

	var borrowerUpdated *models.Borrower
	err := ds.borrowerCollection.FindOneAndUpdate(ctx, filter, update).Decode(&borrowerUpdated)
	if err != nil {
		return nil, errors.New("no matched document found for update")
	}

	return borrowerUpdated, nil
}

func (ds DatastoreBorrowerMG) Delete(ctx context.Context, id *string) error {
	filter := bson.D{primitive.E{Key: "borrower_id", Value: id}}
	result, _ := ds.borrowerCollection.DeleteOne(ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}
