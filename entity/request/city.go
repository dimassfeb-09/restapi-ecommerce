package request

type CreateCityRequest struct {
	Name string `validate:"required" form:"name" json:"name"`
}

type UpdateCityRequest struct {
	ID   int    `validate:"required" form:"id" json:"id"`
	Name string `validate:"required" form:"name" json:"name"`
}
