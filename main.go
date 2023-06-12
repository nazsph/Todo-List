package main

import (
	"context"

	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Task struct {
	ID   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Task string             `json:"task,omitempty" bson:"task,omitempty"`
}

var collection *mongo.Collection

const uri = "mongodb+srv://admin:abcd1234@cluster0.ucqal4p.mongodb.net/?retryWrites=true&w=majority"

func main() {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	_ = client.Connect(context.Background())

	_ = client.Ping(context.Background(), nil)

	db := client.Database("test")
	collection = db.Collection("tasks")

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, DELETE, PATCH, OPTIONS",
	}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})

	app.Get("/getAllTasks", getAllTasks)
	app.Post("/addTask", addTask)
	app.Delete("/deleteTask/:id", deleteTask)

	log.Fatal(app.Listen(":5050"))
}
func getAllTasks(c *fiber.Ctx) error {
	var tasks []Task

	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var task Task
		err := cur.Decode(&task)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
		tasks = append(tasks, task)
	}

	return c.JSON(tasks)
}

func addTask(c *fiber.Ctx) error {
	var task Task
	err := c.BodyParser(&task)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	_, err = collection.InsertOne(context.Background(), task)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.SendStatus(fiber.StatusOK)
}

func deleteTask(c *fiber.Ctx) error {
	id := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.SendStatus(fiber.StatusOK)
}
