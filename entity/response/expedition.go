package response

import "github.com/dimassfeb-09/restapi-ecommerce.git/entity/domain"

type ExpeditionResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ToExpeditionResponse(expedition *domain.Expedition) *ExpeditionResponse {
	return &ExpeditionResponse{
		ID:   expedition.ID,
		Name: expedition.Name,
	}
}

func ToExpeditionResponses(expedition []*domain.Expedition) []*ExpeditionResponse {
	var expeditionResponses []*ExpeditionResponse
	for i := 0; i < len(expedition); i++ {
		expeditionResponses = append(expeditionResponses, ToExpeditionResponse(expedition[i]))
	}
	return expeditionResponses
}
