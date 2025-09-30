package auth

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID       string   `json:"id"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Roles    []string `json:"roles"`
}

type Claims struct {
	UserID   string   `json:"user_id"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
	jwt.RegisteredClaims
}

type AuthService struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	issuer     string
	audience   string
}

func NewAuthService(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey, issuer, audience string) *AuthService {
	return &AuthService{
		privateKey: privateKey,
		publicKey:  publicKey,
		issuer:     issuer,
		audience:   audience,
	}
}

func (as *AuthService) GenerateToken(user *User) (string, error) {
	claims := &Claims{
		UserID:   user.ID,
		Username: user.Username,
		Roles:    user.Roles,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    as.issuer,
			Audience:  []string{as.audience},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(as.privateKey)
}

func (as *AuthService) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return as.publicKey, nil
	})
	
	if err != nil {
		return nil, err
	}
	
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	
	return nil, errors.New("invalid token")
}

type Permission string

const (
	ReadMetrics  Permission = "read:metrics"
	WriteMetrics Permission = "write:metrics"
	ReadAlerts   Permission = "read:alerts"
	WriteAlerts  Permission = "write:alerts"
	Admin        Permission = "admin"
)

type AuthorizationService struct {
	permissions map[string][]Permission
}

func NewAuthorizationService() *AuthorizationService {
	return &AuthorizationService{
		permissions: map[string][]Permission{
			"admin":     {ReadMetrics, WriteMetrics, ReadAlerts, WriteAlerts, Admin},
			"operator":  {ReadMetrics, ReadAlerts},
			"viewer":    {ReadMetrics},
		},
	}
}

func (az *AuthorizationService) HasPermission(role string, permission Permission) bool {
	rolePermissions, exists := az.permissions[role]
	if !exists {
		return false
	}
	
	for _, p := range rolePermissions {
		if p == permission {
			return true
		}
	}
	
	return false
}

func (az *AuthorizationService) CheckPermission(claims *Claims, permission Permission) bool {
	for _, role := range claims.Roles {
		if az.HasPermission(role, permission) {
			return true
		}
	}
	return false
}

func (az *AuthorizationService) RequirePermission(claims *Claims, permission Permission) error {
	if !az.CheckPermission(claims, permission) {
		return fmt.Errorf("insufficient permissions: %s required", permission)
	}
	return nil
}
