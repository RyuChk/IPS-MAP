package floordetailcollectionrepo

import (
	"context"

	wiremongo "git.cie-ips.com/ips/wire-provider/mongodb"
	"github.com/RyuChk/ips-map-service/internal/models"
	"github.com/RyuChk/ips-map-service/internal/repository/mongodb"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//go:generate mockgen -source=floor_detail_collection_repo.go -destination=mock_floorDetailCollectionRepo/mock_floordetailcollectionrepo.go -package=mock_floordetailcollectionrepo
const floorDetailCollectionRepo = "floor-detail-collection"

type Repository interface {
	InsertOne(ctx context.Context, document models.FloorDetail) error
	FindOne(ctx context.Context, filter mongodb.Filter) (models.FloorDetail, error)
	Upsert(ctx context.Context, document models.FloorDetail) error
}

type FloorDetailCollectionRepo struct {
	floorDetailCollection *mongo.Collection
}

func ProvideMapURLCollectionRepo(conn wiremongo.Connection) Repository {
	return &FloorDetailCollectionRepo{
		floorDetailCollection: conn.Database().Collection(floorDetailCollectionRepo),
	}
}

func (r *FloorDetailCollectionRepo) FindOne(ctx context.Context, filter mongodb.Filter) (models.FloorDetail, error) {
	var result models.FloorDetail
	err := r.floorDetailCollection.FindOne(ctx, filter, options.FindOne()).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *FloorDetailCollectionRepo) InsertOne(ctx context.Context, document models.FloorDetail) error {
	_, err := r.floorDetailCollection.InsertOne(ctx, document)
	if err != nil {
		return err
	}
	log.Debug().Msg("append stat to server")
	return nil
}

func (r *FloorDetailCollectionRepo) Upsert(ctx context.Context, document models.FloorDetail) error {
	filter := bson.D{{Key: "key", Value: document.Key}}
	update := bson.D{{Key: "$set", Value: document}}
	opts := options.Update().SetUpsert(true)

	_, err := r.floorDetailCollection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}
