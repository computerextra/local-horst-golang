package main

import (
	"net/http"

	callbacks "github.com/computerextra/local-horst-golang/Callbacks"
	"github.com/computerextra/local-horst-golang/gintemplrenderer"
	pages "github.com/computerextra/local-horst-golang/template/Pages"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

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
		c.HTML(http.StatusOK, "", pages.Einkauf(Einkauf))
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
