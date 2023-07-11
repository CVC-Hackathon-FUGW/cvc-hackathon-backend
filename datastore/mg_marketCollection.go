package datastore

import (
	"context"
	"errors"
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DatastoreMarketCollectionMG struct {
	MarketCollectionCollection *mongo.Collection
}

func NewDatastoreMarketCollectionMG(MarketCollectionCollection *mongo.Collection) *DatastoreMarketCollectionMG {
	return &DatastoreMarketCollectionMG{MarketCollectionCollection}
}

var _ models.DatastoreMarketCollection = (*DatastoreMarketCollectionMG)(nil)

func (ds DatastoreMarketCollectionMG) Create(ctx context.Context, params *models.MarketCollection) (*models.MarketCollection, error) {
	_, err := ds.MarketCollectionCollection.InsertOne(ctx, params)
	if err != nil {
		return nil, err
	}

	return params, nil
}

func (ds DatastoreMarketCollectionMG) FindByID(ctx context.Context, id *string) (*models.MarketCollection, error) {
	var MarketCollection *models.MarketCollection
	query := bson.D{bson.E{Key: "collection_id", Value: id}}
	err := ds.MarketCollectionCollection.FindOne(ctx, query).Decode(&MarketCollection)
	return MarketCollection, err
}

func (ds DatastoreMarketCollectionMG) List(ctx context.Context) ([]*models.MarketCollection, error) {
	var MarketCollections []*models.MarketCollection
	cursor, err := ds.MarketCollectionCollection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var MarketCollection models.MarketCollection
		err := cursor.Decode(&MarketCollection)
		if err != nil {
			return nil, err
		}
		MarketCollections = append(MarketCollections, &MarketCollection)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	if len(MarketCollections) == 0 {
		return nil, errors.New("documents not found")
	}
	return MarketCollections, nil
}

func (ds DatastoreMarketCollectionMG) Update(ctx context.Context, params *models.MarketCollection) (*models.MarketCollection, error) {
	filter := bson.D{primitive.E{Key: "collection_id", Value: params.CollectionId}}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "item_id", Value: params.CollectionId},
			primitive.E{Key: "token_address", Value: params.TokenAddress},
			primitive.E{Key: "collection_name", Value: params.CollectionName},
			primitive.E{Key: "image", Value: params.Image},
		}}}

	var MarketCollectionUpdated *models.MarketCollection
	err := ds.MarketCollectionCollection.FindOneAndUpdate(ctx, filter, update).Decode(&MarketCollectionUpdated)
	if err != nil {
		return nil, errors.New("no matched document found for update")
	}

	return MarketCollectionUpdated, nil
}

func (ds DatastoreMarketCollectionMG) Delete(ctx context.Context, id *string) error {
	filter := bson.D{primitive.E{Key: "collection_id", Value: id}}
	result, _ := ds.MarketCollectionCollection.DeleteOne(ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}
