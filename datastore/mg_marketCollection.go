package datastore

import (
	"context"
	"errors"
	"strconv"

	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/enum"
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DatastoreMarketCollectionMG struct {
	marketCollectionCollection *mongo.Collection
}

func NewDatastoreMarketCollectionMG(MarketCollectionCollection *mongo.Collection) *DatastoreMarketCollectionMG {
	return &DatastoreMarketCollectionMG{MarketCollectionCollection}
}

var _ models.DatastoreMarketCollection = (*DatastoreMarketCollectionMG)(nil)

func (ds DatastoreMarketCollectionMG) Create(ctx context.Context, params *models.MarketCollection) (*models.MarketCollection, error) {
	count, err := ds.marketCollectionCollection.CountDocuments(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	params.CollectionId = int(count) + 1
	params.IsActive = true
	_, err = ds.marketCollectionCollection.InsertOne(ctx, params)
	if err != nil {
		return nil, err
	}

	return params, nil
}

func (ds DatastoreMarketCollectionMG) FindByID(ctx context.Context, id *string) (*models.MarketCollection, error) {
	var MarketCollection *models.MarketCollection
	idInt, err := strconv.Atoi(*id)
	if err != nil {
		return nil, err
	}
	query := bson.D{bson.E{Key: "collection_id", Value: idInt}}
	err = ds.marketCollectionCollection.FindOne(ctx, query).Decode(&MarketCollection)
	return MarketCollection, err
}

func (ds DatastoreMarketCollectionMG) List(ctx context.Context, params enum.MarketCollectionsParams) ([]*models.MarketCollection, error) {
	var marketCollections []*models.MarketCollection
	filter := bson.D{{}}
	if params.Name != "" {
		filter = bson.D{{Key: "collection_name", Value: primitive.Regex{Pattern: params.Name, Options: ""}}}
	}

	cursor, err := ds.marketCollectionCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var MarketCollection models.MarketCollection
		err := cursor.Decode(&MarketCollection)
		if err != nil {
			return nil, err
		}
		if MarketCollection.IsActive == true {
			marketCollections = append(marketCollections, &MarketCollection)
		}
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	return marketCollections, nil
}

func (ds DatastoreMarketCollectionMG) Update(ctx context.Context, params *models.MarketCollection) (*models.MarketCollection, error) {
	var marketCollectionDB *models.MarketCollection

	query := bson.D{bson.E{Key: "collection_id", Value: params.CollectionId}}
	err := ds.marketCollectionCollection.FindOne(ctx, query).Decode(&marketCollectionDB)
	if err != nil {
		return nil, err
	}

	if params.CollectionName != nil {
		marketCollectionDB.CollectionName = params.CollectionName
	}

	if params.TokenAddress != nil {
		marketCollectionDB.TokenAddress = params.TokenAddress
	}

	if params.Image != nil {
		marketCollectionDB.Image = params.Image
	}

	filter := bson.D{primitive.E{Key: "collection_id", Value: params.CollectionId}}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "token_address", Value: marketCollectionDB.TokenAddress},
			primitive.E{Key: "collection_name", Value: marketCollectionDB.CollectionName},
			primitive.E{Key: "image", Value: marketCollectionDB.Image},
		}}}

	var MarketCollectionUpdated *models.MarketCollection
	err = ds.marketCollectionCollection.FindOneAndUpdate(ctx, filter, update).Decode(&MarketCollectionUpdated)
	if err != nil {
		return nil, errors.New("no matched document found for update")
	}

	return MarketCollectionUpdated, nil
}

func (ds DatastoreMarketCollectionMG) Delete(ctx context.Context, id *string) error {
	idInt, err := strconv.Atoi(*id)
	if err != nil {
		return err
	}

	filter := bson.D{primitive.E{Key: "collection_id", Value: idInt}}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "collection_id", Value: idInt},
			primitive.E{Key: "is_active", Value: false},
		}}}

	result, _ := ds.marketCollectionCollection.UpdateOne(ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no matched document found for delete")
	}

	return nil
}

func (ds DatastoreMarketCollectionMG) FindByAddress(ctx context.Context, tokenAddress *string) ([]*models.MarketCollection, error) {
	filter := bson.D{{Key: "token_address", Value: tokenAddress}}

	var marketCollections []*models.MarketCollection
	cursor, err := ds.marketCollectionCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var marketcollection models.MarketCollection
		err := cursor.Decode(&marketcollection)
		if err != nil {
			return nil, err
		}
		if marketcollection.IsActive == true {
			marketCollections = append(marketCollections, &marketcollection)
		}
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	return marketCollections, nil
}
