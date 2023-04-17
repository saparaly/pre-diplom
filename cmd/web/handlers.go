package main

import (
	"errors"
	"net/http"
	"strconv"
)

// userIdPost is loged in users id
var (
	userIdPost int
	userName   string
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		ErrorHandler(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		app.clienrError(w, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		// app.notFound(w)
		ErrorHandler(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// fmt.Println(userid)

	s, err := app.posts.Latest()
	if err != nil {
		// fmt.Println("444")
		app.serverError(w, err)
		ErrorHandler(w, http.StatusText(http.StatusInternalServerError), 500)
		return
	}
	// fmt.Println("555")

	files := []string{
		"./ui/html/home.page.html",
		"./ui/html/base.layout.html",
		"./ui/html/footer.html",
	}
	// fmt.Println("222")

	// Use the new render helper.
	userid := app.GetUserIDForUse(w, r)
	app.render(w, r, files, &templateData{
		Posts:  s,
		UserID: userid,
	})
	// fmt.Println("666")
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		ErrorHandler(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		app.clienrError(w, http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	post, err := app.posts.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			// app.notFound(w)
			ErrorHandler(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			ErrorHandler(w, http.StatusText(http.StatusInternalServerError), 500)
			// app.serverError(w, err)
		}
		return
	}

	if post == nil {
		// app.notFound(w)
		ErrorHandler(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	comments, err := app.posts.GetComments(post.Id)
	if err != nil {
		// app.serverError(w, err)
		ErrorHandler(w, http.StatusText(http.StatusInternalServerError), 500)
		return
	}

	files := []string{
		"./ui/html/show.page.html",
		"./ui/html/base.layout.html",
		"./ui/html/footer.html",
	}

	userid := app.GetUserIDForUse(w, r)
	// Use the new render helper.
	app.render(w, r, files, &templateData{
		Post:     post,
		Comments: comments,
		UserID:   userid,
	})
}
