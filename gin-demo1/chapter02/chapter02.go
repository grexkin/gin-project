package chapter02

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadNews(ctx *gin.Context)  {
	ctx.HTML(200,"news/news.html",nil)
}

// 结构体渲染
type UserInfo struct {
	Id   int
	Name string
	Age  int
	Address string
}

func LoadUser(ctx *gin.Context) {
	name := "hello"
	ctx.HTML(200,"user/user.html",name)
}

func LoadUseInfoStruct(ctx *gin.Context)  {
	//自动推导
	userinfo := UserInfo{
		Id: 0,
		Name: "hello",
		Age: 20,
		Address: "Beijing",
	}
	//变量声明形式
	var userinfo1 UserInfo
	userinfo1.Id = 1
	userinfo1.Name = "zhansan"
	userinfo1.Age = 18
	userinfo1.Address = "Shanghai"
	fmt.Println(userinfo1)
	ctx.HTML(http.StatusOK,"chapter02/user_info.html",userinfo)
}

//数组渲染
func LoadArray(ctx *gin.Context)  {
	arr := [5]int{1,2,3,4,5}
	ctx.HTML(http.StatusOK,"chapter02/array.html",arr)
}

//结构体数组渲染
func ArrayStructController(ctx *gin.Context) {
	arr_struct := [3]UserInfo{
		{Id:1,Name: "hello",Age: 18,Address: "xxx"},
		{Id:2,Name: "hello1",Age: 19,Address: "xxx1"},
		{Id:3,Name: "hello2",Age: 20,Address: "xxx2"},
	}
	ctx.HTML(http.StatusOK,"chapter02/arr_struct.html",arr_struct)
}

//map数据渲染
func MapController(ctx *gin.Context)  {
	map_data := map[string]string{"name":"map_data","age":"22"}
	map_data1 := map[string]int{"age":11,"age1":12}
	map_data2 := map[string]interface{}{"map_data":map_data,"map_data1":map_data1}
	ctx.HTML(http.StatusOK,"chapter02/map.html",map_data2)
}

//map结构体渲染
func MapStructController(ctx *gin.Context) {
	map_struct := map[string]UserInfo{
		"data01":{Id: 0,Name: "zhangsanfeng",Age: 22,Address: "武当山"},
		"data02":{Id: 1,Name: "miejueshitai",Age: 88,Address: "峨眉山"},
		"data03":{Id: 2,Name: "zhangwuji",Age: 22,Address: "光明顶"},
	}
	ctx.HTML(http.StatusOK,"chapter02/map_struct.html",map_struct)
}

//slice渲染
func SliceController(ctx *gin.Context) {
	slice_data := []int{1,2,3,4,5,6,7}
	ctx.HTML(http.StatusOK,"chapter02/slice.html",slice_data)
}

//前端路径上直接加参数
func Param1(ctx *gin.Context)  {
	id := ctx.Param("id")
	ctx.String(http.StatusOK,"hello,%s",id)
}

func Param2(ctx *gin.Context)  {
	id := ctx.Param("id")
	ctx.String(http.StatusOK,"hello,%s",id)
}

func GetQueryData(ctx *gin.Context)  {
	id := ctx.Query("id")
	//http://127.0.0.1:9000/query?id=123&name=hhh
	name := ctx.DefaultQuery("name","zhangsan")
	fmt.Println(name)
	ctx.String(http.StatusOK,"hello %s",id)
}

func GetQueryArrayData(ctx *gin.Context)  {
	//http://127.0.0.1:9000/queryarr?ids=1,2,3,4
	ids := ctx.QueryArray("ids")
	ctx.String(http.StatusOK,"hello %s",ids)
}

func GetQueryMapData(ctx *gin.Context)  {
	//http://127.0.0.1:9000/querymap?user[name]=hello&user[age]=18
	user := ctx.QueryMap("user")
	ctx.String(http.StatusOK,"hello,%s",user)
}

//POST请求
func ToUserAdd(ctx *gin.Context) {
	ctx.HTML(http.StatusOK,"chapter02/user_add.html",nil)
}

func DoUserAdd(ctx *gin.Context)  {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	fmt.Println(username,password)
	ctx.String(http.StatusOK,"user: %s 添加成功",username)
}

func ToUserAdd2(ctx *gin.Context) {
	ctx.HTML(http.StatusOK,"chapter02/user_add2.html",nil)
}

func DoUserAdd2(ctx *gin.Context)  {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	age := ctx.DefaultPostForm("age","18")
	fmt.Println(username,password,age)
	hobbys := ctx.PostFormArray("hobby")
	fmt.Println(hobbys)
	user := ctx.PostFormMap("user")
	fmt.Println(user)
	ctx.String(http.StatusOK,"user: %s 添加成功",username)
}

//ajax 信息提交
func ToUserAdd3(ctx *gin.Context) {
	ctx.HTML(http.StatusOK,"chapter02/user_add3.html",nil)
}

func DoUserAdd3(ctx *gin.Context)  {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	age := ctx.DefaultPostForm("age","18")
	fmt.Println(username,password,age)
	//map_data := map[string]interface{}{
	//	"code": 200,
	//	"msg": "success",
	//}
	//ctx.JSON(http.StatusOK,map_data)
	ctx.JSON(http.StatusOK,gin.H{"code":200,"msg":"成功"})
}

//参数绑定
func ToUserAdd4(ctx *gin.Context) {
	ctx.HTML(http.StatusOK,"chapter02/user_add4.html",nil)
}

type Users struct {
	UserName string `form:"username" json:"username"`
	PassWord string `form:"password" json:"password"`
	Age      int    `form:"age" json:"age"`
}

func DoUserAdd4(ctx *gin.Context)  {
	var user_info Users
	err := ctx.ShouldBind(&user_info)
	fmt.Println(err)
	fmt.Println(user_info)
	ctx.String(http.StatusOK,"ok")
}
