package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func getCharacter(collection *mongo.Collection, ctx context.Context) (map[string]interface{}, error) {

	cur, err := collection.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	var characters []bson.M

	for cur.Next(ctx) {

		var character bson.M

		if err = cur.Decode(&character); err != nil {
			return nil, err
		}

		characters = append(characters, character)

	}

	res := map[string]interface{}{}

	res = map[string]interface{}{
		"data": characters,
	}

	return res, nil
}
