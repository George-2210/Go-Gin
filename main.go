// main.go

package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	// 将 Gin 提供的默认路由设置为 router
	router = gin.Default()

	// 在启动时处理模板，这样它们就不需要重新从磁盘加载。
	// 这样可以非常快速地提供 HTML 页面。
	router.LoadHTMLGlob("templates/*")

	// 定义首页的路由，并显示 index.html 模板
	// 首先，我们将使用内联路由处理程序。稍后，我们将创建独立的函数作为路由处理程序使用。
	router.GET("/", showIndexPage)

	router.GET("/article/view/:article_id", getArticle)

	// 启动应用程序
	router.Run()

}
