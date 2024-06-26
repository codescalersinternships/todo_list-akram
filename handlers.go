package main

import (
	"encoding/json"
	_ "errors"
	"fmt"
	_ "html/template"
	_ "log"
	"net/http"
	_ "os"
	_ "regexp"
	_ "time"

	"github.com/julienschmidt/httprouter"
)

// func ViewHandler(w http.ResponseWriter, r *http.Request) {
// 	// title, err := getTitle(w, r)
// 	// if err != nil {
// 	// 	return
// 	// }

// 	// p, err := loadPage()

// 	// if err != nil {
// 	// 	http.Redirect(w, r, "/edit/"+title, http.StatusFound)
// 	// 	return
// 	// }
// 	renderTemplate(w)
// }

// func EditHandler(w http.ResponseWriter, r *http.Request) {
// 	// title, err := getTitle(w, r)
// 	// if err != nil {
// 	// 	return
// 	// }
// 	// p, err := loadPage(title)
// 	// if err != nil {
// 	// 	p = &Page{Title: title}
// 	// }

// 	// renderTemplate(w, "edit", p)
// 	fmt.Println("Inside edit Handler,Note: editHandler is currently empty")
// }

// func SaveHandler(w http.ResponseWriter, r *http.Request) {
// 	// title, err := getTitle(w, r)
// 	// if err != nil {
// 	// 	return
// 	// }
// 	// body := r.FormValue("body")
// 	// p := &Page{Title: title, Body: []byte(body)}
// 	// err = p.save()
// 	// if err != nil {
// 	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
// 	// 	return
// 	// }
// 	// http.Redirect(w, r, "/view/"+title, http.StatusFound)
// }

// func GetTodos() () {
// 	singleTodo := &Todo{Title: "first single todo", CreatedAt: time.Now()}
// 	jsonReady, err := json.Marshal(singleTodo)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return json.MarshalIndent(jsonReady, "", "  ")

// }

// func AddTodo() {

// }

// func DeleteTodo() {

// }

// func UpdateTodo() {

// }

// func Index(w http.ResponseWriter, r *http.Request) {
// 	response, err := getJsonResponse()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Fprintf(w, string(response))
// }

// func getJsonResponse() ([]byte, error) {

// 	singleTodo := &Todo{Title: "first single todo", CreatedAt: time.Now()}
// 	jsonReady, err := json.Marshal(singleTodo)

// 	if err != nil {
// 		panic(err)
// 	}
// 	return json.MarshalIndent(jsonReady, "", "  ")
// }

func RootHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("Root handler called")

	renderTemplate(w, "home")
}

func GetHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	response, err := fetchTodos()
	if err != nil {
		panic(err)
	}
	fmt.Println("GET handler called")
	w.Header().Set("Content-Type", "application/json")
	// fmt.Fprintf(w, string(response))
	json.NewEncoder(w).Encode(response)
}

func PostHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	response, err := fetchTodos()
	if err != nil {
		panic(err)
	}
	// fmt.Fprintf(w, string(response))
	// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
	fmt.Println("Post handler called")
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	json.NewEncoder(w).Encode(response)

	// json.Unmarshal()

}

func PutHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// response, err := fetchTodos()
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println("PUT handler called")

	// json.NewEncoder(w).Encode(response)

}

func DeleteHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// response, err := fetchTodos()
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println("DELETE handler called")
	// json.NewEncoder(w).Encode(response)

	// fmt.Fprintf(w, string(response))
}
