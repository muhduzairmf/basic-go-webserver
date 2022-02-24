package handler

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

func ContactUsFormHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Error occured. Request not allowed.", http.StatusBadRequest)
		return
	}

	page, err := template.ParseFiles(path.Join("views", "contact-us.html"))

	if err != nil {
		log.Println(err)
		http.Error(w, "Error was occured :(\nKeep calm, we will fix it later...", http.StatusInternalServerError)
		return
	}

	err = page.Execute(w, nil)

	if err != nil {
		log.Println(err)
		http.Error(w, "Error was occured :(\nKeep calm, we will fix it later...", http.StatusInternalServerError)
		return
	}

}

func SubmitFormHandler(w http.ResponseWriter, r *http.Request) {
	
	if r.Method != "POST" {
		http.Error(w, "Error occured. Request not allowed.", http.StatusBadRequest)
		return
	}

	err := r.ParseForm()
	// ParseForm function is getting the incoming data from form tag html of the client browser
	// It will help to encode the incoming data, and we can get and use in Go

	if err != nil {
		log.Println(err)
		http.Error(w, "Error was occured :(\nKeep calm, we will fix it later...", http.StatusInternalServerError)
		return
	}

	// name := r.Form.Get("name")
	// email := r.Form.Get("email")
	// queries := r.Form.Get("queries")
	formDetail := map[string]string {
		"name": r.Form.Get("name"),
		"email": r.Form.Get("email"),
		"queries": r.Form.Get("queries"),
	}

	page, err := template.ParseFiles(path.Join("views", "submitted.html"))

	if err != nil {
		log.Println(err)
		http.Error(w, "Error was occured :(\nKeep calm, we will fix it later...", http.StatusInternalServerError)
		return
	}

	err = page.Execute(w, formDetail)

	if err != nil {
		log.Println(err)
		http.Error(w, "Error was occured :(\nKeep calm, we will fix it later...", http.StatusInternalServerError)
		return
	}

}