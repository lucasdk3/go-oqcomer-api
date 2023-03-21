package services

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

type jwtService struct {
	accessKey  string
	refreshKey string
	issure     string
}

type TokenDetails struct {
	AccessToken    string
	RefreshToken   string
	AccessUuid     string
	RefreshUuid    string
	AccessExpires  int64
	RefreshExpires int64
}

func JWTService() *jwtService {
	return &jwtService{
		accessKey:  "secret-key",
		refreshKey: "refresh-key",
		issure:     "book-api",
	}
}

type Claim struct {
	Sum uint `json:"sum"`
	jwt.StandardClaims
}

func (s *jwtService) GenerateToken(id uint) (*TokenDetails, error) {
	var err error

	tokenDetails := &TokenDetails{}
	tokenDetails.AccessExpires = time.Now().Add(time.Hour * 2).Unix()
	tokenDetails.AccessUuid = uuid.NewV4().String()

	tokenDetails.RefreshExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	tokenDetails.RefreshUuid = uuid.NewV4().String()

	accessClaim := &Claim{
		id,
		jwt.StandardClaims{
			Id:        tokenDetails.AccessUuid,
			ExpiresAt: tokenDetails.AccessExpires,
			Issuer:    s.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaim)

	tokenDetails.AccessToken, err = token.SignedString([]byte(s.accessKey))
	if err != nil {
		return nil, err
	}

	refreshClaim := &Claim{
		id,
		jwt.StandardClaims{
			Id:        tokenDetails.RefreshUuid,
			ExpiresAt: tokenDetails.RefreshExpires,
			Issuer:    s.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaim)

	tokenDetails.RefreshToken, err = refreshToken.SignedString([]byte(s.refreshKey))
	if err != nil {
		return nil, err
	}

	return tokenDetails, nil
}

func (s *jwtService) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %v", token)
		}

		return []byte(s.accessKey), nil
	})

	return err == nil
}

func (s *jwtService) ValidateRefreshToken(token string) (*jwt.Token, error) {
	validRefreshToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %v", token)
		}

		return []byte(s.refreshKey), nil
	})

	if err != nil {
		return nil, err
	}

	return validRefreshToken, nil
}

func (s *jwtService) GetSecretKey() string {
	secretKey := s.accessKey

	return secretKey
}

func (s *jwtService) GetJWTClaim(tokenString string) (uint, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.accessKey), nil
	})

	var fakeUint uint = 0
	if err != nil {
		return fakeUint, err
	}

	if claims, ok := token.Claims.(*Claim); ok && token.Valid {
		return claims.Sum, nil
	} else {
		return fakeUint, nil
	}
}

func (s *jwtService) GetJWTRefreshClaim(refreshTokenString string) (uint, error) {
	token, err := jwt.ParseWithClaims(refreshTokenString, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.refreshKey), nil
	})

	var fakeUint uint = 0
	if err != nil {
		return fakeUint, err
	}

	if claims, ok := token.Claims.(*Claim); ok && token.Valid {
		return claims.Sum, nil
	} else {
		return fakeUint, nil
	}
}

