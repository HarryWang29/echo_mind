package util

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

const (
	SHA1   = "sha1"
	SHA256 = "sha256"
	MD5    = "md5"
)

func Sha256(raw []byte) []byte {
	bs := sha256.Sum256(raw)
	return bs[:]
}

func Sha1(raw []byte) []byte {
	bs := sha1.Sum(raw)
	return bs[:]
}

func HmacSha256(raw []byte, secret []byte) []byte {
	h := hmac.New(sha256.New, secret)
	_, _ = h.Write(raw)
	return h.Sum(nil)
}

func HmacRipeMD160(message, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	_, _ = h.Write(message)
	return h.Sum(nil)
}

func Md5(row []byte) []byte {
	h := md5.New()
	h.Write(row)
	return h.Sum(nil)
}

func Hash(t string, row []byte) []byte {
	switch t {
	case SHA1:
		return Sha1(row)
	case SHA256:
		return Sha256(row)
	case MD5:
		return Md5(row)
	}
	return nil
}

func HashHex(t string, row string) string {
	return hex.EncodeToString(Hash(t, []byte(row)))
}
