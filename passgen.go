package passgen

import (
	"crypto/rand"
	"io"
)

var stdChars map[string][]byte

func init() {
	stdChars["CapsChars"] = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	stdChars["LowerChars"] = []byte("abcdefghijklmnopqrstuvwxyz")
	stdChars["NumberChars"] = []byte("0123456789")
	stdChars["SymbolChars"] = []byte("!@#$%^&*()-_=+,.?/:;{}[]`~")
}

// NewPassword will generate random string equal to length and options provided
func NewPassword(length int, options []string) string {
	chars := make([]byte, 0)
	if len(options) > 0 {
		for _, v := range options {
			if stdChars[v] != nil {
				chars = append(chars, stdChars[v]...)
			}
		}
	} else {
		for _, v := range stdChars {
			chars = append(chars, v...)
		}
	}
	return randChar(length, chars)
}

func randChar(length int, chars []byte) string {
	newPword := make([]byte, length)
	randomData := make([]byte, length+(length/4))
	clen := byte(len(chars))
	maxrb := byte(256 - (256 % len(chars)))
	i := 0
	for {
		if _, err := io.ReadFull(rand.Reader, randomData); err != nil {
			panic(err)
		}
		for _, c := range randomData {
			if c >= maxrb {
				continue
			}
			newPword[i] = chars[c%clen]
			i++
			if i == length {
				return string(newPword)
			}
		}
	}
}
