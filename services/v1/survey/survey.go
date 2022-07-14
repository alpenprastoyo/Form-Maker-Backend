package services

import (
	"fmt"
	"net/http"

	"form-api/database"
	"form-api/models/survey"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Main(c *gin.Context) {

}

func Surveys(c *gin.Context) {

	db := database.GetDB()

	var surveys []survey.Survey

	err := db.Debug().Find(&surveys).Error

	if err != nil {
		fmt.Println("Error finding survey record")
	}

	for _, b := range surveys {
		fmt.Println("Title :", b.NameSurvey)
	}

	c.JSON(http.StatusOK, surveys)
}

func SurveyFind(c *gin.Context) {

}

type surveyStore struct {
	NameSurvey  string `json:"name_survey" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func SurveyStore(c *gin.Context) {

	var surveyStore surveyStore

	db := database.GetDB()

	err := c.ShouldBindJSON(&surveyStore)

	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})

		return

	}

	survey := survey.Survey{}
	survey.NameSurvey = surveyStore.NameSurvey
	survey.Description = surveyStore.Description

	err = db.Create(&survey).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, survey)

}

type surveyUpdate struct {
	NameSurvey  string `json:"name_survey" binding:"required"`
	Description string `json:"description" binding:"required"`
	Id          string `json:"id" binding:"required,number"`
}

func SurveyUpdate(c *gin.Context) {
	var surveyUpdate surveyUpdate

	db := database.GetDB()

	err := c.ShouldBindJSON(&surveyUpdate)

	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})

		return

	}

	survey := survey.Survey{}
	err = db.Debug().First(&survey, surveyUpdate.Id).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": "Error finding survey record",
		})
	}

	survey.NameSurvey = surveyUpdate.NameSurvey
	survey.Description = surveyUpdate.Description

	err = db.Save(&survey).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": "Error update survey record",
		})
	}

	c.JSON(http.StatusOK, survey)
}

type surveyDestroy struct {
	Id string `json:"id" binding:"required,number"`
}

func SurveyDestroy(c *gin.Context) {
	var surveyDestroy surveyDestroy

	db := database.GetDB()

	err := c.ShouldBindJSON(&surveyDestroy)

	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})

		return

	}

	survey := survey.Survey{}
	err = db.Debug().First(&survey, surveyDestroy.Id).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": "Error finding survey record",
		})
	}

	err = db.Delete(&survey).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": "Error destroy survey record",
		})
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"mesasages": "destroy survey has been successfull",
	})
}
