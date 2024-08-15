package routes

import (
	"net/http"
	"time"

	callbacks "github.com/computerextra/local-horst-golang/Callbacks"
	"github.com/computerextra/local-horst-golang/template/Pages/geburtstag"
	"github.com/gin-gonic/gin"
)

func SortDemArray(items []callbacks.GebRes, zustand string) []callbacks.GebRes {
	var alled []callbacks.GebRes
	heute := time.Now()

	for _, item := range items {
		if zustand == "Vergangen" {
			if item.Geburtstag.Before(heute) {
				alled = append(alled, item)
			}
		}

		if zustand == "Heute" {
			if item.Geburtstag.Equal(heute) {
				alled = append(alled, item)
			}

		}

		if zustand == "Zukunft" {
			if item.Geburtstag.After(heute) {
				alled = append(alled, item)
			}
		}
	}

	return alled
}

func GetGeburtstagRoute(r *gin.Engine) {

	r.GET("/Geburtstag", func(c *gin.Context) {
		// Get Einkauf
		Geburtstage, err := callbacks.GetGeburtstage()
		if err != nil {
			panic(err)
		}
		Vergangen := SortDemArray(Geburtstage, "Vergangen")
		Heute := SortDemArray(Geburtstage, "Heute")
		Zukunft := SortDemArray(Geburtstage, "Zukunft")

		c.HTML(http.StatusOK, "", geburtstag.Geburtstag(Vergangen, Heute, Zukunft))
	})
}
