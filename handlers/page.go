package handlers

// handlers for the creation of pages

import (
	"fmt"
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

func RenderTemplate(w http.ResponseWriter, templatePath string, page *Page, posts Posts) {
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

func HomeHandler(w http.ResponseWriter, r *http.Request) {
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

	RenderTemplate(w, "../../templates/index.html", pageData, postData)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var posts Posts
	pathToTemplates := "../../templates/posts/"
	postName := strings.TrimPrefix(r.URL.Path, "/post/")
	postData := &Page{Title: postName, Heading: postName}
	if postName == "" {
		http.NotFound(w, r)
		return
	}
	postPath := pathToTemplates + postName
	RenderTemplate(w, postPath, postData, posts)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var postData Posts
	pageData := &Page{Title: "File-Serving | Login", Heading: "Login to Admin Account"}
	RenderTemplate(w, "../../templates/login.html", pageData, postData)
}

// func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
// 	//logic to log the user in if valid.
// 	err := r.ParseForm()
// 	if err != nil {
// 		fmt.Println("error parsing form")
// 		return
// 	}
// 	username := r.FormValue("username")
// 	password := r.FormValue("password")

// 	HomeHandler(w, r) // send the logged in user to the home page.
// }
