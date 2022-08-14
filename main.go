package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "log"
	"net/http"
)

const (
	dbUri  = "mongodb+srv://gilang:afxvvBf95fSn2Ycj@cluster0.fzfsaqr.mongodb.net/?retryWrites=true&w=majority"
	dbName = "disney_character"
)

func main() {
	http.HandleFunc("/character", requestHandler)
	http.ListenAndServe(":8080", nil)
}

func requestHandler(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	response := map[string]interface{}{}

	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUri))

	if err != nil {
		fmt.Println(err.Error())
	}

	collection := client.Database(dbName).Collection("characters")

	data := map[string]interface{}{}

	err = json.NewDecoder(req.Body).Decode(&data)

	if err != nil {
		fmt.Println(err.Error())
	}

	switch req.Method {
	case "POST":
		response, err = addCharacter(collection, ctx, data)
	case "GET":
		response, err = getCharacter(collection, ctx)
	case "PUT":
		response, err = updateCharacter(collection, ctx, data)
	case "DELETE":
		response, err = deleteCharacter(collection, ctx, data)
	}

	if err != nil {
		response = map[string]interface{}{"error": err.Error()}
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")

	if err := enc.Encode(response); err != nil {
		fmt.Println(err.Error())
	}
}
