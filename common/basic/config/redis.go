package config

// RedisConfig redis 配置
type RedisConfig interface {
	GetEnabled() bool
	GetConn() string
	GetPassword() string
	GetDBNum() int
}

type defaultRedisConfig struct {
	Enabled  bool   `json:"enabled"`
	Conn     string `json:"conn"`
	Password string `json:"password"`
	DBNum    int    `json:"dbNum"`
	Timeout  int    `json:"timeout"`
}

// GetEnabled redis 配置是否激活
func (r defaultRedisConfig) GetEnabled() bool {
	return r.Enabled
}

// GetConn redis 地址
func (r defaultRedisConfig) GetConn() string {
	return r.Conn
}

// GetPassword redis 密码
func (r defaultRedisConfig) GetPassword() string {
	return r.Password
}

// GetDBNum redis 数据库分区序号
func (r defaultRedisConfig) GetDBNum() int {
	return r.DBNum
}
