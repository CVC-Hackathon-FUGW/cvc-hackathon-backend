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

type DatastorePoolMG struct {
	poolCollection *mongo.Collection
}

func NewDatastorePoolMG(poolCollection *mongo.Collection) *DatastorePoolMG {
	return &DatastorePoolMG{poolCollection}
}

var _ models.DatastorePool = (*DatastorePoolMG)(nil)

func (ds DatastorePoolMG) Create(ctx context.Context, params *models.Pool) (*models.Pool, error) {
	count, err := ds.poolCollection.CountDocuments(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	params.PoolId = int(count) + 1
	params.IsActive = true

	_, err = ds.poolCollection.InsertOne(ctx, params)
	if err != nil {
		return nil, err
	}

	return params, nil
}

func (ds DatastorePoolMG) FindByID(ctx context.Context, id *string) (*models.Pool, error) {
	var pool *models.Pool

	idInt, err := strconv.Atoi(*id)
	if err != nil {
		return nil, err
	}
	query := bson.D{bson.E{Key: "pool_id", Value: idInt}}
	err = ds.poolCollection.FindOne(ctx, query).Decode(&pool)
	return pool, err
}

func (ds DatastorePoolMG) List(ctx context.Context, params enum.PoolParams) ([]*models.Pool, error) {
	var pools []*models.Pool
	filter := bson.D{{}}
	if params.Name != "" {
		filter = bson.D{{Key: "collection_name", Value: primitive.Regex{Pattern: params.Name, Options: ""}}}
	}

	cursor, err := ds.poolCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var pool models.Pool
		err := cursor.Decode(&pool)
		if err != nil {
			return nil, err
		}
		if pool.IsActive == true {
			pools = append(pools, &pool)
		}
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	return pools, nil
}

func (ds DatastorePoolMG) Update(ctx context.Context, params *models.Pool) (*models.Pool, error) {
	filter := bson.D{primitive.E{Key: "pool_id", Value: params.PoolId}}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "pool_id", Value: params.PoolId},
			primitive.E{Key: "total_pool_amount", Value: params.TotalPoolAmount},
			primitive.E{Key: "token_address", Value: params.TokenAddress},
			primitive.E{Key: "collection_name", Value: params.CollectionName},
			primitive.E{Key: "apy", Value: params.APY},
			primitive.E{Key: "duration", Value: params.Duration},
			primitive.E{Key: "state", Value: params.State},
			primitive.E{Key: "image", Value: params.Image},
		}}}

	var poolUpdated *models.Pool
	err := ds.poolCollection.FindOneAndUpdate(ctx, filter, update).Decode(&poolUpdated)
	if err != nil {
		return nil, errors.New("no matched document found for update")
	}

	return poolUpdated, nil
}

func (ds DatastorePoolMG) Delete(ctx context.Context, id *string) error {
	idInt, err := strconv.Atoi(*id)
	if err != nil {
		return err
	}

	filter := bson.D{primitive.E{Key: "pool_id", Value: idInt}}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "pool_id", Value: idInt},
			primitive.E{Key: "is_active", Value: false},
		}}}

	result, _ := ds.poolCollection.UpdateOne(ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}
