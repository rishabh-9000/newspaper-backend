package helper

import (
	"crypto/rand"
	"io"
)

// GenerateOTP : Creates OTP
// length : Length of digits in OTP
func GenerateOTP(length int) (string, error) {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

	b := make([]byte, length)
	n, e := io.ReadAtLeast(rand.Reader, b, length)
	if n != length {
		return "", e
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b), nil
}
