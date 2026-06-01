package util_jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// CustomClaims 载荷，可以加一些自己需要的信息
//type MapClaims struct {
//	apiKey               string `json:"api_key"`
//	jwt.RegisteredClaims        //注意!这是jwt-go的v4版本新增的，原先是jwt.StandardClaims
//}

// JwtStruct 签名结构
type JwtStruct struct {
	SignKey string
}

var Jwt *JwtStruct

func init() {
	Jwt = &JwtStruct{SignKey: "1232"}
}

// CreateToken 创建一个 JWT Token
func (obj JwtStruct) CreateToken() (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"api_key":   "db5a77061367b3e57bbdc4c127993382",
		"exp":       now.Add(time.Duration(3600) * time.Second).UnixMilli(),
		"timestamp": now.UnixMilli(),
	}
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token := jwt.Token{
		Header: map[string]interface{}{
			"sign_type": "SIGN",
			"alg":       jwt.SigningMethodHS256.Alg(),
		},
		Claims: claims,
		Method: jwt.SigningMethodHS256,
	}
	tokenString, err := token.SignedString([]byte("GKixqiF3R6xSlhVt"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken 解析 JWT Token
func (obj JwtStruct) ParseToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		secret := obj.SignKey
		return []byte(secret), nil
	})
	if err != nil {
		//if errors.Is(err, jwt.ErrInvalidKey) {
		//	// 密钥无效
		//	//log.Logger.Error(fmt.Sprintf("密钥无效:%s", tokenString))
		//} else if errors.Is(err, jwt.ErrInvalidKeyType) {
		//	//密钥类型无效
		//	//log.Logger.Error(fmt.Sprintf("密钥类型无效:%s", tokenString))
		//} else if errors.Is(err, jwt.ErrHashUnavailable) {
		//	//请求的哈希函数不可用
		//	//log.Logger.Error(fmt.Sprintf("请求的哈希函数不可用:%s", tokenString))
		//} else if errors.Is(err, jwt.ErrTokenMalformed) {
		//	//令牌格式不正确
		//	//log.Logger.Error(fmt.Sprintf("令牌格式不正确:%s", tokenString))
		//} else if errors.Is(err, jwt.ErrTokenUnverifiable) {
		//	//令牌不可验证
		//	log.Logger.Error(fmt.Sprintf("令牌不可验证:%s", tokenString))
		//} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
		//	//令牌签名无效
		//	log.Logger.Error(fmt.Sprintf("令牌签名无效:%s", tokenString))
		//} else if errors.Is(err, jwt.ErrTokenRequiredClaimMissing) {
		//	//令牌缺少所需的声明
		//	log.Logger.Error(fmt.Sprintf("令牌缺少所需的声明:%s", tokenString))
		//} else if errors.Is(err, jwt.ErrTokenInvalidAudience) {
		//	//令牌的受众无效
		//	log.Logger.Error(fmt.Sprintf("令牌的受众无效:%s", tokenString))
		//} else if errors.Is(err, jwt.ErrTokenExpired) {
		//	//令牌过期
		//	log.Logger.Error(fmt.Sprintf("令牌过期:%s", tokenString))
		//} else if errors.Is(err, jwt.ErrTokenUsedBeforeIssued) {
		//	//发出前使用的令牌
		//	log.Logger.Error(fmt.Sprintf("发出前使用的令牌:%s", tokenString))
		//} else if errors.Is(err, jwt.ErrTokenInvalidIssuer) {
		//	//令牌的颁发者无效
		//	log.Logger.Error(fmt.Sprintf("令牌的颁发者无效:%s", tokenString))
		//} else if errors.Is(err, jwt.ErrTokenInvalidSubject) {
		//	//令牌的主题无效
		//	log.Logger.Error(fmt.Sprintf("令牌的主题无效:%s", tokenString))
		//} else if errors.Is(err, jwt.ErrTokenNotValidYet) {
		//	//令牌还无效
		//	log.Logger.Error(fmt.Sprintf("令牌还无效:%s", tokenString))
		//} else if errors.Is(err, jwt.ErrTokenInvalidId) {
		//	//令牌id无效
		//	log.Logger.Error(fmt.Sprintf("令牌id无效:%s", tokenString))
		//} else if errors.Is(err, jwt.ErrTokenInvalidClaims) {
		//	//令牌声明无效
		//	log.Logger.Error(fmt.Sprintf("令牌声明无效:%s", tokenString))
		//} else if errors.Is(err, jwt.ErrInvalidType) {
		//	//索赔类型无效
		//	log.Logger.Error(fmt.Sprintf("索赔类型无效:%s", tokenString))
		//} else {
		//	log.Logger.Error(fmt.Sprintf("非法的令牌:%s", tokenString))
		//}
		return nil, err
	}

	if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("未知异常")
	}

}
