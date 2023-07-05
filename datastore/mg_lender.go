package datastore

import (
	"context"
	"errors"

	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DatastoreLenderMG struct {
	lenderCollection *mongo.Collection
}

func NewDatastoreLenderMG(LenderCollection *mongo.Collection) *DatastoreLenderMG {
	return &DatastoreLenderMG{LenderCollection}
}

var _ models.DatastoreLender = (*DatastoreLenderMG)(nil)

func (ds DatastoreLenderMG) Create(ctx context.Context, params *models.Lender) (*models.Lender, error) {
	_, err := ds.lenderCollection.InsertOne(ctx, params)
	if err != nil {
		return nil, err
	}

	return params, nil
}

func (ds DatastoreLenderMG) FindByID(ctx context.Context, id *string) (*models.Lender, error) {
	var lender *models.Lender
	query := bson.D{bson.E{Key: "lender_id", Value: id}}
	err := ds.lenderCollection.FindOne(ctx, query).Decode(&lender)
	return lender, err
}

func (ds DatastoreLenderMG) List(ctx context.Context) ([]*models.Lender, error) {
	var Lenders []*models.Lender
	cursor, err := ds.lenderCollection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var Lender models.Lender
		err := cursor.Decode(&Lender)
		if err != nil {
			return nil, err
		}
		Lenders = append(Lenders, &Lender)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	if len(Lenders) == 0 {
		return nil, errors.New("documents not found")
	}
	return Lenders, nil
}

func (ds DatastoreLenderMG) Update(ctx context.Context, params *models.Lender) (*models.Lender, error) {
	filter := bson.D{primitive.E{Key: "lender_id", Value: params.LenderID}}
	update := bson.D{
		primitive.E{Key: "offers", Value: params.Offers},
	}

	var LenderUpdated *models.Lender
	err := ds.lenderCollection.FindOneAndUpdate(ctx, filter, update).Decode(&LenderUpdated)
	if err != nil {
		return nil, errors.New("no matched document found for update")
	}

	return LenderUpdated, nil
}

func (ds DatastoreLenderMG) Delete(ctx context.Context, id *string) error {
	filter := bson.D{primitive.E{Key: "lender_id", Value: id}}
	result, _ := ds.lenderCollection.DeleteOne(ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}
