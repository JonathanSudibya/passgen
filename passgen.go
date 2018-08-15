package passgen

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
	"unsafe"

	"github.com/jonathansudibya/passgen/buffer"
)

func byteToString(bytes []byte) string {
	hdr := *(*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	return *(*string)(unsafe.Pointer(&reflect.StringHeader{
		Data: hdr.Data,
		Len:  hdr.Len,
	}))
}

func stringToByte(str string) []byte {
	hdr := *(*reflect.StringHeader)(unsafe.Pointer(&str))
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: hdr.Data,
		Len:  hdr.Len,
		Cap:  hdr.Len,
	}))
}

var stdChars = [][]byte{stringToByte("ABCDEFGHIJKLMNOPQRSTUVWXYZ"), stringToByte("abcdefghijklmnopqrstuvwxyz"), stringToByte("0123456789"), stringToByte("!@#$%^&*()-_=+,.?/:;{}[]`~")}
var bufPool = buffer.NewPool()
var r = rand.New(rand.NewSource(time.Now().UnixNano()))

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
	return byteToString(randChar(length, chars))
}

func randChar(length int, chars *buffer.Buffer) []byte {
	newPword := bufPool.Get()
	defer newPword.Free()

	randomData := bufPool.Get()
	defer randomData.Free()

	clen := byte(chars.Len())

	randomData.Grow(length)
	r.Read(randomData.BS)
	for _, c := range randomData.Bytes() {
		newPword.WriteByte(chars.BS[c%clen])
	}
	return newPword.Bytes()
}
