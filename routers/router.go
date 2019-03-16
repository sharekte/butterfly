package routers

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	//user group

	//admin group

	return r

}
