package datastore

import (
	"context"
	"errors"
	"github.com/CVC-Hackathon-FUGW/cvc-hackathon-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
)

type DatastoreProjectMG struct {
	projectCollection *mongo.Collection
}

func (ds DatastoreProjectMG) Create(ctx context.Context, params *models.Project) (*models.Project, error) {
	count, err := ds.projectCollection.CountDocuments(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	params.IsActive = true
	params.ProjectId = int(count) + 1

	_, err = ds.projectCollection.InsertOne(ctx, params)
	if err != nil {
		return nil, err
	}

	return params, nil
}

func (ds DatastoreProjectMG) FindByID(ctx context.Context, id *string) (*models.Project, error) {
	var project *models.Project
	idInt, err := strconv.Atoi(*id)
	if err != nil {
		return nil, err
	}

	query := bson.D{bson.E{Key: "project_id", Value: idInt}}
	err = ds.projectCollection.FindOne(ctx, query).Decode(&project)
	return project, err
}

func (ds DatastoreProjectMG) List(ctx context.Context) ([]*models.Project, error) {
	var projects []*models.Project
	cursor, err := ds.projectCollection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var project models.Project
		err := cursor.Decode(&project)
		if err != nil {
			return nil, err
		}
		if project.IsActive == true {
			projects = append(projects, &project)
		}
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	return projects, nil
}

func (ds DatastoreProjectMG) Update(ctx context.Context, params *models.Project) (*models.Project, error) {
	var project *models.Project
	query := bson.D{bson.E{Key: "project_id", Value: params.ProjectId}}
	err := ds.projectCollection.FindOne(ctx, query).Decode(&project)
	if err != nil {
		return nil, err
	}

	if params.DueTime != nil {
		project.DueTime = params.DueTime
	}
	if params.ProjectName != nil {
		project.ProjectName = params.ProjectName
	}
	if params.ProjectDescription != nil {
		project.ProjectDescription = params.ProjectDescription
	}
	if params.ProjectImage != nil {
		project.ProjectImage = params.ProjectImage
	}
	if params.ProjectOwner != nil {
		project.ProjectOwner = params.ProjectOwner
	}
	if params.ProjectAddress != nil {
		project.ProjectAddress = params.ProjectAddress
	}
	if params.TotalRaiseAmount != nil {
		project.TotalRaiseAmount = params.TotalRaiseAmount
	}
	if params.TotalFundRaised != nil {
		project.TotalFundRaised = params.TotalFundRaised
	}

	filter := bson.D{primitive.E{Key: "project_id", Value: params.ProjectId}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "due_time", Value: project.DueTime},
		primitive.E{Key: "project_name", Value: project.ProjectName},
		primitive.E{Key: "project_description", Value: project.ProjectDescription},
		primitive.E{Key: "project_image", Value: project.ProjectImage},
		primitive.E{Key: "project_owner", Value: project.ProjectOwner},
		primitive.E{Key: "project_address", Value: project.ProjectAddress},
		primitive.E{Key: "total_raise_amount", Value: project.TotalRaiseAmount},
		primitive.E{Key: "total_fund_raised", Value: project.TotalFundRaised},
	}}}

	var updatedProject *models.Project
	err = ds.projectCollection.FindOneAndUpdate(ctx, filter, update).Decode(&updatedProject)
	if err != nil {
		return nil, errors.New("no matched document found for update")
	}

	return params, nil
}

func (ds DatastoreProjectMG) Delete(ctx context.Context, id *string) error {
	idInt, err := strconv.Atoi(*id)
	if err != nil {
		return err
	}

	filter := bson.D{primitive.E{Key: "project_id", Value: idInt}}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "project_id", Value: idInt},
			primitive.E{Key: "is_active", Value: false},
		}}}

	result, _ := ds.projectCollection.UpdateOne(ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}

func NewDatastoreProjectMG(projectCollection *mongo.Collection) *DatastoreProjectMG {
	return &DatastoreProjectMG{projectCollection}
}

var _ models.DatastoreProject = (*DatastoreProjectMG)(nil)
