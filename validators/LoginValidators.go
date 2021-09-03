package validators

type LoginPostValidator struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
