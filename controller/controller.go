package controller

import (
	"encoding/json"
	"net/http"

	"github.com/dayekim30/RESTdemoGO/model"
	"github.com/dayekim30/RESTdemoGO/view"

	"github.com/gorilla/mux"
)

type MemberController struct {
}

func Homepage(w http.ResponseWriter, r *http.Request) {
	view.ViewParse(w, "static/welcome.html")
}

func Getbyid(w http.ResponseWriter, r *http.Request) {
	pathVariables := mux.Vars(r)
	result := model.Getbyid(pathVariables["id"])
	test, _ := json.Marshal(result)
	w.Write(test)
}
func GetAllMebers(w http.ResponseWriter, r *http.Request) {

	memberlist := model.Getall()
	test, _ := json.Marshal(memberlist)
	w.Write(test)
}
func PostMember(w http.ResponseWriter, r *http.Request) {

	newMember := model.Member{}

	err := r.ParseForm()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	newMember.ID = r.Form.Get("id")

	newMember.Name = r.Form.Get("name")
	newMember.FavoriteTeam = r.Form.Get("team")

	memberlist := model.Addmember(newMember)
	test, _ := json.Marshal(memberlist)
	w.Write(test)
}
func DeleteMember(w http.ResponseWriter, r *http.Request) {

	pathVariables := mux.Vars(r)
	result := model.Deletemember(pathVariables["id"])
	test, _ := json.Marshal(result)
	w.Write(test)

}
