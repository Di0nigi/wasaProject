package api
import (
	"encoding/json"
	"net/http"
	
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	//"fmt"
	//"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
)
type CustomError struct {
	message string
}

//etHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	id:=ps.ByName("AccountId")
	

	res, errs:= rt.db.GetUserID(id)
	if errs !=nil{
		http.Error(w, errs.Error(), http.StatusBadRequest)
		return
	}

	if (res==true){
	err, userP := rt.db.GetUserProfile(id)
	if err !=nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(userP)
	}else{
		er := NewCustomError("User not found")
		http.Error(w, er.Error(), http.StatusBadRequest)
	}

}
func (ce *CustomError) Error() string {
    return ce.message
}

// NewCustomError creates a new CustomError instance with the provided message
func NewCustomError(message string) error {
    return &CustomError{message: message}
}