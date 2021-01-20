package utils

import (
	"context"
	"crypto/hmac"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gichohi/go-rest.git/models"
	"hash"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

func GenerateHash(hash func() hash.Hash, password []byte, salt []byte) string {
	iterations := 256
	outlen := 32
	out := make([]byte, outlen)
	hashSize := hash().Size()
	ibuf := make([]byte, 4)
	block := 1
	p := out
	for outlen > 0 {
		clen := outlen
		if clen > hashSize {
			clen = hashSize
		}

		ibuf[0] = byte((block >> 24) & 0xff)
		ibuf[1] = byte((block >> 16) & 0xff)
		ibuf[2] = byte((block >> 8) & 0xff)
		ibuf[3] = byte((block) & 0xff)

		hmac := hmac.New(hash, password)
		hmac.Write(salt)
		hmac.Write(ibuf)
		tmp := hmac.Sum(nil)
		for i := 0; i < clen; i++ {
			p[i] = tmp[i]
		}

		for j := 1; j < iterations; j++ {
			hmac.Reset()
			hmac.Write(tmp)
			tmp = hmac.Sum(nil)
			for k := 0; k < clen; k++ {
				p[k] ^= tmp[k]
			}
		}
		outlen -= clen
		block++
		p = p[clen:]
	}

	s := string(out[:])
	hashedpassword := b64.StdEncoding.EncodeToString([]byte(s))
	return hashedpassword
}

func GenerateCrypt(passkey string) string{
	password := []byte(passkey)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(hashedPassword))

	return string(hashedPassword)
}

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var resp = models.Response{}
		var header = r.Header.Get("x-access-token") //Grab the token from the header

		header = strings.TrimSpace(header)

		if header == "" {
			//Token is missing, returns with error code 403 Unauthorized
			w.WriteHeader(http.StatusForbidden)
			resp.Code = http.StatusForbidden;
			resp.Message = http.StatusText(http.StatusForbidden)
			json.NewEncoder(w).Encode(resp)
			return
		}
		token := &models.Token{}

		_, err := jwt.ParseWithClaims(header, token, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			resp.Code = http.StatusForbidden;
			resp.Message = http.StatusText(http.StatusForbidden)
			json.NewEncoder(w).Encode(resp)
			return
		}

		ctx := context.WithValue(r.Context(), "user", token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

