package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var toFollow UserId
	id := ctx.User
	err := json.NewDecoder(r.Body).Decode(&toFollow)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err, _ = rt.db.GetBanned(toFollow.IDUser, id)
	if err == nil {
		err = errors.New("Banned")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = rt.db.AddFollower(toFollow.IDUser, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	_, err = w.Write([]byte("User followed"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

	}

}
