package routes

import (
	"fmt"
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
	r.POST("/Archive/Search", func(c *gin.Context) {
		var bindData SearchFormData
		if err := c.ShouldBind(&bindData); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
			return
		}
		Results, err := callbacks.SearchArchive(bindData.Search)
		if err != nil {
			panic(err)
		}
		var Response string = ""

		for _, x := range Results {
			Response = fmt.Sprintf("%s<li><a href=\"\">%s</a></li>", Response, x.Title)
		}

		c.String(http.StatusOK, Response)
	})

}
