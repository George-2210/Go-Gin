// handlers.article.go

package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// showIndexPage 函数用于显示首页
func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	// 使用要渲染的模板名称调用 render 函数
	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles}, "index.html")
}

func getArticle(c *gin.Context) {
	// 检查文章ID是否有效
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// 检查文章是否存在
		if article, err := getArticleByID(articleID); err == nil {
			// 调用 Context 的 HTML 方法来渲染模板
			c.HTML(
				// 将 HTTP 状态设置为 200 (OK)
				http.StatusOK,
				// 使用 article.html 模板
				"article.html",
				// 传递页面使用的数据
				gin.H{
					"title":   article.Title,
					"payload": article,
				},
			)

		} else {
			// 如果文章未找到，则终止并报错
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// 如果在 URL 中指定了无效的文章ID，则终止并报错
		c.AbortWithStatus(http.StatusNotFound)
	}
}

// 根据请求的 'Accept' 头部信息渲染 HTML、JSON 或 CSV 中的一种
// 如果请求头部未指定，则默认渲染 HTML，前提是模板名称已存在
func render(c *gin.Context, data gin.H, templateName string) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// 使用 JSON 响应
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// 使用 XML 响应
		c.XML(http.StatusOK, data["payload"])
	default:
		// 使用 HTML 响应
		c.HTML(http.StatusOK, templateName, data)
	}

}
