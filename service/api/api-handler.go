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

	rt.router.GET("/userActions/:UserId/interactions/Profile/:AccountId", rt.wrap(rt.getUser))

	rt.router.POST("/userActions/:UserId/interactions/manageBan", rt.wrap(rt.banUser))
	rt.router.PUT("/userActions/:UserId/interactions/manageBan", rt.wrap(rt.unBanUser))

	rt.router.POST("/userActions/:UserId/interactions/followingActions", rt.wrap(rt.followUser))
	rt.router.PUT("/userActions/:UserId/interactions/followingActions", rt.wrap(rt.unFollowUser))
	
	
	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
