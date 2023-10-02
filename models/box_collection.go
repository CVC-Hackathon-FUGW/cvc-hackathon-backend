package models

type BoxCollection struct {
	BoxCollectionId      int     `bson:"box_collection_id" json:"box_collection_id"`
	BoxCollectionAddress *string `bson:"box_collection_address" json:"box_collection_address"`
	Image                *string `json:"image" bson:"image"`
	IsActive             bool    `json:"is_active" bson:"is_active"`
}
