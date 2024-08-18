package telefonlisten

import (
	"fmt"
	"net/http"
	"strings"

	callbacks "github.com/computerextra/local-horst-golang/Callbacks"
	"github.com/computerextra/local-horst-golang/template/Pages/telefonlistenpages"
	"github.com/gin-gonic/gin"
)

func GetMitarbeiterListe(r *gin.Engine) {
	r.GET("/Telefonlisten/Mitarbeiter", func(c *gin.Context) {
		Mitarbeiter, err := callbacks.GetMitarbeiter()
		if err != nil {
			panic(err)
		}
		c.HTML(http.StatusOK, "", telefonlistenpages.MitarbeiterListe(Mitarbeiter))
	})

	type BindFormData struct {
		Search string `form:"name"`
	}
	r.POST("/mitarbeiter-search", func(c *gin.Context) {
		var bindData BindFormData
		if err := c.ShouldBind(&bindData); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
			return
		}
		var SortedMitarbeiter []callbacks.Mitarbeiter
		Mitarbeiter, err := callbacks.GetMitarbeiter()
		if err != nil {
			panic(err)
		}
		for _, ma := range Mitarbeiter {
			if strings.Contains(ma.Name.String, bindData.Search) {
				SortedMitarbeiter = append(SortedMitarbeiter, ma)
			}
		}
		c.JSON(http.StatusOK, SortedMitarbeiter)
	})

}
