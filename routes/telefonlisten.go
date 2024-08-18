package routes

import (
	"github.com/computerextra/local-horst-golang/routes/telefonlisten"
	"github.com/gin-gonic/gin"
)

func GetTelefonlistenRoutes(r *gin.Engine) {
	telefonlisten.GetMitarbeiterListe(r)
}
