package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/zoowen/latbeckend/model"
)

func main() {
	// init data store
	store := model.NewArticleStoreInmemory()

	e := echo.New()
	e.GET("/articles", func(c echo.Context) error {
		articles := store.ArticleMap
		return c.JSON(http.StatusOK, articles)
	})
	e.GET("/articles/:id", func(c echo.Context) error {
		id := c.Param("id")
		idconv, _ := strconv.Atoi(id)
		articles := store.ArticleMap
		return c.JSON(http.StatusOK, articles[idconv-1])
	})

	e.POST("/articles", func(c echo.Context) error {
		title := c.FormValue("title")
		body := c.FormValue("body")

		// create article instance
		article, _ := model.CreateArticle(title, body)

		store.Save(article)

		return c.JSON(http.StatusOK, article)
	})

	e.PUT("/articles/:id", func(c echo.Context) error {
		title := c.FormValue("title")
		body := c.FormValue("body")
		id := c.Param("id")
		idconv, _ := strconv.Atoi(id)
		var articles = store.ArticleMap
		articles[idconv-1] = model.Article{ID: idconv, Title: title, Body: body}
		return c.JSON(http.StatusOK, store.ArticleMap)
	})
	e.DELETE("/articles/:id", func(c echo.Context) error {
		id := c.Param("id")
		idconv, _ := strconv.Atoi(id)
		store.ArticleMap = append(store.ArticleMap[:idconv-1], store.ArticleMap[idconv:]...)
		return c.JSON(http.StatusOK, store.ArticleMap)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
