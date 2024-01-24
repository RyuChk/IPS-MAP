package mapurlcollectionrepo

import (
	"context"

	wiremongo "git.cie.com/ips/wire-provider/mongodb"
	"github.com/RyuChk/ips-map-service/apps/map/models"
	"github.com/RyuChk/ips-map-service/internal/repository/mongodb"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//go:generate mockgen -source=map_url_collection_repo.go -destination=mock_mapURLCollectionRepo/mock_mapurlcollectionrepo.go -package=mock_mapurlcollectionrepo

const mapURLCollectionRepo = "map-url-collection"

type Repository interface {
	InsertOne(ctx context.Context, document models.MapImageURL) error
	FindOne(ctx context.Context, filter mongodb.Filter) (models.MapImageURL, error)
}

type DataCollectionRepo struct {
	mapURLCollection *mongo.Collection
}

func ProvideMapURLCollectionRepo(conn wiremongo.Connection) Repository {
	return &DataCollectionRepo{
		mapURLCollection: conn.Database().Collection(mapURLCollectionRepo),
	}
}

func (r *DataCollectionRepo) FindOne(ctx context.Context, filter mongodb.Filter) (models.MapImageURL, error) {
	var result models.MapImageURL
	err := r.mapURLCollection.FindOne(ctx, filter, options.FindOne()).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *DataCollectionRepo) InsertOne(ctx context.Context, document models.MapImageURL) error {

	log.Debug().Any("RSSIStatModel", document).Msg("Inserting into db")

	_, err := r.mapURLCollection.InsertOne(ctx, document)
	if err != nil {
		return err
	}
	log.Debug().Msg("append stat to server")
	return nil
}
