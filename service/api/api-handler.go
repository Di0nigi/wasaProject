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

	rt.router.DELETE("/userActions/:UserId/interactions/manageBan/:banned", rt.wrap(rt.unBanUser))

	rt.router.POST("/userActions/:UserId/interactions/followingActions", rt.wrap(rt.followUser))
	
	rt.router.DELETE("/userActions/:UserId/interactions/followingActions/:followers", rt.wrap(rt.unFollowUser))
	
	rt.router.POST("/userActions/:UserId/photoManager", rt.wrap(rt.postImage))

	rt.router.DELETE("/userActions/:UserId/photoManager/:ObjId", rt.wrap(rt.deleteImage))

	rt.router.POST("/userActions/:UserId/interactions/comments", rt.wrap(rt.commentPhoto))
	
	rt.router.DELETE("/userActions/:UserId/interactions/comments/:photo/:commentId", rt.wrap(rt.unCommentPhoto))

	rt.router.POST("/userActions/:UserId/interactions/likes",rt.wrap(rt.likePost))

	rt.router.DELETE("/userActions/:UserId/interactions/likes/:photo/:likeId",rt.wrap(rt.unLikePost))
	
	
	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
