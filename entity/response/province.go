package response

import "github.com/dimassfeb-09/restapi-ecommerce.git/entity/domain"

type ProvinceResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ToProvinceResponse(province *domain.Province) *ProvinceResponse {
	return &ProvinceResponse{
		ID:   province.ID,
		Name: province.Name,
	}
}

func ToProvinceResponses(province []*domain.Province) []*ProvinceResponse {
	var provinces []*ProvinceResponse
	for i := 0; i < len(province); i++ {
		provinceResponse := ToProvinceResponse(province[i])
		provinces = append(provinces, provinceResponse)
	}
	return provinces
}
