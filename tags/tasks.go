package tags

import (
	"sort"

	"github.com/marcosvto1/fiber-todo-api/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func findOrCreate(name string) (doc TagEntity, err error) {
	filter := bson.M{"name": name}

	err = db.FindOne("tags", filter, &doc)
	if err != nil && err != mongo.ErrNoDocuments {
		return
	}

	if doc.Name != "" {
		return
	}

	doc.Name = name

	id, err := db.Insert("tags", &doc)
	if err != nil {
		return
	}

	doc.ID = id

	return
}

func AddTask(taskId string, names []string) error {
	for _, name := range names {
		tag, err := findOrCreate(name)
		if err != nil {
			return err
		}
		i := sort.SearchStrings(tag.Tasks, taskId)

		if i < len(tag.Tasks) && tag.Tasks[i] == taskId {
			continue
		}

		tag.Tasks = append(tag.Tasks, taskId)

		sort.Strings(tag.Tasks)

		result := new(TagEntity)

		err = db.UpdateOne("tags", tag.ID.Hex(), tag, &result)
		if err != nil {
			return err
		}
	}

	return nil
}

func RemoveTasks(taskId string, names ...string) error {
	filter := bson.M{"tasks": taskId}
	if len(names) > 0 {
		filter["name"] = bson.M{"$in": names}
	}

	var documentes []TagEntity

	err := db.Find("tags", filter, &documentes)
	if err != nil {
		return err
	}

	var result TagEntity

	for _, document := range documentes {
		i := sort.SearchStrings(document.Tasks, taskId)

		document.Tasks = append(document.Tasks[:i], document.Tasks[i+1:]...)

		err := db.UpdateOne("tags", document.ID.Hex(), document, &result)
		if err != nil {
			return err
		}
	}

	return nil
}
