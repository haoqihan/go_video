package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video_server/api/session"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 检查session
	validateUserSession(r)
	m.r.ServeHTTP(w, r)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	// 注册用户
	router.POST("/user", CreateUser)
	// 登录
	router.POST("/user/:username", Login)
	// 获取当前用户信息
	router.GET("/user/:username", GetUserInfo)
	// 添加视频
	router.POST("/user/:username/videos", AddNewVideo)
	// 查看所有视频
	router.GET("/user/:username/videos", ListAllVideos)
	// 删除视频
	router.DELETE("/user/:username/videos/:vid-id", DeleteVideo)
	// 评论内容
	router.POST("/videos/:vid-id/comments", PostComment)
	// 获取评论
	router.GET("/videos/:vid-id/comments", ShowComments)
	return router
}

func Prepare() {
	session.LoadSessionsFromDB()
}

func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", mh)

}
