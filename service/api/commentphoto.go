package api


import (
	"encoding/json"
	"net/http"
	//"fmt"
	
	
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	//"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
)

//etHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var comm database.Comment
	id:=ctx.User
	err := json.NewDecoder(r.Body).Decode(&comm)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err= rt.db.AddComment(id,comm)
	//mt.Printf(err.Error())
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