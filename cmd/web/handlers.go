package main

import (
	"net/http"
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
