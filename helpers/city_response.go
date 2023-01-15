package helpers

import (
	"github.com/dimassfeb-09/restapi-ecommerce.git/entity/domain"
	"github.com/dimassfeb-09/restapi-ecommerce.git/entity/response"
)

func ToCityResponse(city *domain.City) *response.CityResponse {
	return &response.CityResponse{
		ID:   city.ID,
		Name: city.Name,
	}
}

func ToCityResponses(city []*domain.City) []*response.CityResponse {
	var cityResponses []*response.CityResponse
	for i := 0; i < len(city); i++ {
		toCityResponse := ToCityResponse(city[i])
		cityResponses = append(cityResponses, toCityResponse)
	}
	return cityResponses
}
