package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/internal/service"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/go-programming-tour-book/blog-service/pkg/convert"
	"github.com/go-programming-tour-book/blog-service/pkg/errcode"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

// func (t Tag) Get(c *gin.Context) {

// }

// @Summary Get all tags
// @Description
// @Produce json
// @Param name query string false "tag name" maxlength (100)
// @Param state query int false "status" Enums(0,1) default(1)
// @Param page query int false "page num"
// @Param page_size query string false "page size"
// @Success 200 {object} model.TagSwagger "Sucess"
// @Failure 400 {object} errcode.Error "Request Error"
// @Failure 500 {object} errcode.Error "Server Error"
// @Router /apu/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	param := service.TagListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}

	totalRows, err := svc.CountTag(&service.CountArticleRequest{Name: param.Name, State: param.State})
	if err != nil {
		global.Logger.Errorf("svc.CountTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}
	tags, err := svc.GetTagList(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svc.GetTagList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}
	response.ToResponseList(tags, totalRows)
	return
}

// @Summary Create new Tag
// @Description
// @Produce json
// @Param name body string false "tag name" minlength(3) maxlength (100)
// @Param state body int false "status" Enums(0,1) default(1)
// @Param created_by body string true "create author" minlength(3) maxlength(100)
// @Success 200 {object} model.TagSwagger "Success"
// @Failure 400 {object} errcode.Error "Request Error"
// @Failure 500 {object} errcode.Error "Server Error"
// @Router /apu/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	param := service.CreateTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)

	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateTag(&param)

	if err != nil {
		global.Logger.Errorf("svc.CreateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return

}

// @Summary Update exist tag
// @Description
// @Produce json
// @Param id path int true "tag Id"
// @Param name body string false "tag name" minlength(3) maxlength (100)
// @Param state body int false "status" Enums(0,1) default(1)
// @Param modified_by body string true "modified author" minlength(3) maxlength(100)
// @Success 200 {array} model.TagSwagger "Sucess"
// @Failure 400 {object} errcode.Error "Request Error"
// @Failure 500 {object} errcode.Error "Server Error"
// @Router /apu/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {
	param := service.UpdateTagRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &param)

	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&param)

	if err != nil {
		global.Logger.Errorf("svc.UpdateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary Delete tag
// @Description
// @Produce json
// @Param id path int true "tag Id"
// @Success 200 {string} string "Sucess"
// @Failure 400 {object} errcode.Error "Request Error"
// @Failure 500 {object} errcode.Error "Server Error"
// @Router /apu/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {
	param := service.DeleteTagRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)

	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteTag(&param)
	if err != nil {
		global.Logger.Errorf("svc DeleteTag err : %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}
