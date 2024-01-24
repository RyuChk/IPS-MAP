package mapcollectionrepo

import (
	"context"

	wiremongo "git.cie.com/ips/wire-provider/mongodb"
	"github.com/RyuChk/ips-map-service/apps/map/models"
	"github.com/RyuChk/ips-map-service/internal/repository/mongodb"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//go:generate mockgen -source=map_collection_repo.go -destination=mock_mapCollectionRepo/mock_mapcollectionrepo.go -package=mock_mapcollectionrepo
const mapURLCollectionRepo = "map-collection"

type Repository interface {
	InsertOne(ctx context.Context, document models.Map) error
	Find(ctx context.Context, filter mongodb.Filter) ([]models.Map, error)
	FindOne(ctx context.Context, filter mongodb.Filter) (models.Map, error)
}

type DataCollectionRepo struct {
	mapURLCollection *mongo.Collection
}

func ProvideMapURLCollectionRepo(conn wiremongo.Connection) Repository {
	return &DataCollectionRepo{
		mapURLCollection: conn.Database().Collection(mapURLCollectionRepo),
	}
}

func (r *DataCollectionRepo) Find(ctx context.Context, filter mongodb.Filter) ([]models.Map, error) {
	var result []models.Map
	cursor, err := r.mapURLCollection.Find(context.TODO(), filter)
	if err != nil {
		return []models.Map{}, err
	}

	err = cursor.All(ctx, &result)
	if err != nil {
		return []models.Map{}, err
	}
	return result, nil
}

func (r *DataCollectionRepo) FindOne(ctx context.Context, filter mongodb.Filter) (models.Map, error) {
	var result models.Map
	err := r.mapURLCollection.FindOne(ctx, filter, options.FindOne()).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *DataCollectionRepo) InsertOne(ctx context.Context, document models.Map) error {
	_, err := r.mapURLCollection.InsertOne(ctx, document)
	if err != nil {
		log.Error().Err(err).Msg("error adding new map information")
		return err
	}
	log.Debug().Any("models.Map", document).Msg("addming map information into database")
	return nil
}
