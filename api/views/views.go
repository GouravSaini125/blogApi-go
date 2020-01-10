package views

import (
	"encoding/json"
	"fmt"
	"github.com/GouravSaini125/blog/db"
	"github.com/GouravSaini125/blog/internal/models"
	"net/http"
	"strconv"
)

// Index homepage
func Index(w http.ResponseWriter, r *http.Request) {
	var objs []models.Blog

	db := db.DbConn()
	selDB, err := db.Query("SELECT * FROM blogs ORDER BY id")
	if err != nil {
		panic(err.Error())
	}
	for selDB.Next() {
		var id int
		var title, text string
		err = selDB.Scan(&id, &title, &text)
		if err != nil {
			panic(err.Error())
		}
		objs = append(objs, models.Blog{
			ID:    id,
			Title: title,
			Text:  text,
		})
	}
	defer db.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	d, err := json.Marshal(objs)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Write(d)
}

// Add Post blog
func Add(w http.ResponseWriter, r *http.Request) {

	var obj models.Blog

	// decoder := json.NewDecoder(r.Body)

	// var blog models.Blog
	// err := decoder.Decode(&blog)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(blog.Title)
	// fmt.Println(blog.Text)

	db := db.DbConn()
	title := r.FormValue("title")
	text := r.FormValue("text")
	insForm, err := db.Prepare("INSERT INTO blogs(title, text) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}
	s, err := insForm.Exec(title, text)
	if err != nil {
		panic(err.Error())
	}

	id, err := s.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	obj = models.Blog{
		ID:    int(id),
		Title: title,
		Text:  text,
	}

	defer db.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	d, err := json.Marshal(obj)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Write(d)

}

// GetBlog Sing Blog
func GetBlog(w http.ResponseWriter, r *http.Request) {

	var obj models.Blog

	db := db.DbConn()
	nID := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM blogs WHERE id=?", nID)
	if err != nil {
		panic(err.Error())
	}
	for selDB.Next() {
		var id int
		var title, text string
		err = selDB.Scan(&id, &title, &text)
		if err != nil {
			panic(err.Error())
		}
		obj = models.Blog{
			ID:    id,
			Title: title,
			Text:  text,
		}
	}
	defer db.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	d, err := json.Marshal(obj)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Write(d)

}

// DestroyBlog delete blog
func DestroyBlog(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	emp := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM blogs WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)
	defer db.Close()
}

//UpdateBlog edit uur blogs
func UpdateBlog(w http.ResponseWriter, r *http.Request) {

	var obj models.Blog

	db := db.DbConn()
	if r.Method == "POST" {
		title := r.FormValue("title")
		text := r.FormValue("text")
		nID := r.FormValue("id")
		id, err := strconv.Atoi(nID)
		insForm, err := db.Prepare("UPDATE blogs SET title=?, text=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(title, text, id)
		obj = models.Blog{
			ID:    id,
			Title: title,
			Text:  text,
		}
	}
	defer db.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	d, err := json.Marshal(obj)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Write(d)

}
