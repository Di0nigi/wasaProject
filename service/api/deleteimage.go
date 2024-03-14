package api

import (
	//"encoding/json"
	"net/http"
	
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	//"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	//"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
)

//etHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) deleteImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	id:=ctx.User
	photoid:=ps.ByName("ObjId")

	err:= rt.db.DelPost(id,photoid)
	if err!=nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
    _, err= w.Write([]byte("Photo deleted successfully"))
	if err!= nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		
		return
	}
	

}