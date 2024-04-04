package api

import (
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unLikePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	lk := ps.ByName("likeId")
	id := ctx.User
	ph := ps.ByName("PhotoId")

	err := rt.db.DelLike(id, lk, ph)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	_, err = w.Write([]byte("unliked succesfully"))

}
