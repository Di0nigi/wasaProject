package api


import (
	"encoding/json"
	"net/http"
	
	
	
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	
)


func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	
	//_,errdb := rt.db.GetBanned(comm.Owner.IDUser,id)

	if errdb!=nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	
	var comm database.Comment
	id:=ctx.User
	err := json.NewDecoder(r.Body).Decode(&comm)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err= rt.db.AddComment(id,comm)
	
	if err!=nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}
	w.WriteHeader(http.StatusNoContent)
    _, err= w.Write([]byte("Comment added"))
	/*if err!= nil {
		http.Error(w, err.Error(), http.StatusBadRequest)	
		return
	}*/

}