package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/go-programming-tour-book/blog-service/internal/routers/api/v1"
)

// 路由管理
func NewRouter() *gin.Engine {
	r := gin.New()
	// Use()函数将全局中间件附加到路由器，通过Use()附加的中间件将包含在每个请求的处理程序链中，包括404、405、静态文件等。
	r.Use(gin.Logger())   // Logger实例化一个Logger中间件，它将日志写入gin.DefaultWriter。默认情况下，gin.DefaultWriter = os.Stdout。
	r.Use(gin.Recovery()) // Recovery返回一个中间件，用于从任何panic中恢复，并在出现panic时写入500。

	article := v1.NewArticle()
	tag := v1.NewTag()

	// Group创建一个新的路由器组。您应该添加所有具有公共中间件或相同路径前缀的路由。
	apiv1 := r.Group("/api/v1")

	// 注册路由
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}
	return r
}
