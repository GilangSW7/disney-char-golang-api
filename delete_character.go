package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func deleteCharacter(collection *mongo.Collection, ctx context.Context, data map[string]interface{}) (map[string]interface{}, error) {

	_, err := collection.DeleteOne(ctx, bson.M{"id": data["id"]})

	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"data": "Character deleted successfully.",
	}

	return res, nil
}
