package routes

import (
	"net/http"

	callbacks "github.com/computerextra/local-horst-golang/Callbacks"
	"github.com/computerextra/local-horst-golang/template/Pages/archiv"
	"github.com/gin-gonic/gin"
)

type SearchFormData struct {
	Search string `form:"search"`
}

func GetArchiveRoute(r *gin.Engine) {
	r.GET("/Archive/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "", archiv.Archiv([]callbacks.Pdf{}))
	})
	r.GET("/Archive/:search", func(c *gin.Context) {
		search := c.Param("search")

		Results, err := callbacks.SearchArchive(search)
		if err != nil {
			panic(err)
		}

		c.HTML(http.StatusOK, "", archiv.Archiv(Results))

	})

}
