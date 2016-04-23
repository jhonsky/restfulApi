package myapi

import (
	"encoding/json"
	_ "fmt"
	"io/ioutil"
	"log"
	"net/http"
	"postgresql"
	_ "sync"
	"util"

	"github.com/ant0ine/go-json-rest/rest"
)

func GetAllUsers(w rest.ResponseWriter, r *rest.Request) {
	var users []util.UserType = make([]util.UserType, 10)
	val, err := postgresql.GetAllUsers()
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	users, ok := val.([]util.UserType)
	if !ok {
		rest.Error(w, "get users data error", http.StatusInternalServerError)
		return
	}
	w.WriteJson(&users)
}

func PostUser(w rest.ResponseWriter, r *rest.Request) {
	user := util.User{}
	err := r.DecodeJsonPayload(&user)
	if err != nil {
		log.Fatal(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if user.Name == "" {
		rest.Error(w, "post json error", http.StatusInternalServerError)
		return
	}
	user.Id, err = postgresql.InsertUser(user)
	if err != nil {
		log.Fatal(err)
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//	fmt.Println(user)
	userType := util.UserType{User: util.User{Id: user.Id, Name: user.Name}, Type: "user"}
	w.WriteJson(&userType)
}

func GetUserRelationships(w rest.ResponseWriter, r *rest.Request) {
	userid := r.PathParam("userid")
	var relationShips []util.RelationshipType
	val, err := postgresql.GetRelationships(userid)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	relationShips, ok := val.([]util.RelationshipType)
	if !ok {
		rest.Error(w, "get relationsips data error", http.StatusInternalServerError)
		return
	}
	w.WriteJson(&relationShips)
}

func PutUserRelationships(w rest.ResponseWriter, r *rest.Request) {
	usera := r.PathParam("userid")
	userb := r.PathParam("other_user")

	r.ParseForm()
	result, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()
	var rr interface{}
	json.Unmarshal(result, &rr)
	putData, ok := rr.(map[string]interface{})
	var state string
	if ok {
		for k, v := range putData {
			if k != "state" {
				rest.Error(w, "json key err, key should be state", http.StatusBadRequest)
				return
			}
			state, ok = v.(string)
			if !ok {
				rest.Error(w, "put param error", http.StatusBadRequest)
				return
			}
		}
	}
	if state != "disliked" && state != "liked" {
		rest.Error(w, "put param error", http.StatusBadRequest)
		return
	}
	state_new, err := postgresql.PutRelationships(usera, userb, state)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	relationShip := util.RelationshipType{util.Relationship{Id: userb, State: state_new}, "relationship"}
	w.WriteJson(&relationShip)
}
