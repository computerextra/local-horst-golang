package api

import "github.com/gin-gonic/gin"

type Config struct {
	Router *gin.Engine
}

func (app *Config) Routes() {
	// Routes will come here
}
