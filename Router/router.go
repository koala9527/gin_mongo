package Router

import (
	"gin_mongodb/Controllers"
	"gin_mongodb/Middlewares"
	"gin_mongodb/global"
	"log"
	"runtime/debug"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	router.Use(Middlewares.Cors())
	router.NoRoute(HandleNotFound)
	router.NoMethod(HandleNotFound)
	router.Use(Recover)
	// 使用 session(cookie-based)
	// router.Use(sessions.Sessions("myyyyysession", Sessions.Store))
	v1 := router.Group("v1")
	{
		v1.POST("/testinsert", Controllers.TestInsert)
		v1.POST("/log", Controllers.Log)
	}

	router.Run(":8080")
}

//HandleNotFound 404
func HandleNotFound(c *gin.Context) {
	global.NewResult(c).Error(404, "资源未找到", nil)
	return
}

//Recover 500
func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			global.NewResult(c).Error(500, "服务器内部错误", nil)
		}
	}()
	//继续后续接口调用
	c.Next()
}