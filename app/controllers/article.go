package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/zanjs/y-mugg-v3/app/models"
)

// ArticlesController is
type ArticlesController struct {
	Controller
}

// GetAll is get all articles
func (ctl ArticlesController) GetAll(c echo.Context) error {

	var (
		articles    models.ArticlePage
		queryparams models.QueryParams
		err         error
	)
	queryparams = ctl.GetQueryParams(c)

	articles, err = models.GetArticles(queryparams)
	if err != nil {
		return ctl.ResponseError(c, http.StatusForbidden, err.Error())
	}

	return ctl.ResponseSuccess(c, articles)
}

// Get is get one article
func (ctl ArticlesController) Get(c echo.Context) error {
	var (
		article    models.Article
		pathparams models.PathParams
		err        error
	)
	pathparams = ctl.GetPathParam(c)
	article, err = models.GetArticleById(pathparams.ID)
	if err != nil {
		return ctl.ResponseError(c, http.StatusForbidden, err.Error())
	}
	return ctl.ResponseSuccess(c, article)
}

// Create is create article
func (ctl ArticlesController) Create(c echo.Context) error {
	user := ctl.GetUser(c)
	article := new(models.Article)
	article.UserID = user.ID
	article.Title = c.FormValue("title")
	article.Content = c.FormValue("content")

	err := models.CreateArticle(article)

	if err != nil {
		return ctl.ResponseError(c, http.StatusUnprocessableEntity, err.Error())
	}

	return ctl.ResponseSuccess(c, article)
}

// Update is update article
func (ctl ArticlesController) Update(c echo.Context) error {
	// Parse the content
	article := new(models.Article)

	article.Title = c.FormValue("title")
	article.Content = c.FormValue("content")

	var (
		pathParams models.PathParams
	)
	pathParams = ctl.GetPathParam(c)
	m, err := models.GetArticleById(pathParams.ID)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}
	// update article data
	err = m.UpdateArticle(article)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	return c.JSON(http.StatusOK, m)
}

// Delete is delete article
func (ctl ArticlesController) Delete(c echo.Context) error {
	var err error

	// get the param id
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	m, err := models.GetArticleById(id)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	err = m.DeleteArticle()
	return c.JSON(http.StatusNoContent, err)
}
