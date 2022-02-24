package handler

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

type feedback struct {
	ID int
	UserName string
	Description string
	Rating int
}

var feedbackList = []feedback{
	{ ID: 1, UserName: "Kevin", Description: "This service is so good!", Rating: 5 },
	{ ID: 2, UserName: "Kate", Description: "Thank you!", Rating: 4 },
	{ ID: 3, UserName: "Ben", Description: "Not too bad, but all is okay.", Rating: 3 },
}

func FeedbackHandler(w http.ResponseWriter, r *http.Request) {

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

// Basic CRUD API
func APIFeedbackHandler(w http.ResponseWriter, r *http.Request)  {
	
	w.Header().Set("Content-Type", "application/json")

	// r.Method is the variable to get http request method of client request
	switch r.Method {
		// GET request
		case "GET":
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(feedbackList)
			
		// POST request
		case "POST":
			var newFeedback feedback
			
			jsonDecoder := json.NewDecoder(r.Body)
			jsonDecoder.DisallowUnknownFields()
			// Catch unwanted fields

			err := jsonDecoder.Decode(&newFeedback)
			
			if err != nil {
				// Bad JSON or unrecognized json field
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if newFeedback.ID <= 0 || newFeedback.UserName == "" || newFeedback.Description == "" || newFeedback.Rating <= 0 {
				// Check missing field
				http.Error(w, "Missing field. Please ensure that ID (int), UserName (string), Description (string) and Rating (int) was included.", http.StatusBadRequest)
    			return
			}

			if newFeedback.ID <= len(feedbackList) {
				http.Error(w, "ID must be more than " + strconv.Itoa(len(feedbackList)), http.StatusBadRequest)
				return
			}

			if newFeedback.Rating > 5 {
				http.Error(w, "Invalid rating. Maximum is 5", http.StatusBadRequest)
				return
			}

			if jsonDecoder.More() {
				http.Error(w, "Extraneous data after JSON object, only ID (int), UserName (string), Description (string) and Rating (int).", http.StatusBadRequest)
    			return
			}

			feedbackList = append(feedbackList, newFeedback)

			log.Println(newFeedback)
			log.Println(feedbackList)

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(newFeedback)

		// PATCH request
		case "PATCH":
			var updatedFeedback feedback

			jsonDecoder := json.NewDecoder(r.Body)
			jsonDecoder.DisallowUnknownFields()

			err := jsonDecoder.Decode(&updatedFeedback)

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if updatedFeedback.ID <= 0 || updatedFeedback.UserName == "" || updatedFeedback.Description == "" || updatedFeedback.Rating <= 0 {
				http.Error(w, "Missing field. Please ensure that ID (int), UserName (string), Description (string) and Rating (int) was included.", http.StatusBadRequest)
    			return
			}

			if updatedFeedback.Rating > 5 {
				http.Error(w, "Invalid rating. Maximum is 5", http.StatusBadRequest)
				return
			}

			if jsonDecoder.More() {
				http.Error(w, "Extraneous data after JSON object, only ID (int), UserName (string), Description (string) and Rating (int).", http.StatusBadRequest)
    			return
			}

			var isFound bool = false
			var theIndex int = -1

			for index, theFeedback := range feedbackList {
				if updatedFeedback.ID == theFeedback.ID && updatedFeedback.UserName == theFeedback.UserName  {
					theIndex = index
					isFound = true
				}
			}

			if !isFound {
				http.Error(w, "ID or the UserName is invalid. Try again.", http.StatusBadRequest)
				return
			}

			feedbackList = deleteFeedback(feedbackList, theIndex)
			feedbackList = append(feedbackList, updatedFeedback)

			log.Println(updatedFeedback)
			log.Println(feedbackList)

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(updatedFeedback)
			
		// DELETE request
		case "DELETE":
			type deletedFeedback struct {
				ID int
				UserName string
			}

			var toBeDeleted deletedFeedback

			jsonDecoder := json.NewDecoder(r.Body)
			jsonDecoder.DisallowUnknownFields()

			err := jsonDecoder.Decode(&toBeDeleted)

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if toBeDeleted.ID <= 0 || toBeDeleted.UserName == "" {
				http.Error(w, "Missing field. Please ensure that ID (int) and UserName (string).", http.StatusBadRequest)
    			return
			}

			if jsonDecoder.More() {
				http.Error(w, "Extraneous data after JSON object, only ID (int) and UserName (string)", http.StatusBadRequest)
					return
			}

			var isFound bool = false
			var theIndex int = -1

			for index, theFeedback := range feedbackList {
				if toBeDeleted.ID == theFeedback.ID && toBeDeleted.UserName == theFeedback.UserName  {
					theIndex = index
					isFound = true
				}
			}

			if !isFound {
				http.Error(w, "ID or the UserName is invalid. Try again.", http.StatusBadRequest)
				return
			}

			feedbackList = deleteFeedback(feedbackList, theIndex)

			log.Println(feedbackList)

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(feedbackList)

		// If the user send other http request method
		default:
			http.Error(w, "Invalid request method", http.StatusBadRequest)

	}

}

func deleteFeedback(theFeedbackList []feedback, indexToDelete int) []feedback  {
	index := 0
	for i, theFeedback := range theFeedbackList {
		if i != indexToDelete {
			theFeedbackList[index] = theFeedback
			index++
		}
	}

	return theFeedbackList[:index]
}
