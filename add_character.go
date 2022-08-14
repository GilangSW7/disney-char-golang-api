package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

func addCharacter(collection *mongo.Collection, ctx context.Context, data map[string]interface{}) (map[string]interface{}, error) {

	_, err := collection.InsertOne(ctx, data)

	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"data": map[string]interface{}{
			"insertedId": data,
		},
	}

	return res, nil
}
