package chapter01

import "github.com/gin-gonic/gin"

func HelloGin(ctx *gin.Context)  {
	ctx.String(200,"hello gin")
}

func LoadIndex(ctx *gin.Context)  {
	name := "abcd"
	ctx.HTML(200,"index/index.html",name)
}
