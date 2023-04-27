package tasks

import (
	"github.com/marcosvto1/fiber-todo-api/tags"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskEntity struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Tags        []tags.TagEntity   `json:"tags" bson:"tags"`
	Assign      primitive.ObjectID `json:"assign" bson:"assign"`
	Done        bool               `json:"done" bson:"done"`
}
