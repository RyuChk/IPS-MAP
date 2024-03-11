package floorcollectionrepo

import (
	"context"

	wiremongo "git.cie-ips.com/ips/wire-provider/mongodb"
	"github.com/RyuChk/ips-map-service/internal/models"
	"github.com/RyuChk/ips-map-service/internal/repository/mongodb"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//go:generate mockgen -source=floor_collection_repo.go -destination=mock_floorCollectionRepo/mock_floorcollectionrepo.go -package=mock_floorcollectionrepo
const floorCollectionRepo = "floor-collection"

type Repository interface {
	InsertOne(ctx context.Context, document models.Floor) error
	Find(ctx context.Context, filter mongodb.Filter) ([]models.Floor, error)
	FindOne(ctx context.Context, filter mongodb.Filter) (models.Floor, error)
}

type FloorCollectionRepo struct {
	floorcollectionrepo *mongo.Collection
}

func ProvideFloorCollectionRepo(conn wiremongo.Connection) Repository {
	return &FloorCollectionRepo{
		floorcollectionrepo: conn.Database().Collection(floorCollectionRepo),
	}
}

func (r *FloorCollectionRepo) Find(ctx context.Context, filter mongodb.Filter) ([]models.Floor, error) {
	var result []models.Floor
	cursor, err := r.floorcollectionrepo.Find(context.TODO(), filter)
	if err != nil {
		return []models.Floor{}, err
	}

	err = cursor.All(ctx, &result)
	if err != nil {
		return []models.Floor{}, err
	}
	return result, nil
}

func (r *FloorCollectionRepo) FindOne(ctx context.Context, filter mongodb.Filter) (models.Floor, error) {
	var result models.Floor
	err := r.floorcollectionrepo.FindOne(ctx, filter, options.FindOne()).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *FloorCollectionRepo) InsertOne(ctx context.Context, document models.Floor) error {
	_, err := r.floorcollectionrepo.InsertOne(ctx, document)
	if err != nil {
		return err
	}
	log.Debug().Msg("append stat to server")
	return nil
}
