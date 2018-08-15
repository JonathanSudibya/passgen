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
BenchmarkNewPasswordAllCharsIterated/allchar-2-2         	 5000000	       300 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewPasswordAllCharsIterated/allchar-8-2         	 3000000	       353 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewPasswordAllCharsIterated/allchar-32-2        	 3000000	       506 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewPasswordAllCharsIterated/allchar-128-2       	 1000000	      1130 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewPasswordAllCharsIterated/allchar-1024-2      	  200000	      6569 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewPasswordAllChars-2                           	 3000000	       501 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewPasswordCapsChars-2                          	 3000000	       494 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewPasswordLowerChars-2                         	 3000000	       503 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewPasswordNumberChars-2                        	 3000000	       488 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewPasswordSymbolChars-2                        	 3000000	       499 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewPasswordMixedChars-2                         	 3000000	       507 ns/op	       0 B/op	       0 allocs/op
```

# LICENSE

Please see [LICENSE](https://github.com/JonathanSudibya/passgen/blob/master/LICENSE)