package main

import (
	"errors"
	"fmt"
	"nebil/golang/internal/models"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// http.NotFound(w, r)
		app.notFound(w)
		return
	}

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, snippets := range snippets {
		fmt.Fprintf(w, "%+v\n", snippets)
	}

	// files := []string{
	// 	"./ui/html/base.tmpl",
	// 	"./ui/html/pages/home.tm",
	// 	"./ui/html/partials/nav.tmpl",
	// }

	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	// log.Println(err.Error())
	// 	// app.errorLog.Println(err.Error())
	// 	// http.Error(w, "internal server error", http.StatusInternalServerError)
	// 	app.serverError(w, err)
	// 	return
	// }

	// err = ts.ExecuteTemplate(w, "base", nil)
	// if err != nil {
	// 	// log.Println(err.Error())
	// 	// app.errorLog.Println(err.Error())
	// 	// http.Error(w, "internal server error", http.StatusInternalServerError)
	// 	app.serverError(w, err)
	// 	return
	// }
	// w.Write([]byte("Hello Wordld"))
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	val, err := strconv.Atoi(id)
	if err != nil || val <= 0 {
		// http.NotFound(w, r)
		app.notFound(w)
		return
	}
	// w.Write([]byte("Snippet view"))
	snippet, err := app.snippets.Get(val)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
			return
		}
		app.serverError(w, err)
	}
	fmt.Fprintf(w, "%+v", snippet)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader(http.StatusMethodNotAllowed)
		// w.Write([]byte("method not allowed"))
		// http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7
	fmt.Println("Hello")
	id, err := app.snippets.Insert(title, content, expires)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		app.serverError(w, err)
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
