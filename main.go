// package main

// import (
// 	"context"
// 	"encoding/json"
// 	"log"
// 	"net/http"
// 	"os"
// 	"os/signal"
// 	"strings"
// 	"time"

// 	"github.com/go-chi/chi"
// 	"github.com/go-chi/chi/middleware"
// 	"github.com/thedevsaddam/renderer"
// 	mgo "gopkg.in/mgo.v2"
// 	"gopkg.in/mgo.v2/bson"
// )

// var rnd *renderer.Render
// var db *mgo.Database

// const (
// 	hostname       string = "localhost:27017"
// 	dbname         string = "todo_tut"
// 	port           string = ":9000" //port of the project not the db
// 	collectionName string = "todos"
// )

// type (
// 	todoModel struct {
// 		ID        bson.ObjectId `bson:"_id,omitempty"`
// 		Title     string        `bson:"title"`
// 		Completed bool          `bson:"completed"`
// 		CreatedAt time.Time     `bson:"createAt"`
// 	}
// 	todo struct {
// 		ID        string    `json:"id"`
// 		Title     string    `json:"title"`
// 		Completed bool      `json:"completed"`
// 		CreatedAt time.Time `json:"create_at"`
// 	}
// )

// func init() {
// 	rnd = renderer.New()
// 	sess, err := mgo.Dial(hostname)
// 	checkErr(err)
// 	sess.SetMode(mgo.Monotonic, true)
// 	db = sess.DB(dbname)
// }

// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	err := rnd.Template(w, 200, []string{"./static/home.tpl"}, nil)
// 	checkErr(err)
// }

// func fetchTodos(w http.ResponseWriter, r *http.Request) {
// 	todos := []todoModel{}

// 	if err := db.C(collectionName).Find(bson.M{}).All(&todos); err != nil {
// 		rnd.JSON(w, http.StatusProcessing, renderer.M{
// 			"message": "Failed to fetch todo",
// 			"error":   err,
// 		})
// 		return
// 	}

// 	todoList := []todo{}

// 	for _, t := range todos {
// 		todoList = append(todoList, todo{
// 			ID:        t.ID.Hex(),
// 			Title:     t.Title,
// 			Completed: t.Completed,
// 			CreatedAt: t.CreatedAt,
// 		})
// 	}

// 	rnd.JSON(w, http.StatusOK, renderer.M{
// 		"data": todoList,
// 	})

// }

// func createTodo(w http.ResponseWriter, r *http.Request) {
// 	var t todo

// 	if err := json.NewDecoder(r.Body).Decode(t); err != nil {
// 		rnd.JSON(w, http.StatusProcessing, err)
// 		return
// 	}

// 	if t.Title == "" {
// 		rnd.JSON(w, http.StatusBadRequest, renderer.M{
// 			"message": "The title is required",
// 		})
// 	}

// 	tm := todoModel{
// 		ID:        bson.NewObjectId(),
// 		Title:     t.Title,
// 		Completed: false,
// 		CreatedAt: time.Now(),
// 	}

// 	if err := db.C(collectionName).Insert(tm); err != nil {
// 		rnd.JSON(w, http.StatusProcessing, renderer.M{
// 			"message": "failed to save todo",
// 			"error":   err,
// 		})
// 		return
// 	}

// 	rnd.JSON(w, http.StatusCreated, renderer.M{
// 		"message": "todo created successfully",
// 		"todo_id": tm.ID.Hex(),
// 	})

// }

// func deleteTodo(w http.ResponseWriter, r *http.Request) {
// 	id := strings.TrimSpace(chi.URLParam(r, "id"))
// 	if !bson.IsObjectIdHex(id) {
// 		rnd.JSON(w, http.StatusBadRequest, renderer.M{
// 			"message": "The id is invalid",
// 		})
// 		return
// 	}

