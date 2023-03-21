package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 中间件，拦截器
func myHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("usersession", "userid_1")
		context.Next() //放行
		//context.Abort() //阻止
	}
}

func main() {
	//创建服务
	ginServer := gin.Default()
	//注册，全局使用中间件
	ginServer.Use(myHandler())
	//ginServer.Use()
	//连接数据库

	//访问地址
	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "test server"})
	})

	//加载静态目录
	//ginServer.Static()

	//记载静态页面
	ginServer.LoadHTMLGlob("templates/*")
	//响应页面给前端
	ginServer.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "doWeb.html", gin.H{
			"msg": "后台数据",
		})
	})

	//接受前端传递的参数
	//url?userid=xxx&username=zhang
	ginServer.GET("/user/info", myHandler(), func(context *gin.Context) {

		//加中间件
		usersession_1 := context.MustGet("usersession").(string)
		log.Println("========>", usersession_1)

		userid := context.Query("userid")
		username := context.Query("username")
		context.JSON(http.StatusOK, gin.H{
			"userid":   userid,
			"username": username,
		})
	})
	// /user/info/zhang
	ginServer.GET("/user/info/:userid/:username", func(context *gin.Context) {
		userid := context.Param("userid")
		username := context.Param("username")
		context.JSON(http.StatusOK, gin.H{
			"userid":   userid,
			"username": username,
		})
	})

	// 前端给后端传json
	ginServer.POST("/json", func(context *gin.Context) {
		//requset.body
		data, _ := context.GetRawData()
		var m map[string]interface{}
		_ = json.Unmarshal(data, &m)
		context.JSON(http.StatusOK, m)
	})
	ginServer.POST("user/add", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")
		context.JSON(http.StatusOK, gin.H{
			"msg":      "ok",
			"username": username,
			"password": password,
		})
	})
	//路由
	ginServer.GET("/redirect", func(context *gin.Context) {
		//重定向
		context.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})
	//404
	ginServer.NoRoute(func(context *gin.Context) {
		context.HTML(http.StatusNotFound, "error.html", nil)
	})

	//路由组
	userGroup := ginServer.Group("/user")
	{
		userGroup.GET("/add")
		userGroup.GET("/delete")
	}

	//服务器
	ginServer.Run(":8082")

}
