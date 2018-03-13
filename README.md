# passgen 
[![Build Status](https://travis-ci.org/JonathanSudibya/passgen.svg?branch=master)](https://travis-ci.org/JonathanSudibya/passgen)
[![codecov](https://codecov.io/gh/JonathanSudibya/passgen/branch/master/graph/badge.svg)](https://codecov.io/gh/JonathanSudibya/passgen)
[![Go Report Card](https://goreportcard.com/badge/github.com/JonathanSudibya/passgen)](https://goreportcard.com/report/github.com/JonathanSudibya/passgen)

Simple password (or string generator) library for golang

# Disclaimer

This project is inspired by cmiceli's [password-generator-go](https://github.com/cmiceli/password-generator-go)

# Benchmark
```
goos: windows
goarch: amd64
pkg: github.com/jonathansudibya/passgen
BenchmarkNewPasswordAllChars-4           2000000               808 ns/op               0 B/op          0 allocs/op
BenchmarkNewPasswordCapsChars-4          2000000               795 ns/op               0 B/op          0 allocs/op
BenchmarkNewPasswordLowerChars-4         2000000               786 ns/op               0 B/op          0 allocs/op
BenchmarkNewPasswordNumberChars-4        2000000               795 ns/op               0 B/op          0 allocs/op
BenchmarkNewPasswordSymbolChars-4        2000000               831 ns/op               0 B/op          0 allocs/op
BenchmarkNewPasswordMixedChars-4         2000000               822 ns/op               0 B/op          0 allocs/op
```

# LICENSE

Please see [LICENSE](https://github.com/JonathanSudibya/passgen/blob/master/LICENSE)