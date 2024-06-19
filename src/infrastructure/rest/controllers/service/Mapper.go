package service

import (
	"pausalac/src/domain"
	"time"
)

// ToDomainArray maps an array of Service to ServiceResponse array
func ToResponseArray(services *[]domain.Service) []ServiceResponse {
	var serviceResponses []ServiceResponse
	for _, service := range *services {
		serviceResponses = append(serviceResponses, ToResponse(&service))
	}
	return serviceResponses
}

// ToResponse maps Service to ServiceResponse
func ToResponse(service *domain.Service) ServiceResponse {
	return ServiceResponse{
		Id:     service.Id.Hex(),
		Author: service.Author,
		Name:   service.Name,
	}
}

func ToDomain(req *CreateServiceRequest) *domain.NewService {
	return &domain.NewService{
		Author: req.Author,
		Name:   req.Name,
	}
}

func ToDomainArray(req *CreateServiceArrayRequest) domain.NewServiceArray {
	var services []domain.NewService
	if *req == nil {
		return services
	}

	for _, service := range *req {
		s := ToDomain(&service)
		services = append(services, *s)
	}
	return services
}

func ToDomainUpdateArray(req *UpdateServiceArrayRequest) []map[string]interface{} {
	var services []map[string]interface{}
	if *req == nil {
		return services
	}
	for _, service := range *req {
		s := ToDomainUpdate(&service)
		services = append(services, s)
	}
	return services
}

// ToMap maps UpdateServiceRequest to a map for updating the service
func ToDomainUpdate(req *UpdateServiceRequest) map[string]interface{} {
	serviceMap := make(map[string]interface{})
	if req.Name != "" {
		serviceMap["name"] = req.Name
	}
	serviceMap["updated_at"] = time.Now()
	return serviceMap
}
