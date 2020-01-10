package main

import (
	"net/http"

	"github.com/GouravSaini125/blog/api/views"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", views.Index)
	r.HandleFunc("/add", views.Add).Methods("POST")
	r.HandleFunc("/getblog", views.GetBlog)
	r.HandleFunc("/destroyblog", views.DestroyBlog)
	r.HandleFunc("/updateblog", views.UpdateBlog).Methods("POST")

	http.ListenAndServe(":8080", r)

}
