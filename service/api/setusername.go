package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var User UserId
	oldUser := ctx.User
	err := json.NewDecoder(r.Body).Decode(&User)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = rt.db.ChangeUsername(oldUser, User.IDUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	_, err = w.Write([]byte("Username created successfully"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

}
