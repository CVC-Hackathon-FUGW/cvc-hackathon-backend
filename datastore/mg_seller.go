package datastore

import (
	"context"

	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DatastoreSellerMG struct {
	SellerCollection *mongo.Collection
}

func NewDatastoreSellerMG(SellerCollection *mongo.Collection) *DatastoreSellerMG {
	return &DatastoreSellerMG{SellerCollection}
}

var _ models.DatastoreSeller = (*DatastoreSellerMG)(nil)

func (ds DatastoreSellerMG) Create(ctx context.Context, params *models.Seller) (*models.Seller, error) {
	count, err := ds.SellerCollection.CountDocuments(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	params.SellerID = int(count) + 1
	params.IsActive = true

	_, err = ds.SellerCollection.InsertOne(ctx, params)
	if err != nil {
		return nil, err
	}

	return params, nil
}

func (ds DatastoreSellerMG) FindByAddress(ctx context.Context, address *string) (*models.Seller, error) {
	var seller *models.Seller

	query := bson.D{bson.E{Key: "address", Value: address}}
	err := ds.SellerCollection.FindOne(ctx, query).Decode(&seller)
	return seller, err
}

func (ds DatastoreSellerMG) List(ctx context.Context) ([]*models.Seller, error) {
	var sellers []*models.Seller
	filter := bson.D{{}}

	cursor, err := ds.SellerCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var seller models.Seller
		err := cursor.Decode(&seller)
		if err != nil {
			return nil, err
		}
		if seller.IsActive == true {
			sellers = append(sellers, &seller)
		}
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	return sellers, nil
}
