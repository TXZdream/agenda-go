package tools

import (
	"fmt"
	"io"
	"crypto/md5"
)

func MD5Encryption(in string) string {
	h := md5.New()
	io.WriteString(h, in)
	return fmt.Sprintf("%x", h.Sum(nil))
}