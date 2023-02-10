package request

type AuthLoginRequest struct {
	Username string `binding:"required" json:"username"`
	Password string `binding:"required" json:"password"`
}

type AuthRegisterRequest struct {
	Name     string `binding:"required" json:"name"`
	Username string `binding:"required" json:"username"`
	Password string `binding:"required" json:"password"`
}
