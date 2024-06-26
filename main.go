package main

// import (
// 	_ "errors"
// 	_ "fmt"
// 	"html/template"
// 	"log"
// 	"net/http"
// 	_ "os"
// 	_ "regexp"
// 	"time"

// 	"github.com/gorilla/mux"
// )

// var templates = template.Must(template.ParseFiles("home.html"))

// // var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// // type Page struct {
// // 	Title string
// // 	Body  []byte
// // }

// var todos []*Todo = []*Todo{}

// // func (p *Page) save() error {
// // 	filename := p.Title + ".txt"
// // 	return os.WriteFile(filename, p.Body, 0600)
// // }

// // func loadPage(title string) (*Page, error) {
// // 	// filename := "home.html"
// // 	// body, err := os.ReadFile(filename)
// // 	// if err != nil {
// // 	// 	return nil, err
// // 	// }
// // 	return &Page{Title: title, Body: body}, nil
// // }

// // func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
// // 	m := validPath.FindStringSubmatch(r.URL.Path)
// // 	if m == nil {
// // 		http.NotFound(w, r)
// // 		return "", errors.New("invalid Page Title")
// // 	}
// // 	return m[2], nil // The title is the second subexpression.
// // }

// func renderTemplate(w http.ResponseWriter) {
// 	// p := &Page{Title: "the title", Body: []byte("The body")}
// 	err := templates.ExecuteTemplate(w, "home.html", todos)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }

// // type (
// // 	todoModel struct {
// // 	 ID        bson.ObjectId bson:"_id,omitempty"
// // 	 Title     string        bson:"title"
// // 	 Completed bool          bson:"completed"
// // 	 CreatedAt time.Time     bson:"createAt"
// // 	}

// // 	todo struct {
// // 	 ID        string    json:"id"
// // 	 Title     string    json:"title"
// // 	 Completed bool      json:"completed"
// // 	 CreatedAt time.Time json:"created_at"
// // 	}
// // )

// func main() {
// 	todos = append(todos, &Todo{Title: "first todo", CreatedAt: time.Now()})
// 	r := mux.NewRouter()
// 	r.HandleFunc("/todo", GetTodos).Methods("GET")
// 	r.HandleFunc("/todo", AddTodo).Methods("POST")
// 	r.HandleFunc("/todo/{id}", UpdateTodo).Methods("PUT")
// 	r.HandleFunc("/todo/{id}", DeleteTodo).Methods("DELETE")
// 	log.Fatal(http.ListenAndServe(":8080", nil))

// }

import (
	_ "encoding/json"
	_ "fmt"
	"html/template"
	"log"
	"net/http"
	_ "time"

	"github.com/julienschmidt/httprouter"
)

var templates = template.Must(template.ParseFiles("home.html"))

func main() {
	router := httprouter.New()
	// done
	router.GET("/", RootHandler)
	// done
	router.GET("/todo", GetHandler)

	// ...
	router.POST("/todo", PostHandler)
	// ...
	router.PUT("/todo/{id}", PutHandler)
	// ...
	router.DELETE("/todo/{id}", DeleteHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
