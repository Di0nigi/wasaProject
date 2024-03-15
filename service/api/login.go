package api

import (
	"encoding/json"
	"net/http"
	//"fmt"
	
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	//"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
)

//etHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) logIn(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var User UserId
	err := json.NewDecoder(r.Body).Decode(&User)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ret, err := rt.db.GetUserID(User.IDUser)
	//str := fmt.Sprintf("%t", ret)
	//fmt.Printf(str)
	if !ret{
		err = rt.db.CreateUser(User.IDUser)
		if err!= nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(User)
	
}
