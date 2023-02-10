package response

import "github.com/dimassfeb-09/restapi-ecommerce.git/entity/domain"

type CityResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	ProvinceID int    `json:"province_id"`
}

func ToCityResponse(city *domain.City) *CityResponse {
	return &CityResponse{
		ID:         city.ID,
		Name:       city.Name,
		ProvinceID: city.ProvinceID,
	}
}

func ToCityResponses(city []*domain.City) []*CityResponse {
	var cityResponses []*CityResponse
	for i := 0; i < len(city); i++ {
		toCityResponse := ToCityResponse(city[i])
		cityResponses = append(cityResponses, toCityResponse)
	}
	return cityResponses
}
