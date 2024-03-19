package api


import (
	//"encoding/json"
	"net/http"
	
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	//"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
)

//etHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) unCommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	comm :=ps.ByName("commentId")
	id:=ctx.User
	ph:=ps.ByName("photo")
	
	
	
	err := rt.db.DelComment(id,comm,ph)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
    /*_, err=w.Write([]byte("Comment deleted"))
	if err!= nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		
		return
	}*/

}	
