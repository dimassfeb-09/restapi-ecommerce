package helpers

import (
	"github.com/dimassfeb-09/restapi-ecommerce.git/entity/domain"
	"github.com/dimassfeb-09/restapi-ecommerce.git/entity/response"
)

func ToProvinceResponse(province *domain.Province) *response.ProvinceResponse {
	return &response.ProvinceResponse{
		ID:   province.ID,
		Name: province.Name,
	}
}

func ToProvinceResponses(province []*domain.Province) []*response.ProvinceResponse {
	var provinces []*response.ProvinceResponse
	for i := 0; i < len(province); i++ {
		provinceResponse := ToProvinceResponse(province[i])
		provinces = append(provinces, provinceResponse)
	}
	return provinces
}
