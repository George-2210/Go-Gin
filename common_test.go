package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var tmpArticleList []article

// 在执行测试函数之前进行设置的函数
func TestMain(m *testing.M) {
	// 将 Gin 设置为测试模式
	gin.SetMode(gin.TestMode)

	// 运行其他测试
	os.Exit(m.Run())
}

// 在测试期间创建路由的辅助函数
func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("templates/*")
	}
	return r
}

// 处理请求并测试其响应的辅助函数
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {

	// 创建响应记录器
	w := httptest.NewRecorder()

	// 创建服务并处理上述请求
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

// 用于将主要列表存储到临时列表中以进行测试的函数
func saveLists() {
	tmpArticleList = articleList
}

// 用于从临时列表中恢复主要列表的函数
func restoreLists() {
	articleList = tmpArticleList
}
