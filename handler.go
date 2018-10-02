package cms

import (
	"net/http"
	"strings"
	"time"
)

func ServePage(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimLeft(r.URL.Path, "/page/")

	if path == "" {
		http.NotFound(w, r)
		return
	}

	p := &Page{
		Title:   strings.ToTitle(path),
		Content: "here is my page",
	}

	Tmpl.ExecuteTemplate(w, "page", p)
}

func ServePost(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimLeft(r.URL.Path, "/post/")

	if path == "" {
		http.NotFound(w, r)
		return
	}

	p := &Post{
		Title:   strings.ToTitle(path),
		Content: "here is my post, no comments",
	}

	Tmpl.ExecuteTemplate(w, "post", p)
}

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "Go Projects CMS",
		Content: "welcome to our home page",
		Posts: []*Post{
			{
				Title:         "Hello World",
				Content:       "Hello World! Thanks for coming to the site",
				DatePublished: time.Now(),
			},
			{
				Title:         "A post with comments",
				Content:       "Here is a post with comments.",
				DatePublished: time.Now().Add(-time.Hour),
				Comments: []*Comment{
					&Comment{
						Author:        "Tim",
						Comment:       "I like this post",
						DatePublished: time.Now().Add(time.Hour / 2),
					},
				},
			},
		},
	}

	Tmpl.ExecuteTemplate(w, "page", p)
}

func HandleNew(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		Tmpl.ExecuteTemplate(w, "new", nil)

	case "POST":
		title := r.FormValue("title")
		content := r.FormValue("content")
		contentType := r.FormValue("content-type")
		r.ParseForm()

		if contentType == "page" {
			Tmpl.ExecuteTemplate(w, "page", &Page{
				Title:   title,
				Content: content,
			})
			return
		}

		if contentType == "post" {
			Tmpl.ExecuteTemplate(w, "post", &Post{
				Title:   title,
				Content: content,
			})
			return
		}
	default:
		http.Error(w, "Method not supported "+r.Method, http.StatusMethodNotAllowed)
	}
}
