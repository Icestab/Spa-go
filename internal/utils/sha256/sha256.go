package sha256

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(data string) string {
	s := []byte(data)
	key := []byte("kjagdlkasj123414")
	m := hmac.New(sha256.New, key)
	m.Write(s)
	singnature := hex.EncodeToString(m.Sum(nil))
	return singnature
}
