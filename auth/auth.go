package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"log"
	"net/http"
	"net/textproto"
	"strings"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func HashPassword(password string) (string, error) {
	myparams := &argon2id.Params{
		Memory:      128 * 1024,
		Iterations:  4,
		Parallelism: uint8(4),
		SaltLength:  16,
		KeyLength:   32,
	}

	// myparams := argon2id.DefaultParams
	hash, err := argon2id.CreateHash(password, myparams)
	if err != nil {
		log.Fatal(err)
	}
	return hash, err
}

func CheckPasswordHash(password, hash string) (bool, error) {
	match, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		log.Fatal(err)
	}
	return match, err
}

func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {
	decodedstr := []byte(tokenSecret)

	myclaims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    "chirpy-access",
		Subject:   userID.String(),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myclaims)
	ss, err := t.SignedString(decodedstr)
	return ss, err
}

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	decodedstr := []byte(tokenSecret)

	type myclaims struct{ jwt.RegisteredClaims }
	token, err := jwt.ParseWithClaims(tokenString, &myclaims{}, func(token *jwt.Token) (any, error) {
		return decodedstr, nil
	})
	if err != nil {
		log.Fatal(err)
	}

	claims, ok := token.Claims.(*myclaims)
	if !ok {
		log.Fatal("unknown claims type, cannot proceed")
	}
	idstr := claims.Subject
	id, err := uuid.Parse(idstr)
	if err != nil {
		log.Fatal(err)
	}
	return id, nil
}

func GetBearerToken(headers http.Header) (string, error) {
	bts := headers.Get(textproto.CanonicalMIMEHeaderKey("Authorization"))
	if bts == "" {
		return bts, errors.New("No Authorization header found")
	}

	before, after, found := strings.Cut(bts, "Bearer ")
	if found != true || before != "" {
		return bts, errors.New("Authorization not Bearer format")
	}

	return after, nil
}

func MakeRefreshToken() string {
	key := make([]byte, 32)
	rand.Read(key)
	encodestr := hex.EncodeToString(key)
	return encodestr
}
