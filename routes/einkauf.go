package routes

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	callbacks "github.com/computerextra/local-horst-golang/Callbacks"
	"github.com/computerextra/local-horst-golang/template/Pages/einkauf"
	"github.com/gin-gonic/gin"
)

const uploadFolder = "/upload/"

func getRoot() string {
	rootDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Cannot fond root")
	}

	return rootDir
}

func GetEinkaufRoutes(r *gin.Engine) {
	dst := getRoot() + uploadFolder

	r.GET("/Liste", func(c *gin.Context) {
		// Get Einkauf
		Einkauf, err := callbacks.GetEinkäufe()
		if err != nil {
			panic(err)
		}
		c.HTML(http.StatusOK, "", einkauf.Einkauf(Einkauf))
	})

	r.GET("/Auswahl", func(c *gin.Context) {
		Mitarbeiter, err := callbacks.GetNames()
		if err != nil {
			panic(err)
		}
		c.HTML(http.StatusOK, "", einkauf.Auswahl(Mitarbeiter))
	})

	r.GET("/Eingabe/:id", func(c *gin.Context) {
		id := c.Param("id")
		Mitarbeiter, err := callbacks.GetOne(id)
		if err != nil {
			panic(err)
		}
		Einkauf, err := callbacks.GetEinkauf(Mitarbeiter.Id.String)
		if err != nil {
			panic(err)
		}

		c.HTML(http.StatusOK, "", einkauf.Eingabe(Mitarbeiter, Einkauf))
	})

	// delete Image from Database
	r.POST("/deleteimage/:id/:image", func(c *gin.Context) {
		MitarbeiterId := c.Param("id")
		image := c.Param("image")

		// get Imagepath
		imagePath, err := callbacks.GetImagePath(image, MitarbeiterId)
		if err != nil {
			panic(err)
		}

		if err := callbacks.DeleteImage(image, MitarbeiterId); err != nil {
			panic(err)
		}

		os.Remove(dst + filepath.Base(imagePath))

		c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/Eingabe/%s", MitarbeiterId))
	})

	type BindFormData struct {
		Pfand string                `form:"Pfand"`
		Geld  string                `form:"Geld"`
		Dinge string                `form:"Dinge"`
		Bild1 *multipart.FileHeader `form:"Bild1"`
		Bild2 *multipart.FileHeader `form:"Bild2"`
		Bild3 *multipart.FileHeader `form:"Bild3"`
		Abo   string                `form:"Abo"`
	}

	// Verarbeite FormData von Einkauf
	r.POST("/EinkaufInput/:id", func(c *gin.Context) {
		MitarbeiterId := c.Param("id")
		var bindFormData BindFormData

		if err := c.ShouldBind(&bindFormData); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
			return
		}

		alterEinkauf, err := callbacks.GetEinkauf(MitarbeiterId)
		if err != nil {
			panic(err)
		}

		var einkauf callbacks.PostEinkauf

		if alterEinkauf.Bild1.Valid && len(alterEinkauf.Bild1.String) > 0 {
			einkauf.Bild1 = alterEinkauf.Bild1.String
		} else {
			if len(bindFormData.Bild1.Filename) > 0 {
				// Upload the file to specific dst.
				if err := c.SaveUploadedFile(bindFormData.Bild1, dst+filepath.Base(bindFormData.Bild1.Filename)); err != nil {
					fmt.Printf("Fehler: %s", err)
				}
				length := len(strings.Split(bindFormData.Bild1.Filename, "."))
				format := strings.Split(bindFormData.Bild1.Filename, ".")[length-1]
				newName := fmt.Sprintf("%s-Bild1.%s", MitarbeiterId, format)
				os.Rename(dst+filepath.Base(bindFormData.Bild1.Filename), dst+filepath.Base(newName))
				einkauf.Bild1 = uploadFolder + filepath.Base(newName)

			} else {
				einkauf.Bild1 = ""
			}
		}
		if alterEinkauf.Bild2.Valid && len(alterEinkauf.Bild2.String) > 0 {
			einkauf.Bild2 = alterEinkauf.Bild2.String
		} else {
			if len(bindFormData.Bild2.Filename) > 0 {
				// Upload the file to specific dst.
				if err := c.SaveUploadedFile(bindFormData.Bild2, dst+filepath.Base(bindFormData.Bild2.Filename)); err != nil {
					fmt.Printf("Fehler: %s", err)
				}
				length := len(strings.Split(bindFormData.Bild2.Filename, "."))
				format := strings.Split(bindFormData.Bild2.Filename, ".")[length-1]
				newName := fmt.Sprintf("%s-Bild2.%s", MitarbeiterId, format)
				os.Rename(dst+filepath.Base(bindFormData.Bild2.Filename), dst+filepath.Base(newName))
				einkauf.Bild2 = uploadFolder + filepath.Base(newName)
			} else {
				einkauf.Bild2 = ""
			}
		}
		if alterEinkauf.Bild3.Valid && len(alterEinkauf.Bild3.String) > 0 {
			einkauf.Bild3 = alterEinkauf.Bild3.String
		} else {
			if len(bindFormData.Bild3.Filename) > 0 {
				// Upload the file to specific dst.
				if err := c.SaveUploadedFile(bindFormData.Bild3, dst+filepath.Base(bindFormData.Bild3.Filename)); err != nil {
					fmt.Printf("Fehler: %s", err)
				}
				length := len(strings.Split(bindFormData.Bild3.Filename, "."))
				format := strings.Split(bindFormData.Bild3.Filename, ".")[length-1]
				newName := fmt.Sprintf("%s-Bild3.%s", MitarbeiterId, format)
				os.Rename(dst+filepath.Base(bindFormData.Bild3.Filename), dst+filepath.Base(newName))
				einkauf.Bild1 = uploadFolder + filepath.Base(newName)
			} else {
				einkauf.Bild3 = ""
			}
		}
		if bindFormData.Abo == "on" {
			einkauf.Abonniert = true
		} else {
			einkauf.Abonniert = false
		}
		einkauf.Dinge = bindFormData.Dinge
		einkauf.Geld = bindFormData.Geld
		einkauf.Pfand = bindFormData.Pfand
		einkauf.MitarbeiterId = MitarbeiterId
		if strings.ToLower(bindFormData.Geld) == "paypal" {
			einkauf.Paypal = true
		} else {
			einkauf.Paypal = false
		}

		if err := callbacks.SaveEinkauf(einkauf); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
			return
		}

		c.Redirect(http.StatusMovedPermanently, "/Liste")

	})
}
