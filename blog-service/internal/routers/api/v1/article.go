package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/go-programming-tour-book/blog-service/pkg/errcode"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

// @Summary Get article
// @Description
// @Produce json
// @Param id path int true "article Id"
// @Success 200 {object} model.ArticleSwagger "Sucess"
// @Failure 400 {object} errcode.Error "Request Error"
// @Failure 500 {object} errcode.Error "Server Error"
// @Router /apu/v1/articles/{id} [get]
func (a Article) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}

// @Summary Create article
// @Description
// @Produce json
// @Param title body string true "article title" minlength(3) maxlength (100)
// @Param Desc body string true "article Describtion" minlength(3) maxlength (255)
// @Param Content body string true "article content"
// @Param ConverImageUrl body string false "article Cover Image Urls"
// @Param state body int false "state" Enums(0, 1) default(1)
// @Success 200 {object} model.ArticleSwagger "Sucess"
// @Failure 400 {object} errcode.Error "Request Error"
// @Failure 500 {object} errcode.Error "Server Error"
// @Router /apu/v1/articles [post]
func (a Article) Create(c *gin.Context) {}

// @Summary Get all Articles
// @Description
// @Produce json
// @Param title query string false "artcile title" maxlength (100)
// @Param state query int false "status" Enums(0,1) default(1)
// @Param page query int false "page num"
// @Param page_size query string false "page size"
// @Success 200 {object} model.ArticleSwagger "Sucess"
// @Failure 400 {object} errcode.Error "Request Error"
// @Failure 500 {object} errcode.Error "Server Error"
// @Router /apu/v1/articles [get]
func (a Article) List(c *gin.Context) {}

// @Summary Update exist tag
// @Description
// @Produce json
// @Param id path int true "article Id"
// @Param title body string false "article title" minlength(3) maxlength (100)
// @Param state body int false "status" Enums(0,1) default(1)
// @Param modified_by body string true "modified author" minlength(3) maxlength(100)
// @Success 200 {array} model.ArticleSwagger "Sucess"
// @Failure 400 {object} errcode.Error "Request Error"
// @Failure 500 {object} errcode.Error "Server Error"
// @Router /apu/v1/articles/{id} [put]
func (a Article) Update(c *gin.Context) {}

// @Summary Delete tag
// @Description
// @Produce json
// @Param id path int true "article Id"
// @Success 200 {string} string "Sucess"
// @Failure 400 {object} errcode.Error "Request Error"
// @Failure 500 {object} errcode.Error "Server Error"
// @Router /apu/v1/articles/{id} [delete]
func (a Article) Delete(c *gin.Context) {}
