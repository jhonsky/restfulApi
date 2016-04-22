package myapi

import (
	"encoding/json"
	_ "fmt"
	"io/ioutil"
	_ "log"
	"net/http"
	"postgresql"
	_ "sync"
	"util"

	"github.com/ant0ine/go-json-rest/rest"
)

func GetAllUsers(w rest.ResponseWriter, r *rest.Request) {
	var users []util.UserType = make([]util.UserType, 10)
	val := postgresql.GetAllUsers()
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
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.Id = postgresql.InsertUser(user)
	//	fmt.Println(user)
	userType := util.UserType{User: util.User{Id: user.Id, Name: user.Name}, Type: "user"}
	w.WriteJson(&userType)
}

func GetUserRelationships(w rest.ResponseWriter, r *rest.Request) {
	userid := r.PathParam("userid")
	var relationShips []util.RelationshipType
	val := postgresql.GetRelationships(userid)
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
		for _, v := range putData {
			state, ok = v.(string)
			if !ok {
				rest.Error(w, "put param error", http.StatusBadRequest)
				return
			}
		}
	}
	if state != "disliked" && state != "liked" && state != "matched" {
		rest.Error(w, "put param error", http.StatusBadRequest)
		return
	}
	state_new := postgresql.PutRelationships(usera, userb, state)
	relationShip := util.RelationshipType{util.Relationship{Id: userb, State: state_new}, "relationship"}
	w.WriteJson(&relationShip)
}
