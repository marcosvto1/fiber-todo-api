package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Find(collection string, filter bson.M, documents any) error {
	client, ctx := getConnections()
	defer client.Disconnect(ctx)

	c := client.Database(dbName).Collection(collection)

	if filter == nil {
		filter = bson.M{}
	}

	cursor, err := c.Find(context.Background(), filter)
	if err != nil {
		return nil
	}

	defer cursor.Close(context.Background())

	return cursor.All(context.Background(), documents)
}

func FindById(collection string, id string, document any) error {
	client, ctx := getConnections()
	defer client.Disconnect(ctx)

	c := client.Database(dbName).Collection(collection)

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{
		"_id": objectId,
	}

	return c.FindOne(context.Background(), filter).Decode(document)
}

func FindOne(collection string, filter bson.M, doc any) error {
	client, ctx := getConnections()
	defer client.Disconnect(ctx)

	c := client.Database(dbName).Collection(collection)

	return c.FindOne(context.Background(), filter).Decode(doc)
}
