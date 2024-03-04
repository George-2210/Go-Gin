// handlers.article_test.go

package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// 测试未经认证的用户对主页的GET请求是否返回主页以及HTTP状态码200
func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/", showIndexPage)

	// 创建请求以发送到上述路由
	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// 检查HTTP状态码是否为200
		statusOK := w.Code == http.StatusOK

		// 检查页面标题是否为"Home Page"
		// 您可以使用能够解析和处理HTML页面的库进行更详细的测试
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOK
	})
}
