package database

import (
	"auth/api/types"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	COLLECTION_DISTRICTS = "districts"
)

func GetDistricts(database *mongo.Database) ([]types.District, error) {
	col := database.Collection(COLLECTION_DISTRICTS)

	cursor, err := col.Find(context.Background(), bson.M{})

	if err != nil {
		return nil, err
	}

	var results []types.District

	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func GetDistrictsByCity(database *mongo.Database, city string) ([]types.District, error) {
	col := database.Collection(COLLECTION_DISTRICTS)

	cursor, err := col.Find(context.Background(), bson.M{"city": city})

	if err != nil {
		return nil, err
	}

	var results []types.District

	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}
