package controllers

import (
	"github.com/gin-gonic/gin"
	"pheasant-api/app/models"
	"pheasant-api/app/requests"
)

type EntityController interface {
	Index(c *gin.Context)
	Show(c *gin.Context)
	Store(c *gin.Context)
	Update(c *gin.Context)
	Destroy(c *gin.Context)
}

type entityControllerDependencies struct{}

func GetEntityController() EntityController {
	return &entityControllerDependencies{}
}

// Index Get a list of Entities of User with Pagination
func (controller *entityControllerDependencies) Index(c *gin.Context) {
	respondOk(c, "OK")
}

func (controller *entityControllerDependencies) Show(c *gin.Context) {
	isAuthorized := requests.GetEntityShowRequest().Authorize(c)
	if !isAuthorized {
		return
	}

	entity, _ := c.Get("entity")
	respondOk(c, entity)
}

func (controller *entityControllerDependencies) Store(c *gin.Context) {
	request := requests.GetEntityStoreRequest()
	storeBody, err := request.Validate(c)
	if err != nil {
		return
	}

	entity := models.Entity{
		UserID:      request.GetUser(c).ID,
		Title:       storeBody.Title,
		Description: storeBody.Description,
	}

	createdEntity, err := models.GetEntityModel().Create(entity)
	if err != nil {
		respondBadRequest(c, gin.H{"error": "Failed to create new Entity"})

		return
	}

	respondCreated(c, createdEntity)
}

func (controller *entityControllerDependencies) Update(c *gin.Context) {
	respondOk(c, "OK")
}

func (controller *entityControllerDependencies) Destroy(c *gin.Context) {
	request := requests.GetEntityDestroyRequest()
	isAuthorized := requests.GetEntityDestroyRequest().Authorize(c)
	if !isAuthorized {
		return
	}

	entity := request.GetEntity(c)
	deletionStatus := models.GetEntityModel().Delete(models.Entity{ID: entity.ID})
	if !deletionStatus {
		respondBadRequest(c, gin.H{"error": "Failed to delete Entity"})

		return
	}

	respondOk(c, gin.H{"uuid": entity.UUID})
}
