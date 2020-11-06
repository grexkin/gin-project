package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-project/gin-demo1/chapter01"
	"github.com/gin-project/gin-demo1/chapter02"
)

func main() {
	router := gin.Default()
	//router := gin.New()
	router.GET("/", func(context *gin.Context) {
		context.String(200,"hello world")
	})
	router.GET("/hello",chapter01.HelloGin)
	router.LoadHTMLGlob("template/**/*")     //正则匹配
	//router.LoadHTMLFiles("template/index.html","template/news.html")   //精确指定html文件
	router.Static("/static", "static")
	//router.StaticFS("/static",http.Dir("static"))  //不推荐使用
	router.GET("/index",chapter01.LoadIndex)
	router.GET("/news",chapter02.LoadNews)
	router.GET("/user",chapter02.LoadUser)
	router.GET("/users_truct",chapter02.LoadUseInfoStruct)  //结构体渲染
	router.GET("/array",chapter02.LoadArray)    //数组渲染
	router.GET("/array_struct",chapter02.ArrayStructController)    //结构体数组渲染
	router.GET("/map",chapter02.MapController)    //map渲染
	router.GET("/map_struct",chapter02.MapStructController)    //map struct渲染
	router.GET("/slice",chapter02.SliceController)    //slice渲染
	//前端参数传入
	router.GET("/param1/:id",chapter02.Param1)  //http://127.0.0.1:9000/param/123
	router.GET("/param2/*id",chapter02.Param2)
	router.GET("/query",chapter02.GetQueryData)
	router.GET("/queryarr",chapter02.GetQueryArrayData)
	router.GET("/querymap",chapter02.GetQueryMapData)
	//表单数据的提交
	router.GET("/to_user_add",chapter02.ToUserAdd)
	router.POST("/do_user_add",chapter02.DoUserAdd)
	//表单数据提交2
	router.GET("/to_user_add2",chapter02.ToUserAdd2)
	router.POST("/do_user_add2",chapter02.DoUserAdd2)
	//ajax 数据提交
	router.GET("/to_user_add3",chapter02.ToUserAdd3)
	router.POST("/do_user_add3",chapter02.DoUserAdd3)
	//form 数据提交3
	router.GET("/to_user_add4",chapter02.ToUserAdd4)
	router.POST("/do_user_add4",chapter02.DoUserAdd4)
	router.Run(":9000")
}
