package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) logIn(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var User UserId
	var tk string
	err := json.NewDecoder(r.Body).Decode(&User)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ret, errs := rt.db.GetUserID(User.IDUser)

	if errs != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}

	if !ret {
		err, tk = rt.db.CreateUser(User.IDUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		err, tk = rt.db.GetUserToken(User.IDUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	jsontk, errs := json.Marshal(tk)
	if errs != nil {
		http.Error(w, errs.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsontk)

}
