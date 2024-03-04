// models.article_test.go

package main

import "testing"

// 测试获取所有文章的函数
func TestGetAllArticles(t *testing.T) {
	// 获取所有文章
	alist := getAllArticles()

	// 检查返回的文章列表长度是否与全局变量中的长度相同
	if len(alist) != len(articleList) {
		t.Fail()
	}

	// 检查每个成员是否相同
	for i, v := range alist {
		if v.Content != articleList[i].Content ||
			v.ID != articleList[i].ID ||
			v.Title != articleList[i].Title {

			t.Fail()
			break
		}
	}
}
