package main

import(
	"fmt"
	"net/http"
	"encoding/json"
)

type article struct {
	ID		int
	Title	string	
	Content string
}

var data = []article{
	{1, "Rois Ganteng", "Rois Ganteng banget"},
	{2, "Rois baik hati", "Rois baik hati banget"},
}

func articles(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var result, err = json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "Invalid", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/articles", articles)
	fmt.Println("Starting Web Server at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}