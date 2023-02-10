package request

type CreateUserRequest struct {
	Name     string `validate:"required" form:"name" json:"name"`
	Username string `validate:"required" form:"username" json:"username"`
	Password string `validate:"required" form:"password" json:"password"`
}

type UpdateUserRequest struct {
	ID       int    `validate:"numeric" json:"id"`
	Name     string `validate:"required" form:"name" json:"name"`
	Username string `validate:"required" form:"username" json:"username"`
	Password string `validate:"required" form:"password" json:"password"`
	Balance  int    `validate:"required" form:"balance" json:"balance"`
}

type ChangePasswordRequest struct {
	ID               int    `validate:"numeric" json:"id"`
	PreviousPassword string `validate:"required" json:"previous_password"`
	Password         string `validate:"required" json:"password"`
	ConfirmPassword  string `validate:"required" json:"confirm_password"`
}
