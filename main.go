package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/computerextra/local-horst-golang/gintemplrenderer"
	"github.com/computerextra/local-horst-golang/routes"
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

	mode := os.Getenv("MODE")
	if mode == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		fmt.Println("MODE not set in .env, for Production Build use PROD as Value")
		gin.SetMode(gin.DebugMode)
	}

	// Tailwind CSS File
	r.StaticFile("/style.css", "./output.css")
	// Assets
	r.Static("/assets", "./assets")
	// Uploaded Images
	r.Static("/upload", "./upload")

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

	routes.GetEinkaufRoutes(r)
	routes.GetGeburtstagRoute(r)
	routes.GetTelefonlistenRoutes(r)

	return r
}

func main() {
	r := setupRouter()

	var port = ""
	port = os.Getenv("PORT")
	if len(port) < 1 {
		fmt.Println("No Port in .env specified. Using Port 8080 in Dev Mode")
		port = "8080"
	}
	r.Run(fmt.Sprintf(":%s", port))
}
