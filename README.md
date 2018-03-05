# passgen 
[![Build Status](https://travis-ci.org/JonathanSudibya/passgen.svg?branch=master)](https://travis-ci.org/JonathanSudibya/passgen)
[![codecov](https://codecov.io/gh/JonathanSudibya/passgen/branch/master/graph/badge.svg)](https://codecov.io/gh/JonathanSudibya/passgen)
[![Go Report Card](https://goreportcard.com/badge/github.com/JonathanSudibya/passgen)](https://goreportcard.com/report/github.com/JonathanSudibya/passgen)

Simple password (or string generator) library for golang

# Disclaimer

This project is inspired by cmiceli's [password-generator-go](https://github.com/cmiceli/password-generator-go)

# Benchmark
```
BenchmarkNewPasswordAllChars-4   	 2000000	       955 ns/op	      80 B/op	       2 allocs/op
BenchmarkNewPasswordCapsChars-4   	 2000000	       909 ns/op	      80 B/op	       2 allocs/op
BenchmarkNewPasswordLowerChars-4   	 2000000	       919 ns/op	      80 B/op	       2 allocs/op
BenchmarkNewPasswordNumberChars-4   	 2000000	       954 ns/op	      80 B/op	       2 allocs/op
BenchmarkNewPasswordSymbolChars-4   	 2000000	       942 ns/op	      80 B/op	       2 allocs/op
BenchmarkNewPasswordMixedChars-4   	 2000000	       950 ns/op	      80 B/op	       2 allocs/op
```

# LICENSE

Please see [LICENSE](https://github.com/JonathanSudibya/passgen/blob/master/LICENSE)