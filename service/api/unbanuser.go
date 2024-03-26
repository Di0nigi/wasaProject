package api

import (
	
	"net/http"
	
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	
)

func (rt *_router) unBanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	bannedUser :=ps.ByName("banned")
	id:=ctx.User
	
	err := rt.db.DelBlocked(bannedUser,id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
    _, err=w.Write([]byte("User unbanned"))
	
	
}