package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"pheasant-api/app/models"
	"reflect"
	"strings"
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
		routePath := c.FullPath()
		if !strings.Contains(routePath, ":") {
			c.Next()
			return
		}

		params := getParams(routePath)
		for _, routeParam := range params {
			_, exists := routeModelMapping[routeParam]
			if !exists {
				continue
			}

			routeParamModel := routeModelMapping[routeParam]
			routeParamValue := c.Param(routeParam)
			retrieveModel(c, routeParamModel, routeParamValue, routeParam)
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

func getParams(routePath string) []string {
	paramParts := strings.Split(routePath, ":")
	params := []string{}

	for index, paramPart := range paramParts {
		if index == 0 {
			continue
		}

		paramName := strings.Split(paramPart, "/")[0]

		params = append(params, paramName)
	}

	return params
}

func abortNotFound(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"message": "Resource not found",
	})
}
