package admin_user

type AdminUser struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}
