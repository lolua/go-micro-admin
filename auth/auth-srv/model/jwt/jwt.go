package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"micro-admin/common/basic/config"
	"time"
)

// Subject token 持有者
type Subject struct {
	ID   string `json:"id"`
	Data string `json:"data,omitempty"`
}

var (
	// tokenExpiredDate app token过期日期 30天
	tokenExpiredDate = 3600 * 24 * 30 * time.Second

	secretKey string
)

func GenerateToken(subject *Subject) (string, error) {
	claims, err := createTokenClaims(subject)
	if err != nil {
		return "", fmt.Errorf("[GenerateToken] 创建token失败，err: %s", err)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ret, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("[GenerateToken] 创建token失败，err: %s", err)
	}
	return ret, nil
}

func GetSubjectFromToken(tk string) (*Subject, error) {
	claims, err := parseToken(tk)
	if err != nil {
		return nil, err
	}
	fmt.Println("claims", claims)
	return &Subject{ID: claims.Id, Data: claims.Subject}, nil
}

func createTokenClaims(subject *Subject) (m *jwt.StandardClaims, err error) {
	now := time.Now()
	m = &jwt.StandardClaims{
		ExpiresAt: now.Add(tokenExpiredDate).Unix(),
		NotBefore: now.Unix(),
		Id:        subject.ID,
		IssuedAt:  now.Unix(),
		Issuer:    "zhj.micro.admin",
		Subject:   subject.Data,
	}
	return
}
func parseToken(tk string) (c *jwt.StandardClaims, err error) {
	token, err := jwt.Parse(tk, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("未符合的签名方法: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	// jwt 框架自带了一些检测，如过期，发布者错误等
	if err != nil {
		switch e := err.(type) {
		case *jwt.ValidationError:
			switch e.Errors {
			case jwt.ValidationErrorExpired:
				return nil, fmt.Errorf("[parseToken] 过期的token, err:%s", err)
			default:
				break
			}
			break
		default:
			break
		}

		return nil, fmt.Errorf("[parseToken] 不合法的token, err:%s", err)
	}
	// 检测合法
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("[parseToken] 不合法的token")
	}
	jc := &jwt.StandardClaims{
		Id:      claims["jti"].(string),
		Subject: claims["sub"].(string),
	}
	return jc, nil
}

func init() {
	secretKey = config.GetJwtConfig().GetSecretKey()
}
