package request

type CreateExpeditionRequest struct {
	ID   int    `json:"id"`
	Name string `validate:"required" json:"name"`
}

type UpdateExpeditionRequest struct {
	ID   int    `json:"id"`
	Name string `validate:"required" json:"name"`
}
