package dao

import "github.com/go-programming-tour-book/blog-service/internal/model"

func (d *Dao) GetAuth(appKey, AppSecret string) (model.Auth, error) {
	auth := model.Auth{AppKey: appKey, AppSecret: AppSecret}
	return auth.Get(d.engine)
}
