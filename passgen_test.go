package passgen

import (
	"testing"
)

func TestNewPasswordEmpty(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code did panic. %s", r)
		}
	}()
	options := []string{}
	NewPassword(32, options)
}

func TestNewPasswordCapsLower(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code did panic. %s", r)
		}
	}()
	options := []string{"CapsChars", "LowerChars"}
	NewPassword(32, options)
}

func TestNewPasswordNumSymbol(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code did panic. %s", r)
		}
	}()
	options := []string{"NumberChars", "SymbolChars"}
	NewPassword(32, options)
}

func BenchmarkNewPasswordAllChars(b *testing.B) {
	options := []string{}
	for i := 0; i < b.N; i++ {
		NewPassword(32, options)
	}
}

func BenchmarkNewPasswordCapsChars(b *testing.B) {
	options := []string{"CapsChars"}
	for i := 0; i < b.N; i++ {
		NewPassword(32, options)
	}
}

func BenchmarkNewPasswordLowerChars(b *testing.B) {
	options := []string{"LowerChars"}
	for i := 0; i < b.N; i++ {
		NewPassword(32, options)
	}
}

func BenchmarkNewPasswordNumberChars(b *testing.B) {
	options := []string{"NumberChars"}
	for i := 0; i < b.N; i++ {
		NewPassword(32, options)
	}
}

func BenchmarkNewPasswordSymbolChars(b *testing.B) {
	options := []string{"SymbolChars"}
	for i := 0; i < b.N; i++ {
		NewPassword(32, options)
	}
}

func BenchmarkNewPasswordMixedChars(b *testing.B) {
	options := []string{"CapsChars", "LowerChars"}
	for i := 0; i < b.N; i++ {
		NewPassword(32, options)
	}
}
