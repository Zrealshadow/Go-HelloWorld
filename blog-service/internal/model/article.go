package model

import "github.com/go-programming-tour-book/blog-service/pkg/app"

type Article struct {
	*Model
	Title          string `json:"title"`
	Desc           string `json:"desc"`
	Content        string `json:"content"`
	ConverImageUrl string `json:"cover_image_url"`
	State          uint8  `json:"state"`
}

func (a Article) TableName() string {
	return "BlogArticle"
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}
