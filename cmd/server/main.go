package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

type Page struct {
	Title   string
	Heading string
}
type Posts struct {
	Posts []Post
}

type Post struct {
	Title string
}

func renderTemplate(w http.ResponseWriter, templatePath string, page *Page, posts Posts) {
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		PageData *Page
		Posts    Posts
	}{
		PageData: page,
		Posts:    posts,
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	postPath := "../../templates/posts/"
	var posts []Post

	entries, err := os.ReadDir(postPath)
	if err != nil {
		fmt.Println("homehandler could ned read dir path")
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		post := Post{Title: entry.Name()}
		posts = append(posts, post)
	}

	postData := Posts{Posts: posts}
	pageData := &Page{Title: "File-Serving | Home", Heading: "File-Transfer Server"}

	renderTemplate(w, "../../templates/index.html", pageData, postData)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	var posts Posts
	pathToTemplates := "../../templates/posts/"
	postName := strings.TrimPrefix(r.URL.Path, "/post/")
	postData := &Page{Title: postName, Heading: postName}
	if postName == "" {
		http.NotFound(w, r)
		return
	}
	postPath := pathToTemplates + postName
	renderTemplate(w, postPath, postData, posts)
}

func loginhandler(w http.ResponseWriter, r *http.Request) {
	var postData Posts
	pageData := &Page{Title: "File-Serving | Login", Heading: "Login to Admin Account"}
	renderTemplate(w, "../../templates/login.html", pageData, postData)
}

func userLoginHandler(w http.ResponseWriter, r *http.Request) {
	//logic to log the user in if valid.
	err := r.ParseForm()
	if err != nil {
		fmt.Println("error parsing form")
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username != "jeffmuter" || password != "phoenix" {
		// give the user authenticated jwt
	}

	homeHandler(w, r) // send the logged in user to the home page.
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /post/", postHandler)
	mux.HandleFunc("GET /login", loginhandler)
	mux.HandleFunc("GET /", homeHandler)

	mux.HandleFunc("POST /login/", userLoginHandler)

	log.Printf("Started server on 8080:\nhttp://localhost:8080")

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
