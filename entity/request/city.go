package request

type CreateCityRequest struct {
	Name       string `validate:"required" form:"name" json:"name"`
	ProvinceID int    `validate:"required" form:"province_id" json:"province_id"`
}

type UpdateCityRequest struct {
	ID         int    `validate:"required" form:"id" json:"id"`
	Name       string `validate:"required" form:"name" json:"name"`
	ProvinceID int    `validate:"required" form:"province_id" json:"province_id"`
}
