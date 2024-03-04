package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/priyanshupatra02/task-ninja-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// create a collection
var collection *mongo.Collection

func init() {
	loadTheEnv()
	createDbInstance()
}

// loading the environment variables
func loadTheEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// creating the db instance
func createDbInstance() *mongo.Collection {
	connectionString := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	colName := os.Getenv("DB_COLLECTION_NAME")

	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection success ✨✨")

	//creating a reference which can be used everywhere inside the app
	collection = client.Database(dbName).Collection(colName)

	//collection reference
	fmt.Println("Collection instance is ready 💚")

	return collection
}

// Get all tasks
func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	allTasks := getAllTasks()

	fmt.Println("Got all tasks 💙")

	//sending response🚀
	json.NewEncoder(w).Encode(allTasks)

}

// controllers 👇🏻
func CreateATask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Header", "Content-Type")

	var task models.ToDoList
	_ = json.NewDecoder(r.Body).Decode(&task)

	insertOneTask(&task)
	fmt.Println("Task Created 💙")

	//sending response🚀
	json.NewEncoder(w).Encode(task)

}
func TaskComplete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")

	w.Header().Set("Access-Control-Allow-Header", "Content-Type")

	//getting task id
	params := mux.Vars(r)

	taskComplete(params["id"])
	fmt.Println("Task Completed 💙")

	//sending response🚀
	json.NewEncoder(w).Encode("Task " + params["id"] + " is done")
}
func UndoTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Header", "Content-Type")

	//getting task id
	params := mux.Vars(r)

	undoTask(params["id"])
	fmt.Println("Task is undone 💙")

	//sending response🚀
	json.NewEncoder(w).Encode("Task " + params["id"] + " is undone")

}
func DeleteATask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Header", "Content-Type")

	//getting id
	params := mux.Vars(r)

	deleteOneTask(params["id"])
	fmt.Println("Task is deleted 💙")

	//sending response🚀
	json.NewEncoder(w).Encode("Task " + params["id"] + " is deleted")

}
func DeleteAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	count := deleteAllTasks()
	fmt.Println("Deleted All tasks 💙")

	json.NewEncoder(w).Encode(count)
}

//-------- controllers end here --------
//middlewares👇🏻

func getAllTasks() []primitive.M {
	//find all tasks, empty set means " all"
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	checkNilError(err)

	defer cursor.Close(context.Background())

	var listOfTasks []primitive.M

	for cursor.Next(context.Background()) {
		var task primitive.M
		err := cursor.Decode(&task)
		checkNilError(err)

		//else if everything is okay, we're gonna push each "task" to the "listOfTasks" array above
		listOfTasks = append(listOfTasks, task)

	}
	return listOfTasks
}
func insertOneTask(task *models.ToDoList) {
	//filtering using the _id
	insertResult, err := collection.InsertOne(context.Background(), task)
	checkNilError(err)

	task.ID = insertResult.InsertedID.(primitive.ObjectID)

	fmt.Println("Inserted one task 💙 with id: ", insertResult.InsertedID)

}
func taskComplete(taskId string) {
	//converting the string "taskId" to objectId understandable by mongodb
	id, err := primitive.ObjectIDFromHex(taskId)
	checkNilError(err)

	//filtering using the _id
	filter := bson.M{"_id": id}

	//The $set operator replaces the value of a field with the specified value.
	update := bson.M{"$set": bson.M{"status": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	checkNilError(err)

	fmt.Println("Marked as completed 💙 with count", result.ModifiedCount)
}

func undoTask(taskId string) {
	// converting the string "taskId" to objectId understandable by mongodb
	id, err := primitive.ObjectIDFromHex(taskId)
	checkNilError(err)

	//filtering by id
	filter := bson.M{"_id": id}

	//The $set operator replaces the value of a field with the specified value.
	update := bson.M{"$set": bson.M{"status": false}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	checkNilError(err)

	fmt.Println("Undoed a task 💙 with count", result.ModifiedCount)
}
func deleteOneTask(taskId string) {
	//converting the string "taskId" to objectId understandable by mongodb
	id, err := primitive.ObjectIDFromHex(taskId)
	checkNilError(err)

	// filtering by _id
	filter := bson.M{"_id": id}

	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	checkNilError(err)

	fmt.Println("Deleted a task 💙 with count", deleteCount)

}
func deleteAllTasks() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	checkNilError(err)

	fmt.Println("Number of tasks deleted is: ", deleteResult.DeletedCount)

	return deleteResult.DeletedCount
}

//-------- middlewares end here --------

// common functions👇🏻
func checkNilError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
