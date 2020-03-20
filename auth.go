package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

type Key string

func (k Key) Signature(message string) string {
	h := hmac.New(sha256.New, []byte(k))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}
