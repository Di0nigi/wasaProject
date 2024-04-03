package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var bannedUser UserId
	id := ctx.User
	err := json.NewDecoder(r.Body).Decode(&bannedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = rt.db.AddBlocked(bannedUser.IDUser, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusNoContent)
	_, err = w.Write([]byte("User Banned"))
	/*if err!= nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
	*/
}
