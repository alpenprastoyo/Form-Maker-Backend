package routers

import (
	migrationServices "form-api/services/v1/migration"
	surveyServies "form-api/services/v1/survey"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	v1 := r.Group("/v1")
	v1.GET("/", surveyServies.Main)
	v1.GET("/surveys/", surveyServies.Surveys)
	v1.GET("/survey/:id/", surveyServies.SurveyFind)
	v1.POST("/survey/store/", surveyServies.SurveyStore)
	v1.PUT("/survey/update/", surveyServies.SurveyUpdate)
	v1.DELETE("/survey/destroy/", surveyServies.SurveyDestroy)

	v1.GET("/migration/", migrationServices.Migrate)

	return r

}
