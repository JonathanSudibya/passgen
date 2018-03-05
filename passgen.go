package passgen

import (
	"crypto/rand"

	"github.com/JonathanSudibya/passgen/buffer"
)

var stdChars = [][]byte{[]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ"), []byte("abcdefghijklmnopqrstuvwxyz"), []byte("0123456789"), []byte("!@#$%^&*()-_=+,.?/:;{}[]`~")}
var bufPool = buffer.NewPool()

// NewPassword will generate random string equal to length and options provided
func NewPassword(length int, options []string) string {
	chars := bufPool.Get()
	defer chars.Free()
	if len(options) > 0 {
		for _, v := range options {
			switch v {
			case "CapsChars":
				chars.AppendBytes(stdChars[0])
			case "LowerChars":
				chars.AppendBytes(stdChars[1])
			case "NumberChars":
				chars.AppendBytes(stdChars[2])
			case "SymbolChars":
				chars.AppendBytes(stdChars[3])
			}
		}
	} else {
		for _, v := range stdChars {
			chars.AppendBytes(v)
		}
	}
	return randChar(length, chars.Bytes())
}

func randChar(length int, chars []byte) string {
	newPword := bufPool.Get()
	defer newPword.Free()
	randomData := make([]byte, length+(length/4))

	clen := bufPool.Get()
	defer clen.Free()
	clen.AppendByte(byte(len(chars)))

	maxrb := bufPool.Get()
	defer maxrb.Free()
	maxrb.AppendByte(byte(256 - (256 % len(chars))))

	i := 0
	for {
		rand.Read(randomData)
		for _, c := range randomData {
			if clen.Byte() >= maxrb.Byte() {
				continue
			}
			newPword.AppendByte(chars[c%clen.Byte()])
			i++
			if i == length {
				return newPword.String()
			}
		}
	}
}
