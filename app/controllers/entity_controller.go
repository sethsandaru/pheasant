package controllers

import (
	"github.com/gin-gonic/gin"
	"pheasant-api/app/models"
	"pheasant-api/app/requests"
	"strconv"
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
	request := requests.GetEntityIndexRequest()
	user := request.GetUser(c)

	// listing params
	keyword := c.Query("keyword")
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))

	entities, err := models.GetEntityModel().Search(user, keyword, page, pageSize)
	if err != nil {
		respondBadRequest(c, gin.H{"message": "Failed to retrieve the entities"})

		return
	}

	respondOk(c, gin.H{"data": entities, "page": page})
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
	request := requests.GetEntityUpdateRequest()
	isAuthorized := request.Authorize(c)
	if !isAuthorized {
		return
	}

	updateBody, err := request.Validate(c)
	if err != nil {
		return
	}

	entity := request.GetEntity(c)
	entity.Title = updateBody.Title
	entity.Description = updateBody.Description

	updatedEntity, err := models.GetEntityModel().Update(entity)
	if err != nil {
		respondBadRequest(c, gin.H{"error": "Failed to update Entity"})

		return
	}

	respondOk(c, gin.H{"uuid": updatedEntity.UUID})
}

func (controller *entityControllerDependencies) Destroy(c *gin.Context) {
	request := requests.GetEntityDestroyRequest()
	isAuthorized := request.Authorize(c)
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
