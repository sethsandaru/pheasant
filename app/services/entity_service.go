package services

type EntityService interface {
}

type entityServiceDependencies struct{}

func GetEntityService() EntityService {
	return &entityServiceDependencies{}
}
