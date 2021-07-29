package strings

import (
	"crypto/sha1"
	"fmt"
)

func Hash(v string) string {
	hash := sha1.New()
	hash.Write([]byte(v))
	hashBytePass := hash.Sum(nil)
	return fmt.Sprintf("%x", hashBytePass)
}
