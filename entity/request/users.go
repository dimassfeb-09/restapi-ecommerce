package request

type CreateUserRequest struct {
	Name     string `validate:"required" form:"name" json:"name"`
	Username string `validate:"required" form:"username" json:"username"`
	Password string `validate:"required" form:"password" json:"password"`
}

type UpdateUserRequest struct {
	Id       int    `validate:"numeric" json:"id"`
	Name     string `validate:"required" form:"name" json:"name"`
	Username string `validate:"required" form:"username"json:"username"`
	Password string `validate:"required" form:"password" json:"password"`
	Balance  int    `validate:"required" form:"balance" json:"balance"`
}
