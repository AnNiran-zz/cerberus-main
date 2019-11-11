package cryptography

import (
	"crypto/sha1"
	"encoding/base64"
	"crypto/md5"
    "encoding/hex"
)

func ObtainSha(id []byte) string {
	hasher := sha1.New()
	hasher.Write(id)
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	return sha
}

func GetMD5Hash(id string) string {
    hasher := md5.New()
	
    hasher.Write([]byte(id))
    return hex.EncodeToString(hasher.Sum(nil))
}
