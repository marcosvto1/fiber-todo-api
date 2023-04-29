package tasks

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/marcosvto1/fiber-todo-api/db"
	"github.com/marcosvto1/fiber-todo-api/tags"
)

func updateTask(c *fiber.Ctx) error {
	taskBodyData := new(TaskEntity)

	if err := c.BodyParser(taskBodyData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"message": err.Error(),
			"code":    fiber.StatusBadRequest,
		})
	}

	var prev TaskEntity
	err := db.FindById("tasks", c.Params("id"), &prev)
	if err != nil {
		return c.Status(http.StatusBadGateway).JSON(map[string]interface{}{
			"message": err.Error(),
			"code":    fiber.StatusBadGateway,
		})
	}

	var task = TaskEntity{}
	err = db.UpdateOne("tasks", c.Params("id"), taskBodyData, &task)
	if err != nil {
		return c.Status(http.StatusBadGateway).JSON(map[string]interface{}{
			"message": err.Error(),
			"code":    fiber.StatusBadGateway,
		})
	}

	err = updateTagsTasks(c.Params("id"), prev.Tags, task.Tags)
	if err != nil {
		err = tags.RemoveTasks(c.Params("id"))
		if err != nil {
			return c.Status(http.StatusBadGateway).JSON(map[string]interface{}{
				"message": err.Error(),
				"code":    fiber.StatusBadGateway,
			})
		}

		err := db.Delete("tasks", c.Params("id"))
		if err != nil {
			return c.Status(http.StatusBadGateway).JSON(map[string]interface{}{
				"message": err.Error(),
				"code":    fiber.StatusBadGateway,
			})
		}

		_, err = db.Insert("tasks", &task)
		if err != nil {
			return c.Status(http.StatusBadGateway).JSON(map[string]interface{}{
				"message": err.Error(),
				"code":    fiber.StatusBadGateway,
			})
		}

		err = tags.AddTask(task.ID.Hex(), task.Tags)
		if err != nil {
			db.Delete("tasks", c.Params("id"))
			return c.Status(http.StatusBadGateway).JSON(map[string]interface{}{
				"message": err.Error(),
				"code":    fiber.StatusBadGateway,
			})
		}
	}

	return c.JSON(task)
}

func updateTagsTasks(id string, ot, nt []string) error {
	mot := make(map[string]int, len(ot))

	for k, v := range ot {
		mot[v] = k
	}

	var diff []string

	for _, v := range nt {
		if _, key := mot[v]; !key {
			diff = append(diff, v)
		} else {
			delete(mot, v)
		}
	}

	if len(diff) > 0 {
		err := tags.AddTask(id, diff)
		if err != nil {
			return err
		}
	}

	if len(mot) > 0 {
		dt := make([]string, 0, len(mot))
		for k := range mot {
			dt = append(dt, k)
		}

		err := tags.RemoveTasks(id, dt...)
		if err != nil {
			return err
		}
	}

	return nil
}
