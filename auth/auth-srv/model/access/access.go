package access

import "micro-admin/auth/auth-srv/model/admin_user"

// Service 用户服务类
type Service interface {
	// MakeAccessToken 生成token
	CreateAccessToken(a *admin_user.AdminUser) (ret string, err error)

	// GetCachedAccessToken 获取缓存的token
	ValidAccessToken(tk string) (err error)

	// DelUserAccessToken 清除用户token
	DelAccessToken(token string) (err error)
}
