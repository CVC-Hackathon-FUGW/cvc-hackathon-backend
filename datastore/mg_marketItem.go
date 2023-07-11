package datastore

import (
	"context"
	"errors"

	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DatastoreMarketItemMG struct {
	MarketItemCollection *mongo.Collection
}

func NewDatastoreMarketItemMG(MarketItemCollection *mongo.Collection) *DatastoreMarketItemMG {
	return &DatastoreMarketItemMG{MarketItemCollection}
}

var _ models.DatastoreMarketItem = (*DatastoreMarketItemMG)(nil)

func (ds DatastoreMarketItemMG) Create(ctx context.Context, params *models.MarketItem) (*models.MarketItem, error) {
	_, err := ds.MarketItemCollection.InsertOne(ctx, params)
	if err != nil {
		return nil, err
	}

	return params, nil
}

func (ds DatastoreMarketItemMG) FindByID(ctx context.Context, id *string) (*models.MarketItem, error) {
	var marketItem *models.MarketItem
	query := bson.D{bson.E{Key: "item_id", Value: id}}
	err := ds.MarketItemCollection.FindOne(ctx, query).Decode(&marketItem)
	return marketItem, err
}

func (ds DatastoreMarketItemMG) List(ctx context.Context) ([]*models.MarketItem, error) {
	var marketItems []*models.MarketItem
	cursor, err := ds.MarketItemCollection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var marketItem models.MarketItem
		err := cursor.Decode(&marketItem)
		if err != nil {
			return nil, err
		}
		marketItems = append(marketItems, &marketItem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	if len(marketItems) == 0 {
		return nil, errors.New("documents not found")
	}
	return marketItems, nil
}

func (ds DatastoreMarketItemMG) Update(ctx context.Context, params *models.MarketItem) (*models.MarketItem, error) {
	filter := bson.D{primitive.E{Key: "item_id", Value: params.ItemId}}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "item_id", Value: params.ItemId},
			primitive.E{Key: "token_address", Value: params.TokenAddress},
			primitive.E{Key: "price", Value: params.Price},
			primitive.E{Key: "is_offerable", Value: params.IsOfferable},
			primitive.E{Key: "accept_visa_payment", Value: params.AcceptVisaPayment},
			primitive.E{Key: "current_offer_value", Value: params.CurrentOfferValue},
			primitive.E{Key: "current_offerer", Value: params.CurrentOfferer},
			primitive.E{Key: "sold", Value: params.Sold},
		}}}

	var marketItemUpdated *models.MarketItem
	err := ds.MarketItemCollection.FindOneAndUpdate(ctx, filter, update).Decode(&marketItemUpdated)
	if err != nil {
		return nil, errors.New("no matched document found for update")
	}

	return marketItemUpdated, nil
}

func (ds DatastoreMarketItemMG) Delete(ctx context.Context, id *string) error {
	filter := bson.D{primitive.E{Key: "item_id", Value: id}}
	result, _ := ds.MarketItemCollection.DeleteOne(ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}
