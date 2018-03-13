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
	NewPassword(0)
}

func TestNewPasswordAllType(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code did panic. %s", r)
		}
	}()
	NewPassword(0)
}

func TestNewPasswordCapsLower(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code did panic. %s", r)
		}
	}()
	NewPassword(32, "CapsChars", "LowerChars")
}

func TestNewPasswordUndefined(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic. %s", r)
		}
	}()
	NewPassword(32, "CapChars", "LowerCharizard")
}

func TestNewPasswordNumSymbol(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code did panic. %s", r)
		}
	}()
	NewPassword(32, "NumberChars", "SymbolChars")
}

func BenchmarkNewPasswordAllChars(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewPassword(32)
	}
}

func BenchmarkNewPasswordCapsChars(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewPassword(32, "CapsChars")
	}
}

func BenchmarkNewPasswordLowerChars(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewPassword(32, "LowerChars")
	}
}

func BenchmarkNewPasswordNumberChars(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewPassword(32, "NumberChars")
	}
}

func BenchmarkNewPasswordSymbolChars(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewPassword(32, "SymbolChars")
	}
}

func BenchmarkNewPasswordMixedChars(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewPassword(32, "CapsChars", "LowerChars")
	}
}
