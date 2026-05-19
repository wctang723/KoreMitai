package auth

import (
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
)

// NOTE: Some unit test might need to modify if some function has been touched
// (i.e. ValidateJWT)

func TestHashPassword(t *testing.T) {
	password := "testpasswd"
	hash, err := HashPassword(password)
	matchornot, err := CheckPasswordHash(password, hash)
	if err != nil || matchornot != true {
		t.Errorf(`HashPassword(password) = %q, matchornot = %v, err = %v`, hash, matchornot, err)
	}
}

func TestWrongHashPassword(t *testing.T) {
	password := "testpasswd"
	hash, err := HashPassword(password)
	matchornot, err := CheckPasswordHash("wrongpasswd", hash)
	if err != nil || matchornot == true {
		t.Errorf(`HashPassword(password) = %q, matchornot = %v, err = %v`, hash, matchornot, err)
	}
}

func TestMakeJWT(t *testing.T) {
	testtokensecret := "mysecrettoken"
	expirestime := 1 * time.Hour
	myuseruuid, uuiderr := uuid.Parse("63eae2c6-d4fb-42ac-a814-8e8c5669a7a3")
	JwtString, err := MakeJWT(myuseruuid, testtokensecret, expirestime)
	if err != nil || uuiderr != nil {
		t.Errorf(`MakeJWT error: %v`, err)
	}
	t.Logf(`Everything goes well, token string: %v`, JwtString)
}

func TestValidateJWT(t *testing.T) {
	mytokensecret := "mysecrettoken"
	expirestime := 1 * time.Hour
	myuidstr := "63eae2c6-d4fb-42ac-a814-8e8c5669a7a3"
	myuseruuid, _ := uuid.Parse(myuidstr)
	JwtString, _ := MakeJWT(myuseruuid, mytokensecret, expirestime)

	var parseuuid uuid.UUID
	parseuuid, err := ValidateJWT(JwtString, mytokensecret)
	if err != nil || parseuuid != myuseruuid {
		t.Errorf(`ValidateJWT error: %v`, err)
	}
	t.Logf(`Everything works fine, parsed UUID: %v`, parseuuid)
}

func TestGetBearerToken(t *testing.T) {
	testreq, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		t.Errorf(`NewRequest error: %v`, err)
	}
	testreq.Header.Set("Authorization", "Bearer jieksojc=semovhsef3239==")
	bts, err := GetBearerToken(testreq.Header)
	if err != nil {
		t.Errorf(`GetBearerToken error: %v`, err)
	} else {
		t.Logf(`GetBearerToken works well, the bearer token is: %v`, bts)
	}
}

func TestMakeRefreshToken(t *testing.T) {
	str := MakeRefreshToken()
	t.Logf(`Refresh token: %v`, str)
}
