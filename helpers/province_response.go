package response

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
