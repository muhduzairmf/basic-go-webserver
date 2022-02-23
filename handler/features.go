package handler

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

func FeaturesHandler(w http.ResponseWriter, r *http.Request) {

	page, err := template.ParseFiles(path.Join("views", "features.html"))
	// Create a variable to get the specific html file
	// Find the views folder and get the features.html file

	if err != nil {
		// If Go cannot get the files or error something else
		log.Println(err)
		http.Error(w, "Error was occured :(\nKeep calm, we will fix it later...", http.StatusInternalServerError)
		return
	}

	page.Execute(w, nil)
	// Display the html to this path/route

}