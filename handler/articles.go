package handler

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

func ArticleHandler(w http.ResponseWriter, r *http.Request) {

	theArticle := map[string]string{
		"title": "How to survive economy in this Pandemic",
		"content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Augue lacus viverra vitae congue eu consequat ac felis. Vitae purus faucibus ornare suspendisse sed. Feugiat in ante metus dictum at tempor commodo ullamcorper. Id donec ultrices tincidunt arcu non sodales neque sodales. Vitae sapien pellentesque habitant morbi tristique senectus et netus et. At auctor urna nunc id cursus metus aliquam eleifend mi. Massa id neque aliquam vestibulum morbi. Phasellus egestas tellus rutrum tellus pellentesque. Accumsan lacus vel facilisis volutpat est velit. Sagittis id consectetur purus ut faucibus. Faucibus a pellentesque sit amet porttitor. Tempus urna et pharetra pharetra massa massa ultricies mi quis. Et tortor at risus viverra adipiscing. Turpis egestas pretium aenean pharetra magna ac placerat. ",
	}
	// This is the data that to be passing to the html file

	page, err := template.ParseFiles(path.Join("views", "article.html"), path.Join("views/layout", "footer.html"))
	// When we create a layout component, this ParseFiles function
	// must contains the main page and all the layout component
	// For example, the main is the article.html
	// and the layout component is footer.html

	if err != nil {
		log.Println(err)
		http.Error(w, "Error was occured :(\nKeep calm, we will fix it later...", http.StatusInternalServerError)
		return
	}

	err = page.Execute(w, theArticle)
	// Pass the data in the second parameter
	// To get the data in the html, simply use the key (if we use map)
	// or use the attribute (if we use struct)

	if err != nil {
		log.Println(err)
		http.Error(w, "Error was occured :(\nKeep calm, we will fix it later...", http.StatusInternalServerError)
		return
	}

}
