package main

import (
	"log"
	"net/http"

	callbacks "github.com/computerextra/local-horst-golang/Callbacks"
	"github.com/computerextra/local-horst-golang/gintemplrenderer"
	pages "github.com/computerextra/local-horst-golang/template/Pages"
	einkauf "github.com/computerextra/local-horst-golang/template/Pages/einkauf"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const dst = "/upload/"

func setupRouter() *gin.Engine {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	r := gin.Default()

	// Tailwind CSS File
	r.StaticFile("/style.css", "./output.css")
	// Assets
	r.Static("/assets", "./assets")

	// Callbacks
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	ginHtmlRenderer := r.HTMLRender
	r.HTMLRender = &gintemplrenderer.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}

	// Disable trusted proxy warning.
	r.SetTrustedProxies(nil)

	// Sites
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "", pages.Home())
	})

	// Get Einkauf
	Einkauf, err := callbacks.GetEinkauf()
	if err != nil {
		panic(err)
	}
	r.GET("/Liste", func(c *gin.Context) {
		c.HTML(http.StatusOK, "", einkauf.Einkauf(Einkauf))
	})
	r.GET("/Eingabe", func(c *gin.Context) {
		Mitarbeiter, err := callbacks.GetNames()
		if err != nil {
			panic(err)
		}
		c.HTML(http.StatusOK, "", einkauf.Auswahl(Mitarbeiter))
	})
	r.GET("/Eingabe/:id", func(c *gin.Context) {
		id := c.Param("id")
		Mitarbeiter, err := callbacks.GetOne(id)
		var e callbacks.Einkauf
		for x := range Einkauf {
			if Einkauf[x].MitarbeiterId.String == Mitarbeiter.Id.String {
				e = Einkauf[x]
			}
		}
		if err != nil {
			panic(err)
		}
		c.HTML(http.StatusOK, "", einkauf.Eingabe(Mitarbeiter, e))
	})
	// Verarbeite FormData von Einkauf
	r.POST("/Eingabe", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]
		for _, file := range files {
			log.Println(file.Filename)

			// Upload the file to specific dst.
			c.SaveUploadedFile(file, dst)
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "saved",
		})
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Serve on 0.0.0.0:8080
	r.Run(":8080")
}
