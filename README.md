# passgen 
[![Build Status](https://travis-ci.org/JonathanSudibya/passgen.svg?branch=master)](https://travis-ci.org/JonathanSudibya/passgen)
[![codecov](https://codecov.io/gh/JonathanSudibya/passgen/branch/master/graph/badge.svg)](https://codecov.io/gh/JonathanSudibya/passgen)
[![Go Report Card](https://goreportcard.com/badge/github.com/JonathanSudibya/passgen)](https://goreportcard.com/report/github.com/JonathanSudibya/passgen)

Simple password (or string generator) library for golang

# Disclaimer

This project is inspired by cmiceli's [password-generator-go](https://github.com/cmiceli/password-generator-go)

# Benchmark
This benchmark generates string with 32 characters long.
```
BenchmarkNewPasswordAllChars-4           2000000               749 ns/op              32 B/op          1 allocs/op
BenchmarkNewPasswordCapsChars-4          2000000               731 ns/op              32 B/op          1 allocs/op
BenchmarkNewPasswordLowerChars-4         2000000               741 ns/op              32 B/op          1 allocs/op
BenchmarkNewPasswordNumberChars-4        2000000               737 ns/op              32 B/op          1 allocs/op
BenchmarkNewPasswordSymbolChars-4        2000000               739 ns/op              32 B/op          1 allocs/op
BenchmarkNewPasswordMixedChars-4         2000000               760 ns/op              32 B/op          1 allocs/op
```

# LICENSE

Please see [LICENSE](https://github.com/JonathanSudibya/passgen/blob/master/LICENSE)