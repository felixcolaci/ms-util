package ms_util

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"fmt"
)

type RequiredAuthCriteria struct {
	scopes []string
	roles  []string
}

type AuthError struct {
	Message string `json:"message"`
}

func (e AuthError) Error() string {
	return e.Message
}

func VerifyScopes(h http.HandlerFunc, requiredScopes []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if len(authHeader) < 1 {
			writeErr(w, http.StatusUnauthorized, "Missing authorization")
			return
		}
		parts := strings.Split(authHeader, " ")
		tokenType := parts[0]
		token := parts[1]

		if tokenType != "Bearer" {
			writeErr(w, http.StatusUnauthorized, "Invalid token type")
			return
		}

		if len(token) < 1 {
			writeErr(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		claims := jwt.MapClaims{}
		jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("jenlix"), nil
		})

		fmt.Println(claims["scope"].(string))
		
		grantedScopes := strings.Split(claims["scope"].(string), " ")
		hasAllScopes := true
		for _, scope := range requiredScopes {
			hasAllScopes = inSlice(grantedScopes, scope)
			if !hasAllScopes {
				break
			}
		}
		if !hasAllScopes {
			writeErr(w, http.StatusForbidden, "no permission")
			return
		}

		//On Success call handler
		h.ServeHTTP(w, r)
	})
}

func SecureWithOAuthAndRoles(h http.HandlerFunc, requiredScopes []string, requiredRoles []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if len(authHeader) < 1 {
			writeErr(w, http.StatusUnauthorized, "Missing authorization")
			return
		}
		parts := strings.Split(authHeader, " ")
		tokenType := parts[0]
		token := parts[1]

		if tokenType != "Bearer" {
			writeErr(w, http.StatusUnauthorized, "Invalid token type")
			return
		}

		if len(token) < 1 {
			writeErr(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		claims := jwt.MapClaims{}
		jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("jenlix"), nil
		})

		grantedScopes := strings.Split(claims["scope"].(string), " ")
		grantedRoles := strings.Split(claims["roles"].(string), " ")

		hasAllScopes := true
		for _, scope := range requiredScopes {
			hasAllScopes = inSlice(grantedScopes, scope)
			if !hasAllScopes {
				break
			}
		}
		hasAllRoles := true
		for _, role := range requiredRoles {
			hasAllRoles= inSlice(grantedRoles, role)
			if !hasAllRoles {
				break
			}
		}

		if !hasAllScopes || !hasAllRoles {
			writeErr(w, http.StatusForbidden, "no permission")
			return
		}

		//On Success call handler
		h.ServeHTTP(w, r)
	})
}

func writeErr(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := AuthError{Message: message}
	json.NewEncoder(w).Encode(err)
}

func inSlice(slice []string, probe string) bool {
	for _, t := range slice {
		if t == probe {
			return true
		}
	}
	return false
}