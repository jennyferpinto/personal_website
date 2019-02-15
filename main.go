package main

import (
	"html/template"
	"net/http"
	"time"

	"google.golang.org/appengine"
)

//Welcome struct holds information to be displayed in our HTML file
type Welcome struct {
	Time string
}

//Form struct
type Form struct {
	Email   string
	Subject string
	Message string
}

//Go application entrypoint
func main() {

	welcome := Welcome{time.Now().Format(time.Stamp)}

	mainTemplate := template.Must(template.ParseFiles("templates/index.html"))

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))
	// Our html code would therefore be <link rel="stylesheet"  href="/static/stylesheet/...">
	//It is important to note the url in http.Handle can be whatever we like, so long as we are consistent.

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//If errors show an internal server error message
		if err := mainTemplate.ExecuteTemplate(w, "index.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	formTemplate := template.Must(template.ParseFiles("templates/form.html"))
	//This method takes in the URL path "/" and a function that takes in a response writer, and a http request.
	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			formTemplate.Execute(w, nil)
			return
		}

		details := Form{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		// do something with details
		_ = details

		formTemplate.Execute(w, struct{ Success bool }{true})

	})

	appengine.Main()

	//Start the web server, set the port to listen to 8080. Without a path it assumes localhost
	//Print any errors from starting the webserver using fmt
	// fmt.Println("Listening")
	// fmt.Println(http.ListenAndServe(":8000", nil))
}
