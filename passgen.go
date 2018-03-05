package passgen

import (
	"crypto/rand"

	"github.com/JonathanSudibya/passgen/buffer"
)

var stdChars = [][]byte{[]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ"), []byte("abcdefghijklmnopqrstuvwxyz"), []byte("0123456789"), []byte("!@#$%^&*()-_=+,.?/:;{}[]`~")}
var bufPool = buffer.NewPool()

// NewPassword will generate random string equal to length and options provided
func NewPassword(length int, options ...string) string {
	chars := bufPool.Get()
	defer chars.Free()
	if len(options) > 0 {
		for _, v := range options {
			switch v {
			case "CapsChars":
				chars.BS.Write(stdChars[0])
			case "LowerChars":
				chars.BS.Write(stdChars[1])
			case "NumberChars":
				chars.BS.Write(stdChars[2])
			case "SymbolChars":
				chars.BS.Write(stdChars[3])
			}
		}
	} else {
		for _, v := range stdChars {
			chars.BS.Write(v)
		}
	}
	return randChar(length, chars.BS.Bytes())
}

func randChar(length int, chars []byte) string {
	if len(chars) <= 0 {
		return ""
	}

	newPword := bufPool.Get()
	defer newPword.Free()

	randomData := make([]byte, length+(length/4))

	clen := bufPool.Get()
	defer clen.Free()
	clen.BS.WriteByte(byte(len(chars)))

	i := 0
	for {
		rand.Read(randomData)
		for _, c := range randomData {
			newPword.BS.WriteByte(chars[c%clen.BS.Bytes()[0]])
			i++
			if i == length {
				return newPword.BS.String()
			}
		}
	}
}
