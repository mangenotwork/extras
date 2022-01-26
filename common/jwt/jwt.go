package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/logger"
	"reflect"
	"time"
)

type JWT struct {
	Claims jwt.MapClaims
	Secret string
	Expire int
}

func NewJWT() *JWT {
	return &JWT{
		Claims: make(jwt.MapClaims),
		Secret: conf.Arg.Jwt.Secret,
		Expire: conf.Arg.Jwt.Expire,
	}
}

func (j *JWT) Token() (string, error) {
	exp := time.Now().Add(time.Duration(j.Expire) * time.Second).Unix()
	logger.Debug("exp = ", exp)
	j.Claims["Expire"] = exp

	jToken := jwt.New(jwt.SigningMethodHS256)
	jToken.Claims = j.Claims
	return jToken.SignedString([]byte(j.Secret))
}

func (j *JWT) ParseToken(token string) error {
	var ok bool
	JwtToken, err := jwt.Parse(token, j.secret())
	if err != nil {
		return err
	}
	j.Claims, ok = JwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return fmt.Errorf("cannot convert claim to map claim")
	}
	return nil
}

func (j *JWT) secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	}
}

func (j *JWT) IsExpire() bool {
	t := time.Now().Unix()
	t2 := int64(j.Claims["Expire"].(float64))
	if t2 > t {
		return false
	}
	return true
}

func (j *JWT) AddClaims(k string, v interface{}) *JWT {
	j.Claims[k] = v
	return j
}

func (j *JWT) Print() {
	for k, v := range j.Claims {
		logger.Debug(k , v)
	}
}

func (j *JWT) Get(k string) interface{} {
	v, ok := j.Claims[k]
	if !ok {
		return nil
	}
	logger.Debug(reflect.TypeOf(v))
	return v
}

func (j *JWT) GetFloat64(k string) float64 {
	v, ok := j.Claims[k]
	if !ok {
		return 0
	}
	v2, ok2 := v.(float64)
	if !ok2 {
		return 0
	}
	return v2
}

func (j *JWT) GetInt(k string) int {
	i := j.GetFloat64(k)
	return int(i)
}

func (j *JWT) GetInt64(k string) int64 {
	i := j.GetFloat64(k)
	return int64(i)
}

func (j *JWT) GetString(k string) string {
	v, ok := j.Claims[k]
	if !ok {
		return ""
	}
	v2, ok2 := v.(string)
	if !ok2 {
		return ""
	}
	return v2
}

func (j *JWT) GetBool(k string) bool {
	v, ok := j.Claims[k]
	if !ok {
		return false
	}
	v2, ok2 := v.(bool)
	if !ok2 {
		return false
	}
	return v2
}

