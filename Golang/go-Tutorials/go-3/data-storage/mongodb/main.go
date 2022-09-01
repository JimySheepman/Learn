package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/bson"
)

func main() {
	client, err := mongo.NewClient(`mongodb://username:password@localhost:27094/db`)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	collection := client.Database("db").Collection("school")

	res, err := collection.InsertOne(context.Background(), bson.M{"hello": "world"})
	if err != nil {
		panic(err)
	}
	id := res.InsertedID
	fmt.Println(id)

	q1 := bson.M{"firstname": bson.M{"$eq": "John"}}
	q2 := bson.M{"lastname": bson.M{"$eq": "Doe"}}

	clauses := bson.A{q1, q2}

	filter := bson.M{"$and": clauses}

	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		panic(err)
	}

	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		raw, err := cur.DecodeBytes()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", raw)
	}
	if err := cur.Err(); err != nil {
		panic(err)
	}
}
