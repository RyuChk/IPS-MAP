package buildingcollectionrepo

import (
	"context"

	wiremongo "git.cie-ips.com/ips/wire-provider/mongodb"
	"github.com/RyuChk/ips-map-service/internal/models"
	"github.com/RyuChk/ips-map-service/internal/repository/mongodb"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//go:generate mockgen -source=building_collection_repo.go -destination=mock_buildingCollectionRepo/mock_buildingcollectionrepo.go -package=mock_buildingcollectionrepo
const buildingCollectionRepo = "building-collection"

type Repository interface {
	InsertOne(ctx context.Context, document models.Building) error
	Find(ctx context.Context, filter mongodb.Filter) ([]models.Building, error)
	FindOne(ctx context.Context, filter mongodb.Filter) (models.Building, error)
}

type BuildingCollectionRepo struct {
	buildingCollection *mongo.Collection
}

func ProvideBuildingCollectionRepo(conn wiremongo.Connection) Repository {
	return &BuildingCollectionRepo{
		buildingCollection: conn.Database().Collection(buildingCollectionRepo),
	}
}

func (r *BuildingCollectionRepo) Find(ctx context.Context, filter mongodb.Filter) ([]models.Building, error) {
	var result []models.Building
	cursor, err := r.buildingCollection.Find(context.TODO(), filter)
	if err != nil {
		return []models.Building{}, err
	}

	err = cursor.All(ctx, &result)
	if err != nil {
		return []models.Building{}, err
	}
	return result, nil
}

func (r *BuildingCollectionRepo) FindOne(ctx context.Context, filter mongodb.Filter) (models.Building, error) {
	var result models.Building
	err := r.buildingCollection.FindOne(ctx, filter, options.FindOne()).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *BuildingCollectionRepo) InsertOne(ctx context.Context, document models.Building) error {
	_, err := r.buildingCollection.InsertOne(ctx, document)
	if err != nil {
		log.Error().Err(err).Msg("error adding new map information")
		return err
	}
	log.Debug().Any("models.Map", document).Msg("addming map information into database")
	return nil
}
