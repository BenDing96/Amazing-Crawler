package main

import (
	"net/http"

	"learngo/crawler_official/frontend/controller"
)

func main() {
	http.Handle("/", http.FileServer(
		http.Dir("/Users/dingzheyu/Desktop/IdeaProject/learngo/crawler_official/frontend/view")))
	http.Handle("/search",
		controller.CreateSearchResultHandler(
			"/Users/dingzheyu/Desktop/IdeaProject/learngo/crawler_official/frontend/view/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
