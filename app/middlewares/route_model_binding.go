package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"pheasant-api/app/models"
	"reflect"
)

type modelMapping struct {
	model  reflect.Type
	column string
}

var routeModelMapping = map[string]modelMapping{
	"entity": makeModelMapping(models.Entity{}, "uuid"),
}

func makeModelMapping(model interface{}, column string) modelMapping {
	return modelMapping{
		model:  reflect.TypeOf(model),
		column: column,
	}
}

func RouteModelBinding() gin.HandlerFunc {
	return func(c *gin.Context) {
		routeParamValue := ""
		routeParamKey := ""
		var routeParamModel modelMapping

		for routeParam, model := range routeModelMapping {
			routeParamModel = model
			routeParamKey = routeParam
			routeParamValue = c.Param(routeParam)

			if routeParamValue != "" {
				break
			}
		}

		if routeParamValue != "" {
			retrieveModel(c, routeParamModel, routeParamValue, routeParamKey)
		} else {
			c.Next()
		}
	}
}

func retrieveModel(c *gin.Context, modelMapInfo modelMapping, routeValue string, routeParamKey string) {
	model := reflect.New(modelMapInfo.model).Interface()
	queryResult := models.DB.Where(modelMapInfo.column+" = ?", routeValue).First(&model)
	if queryResult.Error != nil {
		log.Print(queryResult.Error)
		abortNotFound(c)

		return
	}

	c.Set(routeParamKey, model)
	c.Next()
}

func abortNotFound(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"message": "Resource not found",
	})
}
