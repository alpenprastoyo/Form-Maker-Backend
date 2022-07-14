package services

import (
	"net/http"

	"form-api/database"
	"form-api/models/survey"

	"github.com/gin-gonic/gin"
)

func Migrate(c *gin.Context) {

	db := database.GetDB()

	db.DropTable(&survey.Survey{})

	db.AutoMigrate(&survey.Survey{})

	c.JSON(http.StatusBadRequest, gin.H{
		"messages": "Data has been migrate",
	})
}
