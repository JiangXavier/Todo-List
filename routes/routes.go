package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"todo_list/api"
)


func NewRouter() *gin.Engine{
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("mysession" , store))
	v1 := r.Group("api/v1")
	{
		//用户操作
		v1.POST("user/register" , api.UserRegister)
		v1.POST("user/login" , api.UserLogin)
	}
	return r
}