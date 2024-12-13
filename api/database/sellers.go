package database

import (
	"auth/api/types"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	COLLECTION_SELLERS = "sellers"
)

func GetSellers(database *mongo.Database) ([]types.Seller, error) {
	col := database.Collection(COLLECTION_SELLERS)

	cursor, err := col.Find(context.Background(), bson.M{})

	if err != nil {
		return nil, err
	}

	var results []types.Seller

	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func GetSellersByDistrict(database *mongo.Database, city string, district string) ([]types.Seller, error) {
	col := database.Collection(COLLECTION_SELLERS)

	cursor, err := col.Find(context.Background(), bson.M{"district": district, "city": city})

	if err != nil {
		return nil, err
	}

	var results []types.Seller

	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func GetSellersByCity(database *mongo.Database, city string) ([]types.Seller, error) {
	col := database.Collection(COLLECTION_SELLERS)

	cursor, err := col.Find(context.Background(), bson.M{"city": city})

	if err != nil {
		return nil, err
	}

	var results []types.Seller

	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}
