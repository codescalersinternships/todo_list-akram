package main

import (
	"net/http"
	_ "time"
)

func renderTemplate(w http.ResponseWriter, tmpl string) {
	p := ""
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func fetchTodos() ([]*Todo, error) {
	i1 := &Todo{title: "first title", id: 0, completed: false}
	i2 := &Todo{title: "second title", id: 1, completed: true}

	var p = []*Todo{}
	p = append(p, i1)
	p = append(p, i2)
	// return json.MarshalIndent(p, "", "  ")
	return p, nil
}
