package main

import (
	"context"
	"log"
	"net/http"
	"time"
	"todo-tasks/database"
	"todo-tasks/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const uri = "mongodb://localhost:27017/"

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.Static("/css", "templates/css")
	r.Static("/images", "templates/images")

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/tasks")
	})

	r.GET("/tasks", func(c *gin.Context) {
		collection := database.GetCollection("task")
		var todoTasks []models.Task
		var doneTasks []models.Task
		ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
		defer cancel()

		todoCur, bad := collection.Find(ctx, bson.M{
			"status": models.TASK_TODO,
		})
		if bad != nil {
			log.Fatal(bad)
		}
		doneCur, bad := collection.Find(ctx, bson.M{
			"status": models.TASK_DONE,
		})
		if bad != nil {
			log.Fatal(bad)
		}
		todoCur.All(ctx, &todoTasks)
		doneCur.All(ctx, &doneTasks)

		c.HTML(200, "tasks.html", gin.H{
			"todoTasks": todoTasks,
			"doneTasks": doneTasks,
		})
	})

	r.POST("/task/new", func(c *gin.Context) {
		name := c.PostForm("name")
		var newTask models.Task
		newTask.Name = name
		newTask.Status = models.TASK_TODO
		newTask.Id = primitive.NewObjectID()
		collection := database.GetCollection("task")
		collection.InsertOne(context.TODO(), newTask)
		c.Redirect(http.StatusFound, "/tasks")
	})

	r.POST("/task/done", func(c *gin.Context) {
		id := c.PostForm("id")

		objectId, _ := primitive.ObjectIDFromHex(id)
		log.Print(objectId)
		collection := database.GetCollection("task")
		collection.FindOneAndUpdate(context.TODO(), bson.M{"id": objectId}, bson.M{
			"$set": bson.M{"status": models.TASK_DONE},
		})
		c.Redirect(http.StatusFound, "/tasks")
	})

	r.POST("/task/todo", func(c *gin.Context) {
		id := c.PostForm("id")

		var task models.Task
		task.Status = models.TASK_DONE
		objectId, _ := primitive.ObjectIDFromHex(id)
		collection := database.GetCollection("task")
		collection.FindOneAndUpdate(context.TODO(), bson.M{"id": objectId}, bson.M{
			"$set": bson.M{"status": models.TASK_TODO},
		})
		c.Redirect(http.StatusFound, "/tasks")
	})

	r.POST("/task/delete", func(c *gin.Context) {
		id := c.PostForm("id")

		var task models.Task
		task.Status = models.TASK_DONE
		objectId, _ := primitive.ObjectIDFromHex(id)
		collection := database.GetCollection("task")
		collection.DeleteOne(context.TODO(), bson.M{"id": objectId})
		c.Redirect(http.StatusFound, "/tasks")
	})

	log.Print("Go to http://localhost:9111")

	http.ListenAndServe(":9111", r)
}
