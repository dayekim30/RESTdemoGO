package main

import (
	"net/http"

	"github.com/dayekim30/RESTdemoGO/controller"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	r.HandleFunc("/", controller.Homepage)
	r.HandleFunc("/members", controller.GetAllMebers).Methods("GET")
	r.HandleFunc("/member/{id}", controller.Getbyid).Methods("GET")
	r.HandleFunc("/member/post", controller.PostMember).Methods("POST")
	r.HandleFunc("/member/delete/{id}", controller.DeleteMember).Methods("DELETE")
	http.ListenAndServe(":8080", r)
}
