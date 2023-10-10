package models

type Package struct {
	PackageId          int     `bson:"package_id" json:"package_id"`
	PackageName        *string `bson:"package_name" json:"package_name"`
	PackageDescription *string `bson:"package_description" json:"package_description"`
	PackageImage       *string `bson:"package_image" json:"package_image"`
	PackagePrice       *int64  `bson:"package_price" json:"package_price"`
	ProjectId          *int    `bson:"project_id" json:"project_id"`
	ProjectName        *string `bson:"project_name" json:"project_name"`
	ProjectAddress     *string `bson:"project_address" json:"project_address"`
	IsActive           bool    `bson:"is_active" json:"is_active"`
}
