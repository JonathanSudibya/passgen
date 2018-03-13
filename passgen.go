package passgen

import (
	"crypto/rand"
	"fmt"

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
				chars.Write(stdChars[0])
			case "LowerChars":
				chars.Write(stdChars[1])
			case "NumberChars":
				chars.Write(stdChars[2])
			case "SymbolChars":
				chars.Write(stdChars[3])
			default:
				panic(fmt.Errorf("character group not found"))
			}
		}
	} else {
		for _, v := range stdChars {
			chars.Write(v)
		}
	}
	return string(randChar(length, chars))
}

func randChar(length int, chars *buffer.Buffer) []byte {
	newPword := bufPool.Get()
	defer newPword.Free()

	randomData := bufPool.Get()
	defer randomData.Free()

	clen := byte(chars.Len())

	randomData.Grow(length)
	rand.Read(randomData.BS)
	for _, c := range randomData.Bytes() {
		newPword.WriteByte(chars.BS[c%clen])
	}
	return newPword.Bytes()
}
