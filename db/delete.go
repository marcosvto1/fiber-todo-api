package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Delete(collection string, id string) error {
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

	result, err := c.DeleteOne(context.Background(), filter)

	if err != nil {
		return err
	}

	if result.DeletedCount != 1 {
		return fmt.Errorf("%d documents deleted", result.DeletedCount)
	}

	return nil
}
