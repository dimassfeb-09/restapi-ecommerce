package request

type CreateProvinceRequest struct {
	ID   int    `json:"id"`
	Name string `validate:"required" form:"name" json:"name"`
}

type UpdateProvinceRequest struct {
	ID   int    `validate:"required" json:"id"`
	Name string `validate:"required" json:"name" form:"name"`
}
