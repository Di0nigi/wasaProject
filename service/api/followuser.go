package api

import (
	"encoding/json"
	"net/http"
	"fmt"
	
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	//"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
)

//etHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var toFollow UserId
	id:=ctx.User
	fmt.Printf(id)
	fmt.Printf("dioo")
	err := json.NewDecoder(r.Body).Decode(&toFollow)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = rt.db.AddFollower(toFollow.IDUser,id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	_, err= w.Write([]byte("User followed"))
	if err!= nil {
		//http.Error(w, err.Error(), http.StatusBadRequest
		return
	}
	

}