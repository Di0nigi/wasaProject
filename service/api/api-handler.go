package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	//rt.router.GET("/", rt.getHelloWorld)
	//rt.router.GET("/context", rt.wrap(rt.getContextReply))
	rt.router.POST("/session",rt.wrap(rt.logIn))
	
	rt.router.POST("/userActions/:UserId", rt.wrap(rt.setUsername))
	rt.router.GET("/userActions/:UserId", rt.wrap(rt.getStream))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