// 	if err := db.C(collectionName).RemoveId(bson.ObjectIdHex(id)); err != nil {
// 		rnd.JSON(w, http.StatusProcessing, renderer.M{
// 			"message": "failed to delete todo",
// 			"error":   err,
// 		})
// 		return
// 	}
// 	rnd.JSON(w, http.StatusOK, renderer.M{
// 		"message": "todo deleted successfully",
// 	})
// }

// func updateTodo(w http.ResponseWriter, r *http.Request) {
// 	id := strings.TrimSpace(chi.URLParam(r, "id"))
// 	if !bson.IsObjectIdHex(id) {
// 		rnd.JSON(w, http.StatusBadRequest, renderer.M{
// 			"message": "The id is invalid",
// 		})
// 		return
// 	}

// 	var t todo

// 	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
// 		rnd.JSON(w, http.StatusProcessing, err)
// 		return
// 	}

// 	if t.Title == "" {
// 		rnd.JSON(w, http.StatusBadRequest, renderer.M{
// 			"message": "The title is required",
// 		})
// 	}

// 	if err := db.C(collectionName).
// 		Update(
// 			bson.M{"_id": bson.ObjectIdHex(id)},
// 			bson.M{"title": t.Title, "completed": t.Completed},
// 		); err != nil {
// 		rnd.JSON(w, http.StatusProcessing, renderer.M{
// 			"message": "failed to update todo",
// 			"error":   err,
// 		})
// 		return
// 	}
// }

// func main() {
// 	stopChan := make(chan os.Signal)
// 	signal.Notify(stopChan, os.Interrupt)
// 	r := chi.NewRouter()
// 	r.Use(middleware.Logger)
// 	r.Get("/", homeHandler)
// 	r.Mount("/todo", todoHandlers())

// 	srv := &http.Server{
// 		Addr:         port,
// 		Handler:      r,
// 		ReadTimeout:  60 * time.Second,
// 		WriteTimeout: 60 * time.Second,
// 		IdleTimeout:  60 * time.Second,
// 	}

// 	go func() {
// 		log.Println("Listening on port ", port)
// 		if err := srv.ListenAndServe(); err != nil {
// 			log.Printf("listen:%s\n", err)
// 		}
// 	}()

// 	<-stopChan
// 	log.Println("shutting down the server ...")
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	srv.Shutdown(ctx)
// 	defer cancel()
// 	log.Println("Server is gratefully")

// }

// func todoHandlers() http.Handler {
// 	rg := chi.NewRouter()
// 	rg.Group(func(r chi.Router) {
// 		r.Get("/", fetchTodos)
// 		r.Post("/", createTodo)
// 		r.Put("/{id}", updateTodo)
// 		r.Delete("/{id}", deleteTodo)
// 	})

// 	return rg
// }

//	func checkErr(err error) {
//		if err != nil {
//			log.Fatal(err)
//		}
//	}
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/thedevsaddam/renderer"
	
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var rnd *renderer.Render
// var db *mgo.Database

const (
	dbName         string = "demo_todo"
	collectionName string = "todo"
	port           string = ":9000"
)

type (
	todoModel struct {
		ID        bson.ObjectId `bson:"_id,omitempty"`
		Title     string        `bson:"title"`
		Completed bool          `bson:"completed"`
		CreatedAt time.Time     `bson:"createAt"`
	}

	todo struct {
		ID        string    `json:"id"`
		Title     string    `json:"title"`
		Completed bool      `json:"completed"`
		CreatedAt time.Time `json:"created_at"`
	}
)

var coll *mongo.Collection

