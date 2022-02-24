package handler

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

type feedback struct {
	UserName string
	Description string
	Rating int
}

func FeedbackHandler(w http.ResponseWriter, r *http.Request) {

	feedbackList := []feedback{
		{ UserName: "Kevin", Description: "This service is so good!", Rating: 5 },
		{ UserName: "Kate", Description: "Thank you!", Rating: 4 },
		{ UserName: "Ben", Description: "Not too bad, but all is okay.", Rating: 3 },
	}
	// Slices of a struct data type

	page, err := template.ParseFiles(path.Join("views", "feedback.html"))

	if err != nil {
		log.Println(err)
		http.Error(w, "Error was occured :(\nKeep calm, we will fix it later...", http.StatusInternalServerError)
		return
	}

	err = page.Execute(w, feedbackList)

	if err != nil {
		log.Println(err)
		http.Error(w, "Error was occured :(\nKeep calm, we will fix it later...", http.StatusInternalServerError)
		return
	}

}