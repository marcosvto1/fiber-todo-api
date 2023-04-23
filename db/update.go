package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateOne(collection string, id string, data any, result any) error {
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

	opts := options.FindOneAndUpdate().SetUpsert(false)

	err = c.FindOneAndUpdate(context.Background(), filter, bson.M{"$set": data}, opts).Err()
	if err != nil {
		return err
	}

	return c.FindOne(context.Background(), filter).Decode(result)
}
