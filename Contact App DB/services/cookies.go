package services

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Claims represents the JWT claims structure
type Claims struct {
	Username string
	jwt.StandardClaims
}

// jwtKey is the secret key used for JWT signing
var jwtKey = []byte("secret_key")

// SetCookie generates a JWT token and sets it as a cookie in the response
func SetCookie(w http.ResponseWriter, r *http.Request, username string) {

	// Set expiration time for the token
	expirationTime := time.Now().Add(time.Minute * 5)

	// Create custom claims with username and expiration time
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Generate JWT token with signing method HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the token as a cookie in the response
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

// isValidCoockie checks if the token in the request cookie is valid
func isValidCoockie(w http.ResponseWriter, r *http.Request) bool {
	// Retrieve the token from the request cookie
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return false
		}
		w.WriteHeader(http.StatusBadRequest)
		return false
	}
	tokenStr := cookie.Value

	// Create claims struct to store the extracted claims from the token
	claims := &Claims{}

	// Parse the token and validate the signature
	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return false
		}
		w.WriteHeader(http.StatusBadRequest)
		return false
	}

	// Check if the token is valid
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}
	return true
}
