package db

import "go.mongodb.org/mongo-driver/bson/primitive"

func Insert(collection string, data any) (primitive.ObjectID, error) {
	client, ctx := getConnections()
	defer client.Disconnect(ctx)

	c := client.Database(dbName).Collection(collection)

	res, err := c.InsertOne(ctx, data)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return res.InsertedID.(primitive.ObjectID), nil
}
