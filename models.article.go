// models.article.go

package main

import "errors"

// 文章结构体定义
type article struct {
	ID      int    `json:"id"`      // 文章ID
	Title   string `json:"title"`   // 文章标题
	Content string `json:"content"` // 文章内容
}

// 在这个演示中，我们将文章列表存储在内存中。
// 在实际应用中，这个列表很可能从数据库或静态文件中获取。
var articleList = []article{
	article{ID: 1, Title: "文章 1", Content: "文章 1 正文"},
	article{ID: 2, Title: "文章 2", Content: "文章 2 正文"},
}

// 返回所有文章的列表
func getAllArticles() []article {
	return articleList
}

// 返回要显示的文章的 ID
func getArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}