func init() {
	// rnd = renderer.New()
	// // fmt.Println("before mgo.dial function")
	// sess, err := mgo.Dial(hostName)
	// checkErr(err)
	// // fmt.Println("after mgo.dial function")
	// sess.SetMode(mgo.Monotonic, true)
	// db = sess.DB(dbName)

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " +
			"www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll = client.Database(dbName).Collection(collectionName)

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	err := rnd.Template(w, http.StatusOK, []string{"static/home.tpl"}, nil)
	checkErr(err)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var t todo

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		rnd.JSON(w, http.StatusProcessing, err)
		return
	}

	// simple validation
	if t.Title == "" {
		rnd.JSON(w, http.StatusBadRequest, renderer.M{
			"message": "The title field is requried",
		})
		return
	}

	// if input is okay, create a todo
	tm := todoModel{
		ID:        bson.NewObjectId(),
		Title:     t.Title,
		Completed: false,
		CreatedAt: time.Now(),
	}
	if err := db.C(collectionName).Insert(&tm); err != nil {
		rnd.JSON(w, http.StatusProcessing, renderer.M{
			"message": "Failed to save todo",
			"error":   err,
		})
		return
	}

	rnd.JSON(w, http.StatusCreated, renderer.M{
		"message": "Todo created successfully",
		"todo_id": tm.ID.Hex(),
	})
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimSpace(chi.URLParam(r, "id"))

	if !bson.IsObjectIdHex(id) {
		rnd.JSON(w, http.StatusBadRequest, renderer.M{
			"message": "The id is invalid",
		})
		return
	}

	var t todo

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		rnd.JSON(w, http.StatusProcessing, err)
		return
	}

	// simple validation
	if t.Title == "" {
		rnd.JSON(w, http.StatusBadRequest, renderer.M{
			"message": "The title field is requried",
		})
		return
	}

	// if input is okay, update a todo
	if err := db.C(collectionName).
		Update(
			bson.M{"_id": bson.ObjectIdHex(id)},
			bson.M{"title": t.Title, "completed": t.Completed},
		); err != nil {
		rnd.JSON(w, http.StatusProcessing, renderer.M{
			"message": "Failed to update todo",
			"error":   err,
		})
		return
	}

	rnd.JSON(w, http.StatusOK, renderer.M{
		"message": "Todo updated successfully",
	})
}

func fetchTodos(w http.ResponseWriter, r *http.Request) {
	todos := []todoModel{}

	if err := db.C(collectionName).
		Find(bson.M{}).
		All(&todos); err != nil {
		rnd.JSON(w, http.StatusProcessing, renderer.M{
			"message": "Failed to fetch todo",
			"error":   err,
		})
		return
	}

	todoList := []todo{}
	for _, t := range todos {
		todoList = append(todoList, todo{
			ID:        t.ID.Hex(),
			Title:     t.Title,
			Completed: t.Completed,
			CreatedAt: t.CreatedAt,
		})
	}

	rnd.JSON(w, http.StatusOK, renderer.M{
		"data": todoList,
	})
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimSpace(chi.URLParam(r, "id"))

	if !bson.IsObjectIdHex(id) {
		rnd.JSON(w, http.StatusBadRequest, renderer.M{
			"message": "The id is invalid",
		})
		return
	}

	if err := db.C(collectionName).RemoveId(bson.ObjectIdHex(id)); err != nil {
		rnd.JSON(w, http.StatusProcessing, renderer.M{
			"message": "Failed to delete todo",
			"error":   err,
		})
		return
	}

	rnd.JSON(w, http.StatusOK, renderer.M{
		"message": "Todo deleted successfully",
	})
}

func main() {
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)

	r.Mount("/todo", todoHandlers())

	srv := &http.Server{
		Addr:         port,
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Println("Listening on port ", port)
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	<-stopChan
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)
	defer cancel()
	log.Println("Server gracefully stopped!")
}

func todoHandlers() http.Handler {
	rg := chi.NewRouter()
	rg.Group(func(r chi.Router) {
		r.Get("/", fetchTodos)
		r.Post("/", createTodo)
		r.Put("/{id}", updateTodo)
		r.Delete("/{id}", deleteTodo)
	})
	return rg
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err) //respond with error page or message
	}
}