package api

import (
	"encoding/json"
	"net/http"
	
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	//"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
)

func (rt *_router) unBanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var bannedUser UserId
	id:=ps.ByName("UserId")
	err := json.NewDecoder(r.Body).Decode(&bannedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = rt.db.DelBlocked(bannedUser.IDUser,id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
    w.Write([]byte("User unbanned"))
	return
}