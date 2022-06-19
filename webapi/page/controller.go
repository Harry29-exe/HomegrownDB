package page

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func Register(router *gin.Engine) {
	router.POST("/api/page/tuple", interpretTuple)
}

func interpretTuple(ctx *gin.Context) {
	data, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.String(404, "bad body")
	}

	println(data)
	ctx.String(200, string(data))
}
