package utils

import (
	"crypto/hmac"
	b64 "encoding/base64"
	"hash"
	"golang.org/x/crypto/bcrypt"
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
