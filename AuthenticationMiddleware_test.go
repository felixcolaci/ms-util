package ms_util_test

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/felixcolaci/ms-util"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func demoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

var testCases = []struct {
	name            string
	expectedStatus  int
	expectedMessage string
	withToken       bool
	generateJwt     bool
	claims          jwt.MapClaims
	tokenType       string
	token           string
}{
	{"without Authorization Header", http.StatusUnauthorized, "Missing authorization", false, false, nil, "", ""},
	{"with invalid token type", http.StatusUnauthorized, "Invalid token type", true, false, nil, "foo", "bar"},
	{"with invalid token", http.StatusUnauthorized, "Invalid token", true, false, nil, "Bearer", ""},
	{"with valid token and valid rights", http.StatusOK, "", true, true, jwt.MapClaims{"scope": "openid profile", "roles": "org_user"}, "Bearer", ""},
	{"with valid token but insufficient scopes", http.StatusForbidden, "no permission", true, true, jwt.MapClaims{"scope": "openid", "roles": "org_user"}, "Bearer", ""},
	{"with valid token but insufficient roles", http.StatusForbidden, "no permission", true, true, jwt.MapClaims{"scope": "openid profile", "roles": "org_admin"}, "Bearer", ""},
	{"with valid token but no scopes", http.StatusForbidden, "no permission", true, true, jwt.MapClaims{"scope": "", "roles": "org_user"}, "Bearer", ""},
	{"with valid token but no roles", http.StatusForbidden, "no permission", true, true, jwt.MapClaims{"scope": "openid profile", "roles": ""}, "Bearer", ""},
}

//Test Oauth Handler
func TestOAuthAndRoles(t *testing.T) {

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			req, err := http.NewRequest("GET", "/demo", nil)
			if err != nil {
				t.Fatalf("unable to construct http request; %v", err)
			}

			if tc.withToken {
				req.Header.Set("Authorization", tc.tokenType+" "+tc.token)
			}

			if tc.generateJwt {

				if len(tc.token) < 1 {
					jwt, err := generateJwt(tc.claims)
					if err != nil {
						t.Fatalf("unable to generate jwt; %v", err)
					}
					tc.token = jwt
				}
				req.Header.Set("Authorization", tc.tokenType+" "+tc.token)

			}

			rr := httptest.NewRecorder()

			http.Handler(ms_util.SecureWithOAuthAndRoles(demoHandler, []string{"openid", "profile"}, []string{"org_user"})).ServeHTTP(rr, req)

			if status := rr.Code; status != tc.expectedStatus {
				t.Errorf("wrong statuscode returned; expected %v but got %v",
					tc.expectedStatus, status)
			}

			var err2 ms_util.AuthError
			decodingErr := json.NewDecoder(rr.Body).Decode(&err2)
			if decodingErr != nil {
				t.Fatal("could not decode response")
			}

			if err2.Error() != tc.expectedMessage {
				t.Errorf("got wrong error message; expected '%v' but got '%v'", tc.expectedMessage, err2.Error())
			}

		})

	}

}

func generateJwt(claims jwt.MapClaims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("jenlix"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
